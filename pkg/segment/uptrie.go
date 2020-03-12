package main

import (
	// "fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp"
)

func main() {
	// t := NewUPTrie()
	// t.Put([]byte{0x11, 0x11, 0x11, 0x21}, 0)
	// t.Put([]byte{0x11, 0x12, 0x11}, 1)
	// r := node{edges: make(edges, 16)}
	// // n := node{prefix: []byte{0x11, 0x11, 0x11}, hmask: 0XF0, tmask: 0x0F, edges: make(edges, 16), props: 0x0001}
	// n := node{prefix: []byte{0x11, 0x22}, hmask: 0xF0, tmask: 0x0F, edges: make(edges, 16), props: 0x0001}
	// r.edges[getNibble(0x11, 0xF0)] = n
	// // n.edges[getNibble(0x11, 0xF0)] = node{prefix: []byte{0x21, 0x11}, hmask: 0XF0, tmask: 0x0F, props: 0x0002}
	// n.edges[getNibble(0x11, 0x0F)] = node{prefix: []byte{0x11}, hmask: 0x0F, tmask: 0x0F, props: 0x0002}

	// // r.get([]byte{0x11, 0x11, 0x11}, 0xF0)
	// r.get([]byte{0x11}, 0xF0)

	t := NewUPTrie()
	t.Put([]byte("remember"), 0x0001)
	t.Put([]byte("remain"), 0x0002)
	// t.Get([]byte("rem"))
	pp.Println(t.Get([]byte("rem")))
}

type UPTrie struct{
	root node
}

type node struct {
	prefix []byte
	props uint64
	hmask byte
	tmask byte
	edges edges
}

type edges []*node

func NewUPTrie() *UPTrie {
	return &UPTrie{root: node{edges: make(edges, 16)}}
}

func (t *UPTrie) Get(key []byte) *node {
	return t.root.get(key, 0xF0)
}

func (t *UPTrie) Put(key []byte, props uint64) {
	t.root.put(&node{prefix: key, hmask: 0xF0, tmask: 0x0F, props: props})
}

func (e edges) add(n *node) {
	idx := getNibble(n.prefix[0], n.hmask)
	e[idx] = n
}

func (head *node) split(brk int, msk byte, off int) (tail *node) {
	if brk < len(head.prefix) {
		tail = &node{prefix: head.prefix[brk + off:], edges: head.edges, hmask: ^msk, tmask: head.tmask, props: head.props}
		head.prefix = head.prefix[:brk + off + 1]
		head.edges = make(edges, 16)
		head.tmask, head.props = msk, 0
	}

	return tail
}

func (n *node) chop(brk int, msk byte, off int) bool {
	if brk < len(n.prefix) {
		n.prefix = n.prefix[brk + off:]
		n.hmask = ^msk
		return true
	}
	return false
}

func (prv *node) put(n *node) {
	for frm, brk, msk, off := walk(prv, n.prefix, n.hmask); brk >= 0 ; frm, brk, msk, off = walk(frm, n.prefix, n.hmask) {
		// SPLIT EXISTING
		if tail := frm.split(brk, msk, off); tail != nil {
			frm.edges.add(tail)
		}
		// MATCH
		if !n.chop(brk, msk, off) {
			frm.props |= n.props
			return
		// NO MATCH
		} else { prv = frm }

	}
	if prv.edges == nil {
		prv.edges = make(edges, 16)
	}
	prv.edges.add(n)
}

func (rt *node) get(key []byte, nib byte) *node {
	for frm, brk, msk, off := walk(rt, key, nib); brk >= 0 ; frm, brk, msk, off = walk(frm, key, msk) {
		// SHORT
		if len(key) < len(frm.prefix) {
			return frm
		// BRANCH
		} else if brk < len(key) {
			key = key[brk + off:]
			msk = ^msk
		// MATCH
		} else {
			return frm
		}
	}

	return nil
}

func walk(from *node, key []byte, nib byte) (*node, int, byte, int) {
	if len(key) == 0 { return nil, -1, nib, 0 }
	idx := getNibble(key[0], nib)
	n := from.edges[idx]
	if n == nil { return nil, -1, nib, 0 }
	// if len(n.prefix) == 0 { return nil, -1, nib, 0 }

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

func minInt(a, b int) int {
	if a < b { return a }
	return b
}

type Index [2]int
func (i *Index) Get(bs []byte) byte {
	x := i[1] % 2
	m := byte((1 - x) * 0x0F + x * 0xF0)
	return getNibble(bs[i[0]], m)
}
