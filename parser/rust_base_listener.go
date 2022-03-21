// Code generated from Rust.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Rust

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRustListener is a complete listener for a parse tree produced by Rust.
type BaseRustListener struct{}

var _ RustListener = &BaseRustListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRustListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRustListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRustListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRustListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseRustListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseRustListener) ExitStart(ctx *StartContext) {}

// EnterGlobal_env is called when production global_env is entered.
func (s *BaseRustListener) EnterGlobal_env(ctx *Global_envContext) {}

// ExitGlobal_env is called when production global_env is exited.
func (s *BaseRustListener) ExitGlobal_env(ctx *Global_envContext) {}

// EnterMain is called when production main is entered.
func (s *BaseRustListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseRustListener) ExitMain(ctx *MainContext) {}

// EnterInstructions is called when production instructions is entered.
func (s *BaseRustListener) EnterInstructions(ctx *InstructionsContext) {}

// ExitInstructions is called when production instructions is exited.
func (s *BaseRustListener) ExitInstructions(ctx *InstructionsContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseRustListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseRustListener) ExitInstruction(ctx *InstructionContext) {}

// EnterListParamsCall is called when production listParamsCall is entered.
func (s *BaseRustListener) EnterListParamsCall(ctx *ListParamsCallContext) {}

// ExitListParamsCall is called when production listParamsCall is exited.
func (s *BaseRustListener) ExitListParamsCall(ctx *ListParamsCallContext) {}

// EnterLoopWhile is called when production loopWhile is entered.
func (s *BaseRustListener) EnterLoopWhile(ctx *LoopWhileContext) {}

// ExitLoopWhile is called when production loopWhile is exited.
func (s *BaseRustListener) ExitLoopWhile(ctx *LoopWhileContext) {}

// EnterLoopBucle is called when production loopBucle is entered.
func (s *BaseRustListener) EnterLoopBucle(ctx *LoopBucleContext) {}

// ExitLoopBucle is called when production loopBucle is exited.
func (s *BaseRustListener) ExitLoopBucle(ctx *LoopBucleContext) {}

// EnterLoopForin is called when production loopForin is entered.
func (s *BaseRustListener) EnterLoopForin(ctx *LoopForinContext) {}

// ExitLoopForin is called when production loopForin is exited.
func (s *BaseRustListener) ExitLoopForin(ctx *LoopForinContext) {}

// EnterTransBreak is called when production transBreak is entered.
func (s *BaseRustListener) EnterTransBreak(ctx *TransBreakContext) {}

// ExitTransBreak is called when production transBreak is exited.
func (s *BaseRustListener) ExitTransBreak(ctx *TransBreakContext) {}

// EnterTransContinue is called when production transContinue is entered.
func (s *BaseRustListener) EnterTransContinue(ctx *TransContinueContext) {}

// ExitTransContinue is called when production transContinue is exited.
func (s *BaseRustListener) ExitTransContinue(ctx *TransContinueContext) {}

// EnterTransReturn is called when production transReturn is entered.
func (s *BaseRustListener) EnterTransReturn(ctx *TransReturnContext) {}

// ExitTransReturn is called when production transReturn is exited.
func (s *BaseRustListener) ExitTransReturn(ctx *TransReturnContext) {}

// EnterCondIf is called when production condIf is entered.
func (s *BaseRustListener) EnterCondIf(ctx *CondIfContext) {}

// ExitCondIf is called when production condIf is exited.
func (s *BaseRustListener) ExitCondIf(ctx *CondIfContext) {}

// EnterCondElseIf is called when production condElseIf is entered.
func (s *BaseRustListener) EnterCondElseIf(ctx *CondElseIfContext) {}

// ExitCondElseIf is called when production condElseIf is exited.
func (s *BaseRustListener) ExitCondElseIf(ctx *CondElseIfContext) {}

// EnterCondElse is called when production condElse is entered.
func (s *BaseRustListener) EnterCondElse(ctx *CondElseContext) {}

// ExitCondElse is called when production condElse is exited.
func (s *BaseRustListener) ExitCondElse(ctx *CondElseContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseRustListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseRustListener) ExitBlock(ctx *BlockContext) {}

// EnterCondMatch is called when production condMatch is entered.
func (s *BaseRustListener) EnterCondMatch(ctx *CondMatchContext) {}

// ExitCondMatch is called when production condMatch is exited.
func (s *BaseRustListener) ExitCondMatch(ctx *CondMatchContext) {}

// EnterListArms is called when production listArms is entered.
func (s *BaseRustListener) EnterListArms(ctx *ListArmsContext) {}

