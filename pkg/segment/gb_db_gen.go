//go:generate go run gb_db_gen.go
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
	"regexp"
	"encoding/binary"

	// "github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp"
)

var db = `
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

var tmpl, err = template.New("db").Parse(db)

var re = regexp.MustCompile(`(?m)^(?:([0-9a-fA-F]{4,8}\.\.[0-9a-fA-F]{4,8})|([0-9a-fA-F]{4,8}))\s*;\s*(\w*)`)

func must(obj interface{}, err error) interface{} {
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

func main() {
	parse()
}

func parse() {
	dir := must(os.Getwd()).(string)
	input := must(os.Open(filepath.Join(dir, "GraphemeBreakProperty.txt"))).(*os.File)
	defer input.Close()
	output := must(os.Create(filepath.Join(dir, "gb_db.go"))).(*os.File)
	defer output.Close()

	if err := tmpl.ExecuteTemplate(output, "Head", "no data"); err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(input)
	// num := 0
	trie := NewUPTrie()
	buf := make([]byte, 4)

	// Get each line in the GraphemeBreakProperty.txt file
	for line, _, err := in.ReadLine(); err == nil; line, _, err = in.ReadLine() {
		// Skip comment only lines (starts with a #)
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		match := re.FindStringSubmatch(string(line))
		fmt.Printf("%q\n", match)
		prop := match[3]

		if len(match[1]) > 0 {
			rng := strings.Split(match[1], "..")
			min, _ := strconv.ParseInt(rng[0], 16, 32)
			max, _ := strconv.ParseInt(rng[1], 16, 32)
			// node := NewNode()
			// fmt.Println(min, max, err)

			for cp := min; cp <= max; cp++ {
				// bs := make([]byte, 4)
				// bs = putBytes(bs, cp)
				// fmt.Printf("%v %v\n", prop, bs)
				// pp.Println(prop, cp, uint32ToBytes(buf, uint32(cp)))
				trie.Put(uint32ToBytes(buf, uint32(cp)), nameToFlag[prop])
			}
		} else {
			// fmt.Println(prop, []byte(string(match[2])))
			cp, _ := strconv.ParseInt(match[2], 16, 32)
			// pp.Println(prop, cp, uint32ToBytes(buf, uint32(cp)))
			trie.Put(uint32ToBytes(buf, uint32(cp)), nameToFlag[prop])
		}


	}

	pp.Println(trie.Get(uint32ToBytes(buf, uint32(0x11D46))))
	// pp.Println(trie)

	if err := tmpl.ExecuteTemplate(output, "Foot", "no data"); err != nil {
		log.Fatal(err)
	}
}

func uint32ToBytes(buf []byte, num uint32) []byte {
	binary.LittleEndian.PutUint32(buf, num)
	switch {
	case num <= 0xFF:
		return buf[:1]
	case num <= 0xFFFF:
		return buf[:2]
	case num <= 0xFFFFFF:
		return buf[:3]
	}
	return buf
}

const (
	CR = 1 << (iota + 2)
	LF
	Control
	Extend
	Regional_Indicator
	Prepend
	SpacingMark
	L
	V
	T
	LV
	LVT
	ExtPict
	ZWJ
)

var nameToFlag = map[string]uint64{
	"CR":CR,
	"LF": LF,
	"Control": Control,
	"Extend": Extend,
	"Regional_Indicator": Regional_Indicator,
	"Prepend": Prepend,
	"SpacingMark": SpacingMark,
	"L": L,
	"V": V,
	"T": T,
	"LV": LV,
	"LVT": LVT,
	"ExtPict": ExtPict,
	"ZWJ": ZWJ,
}
