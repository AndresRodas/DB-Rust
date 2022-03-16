package main

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"OLC2/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {

	//reading file
	is, _ := antlr.NewFileStream("in.txt")
	//creating lexer
	lexer := parser.NewRustLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//creating parser
	p := parser.NewRust(stream)
	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

}

type TreeShapeListener struct {
	*parser.BaseRustListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	//create ast
	var Ast environment.AST
	//get code
	var Code environment.Code
	Code = ctx.GetCode()
	//create environment
	var globalEnv environment.Environment
	globalEnv = environment.NewEnvironment(nil, "Global")

	for _, inst := range Code.Main.ToArray() {
		inst.(interfaces.Instruction).Ejecutar(&Ast, globalEnv)
	}
	//fmt.Println(globalEnv.Tabla)
	//print values
	fmt.Println(Ast.GetPrint())

}
