parser grammar goParser;

options {
  tokenVocab = goScanner;
}

root :                      PACKAGE ID SEMI topDeclarationList                                                          #rootAST
                            ;
topDeclarationList :        (variableDecl | typeDecl | funcDecl)*                                                       #topDeclarationListAST
                            ;
variableDecl :              VAR (singleVarDecl SEMI | LP innerVarDecls RP SEMI | LP RP SEMI)                            #varVDAST
                            ;
innerVarDecls :             singleVarDecl SEMI (singleVarDecl SEMI)*                                                    #innerVarDeclsAST
                            ;
singleVarDecl :             identifierList declType ASOP expressionList                                                 #idListDTSVDAST
                            | identifierList ASOP expressionList                                                        #idListSVDAST
                            | singleVarDeclNoExps                                                                       #singleVarDAST
                            ;
singleVarDeclNoExps	:       identifierList declType                                                                     #singleVarDeclNoExpsAST
                            ;
typeDecl :                  TYPE (singleTypeDecl SEMI | LP innerTypeDecls RP SEMI | LP RP SEMI)                         #typeTDAST
                            ;
innerTypeDecls :            singleTypeDecl SEMI (singleTypeDecl SEMI)*                                                  #innerTypeDeclsAST
                            ;
singleTypeDecl :            ID declType                                                                                 #idSYDAST
                            ;
funcDecl :                  funcFrontDecl block SEMI                                                                    #funcDeclAST
                            ;
funcFrontDecl :             FUNC ID LP (funcArgDecls|epsilon) RP (declType|epsilon)                                     #funcFFDAST
                            ;
funcArgDecls :              singleVarDeclNoExps (CM singleVarDeclNoExps)*                                               #funcArgDeclsAST
                            ;
declType :                  LP declType RP                                                                              #lpDTAST
                            | ID                                                                                        #idDTAST
                            | sliceDeclType                                                                             #sliceDeclTypeDTAST
                            | arrayDeclType                                                                             #arrayDeclTypeDTAST
                            | structDeclType                                                                            #structDeclTypeDTAST
                            ;
sliceDeclType :             LSB RSB declType                                                                            #lsbSDTAST
                            ;
arrayDeclType :             LSB INTLITERAL RSB declType                                                                 #lsbADTAST
                            ;
structDeclType :            STRUCT LCB (structMemDecls|epsilon) RCB                                                     #structSDTAST
                            ;
structMemDecls :            singleVarDeclNoExps SEMI (singleVarDeclNoExps SEMI)*                                        #structMemDeclsAST
                            ;
identifierList :            ID (CM ID)*                                                                                 #identifierListAST
                            ;
expression :                primaryExpression                                                                           #primaryExpressionEAST
                            | expression operator expression                                                            #expressionEAST
                            | operator expression                                                                       #operatorEAST
                            ;
expressionList :            expression (CM expression)*                                                                 #expressionListAST
                            ;
primaryExpression :         operand                                                                                     #operandPEAST
                            | primaryExpression (selector | index | arguments )                                         #primaryExpPEAST
                            | appendExpression                                                                          #appendExpPEAST
                            | lengthExpression                                                                          #lengthExpPEAST
                            | capExpression                                                                             #capExpPEAST
                            ;
operand :                   literal                                                                                     #literalOAST
                            | ID                                                                                        #idOAST
                            | LP expression RP                                                                          #expressionOAST
                            ;
literal :                   INTLITERAL                                                                                  #intliteralAST
                            | FLOATLITERAL                                                                              #floatliteralAST
                            | RUNELITERAL                                                                               #runliteralAST
                            | RAWSTRINGLITERAL                                                                          #rawsliteralAST
                            | INTERPRETEDSTRINGLITERAL                                                                  #interpretedliteralAST
                            ;
index:                      LSB expression RSB                                                                          #idexAST
                            ;
arguments :                 LP (expressionList | epsilon) RP                                                            #argumentsAST
                            ;
selector :                  DOT ID                                                                                      #selectorAST
                            ;
