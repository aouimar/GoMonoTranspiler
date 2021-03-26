package translater

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"aouimar.com/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var tsqlErrorListener *TransactSQLErrorListener

type TransactSQLListener struct {
	*parser.BaseTransactSQLListener
	stack []string
}

type GlobalContext struct {
	*antlr.BaseParserRuleContext
}

func (g *GlobalContext) GetChildren() []antlr.Tree {
	return g.BaseParserRuleContext.GetChildren()
}

func (l *TransactSQLListener) push(s string) {
	l.stack = append(l.stack, s)
}

func (l *TransactSQLListener) pop() string {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func CheckCreateUserContext(ctx *parser.CreateUserStatementContext) (bool, bool) {
	forLogin := parser.NewForLoginExpressionContext(
		ctx.GetParser(),
		ctx.GetParser().GetParserRuleContext(),
		0)

	defaultSchema := parser.NewDefaultSchemaOptionContext(
		ctx.GetParser(),
		ctx.GetParser().GetParserRuleContext(),
		0)

	schema := false
	login := false

	if ContextContainSubContext(antlr.Tree(ctx), reflect.TypeOf(defaultSchema)) {
		schema = true
	}

	if ContextContainSubContext(antlr.Tree(ctx), reflect.TypeOf(forLogin)) {
		login = true
	}

	return login, schema
}

// CREATE USER username FOR LOGIN loginName WITH DEFAULT_SCHEMA=schemaName
// ==
// CREATE USER username WITH LOGIN ROLE roleName
// GRANT ALL PRIVILEGES ON SCHEMA schemaName TO roleName

func (l *TransactSQLListener) ExitCreateUserStatement(ctx *parser.CreateUserStatementContext) {
	schemaName, userName, loginName := "", "", ""
	login, schema := CheckCreateUserContext(ctx)

	if login && schema {
		schemaName, loginName, userName = l.pop(), l.pop(), l.pop()
		l.push(fmt.Sprintf("CREATE USER %s WITH LOGIN ROLE %s;", userName, loginName))
		l.push(fmt.Sprintf("GRANT ALL PRIVILEGES ON SCHEMA %s TO %s;", schemaName, loginName))
	} else if schema {
		schemaName, userName = l.pop(), l.pop()
		l.push(fmt.Sprintf("CREATE USER %s WITH ROLE %s;", userName, userName))
		l.push(fmt.Sprintf("GRANT ALL PRIVILEGES ON SCHEMA %s TO %s;", schemaName, userName))
	} else if login {
		userName, loginName = l.pop(), l.pop()
		l.push(fmt.Sprintf("CREATE USER %s WITH LOGIN ROLE %s;", userName, loginName))
	} else {
		userName = l.pop()
		l.push(fmt.Sprintf("CREATE USER %s", userName))
	}
}

func (l *TransactSQLListener) ExitSchemaName(ctx *parser.SchemaNameContext) {
	var s []string
	if CheckAndFormatSubContext(&s, ctx) {
		l.push(strings.Join(s, ""))
	}
}

func (l *TransactSQLListener) ExitUserName(ctx *parser.UserNameContext) {
	var s []string
	if CheckAndFormatSubContext(&s, ctx) {
		l.push(strings.Join(s, ""))
	}
}

func (l *TransactSQLListener) ExitDataBaseName(ctx *parser.DataBaseNameContext) {
	var s []string
	if CheckAndFormatSubContext(&s, ctx) {
		l.push(strings.Join(s, ""))
	}
}

func (l *TransactSQLListener) ExitUseDatabase(ctx *parser.UseDatabaseContext) {
	databaseName := l.pop()
	l.push("\\c " + databaseName + ";")
}

func ConvertBrackets(s string, tStr string) string {
	ss := strings.ReplaceAll(s, "[", tStr)
	return strings.ReplaceAll(ss, "]", tStr)
}

func PrettyFormat(s string) string {
	ss := strings.ReplaceAll(s, "( ", "(")
	return strings.ReplaceAll(ss, " )", ")")
}

func CheckAndFormatSubContext(s *[]string, tree antlr.Tree) bool {
	var isVerifying bool = false

	switch tree.(type) {
	case *parser.TableNameContext,
		*parser.SchemaNameContext,
		*parser.ConstraintNameContext,
		*parser.DataBaseNameContext,
		*parser.UserNameContext:
		{
			FormatSubContext(s, tree, "", "\"")
			isVerifying = true
		}
	case *parser.ColumnNameContext:
		{
			FormatSubContext(s, tree, " ", "")
			isVerifying = true
		}
	case *parser.DataTypeContext:
		{
			FormatSubContext(s, tree, "", "")
			isVerifying = true
		}
	case *parser.ClusterOptionContext,
		*parser.ConstraintKeyClauseContext,
		*parser.CreateTableOptionsContext,
		*parser.WithCheckOptionContext:
		{
			// we get rid of clauses have no meaning in target
			isVerifying = true
		}
	}

	return isVerifying
}

func FormatSubContext(s *[]string, tree antlr.Tree, jointure string, tbrackets string) {
	var ss []string
	GetName(&ss, tree)
	*s = append(*s, ConvertBrackets(strings.Join(ss, jointure), tbrackets))
}

func GetName(s *[]string, tree antlr.Tree) {
	childs := tree.GetChildren()
	if len(childs) != 0 {
		for _, child := range childs {
			if len(child.GetChildren()) != 0 {
				GetName(s, child)
			} else {
				*s = append(*s, fmt.Sprintf("%s", child))
			}
		}
	}
}

func GetStringsFromChild(s *[]string, tree antlr.Tree) {
	childs := tree.GetChildren()
	if len(childs) != 0 {
		for _, child := range childs {
			if !CheckAndFormatSubContext(s, child) {
				if len(child.GetChildren()) != 0 {
					GetStringsFromChild(s, child)
				} else {
					*s = append(*s, fmt.Sprintf("%s", child))
				}
			}
		}
	}
}

func (l *TransactSQLListener) PrintThroughContext(ctx GlobalContext) {
	tree := ctx.GetChildren()
	var result []string
	for _, leaf := range tree {
		if len(leaf.GetChildren()) != 0 {
			GetStringsFromChild(&result, leaf)
		} else {
			result = append(result, fmt.Sprintf("%s", leaf))
		}
	}
	l.push(strings.Join(result[:], " ") + ";")
}

func ContextContainSubContext(t antlr.Tree, theType reflect.Type) bool {
	var containing bool = false
	tree := t.GetChildren()

	for _, leaf := range tree {
		ttype := reflect.TypeOf(leaf)
		if ttype == theType {
			containing = true
			break
		} else {
			if len(leaf.GetChildren()) != 0 {
				containing = ContextContainSubContext(leaf, theType)
			}
		}

		if containing {
			break
		}
	}

	return containing
}

func (l *TransactSQLListener) ExitStatement(ctx *parser.StatementContext) {

	tree := ctx.GetChild(0)
	switch tree.(type) {
	case *parser.UseDatabaseContext,
		*parser.SetStatementContext,
		*parser.CreateUserStatementContext:
		{
			// statement changed or have no meaning
			break
		}
	case *parser.AlterTableStatementContext:
		{
			//get rid of it if have a specific option otherwise print through normally
			//fake object just for reflect type use
			checkConstraint := parser.NewCheckConstraintContext(
				ctx.GetParser(),
				ctx.GetParser().GetParserRuleContext(),
				0)

			// I dont know a cleaner way than that, it's nasty (go doesnt have dynamic typing)
			// It's too basic too basic, but it is quick
			if !ContextContainSubContext(antlr.Tree(ctx), reflect.TypeOf(checkConstraint)) {
				l.PushStatementContext(ctx)
			}
			break
		}
	case *parser.CreateSchemaStatementContext:
		{
			l.pop() // we pop schema Name
			l.PushStatementContext(ctx)
			break
		}
	default:
		{
			//no change occuring in the statement except formatting or modifying some option
			l.PushStatementContext(ctx)
			break
		}
	}
}

func (l *TransactSQLListener) PushStatementContext(ctx *parser.StatementContext) {
	var c *antlr.BaseParserRuleContext = ctx.BaseParserRuleContext
	var co = GlobalContext{c}
	l.PrintThroughContext(co)
}

func Translate(input string) (string, bool) {
	var returnedString string = ""
	var atError bool = false
	tsqlErrorListener = &TransactSQLErrorListener{0, make(map[int]*TransactSQLError)}
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewTransactSQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTransactSQLParser(stream)

	var listener TransactSQLListener

	p.AddErrorListener(tsqlErrorListener)
	lexer.AddErrorListener(tsqlErrorListener)
	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.CompilationUnit())

	if tsqlErrorListener.nErrors > 0 {
		for i, err := range tsqlErrorListener.errors {
			returnedString += fmt.Sprintf("%d Error in line %d column %d with msg: %s\n",
				i, err.line, err.column, err.msg)
			atError = true
		}
	} else {
		returnedString = PrettyFormat(strings.Join(listener.stack[:], "\n"))
	}

	//print the result
	return returnedString, atError
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TranslateFile(pathFile string, outputFile string) {
	dataFile, err1 := ioutil.ReadFile(pathFile)
	check(err1)
	dataTargetFile, err := Translate(string(dataFile))

	if err {
		fmt.Printf(dataTargetFile + "The input file is not parsable as tsql!! see the grammar")
	} else {
		err2 := ioutil.WriteFile(outputFile, []byte(dataTargetFile), 0644)
		check(err2)
	}
}
