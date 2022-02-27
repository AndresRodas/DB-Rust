parser grammar Rust;

options {
  tokenVocab = RustLexer;
}

@header {
    import arrayList "github.com/colegno/arraylist"
    
}

start returns [string inst]
  : global_env main global_env  { $inst = "mundo" }
  | global_env main             { $inst = "mundo" }
  | main global_env             { $inst = "hola" }
  | main                        { $inst = "hola" }
;

global_env returns []
: global_env declaration
| global_env function
| global_env module
| declaration
| function
| module 
;

main returns []
: FUNC MAIN PARIZQ PARDER LLAVEIZQ instructions LLAVEDER
;

instructions returns[]
: instruction*
;

instruction returns []
: PRINT PYC
;

declaration returns []
: LET MUT ID D_PTS types IGUAL expression PYC 
| LET MUT ID IGUAL expression PYC
| LET ID D_PTS types IGUAL expression PYC
| LET ID IGUAL expression PYC
;

function returns []
: FUNC ID PARIZQ listParams PARDER LLAVEIZQ instructions LLAVEDER
;

module returns[] 
: MODULE ID LLAVEIZQ LLAVEDER
;

types returns[]
: INT
| FLOAT
| BOOL
| CHAR
| STR1
| VECTOR
| STRUCT
;

listParams returns[*arrayList.List l]
: listParams COMA expression { }
| expression { }
;

expression returns[]
: expr_arit
;

expr_arit returns[]
: opIz=expr_arit op=('*'|'/') opDe=expr_arit 
| opIz=expr_arit op=('+'|'-') opDe=expr_arit  
| opIz=expr_arit op=('<'|'<='|'>='|'>') opDe=expr_arit 
| CORIZQ listParams CORDER
| PARIZQ expression PARDER
| primitive   
;

primitive returns[]
: NUMBER
| STRING
;
