// Code generated from TransactSQL.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // TransactSQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TransactSQLListener is a complete listener for a parse tree produced by TransactSQLParser.
type TransactSQLListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatementAux is called when entering the statementAux production.
	EnterStatementAux(c *StatementAuxContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterUseDatabase is called when entering the useDatabase production.
	EnterUseDatabase(c *UseDatabaseContext)

	// EnterCreateUserStatement is called when entering the createUserStatement production.
	EnterCreateUserStatement(c *CreateUserStatementContext)

	// EnterCreateSchemaStatement is called when entering the createSchemaStatement production.
	EnterCreateSchemaStatement(c *CreateSchemaStatementContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterAlterTableStatement is called when entering the alterTableStatement production.
	EnterAlterTableStatement(c *AlterTableStatementContext)

	// EnterAlterTableOption is called when entering the alterTableOption production.
	EnterAlterTableOption(c *AlterTableOptionContext)

	// EnterWithCheckOption is called when entering the withCheckOption production.
	EnterWithCheckOption(c *WithCheckOptionContext)

	// EnterCheckOption is called when entering the checkOption production.
	EnterCheckOption(c *CheckOptionContext)

	// EnterCheckConstraint is called when entering the checkConstraint production.
	EnterCheckConstraint(c *CheckConstraintContext)

	// EnterAddConstraint is called when entering the addConstraint production.
	EnterAddConstraint(c *AddConstraintContext)

	// EnterCreateTableOptions is called when entering the createTableOptions production.
	EnterCreateTableOptions(c *CreateTableOptionsContext)

	// EnterColumnDefinitionList is called when entering the columnDefinitionList production.
	EnterColumnDefinitionList(c *ColumnDefinitionListContext)

	// EnterColumnDefinitions is called when entering the columnDefinitions production.
	EnterColumnDefinitions(c *ColumnDefinitionsContext)

	// EnterOnClause is called when entering the onClause production.
	EnterOnClause(c *OnClauseContext)

	// EnterTextImageOnClause is called when entering the textImageOnClause production.
	EnterTextImageOnClause(c *TextImageOnClauseContext)

	// EnterFileGroup is called when entering the fileGroup production.
	EnterFileGroup(c *FileGroupContext)

	// EnterDefaultOption is called when entering the defaultOption production.
	EnterDefaultOption(c *DefaultOptionContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterTableConstraint is called when entering the tableConstraint production.
	EnterTableConstraint(c *TableConstraintContext)

	// EnterConstraintClause is called when entering the constraintClause production.
	EnterConstraintClause(c *ConstraintClauseContext)

	// EnterTypeKeyClause is called when entering the typeKeyClause production.
	EnterTypeKeyClause(c *TypeKeyClauseContext)

	// EnterForeignKeyClause is called when entering the foreignKeyClause production.
	EnterForeignKeyClause(c *ForeignKeyClauseContext)

	// EnterColumnsTable is called when entering the columnsTable production.
	EnterColumnsTable(c *ColumnsTableContext)

	// EnterColumnNameList is called when entering the columnNameList production.
	EnterColumnNameList(c *ColumnNameListContext)

	// EnterConstraintKeyClause is called when entering the constraintKeyClause production.
	EnterConstraintKeyClause(c *ConstraintKeyClauseContext)

	// EnterWithIndexOption is called when entering the withIndexOption production.
	EnterWithIndexOption(c *WithIndexOptionContext)

	// EnterIndexOptionList is called when entering the indexOptionList production.
	EnterIndexOptionList(c *IndexOptionListContext)

	// EnterIndexOption is called when entering the indexOption production.
	EnterIndexOption(c *IndexOptionContext)

	// EnterLimitedOptionList is called when entering the limitedOptionList production.
	EnterLimitedOptionList(c *LimitedOptionListContext)

	// EnterUserOption is called when entering the userOption production.
	EnterUserOption(c *UserOptionContext)

	// EnterLimitedOption is called when entering the limitedOption production.
	EnterLimitedOption(c *LimitedOptionContext)

	// EnterDefaultSchemaOption is called when entering the defaultSchemaOption production.
	EnterDefaultSchemaOption(c *DefaultSchemaOptionContext)

	// EnterEndSt is called when entering the endSt production.
	EnterEndSt(c *EndStContext)

	// EnterDataBaseName is called when entering the dataBaseName production.
	EnterDataBaseName(c *DataBaseNameContext)

	// EnterPadIndexOption is called when entering the padIndexOption production.
	EnterPadIndexOption(c *PadIndexOptionContext)

	// EnterIgnoreDupKeyOption is called when entering the ignoreDupKeyOption production.
	EnterIgnoreDupKeyOption(c *IgnoreDupKeyOptionContext)

	// EnterStatisticsNoreComputeOption is called when entering the statisticsNoreComputeOption production.
	EnterStatisticsNoreComputeOption(c *StatisticsNoreComputeOptionContext)

	// EnterStatisticsIncrementalOption is called when entering the statisticsIncrementalOption production.
	EnterStatisticsIncrementalOption(c *StatisticsIncrementalOptionContext)

	// EnterAllowRowLocksOption is called when entering the allowRowLocksOption production.
	EnterAllowRowLocksOption(c *AllowRowLocksOptionContext)

	// EnterAllowPageLocksOption is called when entering the allowPageLocksOption production.
	EnterAllowPageLocksOption(c *AllowPageLocksOptionContext)

	// EnterOptimizeForSequentialKeyOption is called when entering the optimizeForSequentialKeyOption production.
	EnterOptimizeForSequentialKeyOption(c *OptimizeForSequentialKeyOptionContext)

	// EnterSetOptions is called when entering the setOptions production.
	EnterSetOptions(c *SetOptionsContext)

	// EnterOnOffOption is called when entering the onOffOption production.
	EnterOnOffOption(c *OnOffOptionContext)

	// EnterKeyOption is called when entering the keyOption production.
	EnterKeyOption(c *KeyOptionContext)

	// EnterClusterOption is called when entering the clusterOption production.
	EnterClusterOption(c *ClusterOptionContext)

	// EnterOrderOption is called when entering the orderOption production.
	EnterOrderOption(c *OrderOptionContext)

	// EnterForLoginExpression is called when entering the forLoginExpression production.
	EnterForLoginExpression(c *ForLoginExpressionContext)

	// EnterWithlimitedOptionList is called when entering the withlimitedOptionList production.
	EnterWithlimitedOptionList(c *WithlimitedOptionListContext)

	// EnterTypeOption is called when entering the typeOption production.
	EnterTypeOption(c *TypeOptionContext)

	// EnterColumnOption is called when entering the columnOption production.
	EnterColumnOption(c *ColumnOptionContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterUserName is called when entering the userName production.
	EnterUserName(c *UserNameContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterPrecision is called when entering the precision production.
	EnterPrecision(c *PrecisionContext)

	// EnterScale is called when entering the scale production.
	EnterScale(c *ScaleContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatementAux is called when exiting the statementAux production.
	ExitStatementAux(c *StatementAuxContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitUseDatabase is called when exiting the useDatabase production.
	ExitUseDatabase(c *UseDatabaseContext)

	// ExitCreateUserStatement is called when exiting the createUserStatement production.
	ExitCreateUserStatement(c *CreateUserStatementContext)

	// ExitCreateSchemaStatement is called when exiting the createSchemaStatement production.
	ExitCreateSchemaStatement(c *CreateSchemaStatementContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitAlterTableStatement is called when exiting the alterTableStatement production.
	ExitAlterTableStatement(c *AlterTableStatementContext)

	// ExitAlterTableOption is called when exiting the alterTableOption production.
	ExitAlterTableOption(c *AlterTableOptionContext)

	// ExitWithCheckOption is called when exiting the withCheckOption production.
	ExitWithCheckOption(c *WithCheckOptionContext)

	// ExitCheckOption is called when exiting the checkOption production.
	ExitCheckOption(c *CheckOptionContext)

	// ExitCheckConstraint is called when exiting the checkConstraint production.
	ExitCheckConstraint(c *CheckConstraintContext)

	// ExitAddConstraint is called when exiting the addConstraint production.
	ExitAddConstraint(c *AddConstraintContext)

	// ExitCreateTableOptions is called when exiting the createTableOptions production.
	ExitCreateTableOptions(c *CreateTableOptionsContext)

	// ExitColumnDefinitionList is called when exiting the columnDefinitionList production.
	ExitColumnDefinitionList(c *ColumnDefinitionListContext)

	// ExitColumnDefinitions is called when exiting the columnDefinitions production.
	ExitColumnDefinitions(c *ColumnDefinitionsContext)

	// ExitOnClause is called when exiting the onClause production.
	ExitOnClause(c *OnClauseContext)

	// ExitTextImageOnClause is called when exiting the textImageOnClause production.
	ExitTextImageOnClause(c *TextImageOnClauseContext)

	// ExitFileGroup is called when exiting the fileGroup production.
	ExitFileGroup(c *FileGroupContext)

	// ExitDefaultOption is called when exiting the defaultOption production.
	ExitDefaultOption(c *DefaultOptionContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitTableConstraint is called when exiting the tableConstraint production.
	ExitTableConstraint(c *TableConstraintContext)

	// ExitConstraintClause is called when exiting the constraintClause production.
	ExitConstraintClause(c *ConstraintClauseContext)

	// ExitTypeKeyClause is called when exiting the typeKeyClause production.
	ExitTypeKeyClause(c *TypeKeyClauseContext)

	// ExitForeignKeyClause is called when exiting the foreignKeyClause production.
	ExitForeignKeyClause(c *ForeignKeyClauseContext)

	// ExitColumnsTable is called when exiting the columnsTable production.
	ExitColumnsTable(c *ColumnsTableContext)

	// ExitColumnNameList is called when exiting the columnNameList production.
	ExitColumnNameList(c *ColumnNameListContext)

	// ExitConstraintKeyClause is called when exiting the constraintKeyClause production.
	ExitConstraintKeyClause(c *ConstraintKeyClauseContext)

	// ExitWithIndexOption is called when exiting the withIndexOption production.
	ExitWithIndexOption(c *WithIndexOptionContext)

	// ExitIndexOptionList is called when exiting the indexOptionList production.
	ExitIndexOptionList(c *IndexOptionListContext)

	// ExitIndexOption is called when exiting the indexOption production.
	ExitIndexOption(c *IndexOptionContext)

	// ExitLimitedOptionList is called when exiting the limitedOptionList production.
	ExitLimitedOptionList(c *LimitedOptionListContext)

	// ExitUserOption is called when exiting the userOption production.
	ExitUserOption(c *UserOptionContext)

	// ExitLimitedOption is called when exiting the limitedOption production.
	ExitLimitedOption(c *LimitedOptionContext)

	// ExitDefaultSchemaOption is called when exiting the defaultSchemaOption production.
	ExitDefaultSchemaOption(c *DefaultSchemaOptionContext)

	// ExitEndSt is called when exiting the endSt production.
	ExitEndSt(c *EndStContext)

	// ExitDataBaseName is called when exiting the dataBaseName production.
	ExitDataBaseName(c *DataBaseNameContext)

	// ExitPadIndexOption is called when exiting the padIndexOption production.
	ExitPadIndexOption(c *PadIndexOptionContext)

	// ExitIgnoreDupKeyOption is called when exiting the ignoreDupKeyOption production.
	ExitIgnoreDupKeyOption(c *IgnoreDupKeyOptionContext)

	// ExitStatisticsNoreComputeOption is called when exiting the statisticsNoreComputeOption production.
	ExitStatisticsNoreComputeOption(c *StatisticsNoreComputeOptionContext)

	// ExitStatisticsIncrementalOption is called when exiting the statisticsIncrementalOption production.
	ExitStatisticsIncrementalOption(c *StatisticsIncrementalOptionContext)

	// ExitAllowRowLocksOption is called when exiting the allowRowLocksOption production.
	ExitAllowRowLocksOption(c *AllowRowLocksOptionContext)

	// ExitAllowPageLocksOption is called when exiting the allowPageLocksOption production.
	ExitAllowPageLocksOption(c *AllowPageLocksOptionContext)

	// ExitOptimizeForSequentialKeyOption is called when exiting the optimizeForSequentialKeyOption production.
	ExitOptimizeForSequentialKeyOption(c *OptimizeForSequentialKeyOptionContext)

	// ExitSetOptions is called when exiting the setOptions production.
	ExitSetOptions(c *SetOptionsContext)

	// ExitOnOffOption is called when exiting the onOffOption production.
	ExitOnOffOption(c *OnOffOptionContext)

	// ExitKeyOption is called when exiting the keyOption production.
	ExitKeyOption(c *KeyOptionContext)

	// ExitClusterOption is called when exiting the clusterOption production.
	ExitClusterOption(c *ClusterOptionContext)

	// ExitOrderOption is called when exiting the orderOption production.
	ExitOrderOption(c *OrderOptionContext)

	// ExitForLoginExpression is called when exiting the forLoginExpression production.
	ExitForLoginExpression(c *ForLoginExpressionContext)

	// ExitWithlimitedOptionList is called when exiting the withlimitedOptionList production.
	ExitWithlimitedOptionList(c *WithlimitedOptionListContext)

	// ExitTypeOption is called when exiting the typeOption production.
	ExitTypeOption(c *TypeOptionContext)

	// ExitColumnOption is called when exiting the columnOption production.
	ExitColumnOption(c *ColumnOptionContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitUserName is called when exiting the userName production.
	ExitUserName(c *UserNameContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitPrecision is called when exiting the precision production.
	ExitPrecision(c *PrecisionContext)

	// ExitScale is called when exiting the scale production.
	ExitScale(c *ScaleContext)
}
