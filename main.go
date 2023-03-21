package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

const testFile = "/Users/andrewwillette/git/leetcode-practice/go-leetcode/canplaceflowers/canplaceflowers_test.go"
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

func main() {
	file := testFile
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("failed to read file: " + file)
	}
	fmt.Printf("%v\n", string(data))
}

func getFuncName(file []byte) string {
	r := regexp.MustCompile(`func \w*`)
	fmt.Printf("%#v\n", r.FindStringSubmatch(string(file)))
	return ""
}
