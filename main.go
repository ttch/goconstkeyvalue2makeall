package main


/*
	go lang const to make all by key value desc
	start here 

	we formate const (
		key value desc
	)

	 wapper to template that make any code

	and a trans is task , support multple task

*/

import (
	"go/ast"
	"go/parser"
	"go/token"
	"fmt"
	//"reflect"
	//"bytes"
	//"go/format"
)

var begin bool = false
var currentkey = "";

type ConstKeyValue struct{
	Key string;
	Value string;
	Desc string;
}

var SymbolList []string = make([]string,0,0)

var ConstMap map[string]*ConstKeyValue = map[string]*ConstKeyValue{}


func IsIn(name string)bool{
	for _,y := range SymbolList{
		if(y == name){
			return true
		}
	}
	return false
}

func main() {

    // Create the AST by parsing src.
    fset := token.NewFileSet() // positions are relative to fset
    f, err := parser.ParseFile(fset, "/Users/zhaonf/github/Server2/src/bms/models/payfee.go", nil, parser.ParseComments)
    if err != nil {
            panic(err)
    }

    ast.Inspect(f, func(n ast.Node) bool {
    	/*
    	fmt.Println(n)
    	fmt.Println(reflect.TypeOf(n))
    	fmt.Println("-------")
		*/
    	switch x := n.(type){
    	case *ast.File:
    		for _,y := range x.Scope.Objects{
    			if(y.Kind == ast.Con){
    				SymbolList = append(SymbolList,y.Name)
    			}
    		}
    		
    	case *ast.Ident :
    		if( IsIn(x.Name)){
				begin = true
				currentkey = x.Name
				ConstMap[currentkey] = new(ConstKeyValue) 
				ConstMap[currentkey].Key = x.Name
			}
    		return true
    	case *ast.BasicLit:
    		if begin {
    			if x.Kind == token.STRING {
    				ConstMap[currentkey].Value = x.Value
    			}
    		}
    		return true
    	case *ast.CommentGroup:
    		if begin {

    		}
    		return true
    	case *ast.Comment:
    		if begin {
    			ConstMap[currentkey].Desc = x.Text
    			begin = false
    		}
    		return true
    	default:
    		return true
    	}

    	return true
    });

    for _ , obj := range ConstMap{
    	fmt.Println(obj.Key)
    	fmt.Println(obj.Value)
    	fmt.Println(obj.Desc)
    }
}