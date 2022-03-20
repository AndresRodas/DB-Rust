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

start returns [environment.Code code]
  : e+=global_env* main e+=global_env*
    {
        global := arrayList.New()
        listInt := localctx.(*StartContext).GetE()
        for _, e := range listInt {
            global.Add(e.GetHi())
        }
        $code = environment.NewCode($main.mainInst, global)
    }
;

global_env returns [string hi]
: declaration PYC {$hi = "declaration"}
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
: impression PYC { $inst = $impression.pr }
| declaration PYC { $inst = $declaration.dec }
| assignment PYC { $inst = $assignment.ass }
| condIf { $inst = $condIf.ifCond }
| condMatch { $inst = $condMatch.mtch }
| loopWhile { $inst = $loopWhile.lw }
| loopBucle { $inst = $loopBucle.lb }
| loopForin { $inst = $loopForin.lfi }
| transBreak PYC { $inst = $transBreak.brk }
| transContinue PYC { $inst = $transContinue.cnt }
| structCreation { $inst = $structCreation.dec }
;

loopWhile returns[interfaces.Instruction lw]
: WHILE expression LLAVEIZQ block LLAVEDER { $lw = instructions.NewWhile($WHILE.line, $WHILE.pos, $expression.p, $block.blk ) }
;

loopBucle returns[interfaces.Instruction lb]
: LOOP LLAVEIZQ block LLAVEDER { $lb = instructions.NewLoop($LOOP.line, $LOOP.pos, $block.blk) }
;

loopForin returns[interfaces.Instruction lfi]
: FOR ID IN expression LLAVEIZQ instructions LLAVEDER { $lfi = instructions.NewForIn($FOR.line, $FOR.pos, $ID.text, $expression.p, $instructions.insts) }
;

transBreak returns[interfaces.Instruction brk]
: BREAK expression { $brk = instructions.NewBreak($BREAK.line, $BREAK.pos, $expression.p) }
| BREAK { $brk = instructions.NewBreak($BREAK.line, $BREAK.pos, nil) }
;

transContinue returns[interfaces.Instruction cnt]
: CONTINUE { $cnt = instructions.NewContinue($CONTINUE.line, $CONTINUE.pos) }
;

condIf returns [ interfaces.Instruction ifCond ]
: IF expression LLAVEIZQ block LLAVEDER e+=condElseIf* condElse{
            elif := arrayList.New()
            listElif := localctx.(*CondIfContext).GetE()
            for _, e := range listElif {
                elif.Add(e.GetElif())
            }
            $ifCond = instructions.NewIf($IF.line, $IF.pos, $expression.p, $block.blk, elif, $condElse.blkelse)
            }
;

condElseIf returns [interfaces.Instruction elif]
: ELSE IF expression LLAVEIZQ block LLAVEDER {
    elif := arrayList.New()
    condelse := arrayList.New()
    $elif = instructions.NewIf($ELSE.line, $ELSE.pos, $expression.p, $block.blk, elif, condelse)
    }
;

condElse returns [*arrayList.List blkelse]
: ELSE LLAVEIZQ block LLAVEDER { $blkelse = $block.blk }
| { $blkelse = arrayList.New() }
;

block returns[*arrayList.List blk]
@init{
    $blk = arrayList.New()
  }
: bloque=block instruction {
                            $bloque.blk.Add($instruction.inst)
                            $blk = $bloque.blk
                            }
| bloque=block expression {
                            $bloque.blk.Add($expression.p)
                            $blk = $bloque.blk
                           }
| instruction { $blk.Add($instruction.inst) }
| expression { $blk.Add($expression.p) }
;

condMatch returns [interfaces.Instruction mtch]
: MATCH expression LLAVEIZQ e+=listArms+ defaultArm LLAVEDER
        {
        arrarms := arrayList.New()
        larms := localctx.(*CondMatchContext).GetE()
        for _, e := range larms {
            arrarms.Add(e.GetArms())
        }
        $mtch = instructions.NewMatch($MATCH.line, $MATCH.pos, $expression.p, arrarms, $defaultArm.defa)
        }
;

listArms returns [interfaces.Instruction arms]
: listMatch ARROW2 block COMA{
         $arms = instructions.NewArm($listMatch.start.GetLine(),$listMatch.start.GetColumn(), $listMatch.ma, $block.blk)
         }
