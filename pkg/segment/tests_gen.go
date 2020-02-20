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

	// "github.com/davecgh/go-spew/spew"
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
t.Run("Name={{.Name}}Num={{.Num}}", func(t *testing.T) {
	{{.Parts}}
})
`

var tmpl, err = template.New("todos").Parse(test)


type fixture struct {
	Num int
	Name, Comment string
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
	output := must(os.Create(filepath.Join(dir, "grapheme_test.go"))).(*os.File)
	defer output.Close()

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

		Loop:
		for i, r := range rs {
			switch {
			case r == '#':
				fix.Num = num
				fix.Name = strings.Join(toks, "")
				fix.Comment = string(rs[i:])
				fix.Parts = toks
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
		// spew.Dump(dir)

		// err = tmpl.Execute(os.Stdout, fix)
		err = tmpl.Execute(output, fix)
		if err != nil {
			panic(err)
		}
		// break
	}
}
