grammar TransactSQL;

//Entry Point

compilationUnit
    : statementList? EOF;

//Statements

statementList: statementAux+;
statementAux : statement GoCommand? endSt?;
statement
    : ( useDatabase
    | createSchemaStatement
    | createUserStatement
    | setStatement
    | createTableStatement
    | alterTableStatement
    ) endSt?;

typeName
    : '[char]' 
    | '[geography]'
    | '[int]'
    | '[money]'
    | '[varchar]'
    ;

useDatabase : 'USE' dataBaseName;

createUserStatement : 'CREATE' 'USER' userName forLoginExpression? withlimitedOptionList?;

createSchemaStatement : 'CREATE' 'SCHEMA' schemaName;

setStatement : 'SET' setOptions; 

createTableStatement : 'CREATE' 'TABLE' tableName'(' columnDefinitionList ')' createTableOptions*;

alterTableStatement : 'ALTER' 'TABLE' tableName alterTableOption;

alterTableOption
    : checkConstraint
    | withCheckOption addConstraint
    ;

withCheckOption : 'WITH' checkOption; 

checkOption
     : 'CHECK'
     |  'NOCHECK' ;


checkConstraint: 'CHECK' tableConstraint;
addConstraint : 'ADD' tableConstraint;

createTableOptions
    : onClause
    | textImageOnClause
    ;

columnDefinitionList
    : columnDefinitions
    | columnDefinitionList ',' tableConstraint
    ;

columnDefinitions
    : columnDefinition
    | columnDefinitions ',' columnDefinition
    ;

onClause :  'ON' (  fileGroup | defaultOption  );  
textImageOnClause : 'TEXTIMAGE_ON' ( fileGroup | defaultOption );

fileGroup : '[PRIMARY]';
defaultOption : '"defaul"';

columnDefinition : columnName dataType columnOption?;

tableName
    : Identifier '.' Identifier
    | Identifier '.' Identifier '.' Identifier
    | Identifier
    ;

dataType : typeName typeOption?; 

//typeColumnClause
     //: Identifier '.' typeName
 //    : typeName;

tableConstraint: 'CONSTRAINT' constraintName constraintClause?;

constraintClause
    : typeKeyClause
    | foreignKeyClause
    ;

typeKeyClause : keyOption clusterOption? '(' columnNameList ')' constraintKeyClause?;

foreignKeyClause : 'FOREIGN' 'KEY' '(' columnNameList ')' 'REFERENCES' tableName columnsTable?;

columnsTable : '(' columnNameList ')';

columnNameList
    : columnName
    | columnNameList ',' columnName
    ;

constraintKeyClause
    : onClause
    | withIndexOption
    | withIndexOption onClause
    ;

withIndexOption : 'WITH' '(' indexOptionList ')'; 

indexOptionList
    : indexOption
    | indexOptionList ',' indexOption
    ;

indexOption
    : padIndexOption
    | ignoreDupKeyOption
    | statisticsNoreComputeOption
    | statisticsIncrementalOption
    | allowPageLocksOption
    | allowRowLocksOption
    | optimizeForSequentialKeyOption
    ;

limitedOptionList
         : limitedOption
         | limitedOptionList ',' limitedOption
         ; 

userOption
    : limitedOption
    | 'SID' '=' CharacterSequence
    ;


limitedOption
    : defaultSchemaOption
    | 'DEFAULT_LANGUAGE' '=' (Identifier | 'NONE')
    ;              

defaultSchemaOption : 'DEFAULT_SCHEMA' '=' schemaName;

endSt : ';';

dataBaseName : Identifier;
//Language : ('NONE' | Identifier);



DigitSequence : DigitNoZero Digit*;

padIndexOption : 'PAD_INDEX' '=' onOffOption;
ignoreDupKeyOption : 'IGNORE_DUP_KEY' '=' onOffOption;
statisticsNoreComputeOption : 'STATISTICS_NORECOMPUTE' '=' onOffOption;
statisticsIncrementalOption: 'STATISTICS_INCREMENTAL' '=' onOffOption;
allowRowLocksOption : 'ALLOW_ROW_LOCKS' '=' onOffOption;
allowPageLocksOption : 'ALLOW_PAGE_LOCKS' '=' onOffOption;
optimizeForSequentialKeyOption : 'OPTIMIZE_FOR_SEQUENTIAL_KEY' '=' onOffOption;
//fillFactorOption : 'FILLFACTOR' '=' fillfactor;

