package main

import "go/ast"

func GetFuncName(file *ast.File) string {
	funcName := ""
	ast.Inspect(file, func(n ast.Node) bool {
		funcDef, ok := n.(*ast.FuncDecl)
		if ok {
			funcName = funcDef.Name.Name
		}
		return true
	})
	return funcName
}

func GetFuncParams(file *ast.File) []FuncParam {
	funcParams := []FuncParam{}
	ast.Inspect(file, func(n ast.Node) bool {
		funcDef, ok := n.(*ast.FuncDecl)
		if ok {
			// funcName = funcDef.Name.Name
			for _, v := range funcDef.Type.Params.List {
				funcParams = append(funcParams,
					FuncParam{
						paramName: v.Names[0].Name,
						paramType: v.Type.(*ast.Ident).Name,
					},
				)
			}
		}
		return true
	})
	return funcParams
}

func GetReturnType(file *ast.File) string {
	returnType := ""
	ast.Inspect(file, func(n ast.Node) bool {
		funcDef, ok := n.(*ast.FuncDecl)
		if ok {
			returnType = funcDef.Type.Results.List[0].Type.(*ast.Ident).Name
		}
		return true
	})
	return returnType
}
