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

global_env returns [interfaces.Instruction hi]
: declaration PYC { $hi = $declaration.dec }
| function { $hi = $function.fun }
| module
;

main returns [*arrayList.List mainInst]
: FUNC MAIN PARIZQ PARDER LLAVEIZQ block LLAVEDER { $mainInst = $block.blk }
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
| transReturn { $inst = $transReturn.rt }
| structCreation { $inst = $structCreation.dec }
| insVectors PYC { $inst = $insVectors.iv }
;

insVectors returns[interfaces.Instruction iv]
: ID PUNTO PUSH PARIZQ expression PARDER { $iv = instructions.NewPush($ID.line, $ID.pos, $ID.text, $expression.p) }
| ID PUNTO INSERT PARIZQ exp1=expression COMA exp2=expression PARDER { $iv = instructions.NewInsert($ID.line, $ID.pos, $ID.text, $exp1.p, $exp2.p) }
;

listParamsCall returns[*arrayList.List l]
: list=listParamsCall COMA expression{
                                         ByRef := environment.NewByReference($expression.p, false)
                                         $list.l.Add(ByRef);
                                         $l = $list.l;
                                      }
| list=listParamsCall COMA AMP MUT expression {
                                            ByRef := environment.NewByReference($expression.p, true)
                                            $list.l.Add(ByRef);
                                            $l = $list.l;
                                         }
| expression {
                 ByRef := environment.NewByReference($expression.p, false)
                 $l = arrayList.New()
                 $l.Add(ByRef)
              }
| AMP MUT expression {
                        ByRef := environment.NewByReference($expression.p, true)
                        $l = arrayList.New()
                        $l.Add(ByRef)
                     }
|      {
          $l = arrayList.New()
       }
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

transReturn returns[interfaces.Instruction rt]
: RETURN expression PYC { $rt = instructions.NewReturn($RETURN.line, $RETURN.pos, $expression.p) }
| RETURN PYC{ $rt = instructions.NewReturn($RETURN.line, $RETURN.pos, nil) }
| RETURN expression { $rt = instructions.NewReturn($RETURN.line, $RETURN.pos, $expression.p) }
| RETURN { $rt = instructions.NewReturn($RETURN.line, $RETURN.pos, nil) }
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
| vectDeclaration { $dec = $vectDeclaration.vec }
;

vectDeclaration returns[interfaces.Instruction vec]
: LET ID D_PTS VECTOR2 MENOR types MAYOR IGUAL VECTOR2 C_PTS NEW PARIZQ PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, nil, false, nil, "") }
| LET ID D_PTS VECTOR2 MENOR types MAYOR IGUAL VECTOR2 C_PTS WCAPACITY PARIZQ expression PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, $expression.p, false, nil, "") }
| LET MUT ID D_PTS VECTOR2 MENOR types MAYOR IGUAL VECTOR2 C_PTS NEW PARIZQ PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, nil, true, nil, "") }
| LET MUT ID D_PTS VECTOR2 MENOR types MAYOR IGUAL VECTOR2 C_PTS WCAPACITY PARIZQ expression PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, $expression.p, true, nil, "") }
| LET ID D_PTS VECTOR2 MENOR types MAYOR IGUAL expression { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, nil, false, $expression.p, "") }
| LET MUT ID D_PTS VECTOR2 MENOR types MAYOR IGUAL expression { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, nil, true, $expression.p, "") }

| LET id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR IGUAL VECTOR2 C_PTS NEW PARIZQ PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $id1.text, environment.WILDCARD, nil, false, nil, $id2.text) }
| LET id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR IGUAL VECTOR2 C_PTS WCAPACITY PARIZQ expression PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $id1.text, environment.WILDCARD, $expression.p, false, nil, $id2.text) }
| LET MUT id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR IGUAL VECTOR2 C_PTS NEW PARIZQ PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $id1.text, environment.WILDCARD, nil, true, nil, $id2.text) }
| LET MUT id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR IGUAL VECTOR2 C_PTS WCAPACITY PARIZQ expression PARDER { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $id1.text, environment.WILDCARD, $expression.p, true, nil, $id2.text) }
| LET id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR IGUAL expression { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $id1.text, environment.WILDCARD, nil, false, $expression.p, $id2.text) }
| LET MUT id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR IGUAL expression { $vec = instructions.NewVectorDeclaration($LET.line, $LET.pos, $id1.text, environment.WILDCARD, nil, true, $expression.p, $id2.text) }

