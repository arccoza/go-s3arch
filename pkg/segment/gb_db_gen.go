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
	// root := Node{}
	node := NewNode()

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
				node.PutRune(rune(cp), nameToFlag[prop])
			}
		} else {
			// fmt.Println(prop, []byte(string(match[2])))
			cp, _ := strconv.ParseInt(match[2], 16, 32)
			node.PutRune(rune(cp), nameToFlag[prop])
		}


	}

	// node.PutRune(0x11D46, nameToFlag["Prepend"])
	fmt.Println(node.GetRune(0x1F1E6).Props, nameToFlag["Regional_Indicator"])

	if err := tmpl.ExecuteTemplate(output, "Foot", "no data"); err != nil {
		log.Fatal(err)
	}
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

type Node struct {
	Value byte
	Props uint64
	Next []Node
}

func NewNode() *Node {
	return &Node{Next: make([]Node, 256)}
}

func (n *Node) PutRune(r rune, prop uint64) {
	bs := make([]byte, 4)
	if r <= 0xFFFF {
		bs = bs[:2]
		binary.LittleEndian.PutUint16(bs, uint16(r))
	} else {
		binary.LittleEndian.PutUint32(bs, uint32(r))
	}

	n.Put(bs, prop)
}

func (n *Node) Put(bs []byte, prop uint64) {
	if len(bs) == 0 { return }
	b := bs[0]

	if n.Next == nil {
		n.Next = make([]Node, 256)
	}

	nn := n.Next[b]
	nn.Value = b

	if len(bs) == 1 {
		nn.Props |= prop
		n.Next[b] = nn
		return
	}

	nn.Put(bs[1:], prop)
	n.Next[b] = nn
}

func (n *Node) GetRune(r rune) *Node {
	bs := make([]byte, 4)
	if r <= 0xFFFF {
		bs = bs[:2]
		binary.LittleEndian.PutUint16(bs, uint16(r))
	} else {
		binary.LittleEndian.PutUint32(bs, uint32(r))
	}

	return n.Get(bs)
}

func (n *Node) Get(bs []byte) *Node {
	if len(bs) == 0 { return n }

	// fmt.Println(bs, bs[0], len(n.Next))
	nn := n.Next[bs[0]]

	if len(bs) == 1 {
		// fmt.Println(nn, Prepend)
		return &nn
	}

	// nn := n.Next[int(bs[0])]
	return nn.Get(bs[1:])
}