appendExpression :          APPEND LP expression CM expression RP                                                       #appendExpressionAST
                            ;
lengthExpression :          LEN LP expression RP                                                                        #lengthExpressionAST
                            ;
capExpression :             CAP LP expression RP                                                                        #capExpressionAST
                            ;
statementList :             statement*                                                                                  #statementListAST
                            ;
block :                     LCB statementList RCB                                                                       #blockAST
                            ;
statement :                 PRINT LP (expressionList | epsilon) RP SEMI                                                 #printSAST
                            | PRINTLN LP (expressionList | epsilon) RP SEMI                                             #printlnSAST
                            | RETURN (expression | epsilon) SEMI                                                        #returnSAST
                            | BREAK SEMI                                                                                #breakSAST
                            | CONTINUE SEMI                                                                             #continueSAST
                            | simpleStatement SEMI                                                                      #simpleStatementSAST
                            | block SEMI                                                                                #blockSAST
                            | switch SEMI                                                                               #switchSAST
                            | ifStatement SEMI                                                                          #ifStatementSAST
                            | loop SEMI                                                                                 #loopSAST
                            | typeDecl                                                                                  #typeDeclSAST
                            | variableDecl                                                                              #variableDeclSAST
                            ;
simpleStatement	:           epsilon                                                                                     #epsilonSSAST
                            | expression (INC | DEC | epsilon)                                                          #expressionSSAST
                            | assignmentStatement                                                                       #assigmentStatementSSAST
                            | expressionList SVD expressionList                                                         #expListSSAST
                            ;
assignmentStatement :       expressionList operator expressionList                                                      #expListASAST
                            |expression operator  expression                                                            #expASAST
                            ;
ifStatement :               IF expression block                                                                         #isExpressionBlockISAST
                            | IF expression block ELSE ifStatement                                                      #isExpBlockIsISAST
                            | IF expression block ELSE block                                                            #isExpBlockISAST
                            | IF simpleStatement  SEMI expression block                                                 #isSSExpBlockISAST
                            | IF simpleStatement SEMI  expression block ELSE ifStatement                                #isSSExpBlockifSAST
                            | IF simpleStatement  SEMI expression block ELSE block                                      #isSSExpBlockBlockAST
                            ;
loop :                      FOR block                                                                                   #fBlockLAST
                            | FOR expression block                                                                      #fExpBlockLAST
                            | FOR simpleStatement SEMI expression SEMI simpleStatement block                            #fSimpStateExpSBlockLAST
                            | FOR simpleStatement SEMI SEMI simpleStatement block                                       #fSimpStateSimpStateBlockLAST
                            ;
switch :                    SWITCH simpleStatement SEMI expression LCB expressionCaseClauseList RCB                     #sSimpleStatSAST
                            | SWITCH expression LCB expressionCaseClauseList RCB                                        #sExpressionSAST
                            | SWITCH simpleStatement SEMI LCB expressionCaseClauseList RCB                              #sSimpleStatExpCCListSAST
                            | SWITCH LCB expressionCaseClauseList RCB                                                   #sBlockSAST
                            ;
expressionCaseClauseList :  epsilon                                                                                     #epsilonECCLAST
                            | expressionCaseClause expressionCaseClauseList                                             #expCaseClauseECCLAST
                            ;
expressionCaseClause 	:   expressionSwitchCase COL statementList                                                      #expSwitchCaseECCAST
                            ;
expressionSwitchCase :      CASE expressionList                                                                         #caseESCAST
                            | DEFAULT                                                                                   #defaultESCAST
                            ;
operator :                  MUL | DIV | MOD | LSH | RSH | AMPER | BC | PL | MIN | VB | CARET | EQ | NEQ | LT |
                            LTE | GT | GTE | LAND | LOR | NEG | ASOP | AAOP | BAOP | SAOP | BOAOP | MAOP |
                            BXOOP |LSAOP |RSAOP |BCAOP |RAOP |DAOP                                                      #operatorAST
                            ;
epsilon :                                                                                                               #epsilonAST
                            ;