// ExitListArms is called when production listArms is exited.
func (s *BaseRustListener) ExitListArms(ctx *ListArmsContext) {}

// EnterListMatch is called when production listMatch is entered.
func (s *BaseRustListener) EnterListMatch(ctx *ListMatchContext) {}

// ExitListMatch is called when production listMatch is exited.
func (s *BaseRustListener) ExitListMatch(ctx *ListMatchContext) {}

// EnterDefaultArm is called when production defaultArm is entered.
func (s *BaseRustListener) EnterDefaultArm(ctx *DefaultArmContext) {}

// ExitDefaultArm is called when production defaultArm is exited.
func (s *BaseRustListener) ExitDefaultArm(ctx *DefaultArmContext) {}

// EnterImpression is called when production impression is entered.
func (s *BaseRustListener) EnterImpression(ctx *ImpressionContext) {}

// ExitImpression is called when production impression is exited.
func (s *BaseRustListener) ExitImpression(ctx *ImpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseRustListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseRustListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterStructCreation is called when production structCreation is entered.
func (s *BaseRustListener) EnterStructCreation(ctx *StructCreationContext) {}

// ExitStructCreation is called when production structCreation is exited.
func (s *BaseRustListener) ExitStructCreation(ctx *StructCreationContext) {}

// EnterListStructDec is called when production listStructDec is entered.
func (s *BaseRustListener) EnterListStructDec(ctx *ListStructDecContext) {}

// ExitListStructDec is called when production listStructDec is exited.
func (s *BaseRustListener) ExitListStructDec(ctx *ListStructDecContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseRustListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseRustListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterListAccessStruct is called when production listAccessStruct is entered.
func (s *BaseRustListener) EnterListAccessStruct(ctx *ListAccessStructContext) {}

// ExitListAccessStruct is called when production listAccessStruct is exited.
func (s *BaseRustListener) ExitListAccessStruct(ctx *ListAccessStructContext) {}

// EnterListAccessArray is called when production listAccessArray is entered.
func (s *BaseRustListener) EnterListAccessArray(ctx *ListAccessArrayContext) {}

// ExitListAccessArray is called when production listAccessArray is exited.
func (s *BaseRustListener) ExitListAccessArray(ctx *ListAccessArrayContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseRustListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseRustListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseRustListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseRustListener) ExitFunction(ctx *FunctionContext) {}

// EnterListParamsFunc is called when production listParamsFunc is entered.
func (s *BaseRustListener) EnterListParamsFunc(ctx *ListParamsFuncContext) {}

// ExitListParamsFunc is called when production listParamsFunc is exited.
func (s *BaseRustListener) ExitListParamsFunc(ctx *ListParamsFuncContext) {}

// EnterModule is called when production module is entered.
func (s *BaseRustListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseRustListener) ExitModule(ctx *ModuleContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseRustListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseRustListener) ExitTypes(ctx *TypesContext) {}

// EnterListParams is called when production listParams is entered.
func (s *BaseRustListener) EnterListParams(ctx *ListParamsContext) {}

// ExitListParams is called when production listParams is exited.
func (s *BaseRustListener) ExitListParams(ctx *ListParamsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRustListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRustListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCallFunction is called when production callFunction is entered.
func (s *BaseRustListener) EnterCallFunction(ctx *CallFunctionContext) {}

// ExitCallFunction is called when production callFunction is exited.
func (s *BaseRustListener) ExitCallFunction(ctx *CallFunctionContext) {}

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseRustListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseRustListener) ExitExpr_arit(ctx *Expr_aritContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseRustListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseRustListener) ExitPrimitive(ctx *PrimitiveContext) {}

// EnterStringTypes is called when production stringTypes is entered.
func (s *BaseRustListener) EnterStringTypes(ctx *StringTypesContext) {}

// ExitStringTypes is called when production stringTypes is exited.
func (s *BaseRustListener) ExitStringTypes(ctx *StringTypesContext) {}

// EnterListArray is called when production listArray is entered.
func (s *BaseRustListener) EnterListArray(ctx *ListArrayContext) {}

// ExitListArray is called when production listArray is exited.
func (s *BaseRustListener) ExitListArray(ctx *ListArrayContext) {}

// EnterListStructExp is called when production listStructExp is entered.
func (s *BaseRustListener) EnterListStructExp(ctx *ListStructExpContext) {}

// ExitListStructExp is called when production listStructExp is exited.
func (s *BaseRustListener) ExitListStructExp(ctx *ListStructExpContext) {}
