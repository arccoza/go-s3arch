// +build ignore

package main

import (
	"os" 
	// "fmt"
	"log"
	"path/filepath"
	"bufio"
	"strings"
	"text/template"

	"github.com/davecgh/go-spew/spew"
)

// var body = `
// func TestGraphemes(t *testing.T) {
//     got := Abs(-1)
//     if got != 1 {
//         t.Errorf("Abs(-1) = %d; want 1", got)
//     }
// }
// `

var test = `
t.Run("A=1", func(t *testing.T) {
	{{.}}
})
`

var tmpl, err = template.New("todos").Parse(test)


func main() {
	if dir, err := os.Getwd(); err != nil {
		log.Fatal(err)
	} else if file, err := os.Open(filepath.Join(dir, "GraphemeBreakTest.txt")); err != nil {
		log.Fatal(err)
	} else {
		f := bufio.NewReader(file)

		for line, _, err := f.ReadLine(); err == nil; line, _, err = f.ReadLine() {
			// fmt.Printf("read %d bytes: %q\n", len(line), line)
			if line[0] == '#' {
				continue
			}
			s := string(line)
			rs := []rune(s)
			tok := []rune{}
			toks := []string{}
			name := ""
			comment := ""

			Loop:
			for i, r := range rs {
				switch {
				case r == '#':
					// toks = append(toks, string(rs[i:]))
					name = strings.Join(toks, "")
					comment = string(rs[i:])
					break Loop
				case r == '÷' || r == '×':
					if len(tok) > 0 {
						toks = append(toks, string(tok))
						tok = tok[:0]
					}
					toks = append(toks, string(r))
				case (0x0030 <= r && r <= 0x0039) || (0x0041 <= r && r <= 0x0046):
					tok = append(tok, r)
				}
			}
			spew.Dump(name, comment)

			err = tmpl.Execute(os.Stdout, toks)
			if err != nil {
				panic(err)
			}
			// break
		}
	}
}
