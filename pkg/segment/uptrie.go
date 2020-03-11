package main

import (
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp"
)

func main() {
	// t := NewUPTrie()
	// t.Put([]byte{0x11, 0x11, 0x11, 0x21}, 0)
	// t.Put([]byte{0x11, 0x12, 0x11}, 1)
	r := node{edges: make(edges, 16)}
	// n := node{prefix: []byte{0x11, 0x11, 0x11}, hmask: 0XF0, tmask: 0x0F, edges: make(edges, 16), props: 0x0001}
	n := node{prefix: []byte{0x11, 0x22}, hmask: 0xF0, tmask: 0x0F, edges: make(edges, 16), props: 0x0001}
	r.edges[getNibble(0x11, 0xF0)] = n
	// n.edges[getNibble(0x11, 0xF0)] = node{prefix: []byte{0x21, 0x11}, hmask: 0XF0, tmask: 0x0F, props: 0x0002}
	n.edges[getNibble(0x11, 0x0F)] = node{prefix: []byte{0x11}, hmask: 0x0F, tmask: 0x0F, props: 0x0002}

	// r.get([]byte{0x11, 0x11, 0x11}, 0xF0)
	r.get([]byte{0x11}, 0xF0)
}

type node struct {
	prefix []byte
	hmask byte
	tmask byte
	edges edges
	props uint64
}

type edges []node

type UPTrie struct{
	root node
}

func NewUPTrie() *UPTrie {
	return &UPTrie{root: node{edges: make(edges, 16)}}
}

func (t *UPTrie) Get(key []byte) *node {
	return t.root.get(key, 0xF0)
}

func (t *UPTrie) Put(key []byte, props uint64) {
	n := t.root

	n.add(&node{prefix: key, hmask: 0xF0, tmask: 0x0F, props: props})
}

func (n *node) add(b *node) {
	idx := b.first()
	a := &n.edges[idx]
	fmt.Println("---")
	fmt.Println("edges.add 0: ", idx)

	if a.prefix == nil {
		fmt.Println("edges.add 1: ")
		n.edges[idx] = *b
		return
	}

	if len(b.prefix) < len(a.prefix) {
		*b, *a = *a, *b
		if a.edges == nil {
			a.edges = make(edges, 16)
		}
	}

	brk, brkMsk := a.comparePrefixes(b)

	off := 0
	if brkMsk != 0xF0 {
		off = 1
	}

	if brk < len(a.prefix) {
		c := &node{prefix: a.prefix[brk + off:], hmask: ^brkMsk, tmask: a.tmask, props: a.props}
		a.prefix = a.prefix[:brk + 1]
		a.props, a.tmask = 0, brkMsk
		a.add(c)
	} else if brk == len(a.prefix) {
		a.props |= b.props
		return
	}

	b.prefix = b.prefix[brk + off:]
	b.hmask = ^brkMsk
	a.add(b)
	fmt.Println("edges.add 4: ", brk, brkMsk, a, b)

}

func (a *node) comparePrefixes(b *node) (int, byte) {
	brk, msk := 0, byte(0x0F)
	for minLen := minInt(len(a.prefix), len(b.prefix)); brk < minLen; brk++ {
		if ai, bi := a.prefix[brk], b.prefix[brk]; ai != bi {
			if (ai ^ bi) & 0xF0 == 0 {
				msk = 0xF0
			} else {
				brk--
			}
			break
		}
	}

	return brk, msk
}

func (rt *node) get(key []byte, nib byte) *node {
	pp.Println(key, nib)
	for frm, brk, msk, off := walk(rt, key, nib); /*brk > 0 || msk == 0xF0*/ ; frm, brk, msk, off = walk(frm, key, msk) {
		pp.Println(frm, brk, msk, off)
		if len(key) < len(frm.prefix) {
			pp.Println("---> SHORT")
			return nil
		} else if brk < len(key) {
			pp.Println("---> BRANCH")
			key = key[brk + off:]
			msk = ^msk
		} else {
			pp.Println("---> MATCH")
			return frm
		}
	}

	pp.Println("---> END")
	return nil
}

func walk(from *node, key []byte, nib byte) (*node, int, byte, int) {
	if len(key) == 0 { return nil, -1, nib, 0 }
	idx := getNibble(key[0], nib)
	n := &from.edges[idx]
	if len(n.prefix) == 0 { return nil, -1, nib, 0 }

	brk, msk, off := 0, byte(0x0F), 1
	for minLen := minInt(len(n.prefix), len(key)); brk < minLen; brk++ {
		if a, b := n.prefix[brk], key[brk]; a != b {
			if (a ^ b) & 0xF0 == 0 {
				msk = 0xF0
				off = 0
			} else {
				brk--
			}
			break
		}
	}

	// If the entire prefix matched, do a tail check to ensure the last element
	// isn't the first nibble of the byte, otherwise we need to break there.
	// You don't need to do this sort of check on the head, which nibble of the byte
	// the head actually starts on because of the nature of prefix tries,
	// the prefixes have to match (ie. the earlier nibbles have to match).
	if brk == len(n.prefix) && n.tmask == 0xF0 {
		brk, msk, off = brk - 1, 0xF0, 0
	}

	return n, brk, msk, off
}

func getNibble(b, m byte) byte {
	n := b & m
	if n > 0x0F { n = n >> 4 }
	return n
}

type Index [2]int
func (i *Index) Get(bs []byte) byte {
	x := i[1] % 2
	m := byte((1 - x) * 0x0F + x * 0xF0)
	return getNibble(bs[i[0]], m)
}

func (n *node) first() byte {
	f := n.prefix[0] & n.hmask
	if f > 0x0F {
		f = f >> 4
	}
	return f
}

func (n *node) last() byte {
	l := n.prefix[len(n.prefix) - 1] & n.tmask
	if l > 0x0F {
		l = l >> 4
	}
	return l
}

func minInt(a, b int) int {
	if a < b { return a }
	return b
}

func nibble(b, m byte) byte {
	if m == 0xF0 {
		return b >> 4
	} else {
		return b & m
	}
}
