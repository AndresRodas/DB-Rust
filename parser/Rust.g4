parser grammar Rust;

options {
  tokenVocab = RustLexer;
}

@header {
    import arrayList "github.com/colegno/arraylist"
    import "OLC2/interfaces"
    import "OLC2/expressions"
    import "OLC2/instructions"
    import "OLC2/environment"
    import "strings"
}

start returns [environment.AST ast]
  : e+=global_env* main e+=global_env*
    {
        global := arrayList.New()
        listInt := localctx.(*StartContext).GetE()
        for _, e := range listInt {
            global.Add(e.GetHi())
        }
        $ast = environment.NewAST($main.mainInst, global, "")
    }
;

global_env returns [string hi]
: declaration {$hi = "declaration"}
| function {$hi = "function"}
| module {$hi = "module"}
;

main returns [*arrayList.List mainInst]
: FUNC MAIN PARIZQ PARDER LLAVEIZQ instructions LLAVEDER { $mainInst = $instructions.insts }
;

instructions returns[*arrayList.List insts]
@init{
    $insts = arrayList.New()
  }
: e+=instruction+
{
        listInt := localctx.(*InstructionsContext).GetE()
        for _, e := range listInt {
            $insts.Add(e.GetInst())
        }
    }
;

instruction returns [interfaces.Instruction inst]
: impression { $inst = $impression.pr }
;

impression returns [interfaces.Instruction pr]
: PRINT PARIZQ listParams PARDER PYC { $pr = instructions.NewPrint(0,0,$listParams.l) }
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
: list=listParams COMA expression   {
                                        $list.l.Add($expression.p)
                                        $l = $list.l
                                     }
| expression {
                $l = arrayList.New()
                $l.Add($expression.p)
             }
;

expression returns[interfaces.Expression p]
: expr_arit { $p = $expr_arit.p }
;

expr_arit returns[interfaces.Expression p]
: opIz=expr_arit op=(MUL|DIV) opDe=expr_arit {$p = expressions.NewOperation(0,0,$opIz.p,$op.text,$opDe.p)}
| opIz=expr_arit op=(ADD|SUB) opDe=expr_arit {$p = expressions.NewOperation(0,0,$opIz.p,$op.text,$opDe.p)}
| opIz=expr_arit op=(MENOR|MENORIGUAL|MAYORIGUAL|MAYOR) opDe=expr_arit {$p = expressions.NewOperation(0,0,$opIz.p,$op.text,$opDe.p)}
| CORIZQ listParams CORDER
| PARIZQ expression PARDER { $p = $expression.p }
| primitive { $p = $primitive.p }
;

primitive returns[interfaces.Expression p]
: NUMBER{
            if (strings.Contains($NUMBER.text,".")){
                num,err := strconv.ParseFloat($NUMBER.text, 64);
                if err!= nil{
                    fmt.Println(err)
                }
                $p = expressions.NewPrimitive(0,0,num,environment.FLOAT)
            }else{
                num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
                $p = expressions.NewPrimitive(0,0,num,environment.INTEGER)
            }
        }
| STRING    {
                str := $STRING.text
                $p = expressions.NewPrimitive(0,0,str[1:len(str)-1],environment.STRING)
            }
;
