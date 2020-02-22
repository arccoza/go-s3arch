// +build ignore

package main

import (
	"os" 
	// "fmt"
	"log"
	"path/filepath"
	"bufio"
	"strings"
	// "strconv"
	"text/template"

	"github.com/davecgh/go-spew/spew"
)

var test = `
{{define "Head"}}
package segment

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
)

func TestGraphemes(t *testing.T) {
{{end}}
{{define "SubTest"}}
	t.Run("Num={{.Num}},Match={{.Match}}", func(t *testing.T) {
		want := []string{ {{range .Expected}}"{{.}}",{{end}} }
		got := make([]string, 0, 3)
		scanner := Graphemes(strings.NewReader("{{.Input}}"))

		for scanner.Scan() {
			got = append(got, scanner.Text())
		}

		assert.Equal(t, want, got, "they should be equal")
	})
{{end}}
{{define "Foot"}}
}
{{end}}
`

var tmpl, err = template.New("tests").Parse(test)

type fixture struct {
	Num int
	Match, Comment, Input string
	Expected []string
	Parts []string
}

func must(obj interface{}, err error) interface{} {
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

func main() {
	dir := must(os.Getwd()).(string)
	input := must(os.Open(filepath.Join(dir, "GraphemeBreakTest.txt"))).(*os.File)
	defer input.Close()
	output := must(os.Create(filepath.Join(dir, "graphemes_test.go"))).(*os.File)
	defer output.Close()

	if err := tmpl.ExecuteTemplate(output, "Head", "no data"); err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(input)
	num := 0

	for line, _, err := in.ReadLine(); err == nil; line, _, err = in.ReadLine() {
		// fmt.Printf("read %d bytes: %q\n", len(line), line)
		if line[0] == '#' {
			continue
		}

		num++
		s := string(line)
		rs := []rune(s)
		tok := []rune{}
		toks := []string{}
		fix := fixture{}
		var prvBrk rune

		Loop:
		for i, r := range rs {
			switch {
			case r == '#':
				fix.Num = num
				fix.Match = strings.Join(toks, "")
				fix.Comment = string(rs[i:])
				fix.Parts = toks
				break Loop
			case r == '÷' || r == '×':
				if len(tok) > 0 {
					strTok := string(tok)
					toks = append(toks, strTok)
					fix.Input += `\u` + strTok
					tok = tok[:0]

					if last := len(fix.Expected) - 1; prvBrk == '×' {
						fix.Expected[last] += `\u` + strTok
					} else {
						fix.Expected = append(fix.Expected, `\u` + strTok)
					}
					prvBrk = r
				}
				toks = append(toks, string(r))
			case (0x0030 <= r && r <= 0x0039) || (0x0041 <= r && r <= 0x0046):
				tok = append(tok, r)
			}
		}
		spew.Dump(fix.Input, fix.Expected)

		// err = tmpl.Execute(os.Stdout, fix)
		err = tmpl.ExecuteTemplate(output, "SubTest", fix)
		if err != nil {
			log.Fatal(err)
		}
		// break
	}

	if err := tmpl.ExecuteTemplate(output, "Foot", "no data"); err != nil {
		log.Fatal(err)
	}
}
