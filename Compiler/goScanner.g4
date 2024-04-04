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
RUNELITERAL : '\'' (UnicodeValue | Byte_value) '\'' ;
Raw_string_lit : { unicode_char | newline } "';
Interpreted_string_lit: '"' { UnicodeValue | byte_value } '"';
Byte_value :  Octal_byte_value | Hex_byte_value ;
Octal_byte_value :  Octal_digit Octal_digit Octal_digit;
Hex_byte_value : BS 'x' HexDigit HexDigit;
Little_u_value : BS 'u' HexDigit HexDigit HexDigit HexDigit;
Big_u_value : BS 'U' HexDigit HexDigit HexDigit HexDigit;

LETTER : 'a'..'z' | 'A'..'Z';
DIGIT : '0'..'9' ;
UnicodeValue : Unicode_char | Little_u_value | Big_u_value | Escaped_char;
HexDigit : [0-9a-fA-F] ;
EscapedChar : BS LP'a'|'b'|'f'|'n'|'r'|'t'|'v'| '\'' | '''' | '"' RP;
//Skiped elements
LINE_COMMENT:   '//' ~[\r\n]* -> skip ;
EPSILON  :   [ \t\n\r]+ -> skip ;

//rune_lit         = "'" ( unicode_value | byte_value ) "'" .
//unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
//¿¿¿¿byte_value       = octal_byte_value | hex_byte_value .
//octal_byte_value = \ octal_digit octal_digit octal_digit .
//hex_byte_value   = \ "x" hex_digit hex_digit .
//little_u_value   = \ "u" hex_digit hex_digit hex_digit hex_digit .
//big_u_value      = \ "U" hex_digit hex_digit hex_digit hex_digit
                           //hex_digit hex_digit hex_digit hex_digit .
//escaped_char     = \ ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | \ | "'" | " ) .

string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "" { unicode_char | newline } "" .
interpreted_string_lit = " { unicode_value | byte_value } " .