setOptions
    : 'ANSI_NULLS' onOffOption
    | 'QUOTED_IDENTIFIER' onOffOption
    ;

onOffOption
    : 'ON'
    | 'OFF';

keyOption
    : 'PRIMARY' 'KEY'
    | 'UNIQUE'
    ;

clusterOption
    : 'CLUSTERED'
    | 'NONCLUSTERED'
    ;

orderOption
    : 'ASC'
    | 'DESC'; 

forLoginExpression : 'FOR' 'LOGIN' userName ;

withlimitedOptionList : 'WITH' limitedOptionList;     

typeOption : '(' precision scale? ')';

columnOption
    : 'NULL'
    | 'NOT' 'NULL'
    ;



schemaName : Identifier;
//simpleTableName : Identifier;     
columnName : Identifier orderOption?;
userName : Identifier;
constraintName: Identifier;
precision : DigitSequence;
scale : ',' DigitSequence;

Add : 'ADD';
AllowPageLocks : 'ALLOW_PAGE_LOCKS';
AllowRowLocks : 'ALLOW_ROW_LOCKS';
Alter : 'ALTER';
AnsiNulls : 'ANSI_NULLS';
Asc : 'ASC';

Check : 'CHECK';
CharType : '[char]';
Clustered : 'CLUSTERED';
Constraint : 'CONSTRAINT';
Create : 'CREATE'; 

Default : 'default';
DefaultLanguage : 'DEFAULT_LANGUAGE';
DefaultSchema : 'DEFAULT_SCHEMA';
Desc : 'DESC'; 

Equal : '=';

For : 'FOR';
Foreign : 'FOREIGN';

GeographyType : '[geography]';

IgnoreDupKey : 'IGNORE_DUP_KEY';
IntType : '[int]';

Key : 'KEY';

LeftParen : '(';
LeftBracket : '[';
Login : 'LOGIN';

MoneyType : '[money]';

NoCheck : 'NOCHECK';
Null : 'NULL';
Not : 'NOT';
NonClustered : 'NONCLUSTERED';

On: 'ON';
OptimizeForSequentialKey : 'OPTIMIZE_FOR_SEQUENTIAL_KEY';

PadIndex : 'PAD_INDEX';
Primary : 'PRIMARY';

References : 'REFERENCES';
RightBracket : ']';
RightParen : ')';

Schema : 'SHCEMA';
Set : 'SET';
Sid : 'SID';
StatisticsIncremental : 'STATISTICS_INCREMENTAL';
StatisticsNoreCompute : 'STATISTICS_NORECOMPUTE';

Table : 'TABLE';
TextImageOn : 'TEXTIMAGE_ON';

Unique : 'UNIQUE';
Use : 'USE';
User : 'USER';

VarcharType : '[varchar]';

With : 'WITH';

//Fragment 


Identifier : '[' CharacterSequence SimpleSpace* CharacterSequence? ']'; 

fragment CharacterSequence: CharacterNoDigit CharacterPart*;

CharacterPart
    : Character
    | SpecialCharacter
    ;

fragment CharacterNoDigit : [a-zA-Z];
fragment Character: [a-zA-Z0-9];
fragment SpecialCharacter : [_];
fragment DigitNoZero : [1-9];
fragment Digit : [0-9];
fragment SimpleEscapeSequence :  '\\' ['"?abfnrtv\\];
fragment SimpleSpace : [ ];
GoCommand : 'GO';
// Escaped lex

IgnoredQuery : GoCommand -> skip;

Whitespace
    :   [ \t]+
        -> skip
    ;

Newline
    :   (   '\r' '\n'?
        |   '\n'
        )
        -> skip
    ;

BlockComment
    :   '/*' .*? '*/'
        -> skip
    ;