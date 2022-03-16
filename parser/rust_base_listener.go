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

// EnterImpression is called when production impression is entered.
func (s *BaseRustListener) EnterImpression(ctx *ImpressionContext) {}

// ExitImpression is called when production impression is exited.
func (s *BaseRustListener) ExitImpression(ctx *ImpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseRustListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseRustListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseRustListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseRustListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseRustListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseRustListener) ExitFunction(ctx *FunctionContext) {}

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

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseRustListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseRustListener) ExitExpr_arit(ctx *Expr_aritContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseRustListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseRustListener) ExitPrimitive(ctx *PrimitiveContext) {}

// EnterListArray is called when production listArray is entered.
func (s *BaseRustListener) EnterListArray(ctx *ListArrayContext) {}

// ExitListArray is called when production listArray is exited.
func (s *BaseRustListener) ExitListArray(ctx *ListArrayContext) {}
