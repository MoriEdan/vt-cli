package yaml

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type T struct {
	data interface{}
	yaml string
}

func Y(s string) string {
	t := make([]string, 0)
	for _, l := range strings.Split(s, "\n") {
		if len(l) > 3 {
			t = append(t, strings.Replace(l[3:], "\t", " ", -1))
		}
	}
	return strings.Join(t, "\n") + "\n"
}

var tests = []T{
	{
		data: "foo",
		yaml: Y(`
			"foo"
		`),
	},
	{
		data: 1,
		yaml: Y(`
			1
		`),
	},
	{
		data: false,
		yaml: Y(`
			false
			`),
	},
	{
		data: true,
		yaml: Y(`
			true
			`),
	},
	{
		data: map[string]string{},
		yaml: Y(``),
	},
	{
		data: map[string]map[string]string{
			"foo": map[string]string{},
		},
		yaml: Y(`
			foo: `),
	},
	{
		data: []string{},
		yaml: Y(`
			[]`),
	},
	{
		data: map[string]string{
			"uno":  "1",
			"dos":  "2",
			"tres": "3",
		},
		yaml: Y(`
			dos: "2"
			tres: "3"
			uno: "1"
			`),
	},
	{
		data: []string{
			"uno",
			"dos",
			"tres",
		},
		yaml: Y(`
			- "uno"
			- "dos"
			- "tres"
			`),
	},
	{
		data: struct {
			Foo string
			Bar string
		}{
			"uno",
			"dos",
		},
		// Struct fields are re-ordered alphabetically.
		yaml: Y(`
			Bar: "dos"
			Foo: "uno"
			`),
	},
	{
		data: struct {
			Foo string
		}{
			"uno\ndos",
		},
		yaml: Y(`
			Foo: |
			  uno
			  dos
			`),
	},
	{
		data: map[string]interface{}{
			"numbers": []interface{}{
				map[string]string{
					"number":  "1",
					"numeral": "first",
				},
				map[string]string{
					"number":  "2",
					"numeral": "second",
				},
			},
		},
		yaml: Y(`
			numbers:` + " " + `
			- number: "1"
			  numeral: "first"
			- number: "2"
			  numeral: "second"
			`),
	},
}

func TestYAML(t *testing.T) {

	var b bytes.Buffer

	for _, test := range tests {
		enc := NewEncoder(&b)
		enc.IndentSize = 1
		assert.NoError(t, enc.Encode(test.data))
		assert.Equal(t, test.yaml, b.String(), "Test %v", test.data)
		b.Reset()
	}

	enc := NewEncoder(&b)
	enc.IndentSize = 1
	assert.NoError(t, enc.Encode(tests[5].data))
	assert.Equal(t, tests[5].yaml, b.String())

}
