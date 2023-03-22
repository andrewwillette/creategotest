package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleGoFile = `
package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	return false
}
`

func TestGetFuncName(t *testing.T) {
	funcName := getFuncName([]byte(exampleGoFile))
	require.Equal(t, "canPlaceFlowers", funcName)
}

func TestGetTestFile(t *testing.T) {
	getTestFile()
}

func TestGetTestString(t *testing.T) {
}
