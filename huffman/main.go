package main

import (
	"fmt"
	"sort"
)

type huffmanTree struct {
	Parent *huffmanTree
	Left   *huffmanTree
	Right  *huffmanTree
	Weight float64
	Value  rune
}

func sortHuffman(hf []*huffmanTree) {
	sort.Slice(hf, func(i, j int) bool {
		return hf[i].Weight < hf[j].Weight
	})
}

func sortedHuffman(hf []huffmanTree) []huffmanTree {
	sort.Slice(hf, func(i, j int) bool {
		return hf[i].Weight < hf[j].Weight
	})
	return hf
}

func build(leaves []*huffmanTree) *huffmanTree {
	sortHuffman(leaves)
	if len(leaves) == 0 {
		return nil
	}
	for len(leaves) > 1 {
		left, right := leaves[0], leaves[1]
		parentWeight := left.Weight + right.Weight
		parent := &huffmanTree{Left: left, Right: right, Weight: parentWeight}
		left.Parent = parent
		right.Parent = parent
		leaves = append(leaves[2:], parent)
		sortHuffman(leaves)
	}
	return leaves[0]
}

func (hf *huffmanTree) String() string {
	return fmt.Sprintf("{%s, %s, %s, %s}", fmt.Sprint(hf.Left), fmt.Sprint(hf.Right),
		fmt.Sprintf("%v", hf.Weight), fmt.Sprintf("%v", hf.Value))
}

func main() {
	hf := []*huffmanTree{
		{Weight: 1.0, Value: 'a'},
		{Weight: 2.0, Value: 'b'},
	}
	for _, value := range hf {
		fmt.Println(value)
	}
	after := build(hf)
	fmt.Println(after)
}