;

structCreation returns[interfaces.Instruction dec]
: STRUCT ID LLAVEIZQ listStructDec LLAVEDER { $dec = instructions.NewStruct($STRUCT.line, $STRUCT.pos, $ID.text, $listStructDec.l) }
;

listStructDec returns[*arrayList.List l]
: list=listStructDec COMA id1=ID D_PTS types {
                                        StrDef := environment.NewStructType($id1.text, $types.ty, "")
                                        $list.l.Add(StrDef);
                                        $l = $list.l;
                                    }
| list=listStructDec COMA id1=ID D_PTS id2=ID {
                                          StrDef := environment.NewStructType($id1.text,environment.WILDCARD, $id2.text)
                                          $list.l.Add(StrDef);
                                          $l = $list.l;
                                      }
| id1=ID D_PTS types{
                    StrDef := environment.NewStructType($id1.text, $types.ty, "")
                    $l = arrayList.New();
                    $l.Add(StrDef);
                }
| id1=ID D_PTS id2=ID {
                      StrDef := environment.NewStructType($id1.text, environment.WILDCARD, $id2.text)
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

function returns [ interfaces.Instruction fun ]
: FUNC ID PARIZQ listParamsFunc PARDER LLAVEIZQ block LLAVEDER {
                        $fun = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.lpf, environment.WILDCARD, $block.blk, "")
                        }
| FUNC ID PARIZQ listParamsFunc PARDER ARROW1 types LLAVEIZQ block LLAVEDER{
                       $fun = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.lpf, $types.ty, $block.blk, "")
                       }
| FUNC id1=ID PARIZQ listParamsFunc PARDER ARROW1 id2=ID LLAVEIZQ block LLAVEDER{
                       $fun = instructions.NewFunction($FUNC.line, $FUNC.pos, $id1.text, $listParamsFunc.lpf, environment.WILDCARD, $block.blk, $id2.text)
                       }
| FUNC id1=ID PARIZQ listParamsFunc PARDER ARROW1 VECTOR2 MENOR id2=ID MAYOR LLAVEIZQ block LLAVEDER{
                       $fun = instructions.NewFunction($FUNC.line, $FUNC.pos, $id1.text, $listParamsFunc.lpf, environment.WILDCARD, $block.blk, $id2.text)
                       }
;

listParamsFunc returns[*arrayList.List lpf]
: list=listParamsFunc COMA ID D_PTS types {
                   newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, $types.ty, "")
                   $list.lpf.Add(newParam)
                   $lpf = $list.lpf
                    }
| list=listParamsFunc COMA ID D_PTS AMP MUT arrayType {
             newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, environment.ARRAY, "")
             $list.lpf.Add(newParam)
             $lpf = $list.lpf
              }
| list=listParamsFunc COMA id1=ID D_PTS AMP MUT VECTOR2 MENOR id2=ID MAYOR {
             newParam := instructions.NewParamsDeclaration($id1.line, $id1.pos, $id1.text, environment.WILDCARD, $id2.text)
             $list.lpf.Add(newParam)
             $lpf = $list.lpf
              }
| list=listParamsFunc COMA ID D_PTS arrayType {
             newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, environment.ARRAY, "")
             $list.lpf.Add(newParam)
             $lpf = $list.lpf
              }
| list=listParamsFunc COMA id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR  {
             newParam := instructions.NewParamsDeclaration($id1.line, $id1.pos, $id1.text, environment.WILDCARD, $id2.text)
             $list.lpf.Add(newParam)
             $lpf = $list.lpf
              }
| ID D_PTS types{
                $lpf = arrayList.New()
                newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, $types.ty, "")
                $lpf.Add(newParam)
             }
| ID D_PTS AMP MUT arrayType {
                 $lpf = arrayList.New()
                 newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, environment.ARRAY, "")
                 $lpf.Add(newParam)
              }
| id1=ID D_PTS AMP MUT VECTOR2 MENOR id2=ID MAYOR {
                 $lpf = arrayList.New()
                 newParam := instructions.NewParamsDeclaration($id1.line, $id1.pos, $id1.text, environment.WILDCARD, $id2.text)
                 $lpf.Add(newParam)
              }
| ID D_PTS arrayType {
                 $lpf = arrayList.New()
                 newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, environment.ARRAY,"")
                 $lpf.Add(newParam)
              }
