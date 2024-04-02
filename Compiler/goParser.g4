parser grammar goParser;

options {
  tokenVocab = goScanner;
}

root :                      PACKAGE ID SEMI topDeclarationList;
topDeclarationList :        (variableDecl | typeDecl | funcDecl)* ;
variableDecl :              VAR (singleVarDecl SEMI | LP innerVarDecls RP SEMI | LP RP SEMI);
innerVarDecls :             singleVarDecl SEMI (singleVarDecl SEMI)*;
singleVarDecl :             identifierList declType ASOP expressionList
                            | identifierList ASOP expressionList
                            | singleVarDeclNoExps;
singleVarDeclNoExps	:       identifierList declType;
typeDecl :                  TYPE (singleTypeDecl SEMI | LP innerTypeDecls RP SEMI | LP RP SEMI);
innerTypeDecls :            singleTypeDecl SEMI (singleTypeDecl SEMI)* ;
singleTypeDecl :            ID declType;
funcDecl :                  funcFrontDecl block SEMI;
funcFrontDecl :             FUNC ID LP (funcArgDecls|EPSILON) RP (declType|EPSILON);
funcArgDecls :              singleVarDeclNoExps (CM singleVarDeclNoExps)* ;
declType :                  LP declType RP
                            | ID
                            | sliceDeclType
                            | arrayDeclType
                            | structDeclType;
sliceDeclType :             LSB RSB declType;
arrayDeclType :             LSB INTLITERAL RSB declType;
structDeclType :            STRUCT LCB (structMemDecls|EPSILON) RCB;
structMemDecls :            singleVarDeclNoExps SEMI (singleVarDeclNoExps SEMI)*;
identifierList :            ID (CM ID)*;
expression :                primaryExpression
                            | expression operator expression
                            | operator expression;
expressionList :            expression (CM expression)*;
primaryExpression :         operand
                            | primaryExpression (selector | index | arguments )
                            | appendExpression
                            | lengthExpression
                            | capExpression;
operand :                   literal
                            | ID
                            | LP expression RP;
literal :                   INTLITERAL
                            | FLOATLITERAL
                            | RUNELITERAL
                            | RAWSTRINGLITERAL
                            | INTERPRETEDSTRINGLITERAL;
index:                      LSB expression RSB;
arguments :                 LP expressionList | EPSILON RP;
selector :                  DOT ID;
appendExpression :          APPEND LP expression CM expression RP;
lengthExpression :          LEN LP expression RP;
capExpression :             CAP LP expression RP;
statementList :             statement* ;
block :                     LCB statementList RCB;
statement :                 PRINT LP expressionList | EPSILON RP SEMI
                            | PRINTLN LP expressionList | EPSILON RP SEMI
                            | RETURN (expression | EPSILON) SEMI
                            | BREAK SEMI
                            | CONTINUE SEMI
                            | simpleStatement SEMI
                            | block SEMI
                            | switch SEMI
                            | ifStatement SEMI
                            | loop SEMI
                            | typeDecl
                            | variableDecl;
simpleStatement	:           EPSILON
                            | expression (INC | DEC | EPSILON)
                            | assignmentStatement
                            | expressionList SVD expressionList;
assignmentStatement :       expressionList operator expressionList
                            |expression operator  expression;
ifStatement :               IF expression block
                            | IF expression block ELSE ifStatement
                            | IF expression block ELSE block
                            | IF simpleStatement  SEMI expression block
                            | IF simpleStatement SEMI  expression block ELSE ifStatement
                            | IF simpleStatement  SEMI expression block ELSE block;
loop :                      FOR block
                            | FOR expression block
                            | FOR simpleStatement SEMI expression SEMI simpleStatement block
                            | FOR simpleStatement SEMI SEMI simpleStatement block;
switch :                    SWITCH simpleStatement SEMI expression LCB expressionCaseClauseList RCB
                            | SWITCH expression LCB expressionCaseClauseList RCB
                            | SWITCH simpleStatement SEMI LCB expressionCaseClauseList RCB
                            | SWITCH LCB expressionCaseClauseList RCB;
expressionCaseClauseList :  EPSILON
                            | expressionCaseClause expressionCaseClauseList;
expressionCaseClause 	:   expressionSwitchCase COL statementList;
expressionSwitchCase :      CASE expressionList
                            | DEFAULT;
operator :                  MUL | DIV | MOD | LSH | RSH | AMPER | BC | PL | MIN | VB | CARET | EQ | NEQ | LT | LTE | GT
| GTE | LAND | LOR | NEG | ASOP | AAOP | BAOP | SAOP | BOAOP | MAOP |BXOOP |LSAOP |RSAOP |BCAOP |RAOP |DAOP;

