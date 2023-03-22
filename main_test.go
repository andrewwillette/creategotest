package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleGoFile = `
package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	return false
}
`

var (
	template   = TestTemplate{template: testTemplate}
	funcToTest = FuncToTest{function: exampleGoFile}
)

func TestGetFuncName(t *testing.T) {
	funcToTest := FuncToTest{function: exampleGoFile}
	funcName := funcToTest.getFuncName()
	require.Equal(t, "canPlaceFlowers", funcName)
}

func TestGetReturnType(t *testing.T) {
	funcToTest := FuncToTest{function: exampleGoFile}
	returnType := funcToTest.getReturnType()
	require.Equal(t, "bool", returnType)
}

func TestGetTestFile(t *testing.T) {
	getTestFile()
}

func TestGetTestFunctionAsString(t *testing.T) {
	funcParams := map[string]string{
		"flowerbed": "[]int",
		"n":         "int",
	}
	result := template.getTestFunctionAsString("canPlaceFlowers", "bool", funcParams)
	require.Contains(t, result, "TestCanPlaceFlowers(t *testing.T)")
	require.Contains(t, template.template, "canPlaceFlowers(")
	require.Contains(t, template.template, "flowerbed []int\nn int")
	require.Contains(t, template.template, "expected bool")
	require.Contains(t, template.template, "(c.flowerbed, c.n)")
}

func TestInsertTestFunctionSplint(t *testing.T) {
	template.insertTestFunctionSplint("canPlaceFlowers")
	require.Contains(t, template.template, "TestCanPlaceFlowers(t *testing.T)")
}

func TestInsertFunctionSplint(t *testing.T) {
	template.insertFunctionSplint("canPlaceFlowers")
	require.Contains(t, template.template, "canPlaceFlowers(")
}

func TestInsertCaseSplint(t *testing.T) {
	funcParams := map[string]string{
		"flowerbed": "[]int",
		"n":         "int",
	}
	template.insertCaseSplint(funcParams)
	require.Contains(t, template.template, "flowerbed []int\nn int")
}

func TestInsertExpectedSplint(t *testing.T) {
	template.insertExpectedSplint("int")
	require.Contains(t, template.template, "expected int")
}

func TestInsertParamsSplint(t *testing.T) {
	funcParams := map[string]string{
		"flowerbed": "[]int",
		"n":         "int",
	}
	template.insertParamsSplint(funcParams)
	require.Contains(t, template.template, "(c.flowerbed, c.n)")
}