| id1=ID D_PTS VECTOR2 MENOR id2=ID MAYOR  {
                 $lpf = arrayList.New()
                 newParam := instructions.NewParamsDeclaration($id1.line, $id1.pos, $id1.text, environment.WILDCARD, $id2.text)
                 $lpf.Add(newParam)
              }
|  { $lpf = arrayList.New() }
;

module returns[] 
: MODULE ID LLAVEIZQ moduleContent LLAVEDER
;

moduleContent returns[]
: module
| function
;

types returns[environment.TipoExpresion ty]
: INT { $ty = environment.INTEGER }
| FLOAT { $ty = environment.FLOAT }
| BOOL { $ty = environment.BOOLEAN }
| CHAR { $ty = environment.CHAR }
| STR1 { $ty = environment.STRING }
| STR2 { $ty = environment.STR }
| VECTOR1 { $ty = environment.VECTOR }
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

listVec returns [interfaces.Expression lv]
: exp1=expression PYC exp2=expression { $lv = expressions.NewVectorList($exp1.start.GetLine(),$exp1.start.GetColumn(), $exp1.p, $exp2.p) }
;

expression returns[interfaces.Expression p]
: expr_arit { $p = $expr_arit.p }
| expuno=expression PUNTO PUNTO expdos=expression { $p = expressions.NewRange($expuno.start.GetLine(),$expuno.start.GetColumn(), $expuno.p, $expdos.p) }
;

callFunction returns[interfaces.Expression cf]
: ID PARIZQ listParamsCall PARDER PYC   { $cf = expressions.NewCallExp($ID.line, $ID.pos, $ID.text, $listParamsCall.l) }
| ID PARIZQ listParamsCall PARDER       { $cf = expressions.NewCallExp($ID.line, $ID.pos, $ID.text, $listParamsCall.l) }
;

expr_arit returns[interfaces.Expression p]
: opIz=expr_arit op=(MUL|DIV|MOD) opDe=expr_arit {$p = expressions.NewOperation($opIz.start.GetLine(),$opIz.start.GetColumn(),$opIz.p,$op.text,$opDe.p)}
| opIz=expr_arit op=(ADD|SUB) opDe=expr_arit {$p = expressions.NewOperation($opIz.start.GetLine(),$opIz.start.GetColumn(),$opIz.p,$op.text,$opDe.p)}
| opIz=expr_arit op=(MENOR|MENORIGUAL|MAYORIGUAL|MAYOR|IG_IG|DIFERENTE) opDe=expr_arit {$p = expressions.NewOperation($opIz.start.GetLine(),$opIz.start.GetColumn(),$opIz.p,$op.text,$opDe.p)}
| SUB opDe=expr_arit {$p = expressions.NewOperation($SUB.line,$SUB.pos,$opDe.p,"MENOS_UNARIO",nil)}
| NOT opDe=expr_arit {$p = expressions.NewOperation($NOT.line,$NOT.pos,$opDe.p,$NOT.text,nil)}
| CORIZQ listParams CORDER { $p = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listParams.l) }
| VECTOR1 NOT CORIZQ listParams CORDER { $p = expressions.NewVector($VECTOR1.line, $VECTOR1.pos, $listParams.l) }
| VECTOR1 NOT CORIZQ listVec CORDER { $p = $listVec.lv }
| PARIZQ expression PARDER { $p = $expression.p }
| ID LLAVEIZQ listStructExp LLAVEDER { $p = expressions.NewStructExp($ID.line, $ID.pos, $ID.text, $listStructExp.l ) }
| callFunction { $p = $callFunction.cf }
| primitive { $p = $primitive.p }
| condIf { $p = $condIf.ifCond }
| condMatch { $p = $condMatch.mtch }
| loopBucle { $p = $loopBucle.lb }
| expVectors { $p = $expVectors.ev }
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

expVectors returns[interfaces.Expression ev]
: ID PUNTO REMOVE PARIZQ expression PARDER { $ev = expressions.NewRemove($ID.line, $ID.pos, $ID.text, $expression.p) }
| ID PUNTO CONTAINS PARIZQ AMP expression PARDER { $ev = expressions.NewContains($ID.line, $ID.pos, $ID.text, $expression.p) }
| ID PUNTO LEN PARIZQ PARDER { $ev = expressions.NewLen($ID.line, $ID.pos, $ID.text) }
| ID PUNTO CAPACITY PARIZQ PARDER { $ev = expressions.NewCapacity($ID.line, $ID.pos, $ID.text) }
;