| listMatch ARROW2 LLAVEIZQ block LLAVEDER COMA {
        $arms = instructions.NewArm($listMatch.start.GetLine(),$listMatch.start.GetColumn(), $listMatch.ma, $block.blk)
        }
;

listMatch returns[*arrayList.List ma]
: lma=listMatch PLEC expression {
                                $lma.ma.Add($expression.p)
                                $ma = $lma.ma
                             }
| expression {
                 $ma = arrayList.New()
                 $ma.Add($expression.p)
              }
;

defaultArm returns[*arrayList.List defa]
: UNDERSCORE ARROW2 block COMA { $defa = $block.blk }
| UNDERSCORE ARROW2 LLAVEIZQ block LLAVEDER COMA { $defa = $block.blk }
| { $defa = arrayList.New() }
;

impression returns [interfaces.Instruction pr]
: PRINT PARIZQ listParams PARDER { $pr = instructions.NewPrint($PRINT.line,$PRINT.pos,$listParams.l) }
;

declaration returns [interfaces.Instruction dec]
: LET MUT ID D_PTS types IGUAL expression   { $dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, $expression.p, true) }
| LET MUT ID IGUAL expression               { $dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, environment.WILDCARD, $expression.p, true) }
| LET ID D_PTS types IGUAL expression       { $dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, $expression.p, false) }
| LET ID IGUAL expression                   { $dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, environment.WILDCARD, $expression.p, false) }
| LET MUT ID D_PTS arrayType IGUAL expression { $dec = instructions.NewArrayDeclaration($LET.line, $LET.pos, $ID.text, $arrayType.t, $expression.p, true) }
| LET ID D_PTS arrayType IGUAL expression   { $dec = instructions.NewArrayDeclaration($LET.line, $LET.pos, $ID.text, $arrayType.t, $expression.p, false) }
;

structCreation returns[interfaces.Instruction dec]
: STRUCT ID LLAVEIZQ listStructDec LLAVEDER { $dec = instructions.NewStruct($STRUCT.line, $STRUCT.pos, $ID.text, $listStructDec.l) }
;

listStructDec returns[*arrayList.List l]
: list=listStructDec COMA ID D_PTS types {
                                        StrDef := environment.NewStructType($ID.text, $types.ty)
                                        $list.l.Add(StrDef);
                                        $l = $list.l;
                                    }
| ID D_PTS types{
                    StrDef := environment.NewStructType($ID.text, $types.ty)
                    $l = arrayList.New();
                    $l.Add(StrDef);
                }
;

assignment returns [interfaces.Instruction ass]
: ID IGUAL expression { $ass = instructions.NewAssignment($ID.line, $ID.pos, $ID.text, $expression.p)}
| listAccessStruct IGUAL expression { $ass = instructions.NewStructAssign($listAccessStruct.start.GetLine(),$listAccessStruct.start.GetColumn(), $listAccessStruct.l, $expression.p) }
| ID listAccessArray IGUAL expression { $ass = instructions.NewArrayAssign($ID.line, $ID.pos, $ID.text, $listAccessArray.l, $expression.p) }
;

listAccessStruct returns[*arrayList.List l]
: list=listAccessStruct PUNTO ID {
                                   $list.l.Add($ID.text)
                                   $l = $list.l
                                  }
| ID {
    $l = arrayList.New()
    $l.Add($ID.text)
}
;

listAccessArray returns[*arrayList.List l]
: list=listAccessArray CORIZQ expression CORDER{
                                                  $list.l.Add($expression.p)
                                                  $l = $list.l
                                                 }
| CORIZQ expression CORDER{
                              $l = arrayList.New()
                              $l.Add($expression.p)
                          }
;

arrayType returns [*arrayList.List t]
: CORIZQ arrayType PYC expression CORDER {
                                            newType := environment.NewArrayType(environment.ARRAY, $expression.p)
                                           $arrayType.t.Add(newType)
                                           $t = $arrayType.t
                                            }
| CORIZQ types PYC expression CORDER{
                                        $t = arrayList.New()
                                        newType := environment.NewArrayType($types.ty, $expression.p)
                                        $t.Add(newType)
                                     }
;

function returns []
: FUNC ID PARIZQ listParams PARDER LLAVEIZQ instructions LLAVEDER
;

module returns[] 
: MODULE ID LLAVEIZQ LLAVEDER
;

