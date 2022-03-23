package main

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"OLC2/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
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
	globalEnv = environment.NewEnvironment(nil, "GLOBAL")
	//evaluando variables globales
	for _, inst := range Code.Instructions.ToArray() {
		inst.(interfaces.Instruction).Ejecutar(&Ast, globalEnv)
	}
	//evaluando el main
	var mainEnv environment.Environment
	mainEnv = environment.NewEnvironment(globalEnv, "MAIN")
	for _, bloc := range Code.Main.ToArray() {
		if strings.Contains(fmt.Sprintf("%T", bloc), "instructions") {
			bloc.(interfaces.Instruction).Ejecutar(&Ast, mainEnv)

		} else if strings.Contains(fmt.Sprintf("%T", bloc), "expressions") {
			bloc.(interfaces.Expression).Ejecutar(&Ast, mainEnv)

		}
		//bloc.(interfaces.Instruction).Ejecutar(&Ast, mainEnv)
	}
	//print values
	fmt.Println(Ast.GetPrint())

}
