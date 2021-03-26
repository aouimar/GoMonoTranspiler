// Code generated from TransactSQL.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // TransactSQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTransactSQLListener is a complete listener for a parse tree produced by TransactSQLParser.
type BaseTransactSQLListener struct{}

var _ TransactSQLListener = &BaseTransactSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTransactSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTransactSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTransactSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTransactSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseTransactSQLListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseTransactSQLListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseTransactSQLListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseTransactSQLListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatementAux is called when production statementAux is entered.
func (s *BaseTransactSQLListener) EnterStatementAux(ctx *StatementAuxContext) {}

// ExitStatementAux is called when production statementAux is exited.
func (s *BaseTransactSQLListener) ExitStatementAux(ctx *StatementAuxContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTransactSQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTransactSQLListener) ExitStatement(ctx *StatementContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseTransactSQLListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseTransactSQLListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterUseDatabase is called when production useDatabase is entered.
func (s *BaseTransactSQLListener) EnterUseDatabase(ctx *UseDatabaseContext) {}

// ExitUseDatabase is called when production useDatabase is exited.
func (s *BaseTransactSQLListener) ExitUseDatabase(ctx *UseDatabaseContext) {}

// EnterCreateUserStatement is called when production createUserStatement is entered.
func (s *BaseTransactSQLListener) EnterCreateUserStatement(ctx *CreateUserStatementContext) {}

// ExitCreateUserStatement is called when production createUserStatement is exited.
func (s *BaseTransactSQLListener) ExitCreateUserStatement(ctx *CreateUserStatementContext) {}

// EnterCreateSchemaStatement is called when production createSchemaStatement is entered.
func (s *BaseTransactSQLListener) EnterCreateSchemaStatement(ctx *CreateSchemaStatementContext) {}

// ExitCreateSchemaStatement is called when production createSchemaStatement is exited.
func (s *BaseTransactSQLListener) ExitCreateSchemaStatement(ctx *CreateSchemaStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseTransactSQLListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseTransactSQLListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseTransactSQLListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseTransactSQLListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseTransactSQLListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseTransactSQLListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterAlterTableOption is called when production alterTableOption is entered.
func (s *BaseTransactSQLListener) EnterAlterTableOption(ctx *AlterTableOptionContext) {}

// ExitAlterTableOption is called when production alterTableOption is exited.
func (s *BaseTransactSQLListener) ExitAlterTableOption(ctx *AlterTableOptionContext) {}

// EnterWithCheckOption is called when production withCheckOption is entered.
func (s *BaseTransactSQLListener) EnterWithCheckOption(ctx *WithCheckOptionContext) {}

// ExitWithCheckOption is called when production withCheckOption is exited.
func (s *BaseTransactSQLListener) ExitWithCheckOption(ctx *WithCheckOptionContext) {}

// EnterCheckOption is called when production checkOption is entered.
func (s *BaseTransactSQLListener) EnterCheckOption(ctx *CheckOptionContext) {}

// ExitCheckOption is called when production checkOption is exited.
func (s *BaseTransactSQLListener) ExitCheckOption(ctx *CheckOptionContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseTransactSQLListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseTransactSQLListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterAddConstraint is called when production addConstraint is entered.
func (s *BaseTransactSQLListener) EnterAddConstraint(ctx *AddConstraintContext) {}

// ExitAddConstraint is called when production addConstraint is exited.
func (s *BaseTransactSQLListener) ExitAddConstraint(ctx *AddConstraintContext) {}

// EnterCreateTableOptions is called when production createTableOptions is entered.
func (s *BaseTransactSQLListener) EnterCreateTableOptions(ctx *CreateTableOptionsContext) {}

// ExitCreateTableOptions is called when production createTableOptions is exited.
func (s *BaseTransactSQLListener) ExitCreateTableOptions(ctx *CreateTableOptionsContext) {}

// EnterColumnDefinitionList is called when production columnDefinitionList is entered.
func (s *BaseTransactSQLListener) EnterColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// ExitColumnDefinitionList is called when production columnDefinitionList is exited.
func (s *BaseTransactSQLListener) ExitColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// EnterColumnDefinitions is called when production columnDefinitions is entered.
func (s *BaseTransactSQLListener) EnterColumnDefinitions(ctx *ColumnDefinitionsContext) {}

// ExitColumnDefinitions is called when production columnDefinitions is exited.
func (s *BaseTransactSQLListener) ExitColumnDefinitions(ctx *ColumnDefinitionsContext) {}

// EnterOnClause is called when production onClause is entered.
func (s *BaseTransactSQLListener) EnterOnClause(ctx *OnClauseContext) {}

// ExitOnClause is called when production onClause is exited.
func (s *BaseTransactSQLListener) ExitOnClause(ctx *OnClauseContext) {}

// EnterTextImageOnClause is called when production textImageOnClause is entered.
func (s *BaseTransactSQLListener) EnterTextImageOnClause(ctx *TextImageOnClauseContext) {}

// ExitTextImageOnClause is called when production textImageOnClause is exited.
func (s *BaseTransactSQLListener) ExitTextImageOnClause(ctx *TextImageOnClauseContext) {}

// EnterFileGroup is called when production fileGroup is entered.
func (s *BaseTransactSQLListener) EnterFileGroup(ctx *FileGroupContext) {}

// ExitFileGroup is called when production fileGroup is exited.
func (s *BaseTransactSQLListener) ExitFileGroup(ctx *FileGroupContext) {}

// EnterDefaultOption is called when production defaultOption is entered.
func (s *BaseTransactSQLListener) EnterDefaultOption(ctx *DefaultOptionContext) {}

// ExitDefaultOption is called when production defaultOption is exited.
func (s *BaseTransactSQLListener) ExitDefaultOption(ctx *DefaultOptionContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseTransactSQLListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseTransactSQLListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseTransactSQLListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseTransactSQLListener) ExitTableName(ctx *TableNameContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseTransactSQLListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseTransactSQLListener) ExitDataType(ctx *DataTypeContext) {}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseTransactSQLListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseTransactSQLListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterConstraintClause is called when production constraintClause is entered.
func (s *BaseTransactSQLListener) EnterConstraintClause(ctx *ConstraintClauseContext) {}

// ExitConstraintClause is called when production constraintClause is exited.
func (s *BaseTransactSQLListener) ExitConstraintClause(ctx *ConstraintClauseContext) {}

// EnterTypeKeyClause is called when production typeKeyClause is entered.
func (s *BaseTransactSQLListener) EnterTypeKeyClause(ctx *TypeKeyClauseContext) {}

// ExitTypeKeyClause is called when production typeKeyClause is exited.
func (s *BaseTransactSQLListener) ExitTypeKeyClause(ctx *TypeKeyClauseContext) {}

// EnterForeignKeyClause is called when production foreignKeyClause is entered.
func (s *BaseTransactSQLListener) EnterForeignKeyClause(ctx *ForeignKeyClauseContext) {}

// ExitForeignKeyClause is called when production foreignKeyClause is exited.
func (s *BaseTransactSQLListener) ExitForeignKeyClause(ctx *ForeignKeyClauseContext) {}

// EnterColumnsTable is called when production columnsTable is entered.
func (s *BaseTransactSQLListener) EnterColumnsTable(ctx *ColumnsTableContext) {}

// ExitColumnsTable is called when production columnsTable is exited.
func (s *BaseTransactSQLListener) ExitColumnsTable(ctx *ColumnsTableContext) {}

// EnterColumnNameList is called when production columnNameList is entered.
func (s *BaseTransactSQLListener) EnterColumnNameList(ctx *ColumnNameListContext) {}

// ExitColumnNameList is called when production columnNameList is exited.
func (s *BaseTransactSQLListener) ExitColumnNameList(ctx *ColumnNameListContext) {}

// EnterConstraintKeyClause is called when production constraintKeyClause is entered.
func (s *BaseTransactSQLListener) EnterConstraintKeyClause(ctx *ConstraintKeyClauseContext) {}

// ExitConstraintKeyClause is called when production constraintKeyClause is exited.
func (s *BaseTransactSQLListener) ExitConstraintKeyClause(ctx *ConstraintKeyClauseContext) {}

// EnterWithIndexOption is called when production withIndexOption is entered.
func (s *BaseTransactSQLListener) EnterWithIndexOption(ctx *WithIndexOptionContext) {}

// ExitWithIndexOption is called when production withIndexOption is exited.
func (s *BaseTransactSQLListener) ExitWithIndexOption(ctx *WithIndexOptionContext) {}

// EnterIndexOptionList is called when production indexOptionList is entered.
func (s *BaseTransactSQLListener) EnterIndexOptionList(ctx *IndexOptionListContext) {}

// ExitIndexOptionList is called when production indexOptionList is exited.
func (s *BaseTransactSQLListener) ExitIndexOptionList(ctx *IndexOptionListContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseTransactSQLListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseTransactSQLListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterLimitedOptionList is called when production limitedOptionList is entered.
func (s *BaseTransactSQLListener) EnterLimitedOptionList(ctx *LimitedOptionListContext) {}

// ExitLimitedOptionList is called when production limitedOptionList is exited.
func (s *BaseTransactSQLListener) ExitLimitedOptionList(ctx *LimitedOptionListContext) {}

// EnterUserOption is called when production userOption is entered.
func (s *BaseTransactSQLListener) EnterUserOption(ctx *UserOptionContext) {}

// ExitUserOption is called when production userOption is exited.
func (s *BaseTransactSQLListener) ExitUserOption(ctx *UserOptionContext) {}

// EnterLimitedOption is called when production limitedOption is entered.
func (s *BaseTransactSQLListener) EnterLimitedOption(ctx *LimitedOptionContext) {}

// ExitLimitedOption is called when production limitedOption is exited.
func (s *BaseTransactSQLListener) ExitLimitedOption(ctx *LimitedOptionContext) {}

// EnterDefaultSchemaOption is called when production defaultSchemaOption is entered.
func (s *BaseTransactSQLListener) EnterDefaultSchemaOption(ctx *DefaultSchemaOptionContext) {}

// ExitDefaultSchemaOption is called when production defaultSchemaOption is exited.
func (s *BaseTransactSQLListener) ExitDefaultSchemaOption(ctx *DefaultSchemaOptionContext) {}

// EnterEndSt is called when production endSt is entered.
func (s *BaseTransactSQLListener) EnterEndSt(ctx *EndStContext) {}

// ExitEndSt is called when production endSt is exited.
func (s *BaseTransactSQLListener) ExitEndSt(ctx *EndStContext) {}

// EnterDataBaseName is called when production dataBaseName is entered.
func (s *BaseTransactSQLListener) EnterDataBaseName(ctx *DataBaseNameContext) {}

// ExitDataBaseName is called when production dataBaseName is exited.
func (s *BaseTransactSQLListener) ExitDataBaseName(ctx *DataBaseNameContext) {}

// EnterPadIndexOption is called when production padIndexOption is entered.
func (s *BaseTransactSQLListener) EnterPadIndexOption(ctx *PadIndexOptionContext) {}

// ExitPadIndexOption is called when production padIndexOption is exited.
func (s *BaseTransactSQLListener) ExitPadIndexOption(ctx *PadIndexOptionContext) {}

// EnterIgnoreDupKeyOption is called when production ignoreDupKeyOption is entered.
func (s *BaseTransactSQLListener) EnterIgnoreDupKeyOption(ctx *IgnoreDupKeyOptionContext) {}

// ExitIgnoreDupKeyOption is called when production ignoreDupKeyOption is exited.
func (s *BaseTransactSQLListener) ExitIgnoreDupKeyOption(ctx *IgnoreDupKeyOptionContext) {}

// EnterStatisticsNoreComputeOption is called when production statisticsNoreComputeOption is entered.
func (s *BaseTransactSQLListener) EnterStatisticsNoreComputeOption(ctx *StatisticsNoreComputeOptionContext) {
}

// ExitStatisticsNoreComputeOption is called when production statisticsNoreComputeOption is exited.
func (s *BaseTransactSQLListener) ExitStatisticsNoreComputeOption(ctx *StatisticsNoreComputeOptionContext) {
}

// EnterStatisticsIncrementalOption is called when production statisticsIncrementalOption is entered.
func (s *BaseTransactSQLListener) EnterStatisticsIncrementalOption(ctx *StatisticsIncrementalOptionContext) {
}

// ExitStatisticsIncrementalOption is called when production statisticsIncrementalOption is exited.
func (s *BaseTransactSQLListener) ExitStatisticsIncrementalOption(ctx *StatisticsIncrementalOptionContext) {
}

// EnterAllowRowLocksOption is called when production allowRowLocksOption is entered.
func (s *BaseTransactSQLListener) EnterAllowRowLocksOption(ctx *AllowRowLocksOptionContext) {}

// ExitAllowRowLocksOption is called when production allowRowLocksOption is exited.
func (s *BaseTransactSQLListener) ExitAllowRowLocksOption(ctx *AllowRowLocksOptionContext) {}

// EnterAllowPageLocksOption is called when production allowPageLocksOption is entered.
func (s *BaseTransactSQLListener) EnterAllowPageLocksOption(ctx *AllowPageLocksOptionContext) {}

// ExitAllowPageLocksOption is called when production allowPageLocksOption is exited.
func (s *BaseTransactSQLListener) ExitAllowPageLocksOption(ctx *AllowPageLocksOptionContext) {}

// EnterOptimizeForSequentialKeyOption is called when production optimizeForSequentialKeyOption is entered.
func (s *BaseTransactSQLListener) EnterOptimizeForSequentialKeyOption(ctx *OptimizeForSequentialKeyOptionContext) {
}

// ExitOptimizeForSequentialKeyOption is called when production optimizeForSequentialKeyOption is exited.
func (s *BaseTransactSQLListener) ExitOptimizeForSequentialKeyOption(ctx *OptimizeForSequentialKeyOptionContext) {
}

// EnterSetOptions is called when production setOptions is entered.
func (s *BaseTransactSQLListener) EnterSetOptions(ctx *SetOptionsContext) {}

// ExitSetOptions is called when production setOptions is exited.
func (s *BaseTransactSQLListener) ExitSetOptions(ctx *SetOptionsContext) {}

// EnterOnOffOption is called when production onOffOption is entered.
func (s *BaseTransactSQLListener) EnterOnOffOption(ctx *OnOffOptionContext) {}

// ExitOnOffOption is called when production onOffOption is exited.
func (s *BaseTransactSQLListener) ExitOnOffOption(ctx *OnOffOptionContext) {}

// EnterKeyOption is called when production keyOption is entered.
func (s *BaseTransactSQLListener) EnterKeyOption(ctx *KeyOptionContext) {}

// ExitKeyOption is called when production keyOption is exited.
func (s *BaseTransactSQLListener) ExitKeyOption(ctx *KeyOptionContext) {}

// EnterClusterOption is called when production clusterOption is entered.
func (s *BaseTransactSQLListener) EnterClusterOption(ctx *ClusterOptionContext) {}

// ExitClusterOption is called when production clusterOption is exited.
func (s *BaseTransactSQLListener) ExitClusterOption(ctx *ClusterOptionContext) {}

// EnterOrderOption is called when production orderOption is entered.
func (s *BaseTransactSQLListener) EnterOrderOption(ctx *OrderOptionContext) {}

// ExitOrderOption is called when production orderOption is exited.
func (s *BaseTransactSQLListener) ExitOrderOption(ctx *OrderOptionContext) {}

// EnterForLoginExpression is called when production forLoginExpression is entered.
func (s *BaseTransactSQLListener) EnterForLoginExpression(ctx *ForLoginExpressionContext) {}

// ExitForLoginExpression is called when production forLoginExpression is exited.
func (s *BaseTransactSQLListener) ExitForLoginExpression(ctx *ForLoginExpressionContext) {}

// EnterWithlimitedOptionList is called when production withlimitedOptionList is entered.
func (s *BaseTransactSQLListener) EnterWithlimitedOptionList(ctx *WithlimitedOptionListContext) {}

// ExitWithlimitedOptionList is called when production withlimitedOptionList is exited.
func (s *BaseTransactSQLListener) ExitWithlimitedOptionList(ctx *WithlimitedOptionListContext) {}

// EnterTypeOption is called when production typeOption is entered.
func (s *BaseTransactSQLListener) EnterTypeOption(ctx *TypeOptionContext) {}

// ExitTypeOption is called when production typeOption is exited.
func (s *BaseTransactSQLListener) ExitTypeOption(ctx *TypeOptionContext) {}

// EnterColumnOption is called when production columnOption is entered.
func (s *BaseTransactSQLListener) EnterColumnOption(ctx *ColumnOptionContext) {}

// ExitColumnOption is called when production columnOption is exited.
func (s *BaseTransactSQLListener) ExitColumnOption(ctx *ColumnOptionContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseTransactSQLListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseTransactSQLListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseTransactSQLListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseTransactSQLListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterUserName is called when production userName is entered.
func (s *BaseTransactSQLListener) EnterUserName(ctx *UserNameContext) {}

// ExitUserName is called when production userName is exited.
func (s *BaseTransactSQLListener) ExitUserName(ctx *UserNameContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseTransactSQLListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseTransactSQLListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterPrecision is called when production precision is entered.
func (s *BaseTransactSQLListener) EnterPrecision(ctx *PrecisionContext) {}

// ExitPrecision is called when production precision is exited.
func (s *BaseTransactSQLListener) ExitPrecision(ctx *PrecisionContext) {}

// EnterScale is called when production scale is entered.
func (s *BaseTransactSQLListener) EnterScale(ctx *ScaleContext) {}

// ExitScale is called when production scale is exited.
func (s *BaseTransactSQLListener) ExitScale(ctx *ScaleContext) {}
