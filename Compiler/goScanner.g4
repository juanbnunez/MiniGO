lexer grammar goScanner;

//Simbols
SEMI    : ';';    //Semicolon
SVD     : ':=';   //Short variable declaration
LP      : '(';    //Left parenthesis
RP      : ')';    //Right parenthesis
TIL     : '~';    //Tilde
COL     : ':';    //Colon
PL      : '+';    //Plus
MIN     : '-';    //Minus
MUL     : '*';    //Multiplication / Asterisk
DIV     : '/';    //Division / Slash
CM      : ',';    //Comma
LSB     : '[';    //Left square bracket
RSB     : ']';    //Right square bracket
LCB     : '{';    //Left curly brace
RCB     : '}';    //Right curly brace
INC     : '++';   //Increment operator
DEC     : '--';   //"Decrement operator
MOD     : '%';    //Modulus
LSH     : '<<';   //Left shift
RSH     : '>>';   //Right shift
AMPER   : '&';    //Ampersand
BC      : '&^';   //Bit clear (AND NOT)
VB      : '|';    //Vertical bar
CARET   : '^';    //Caret
EQ      : '==';   //Equality operator
NEQ     : '!=';   //Inequality operator
LT      : '<';    //Less than
GT      : '>';    //Greater than
LTE     : '<=';   //Less than or equal to
GTE     : '>=';   //Greater than or equal to
LAND    : '&&';   //Logical AND
LOR     : '||';   //Logical OR
NEG     : '!';    //Negation
ASOP    : '=';    //Assignment operator
AAOP    : '+=';   //Addition assignment operator
BAOP    : '&=';   //Bitwise AND assignment operator
SAOP    : '-=';   //Subtraction assignment operator
BOAOP   : '|=';   //Bitwise OR assignment operator
MAOP    : '*=';   //Multiplication assignment operator
BXOOP   : '^=';   //Bitwise XOR assignment operator
LSAOP   : '<<=';  //Left shift assignment operator
RSAOP   : '>>=';  //Right shift assignment operator
BCAOP   : '&^=';  //Bit clear (AND NOT) assignment operator
RAOP    : '%=';   //Remainder assignment operator
DAOP    : '/=';   //Division assignment operator
DOT     : '.';    //DOT
DBS     : '\\';   //Double Backslash
BS      : '\'';   //Back Slash

//Reserved words
PACKAGE : 'package';
IF      : 'if' ;
WHILE   : 'while' ;
LET     : 'let';
THEN    : 'then';
ELSE    : 'else';
IN      : 'in';
DO      : 'do';
BEGIN   : 'begin';
END     : 'end';
CONST   : 'const';
VAR     : 'var';
TYPE    : 'type';
FUNC    : 'func';
STRUCT  : 'struct';
APPEND  : 'append';
LEN     : 'len';
CAP     : 'cap';
PRINT   : 'print';
PRINTLN : 'println';
RETURN  : 'return';
BREAK   : 'break';
CONTINUE: 'continue';
FOR     : 'for';
SWITCH  : 'switch';
CASE    : 'case';
DEFAULT : 'default';

ID : ('_'|) LETTER (LETTER|DIGIT|'_')* ;
INTLITERAL : DIGIT+ ;
FLOATLITERAL : DIGIT+ '.' DIGIT* ([eE] [+\-]? DIGIT+)? ;
RUNELITERAL : '\'' . '\'' ;
RAWSTRINGLITERAL : '`' ( '\\' . | ~('\\'|'`') )* '`';
INTERPRETEDSTRINGLITERAL : '"' ( '\\' . | ~('\\'|'"') )* '"';
LETTER : 'a'..'z' | 'A'..'Z';
DIGIT : '0'..'9' ;

//Skiped elements
LINE_COMMENT:   '//' ~[\r\n]* -> skip ;
SPACES  :   [ \t\n\r]+ -> skip ;

