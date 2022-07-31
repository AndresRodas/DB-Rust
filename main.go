package main

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"OLC2/parser"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"html/template"
	"net/http"
	"strings"
)

type TreeShapeListener struct {
	*parser.BaseRustListener
	Code environment.Code
}
type Interpreter struct {
	Input  string
	Output string
}

type GetPost struct {
	text  string
	valor int
}

func main() {
	t, err := template.ParseFiles("./index.html")
	//POST petition
	http.HandleFunc("/rustdb", func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Set("Content-Type", "application/json")
		Code := req.FormValue("text")
		message := ""
		//validando entrada
		if len(Code) > 0 {
			//creando ast
			var Ast environment.AST
			Ast = Run(Code)
			//seteando salida
			message = Ast.GetPrint()
			//fmt.Println(Ast.Symbols.ToArray())
		}
		//seteando respuesta
		resp := make(map[string]string)
		resp["message"] = message
		resp["testi"] = "[3,2,2,2]"
		resp["symbolTable"] = "jaja"

		//convirtiendo a json
		jsonResp, _ := json.Marshal(resp)
		//respondiendo
		res.Write(jsonResp)
	})

	//HTTP - VIEW
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err = t.Execute(res, nil)
	})
	if err != nil {
		return
	}
	fmt.Println("Server running on http://localhost:8086")
	err = http.ListenAndServe(":8086", nil)
	if err != nil {
		return
	}

}

func Run(code string) environment.AST {
	//reading file
	//is, _ := antlr.NewFileStream("in.txt")
	is := antlr.NewInputStream(code)
	//creating lexer
	lexer := parser.NewRustLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//creating parser
	p := parser.NewRust(stream)
	p.BuildParseTrees = true
	tree := p.Start()
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	//get code
	Code := listener.Code
	//create ast
	var Ast environment.AST
	//create environment
	var globalEnv environment.Environment
	globalEnv = environment.NewEnvironment(nil, "GLOBAL")
	//saving global environments
	for _, inst := range Code.Instructions.ToArray() {
		inst.(interfaces.Instruction).Ejecutar(&Ast, globalEnv)
	}
	//running main
	var mainEnv environment.Environment
	mainEnv = environment.NewEnvironment(globalEnv, "MAIN")
	for _, bloc := range Code.Main.ToArray() {
		if strings.Contains(fmt.Sprintf("%T", bloc), "instructions") {
			bloc.(interfaces.Instruction).Ejecutar(&Ast, mainEnv)
		} else if strings.Contains(fmt.Sprintf("%T", bloc), "expressions") {
			bloc.(interfaces.Expression).Ejecutar(&Ast, mainEnv)
		}
	}
	//returning ast
	return Ast

}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	this.Code = ctx.GetCode()
}
