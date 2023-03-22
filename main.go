package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	file := getTestFile()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("failed to read file: " + file)
	}
	funcToTest := FuncToTest{function: string(data)}
	funcName := funcToTest.getFuncName()
	returnType := funcToTest.getReturnType()
	funcParams := funcToTest.getFuncParams()
	template := TestTemplate{template: testTemplate}
	result := template.getTestFunctionAsString(funcName, returnType, funcParams)
	fmt.Printf("%v\n", result)
}

func getTestFile() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return dirname + "/git/leetcode-practice/go-leetcode/canplaceflowers/canplaceflowers_test.go"
}

var (
	testTemplate = `
func Test<testFunctionSplint>(t *testing.T) {
	var cases = []struct {
		<caseSplint>
		<expectedSplint>
	}{
		{
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := <functionSplint>(<paramsSplint>)
			require.Equal(t, c.expected, result)
		})
	}
}
`
)

const (
	testFuncSplint = "<testFunctionSplint>"
	funcSplint     = "<functionSplint>"
	paramsSplint   = "<paramsSplint>"
	caseSplint     = "<caseSplint>"
	expectedSplint = "<expectedSplint>"
)

type FuncToTest struct {
	function string
}

type FuncParam struct {
	paramName string
	paramType string
}

func (funcToTest *FuncToTest) getFuncName() string {
	r := regexp.MustCompile(`func \w*`)
	funcWithSpace := strings.TrimLeft(r.FindStringSubmatch(funcToTest.function)[0], "func")
	return strings.TrimLeft(funcWithSpace, " ")
}

func (funcToTest *FuncToTest) getFuncParams() []FuncParam {
	r := regexp.MustCompile(`\(.*\)`)
	returnVal := r.FindStringSubmatch(funcToTest.function)[0]
	returnVal = strings.TrimLeft(returnVal, "(")
	returnVal = strings.TrimRight(returnVal, ")")
	params := strings.Split(returnVal, ",")
	funcParams := []FuncParam{}
	for _, v := range params {
		v = strings.Trim(v, " ")
		paramNameVal := strings.Split(v, " ")
		funcParams = append(funcParams,
			FuncParam{
				paramName: paramNameVal[0],
				paramType: paramNameVal[1],
			},
		)
	}
	return funcParams
}

func (funcToTest *FuncToTest) getReturnType() string {
	r := regexp.MustCompile(`\) \w* {`)
	returnVal := r.FindStringSubmatch(funcToTest.function)[0]
	returnVal = strings.TrimLeft(returnVal, ") ")
	returnVal = strings.TrimRight(returnVal, " {")
	return returnVal
}

type TestTemplate struct {
	template string
}

func (template *TestTemplate) getTestFunctionAsString(funcName string, returnType string, funcParams []FuncParam) string {
	template.insertTestFunctionSplint(funcName)
	template.insertFunctionSplint(funcName)
	template.insertParamsSplint(funcParams)
	template.insertCaseSplint(funcParams)
	template.insertExpectedSplint(returnType)
	return template.template
}

func (template *TestTemplate) insertTestFunctionSplint(funcName string) {
	m := regexp.MustCompile(testFuncSplint)
	res := m.ReplaceAllString(template.template, cases.Title(language.English, cases.NoLower).String(funcName))
	template.template = res
}

func (template *TestTemplate) insertFunctionSplint(funcName string) {
	m := regexp.MustCompile(funcSplint)
	res := m.ReplaceAllString(template.template, funcName)
	template.template = res
}

func (template *TestTemplate) insertCaseSplint(funcParams []FuncParam) {
	var sb strings.Builder
	cnt := 0
	for _, v := range funcParams {
		if cnt == 0 {
			sb.WriteString(v.paramName + " " + v.paramType + "\n")
		} else if cnt < len(funcParams)-1 {
			sb.WriteString("\t\t" + v.paramName + " " + v.paramType + "\n")
		} else {
			sb.WriteString("\t\t" + v.paramName + " " + v.paramType)
		}
		cnt++
	}
	m := regexp.MustCompile(caseSplint)
	res := m.ReplaceAllString(template.template, sb.String())
	template.template = res
}

func (template *TestTemplate) insertExpectedSplint(returnType string) {
	m := regexp.MustCompile(expectedSplint)
	var sb strings.Builder
	sb.WriteString("expected " + returnType)
	res := m.ReplaceAllString(template.template, sb.String())
	template.template = res
}

func (template *TestTemplate) insertParamsSplint(params []FuncParam) {
	m := regexp.MustCompile(paramsSplint)
	var sb strings.Builder
	cnt := 0
	for _, v := range params {
		if cnt < len(params)-1 {
			sb.WriteString("c." + v.paramName + ", ")
		} else {
			sb.WriteString("c." + v.paramName)
		}
		cnt++
	}
	res := m.ReplaceAllString(template.template, sb.String())
	template.template = res
}
