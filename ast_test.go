package main

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

const testFile = "./testdata/example_test.go"

func TestGetFuncNameAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, testFile, nil, parser.ParseComments)
	require.NoError(t, err)
	funcName := GetFuncName(file)
	require.Equal(t, "minInsertions", funcName)
}

func TestGetFuncParamsAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, testFile, nil, parser.ParseComments)
	require.NoError(t, err)
	funcParams := GetFuncParams(file)
	expectedParams := []FuncParam{{paramName: "s", paramType: "string"}}
	require.Equal(t, expectedParams, funcParams)
}

func TestGetReturnTypeAst(t *testing.T) {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, testFile, nil, parser.ParseComments)
	require.NoError(t, err)
	returnType := GetReturnType(file)
	require.Equal(t, "int", returnType)
}
