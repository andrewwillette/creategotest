package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	file := getTestFile()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("failed to read file: " + file)
	}
	fmt.Printf("%v\n", string(data))
}

const testTemplate = `
func Test<functionNameTest>(t *testing.T) {
	var cases = []struct {
		<caseValues>
	}{
		{
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := <functionName>()
			require.Equal(t, c.expected, result)
		})
	}
}
`

func getFuncName(file []byte) string {
	r := regexp.MustCompile(`func \w*`)
	funcWithSpace := strings.TrimLeft(r.FindStringSubmatch(string(file))[0], "func")
	return strings.TrimLeft(funcWithSpace, " ")
}

func getTestFile() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return dirname + "/git/leetcode-practice/go-leetcode/canplaceflowers/canplaceflowers_test.go"
}

func getTestFunctionAsString(funcName string, returnType string, funcParams map[string]string) string {
	return ""
}
