// +build ignore

package main

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
	"bufio"
	"strings"
	"strconv"
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
	t.Parallel()
{{end}}
{{define "SubTest"}}
	/* {{.Content}} */
	t.Run("Num={{.Num}},Match={{.Match}}", func(t *testing.T) {
		t.Log("Num={{.Num}},Match={{.Match}},Input={{.Input}}")
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
	Match, Comment, Input, Content string
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

	// Get each line in the GraphemeBreakTest.txt file
	for line, _, err := in.ReadLine(); err == nil; line, _, err = in.ReadLine() {
		// Skip comment only lines (starts with a #)
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
			// This is the beginning of the comment on this line, stop the loop here
			case r == '#':
				fix.Num = num
				fix.Match = strings.Join(toks, "")
				fix.Comment = string(rs[i:])
				fix.Content, _ = strconv.Unquote(`"` + fix.Input + `"`)
				fix.Parts = toks
				break Loop
			// These are grapheme break markers;
			// ÷ means break here
			// × means do not break and merge with next rune
			case r == '÷' || r == '×':
				if len(tok) > 0 {
					// Convert tok (slice of runes) to a string and add it to the
					// list of toks
					strTok := string(tok)
					toks = append(toks, strTok)

					// Turn tok into the appropriate length unicode escape string
					if len(tok) <= 4 {
						strTok = `\u` + fmt.Sprintf("%04s", strTok)
					} else {
						strTok = `\U` + fmt.Sprintf("%08s", strTok)
					}

					// Create an string of the escaped runes as input for testing
					fix.Input += strTok
					// Clear tok for the next tok to gather
					tok = tok[:0]

					// Gather up the expected graphemes as a slice of strings
					// using prvBrk to track the previous break char (÷ or ×)
					if last := len(fix.Expected) - 1; prvBrk == '×' {
						fix.Expected[last] += strTok
					} else {
						fix.Expected = append(fix.Expected, strTok)
					}
					prvBrk = r
				}
				toks = append(toks, string(r))
			// If r is one of the hex characters gather it into tok
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

		if num > 10000 {
			break
		}
	}

	if err := tmpl.ExecuteTemplate(output, "Foot", "no data"); err != nil {
		log.Fatal(err)
	}
}
