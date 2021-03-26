package translater

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TransactSQLError struct {
	line   int
	column int
	msg    string
}

type TransactSQLErrorListener struct {
	nErrors int
	errors  map[int]*TransactSQLError
}

func (l *TransactSQLErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	var tError = TransactSQLError{line, column, msg}
	l.errors[l.nErrors] = &tError
	l.nErrors += 1
}

func (l *TransactSQLErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	var tError = TransactSQLError{startIndex, stopIndex, "Ambiguity"}
	l.errors[l.nErrors] = &tError
	l.nErrors += 1
}
func (l *TransactSQLErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	var tError = TransactSQLError{startIndex, stopIndex, "Attempting Full Context"}
	l.errors[l.nErrors] = &tError
	l.nErrors += 1
}
func (l *TransactSQLErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	var tError = TransactSQLError{startIndex, stopIndex, "Context Sensitivity"}
	l.errors[l.nErrors] = &tError
	l.nErrors += 1
}