types returns[environment.TipoExpresion ty]
: INT { $ty = environment.INTEGER }
| FLOAT { $ty = environment.FLOAT }
| BOOL { $ty = environment.BOOLEAN }
| CHAR { $ty = environment.CHAR }
| STR1 { $ty = environment.STRING }
| STR2 { $ty = environment.STR }
| VECTOR { $ty = environment.VECTOR }
| STRUCT { $ty = environment.STRUCT }
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
| expuno=expression PUNTO PUNTO expdos=expression { $p = expressions.NewRange($expuno.start.GetLine(),$expuno.start.GetColumn(), $expuno.p, $expdos.p) }
;

expr_arit returns[interfaces.Expression p]
: opIz=expr_arit op=(MUL|DIV|MOD) opDe=expr_arit {$p = expressions.NewOperation($opIz.start.GetLine(),$opIz.start.GetColumn(),$opIz.p,$op.text,$opDe.p)}
| opIz=expr_arit op=(ADD|SUB) opDe=expr_arit {$p = expressions.NewOperation($opIz.start.GetLine(),$opIz.start.GetColumn(),$opIz.p,$op.text,$opDe.p)}
| opIz=expr_arit op=(MENOR|MENORIGUAL|MAYORIGUAL|MAYOR|IG_IG|DIFERENTE) opDe=expr_arit {$p = expressions.NewOperation($opIz.start.GetLine(),$opIz.start.GetColumn(),$opIz.p,$op.text,$opDe.p)}
| SUB opDe=expr_arit {$p = expressions.NewOperation($SUB.line,$SUB.pos,$opDe.p,"MENOS_UNARIO",nil)}
| NOT opDe=expr_arit {$p = expressions.NewOperation($NOT.line,$NOT.pos,$opDe.p,$NOT.text,nil)}
| CORIZQ listParams CORDER { $p = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listParams.l) }
| PARIZQ expression PARDER { $p = $expression.p }
| ID LLAVEIZQ listStructExp LLAVEDER { $p = expressions.NewStructExp($ID.line, $ID.pos, $ID.text, $listStructExp.l ) }
| primitive { $p = $primitive.p }
| condIf { $p = $condIf.ifCond }
| condMatch { $p = $condMatch.mtch }
| loopBucle { $p = $loopBucle.lb }
;

primitive returns[interfaces.Expression p]
: NUMBER{
            if (strings.Contains($NUMBER.text,".")){
                num,err := strconv.ParseFloat($NUMBER.text, 64);
                if err!= nil{
                    fmt.Println(err)
                }
                $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
            }else{
                num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
                $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
            }
        }
| stringTypes    { $p = $stringTypes.st  }
| CHARACTER { $p = expressions.NewPrimitive($CHARACTER.line, $CHARACTER.pos,$CHARACTER.text,environment.CHAR) }
| TRU { $p = expressions.NewPrimitive($TRU.line, $TRU.pos,true,environment.BOOLEAN) }
| FAL { $p = expressions.NewPrimitive($FAL.line, $FAL.pos,false,environment.BOOLEAN) }
| list=listArray { $p = $list.p}
;

stringTypes returns[interfaces.Expression st]
: STRING PUNTO fnc=(TOSTR|TOOWN) {
                                     str := $STRING.text
                                     $st = expressions.NewPrimitive($STRING.line, $STRING.pos,str[1:len(str)-1],environment.STRING)
                                 }
| AMP* STRING (PUNTO TOSTR|PUNTO TOOWN)* {
              str := $STRING.text
              $st = expressions.NewPrimitive($AMP.line, $AMP.pos,str[1:len(str)-1],environment.STR)
          }
;

listArray returns[interfaces.Expression p]
: list = listArray CORIZQ expression CORDER { $p = expressions.NewArrayAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $expression.p) }
| list = listArray PUNTO ID { $p = expressions.NewStructAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $ID.text)  }
| ID { $p = expressions.NewCallVar($ID.line, $ID.pos, $ID.text)}
;


listStructExp returns[*arrayList.List l]
: list=listStructExp COMA ID D_PTS expression {
                                        StrExp := environment.NewStructContent($ID.text, $expression.p)
                                        $list.l.Add(StrExp);
                                        $l = $list.l;
                                    }
| ID D_PTS expression{
                    StrExp := environment.NewStructContent($ID.text, $expression.p)
                    $l = arrayList.New();
                    $l.Add(StrExp);
                }
;
