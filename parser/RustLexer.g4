lexer grammar RustLexer;

// --------------- Tokens

// types
INT: 'i64';
FLOAT: 'f64';
BOOL: 'bool';
CHAR: 'char';
STR1: 'String';
STR2: '&str';
USIZE: 'usize';
VECTOR: 'vec';
STRUCT: 'struct';

// reserved words
POW: 'pow';
PRINT: 'println';
LET: 'let';
MUT: 'mut';
FUNC: 'fn';
ARROW1: '->';
ARROW2: '=>';
ABS: 'abs';
SQRT: 'sqrt';
TOSTR: 'to_string';
CLONE: 'clone';
NEW: 'new';
LEN: 'len';
PUSH: 'push';
REMOVE: 'remove';
CONTAINS: 'contains';
INSERT: 'insert';
CAPACITY: 'capacity';
WCAPACITY: 'with_capacity';
MAIN: 'main';
IF: 'if';
MATCH: 'match';
LOOP: 'loop';
WHILE: 'while';
FOR: 'for';
IN: 'in';
BREAK: 'break';
CONTINUE: 'continue';
RETURN: 'return';
MODULE: 'mod';
PUB: 'pub';

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"'~["]*'"';
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

// symbols
PUNTO:          '.';
C_PTS:          '::';
D_PTS:          ':';
PYC:            ';';
COMA:           ',';
DIFERENTE:      '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IGUAL:          '=';
MAYORIGUAL:     '>=';
MENORIGUAL:     '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
MOD:            '%';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
CORIZQ:         '[';
CORDER:         ']';


// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;
