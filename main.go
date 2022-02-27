package main

import (
	/* "OLC2/environment"
	"OLC2/interfaces"*/
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
	fmt.Println("end main")

}

type TreeShapeListener struct {
	*parser.BaseRustListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetInst()
	fmt.Println(result)
}
