package main

import (
	"crypto/sha256"
	"fmt"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

type MerkleTree struct {
	RootNode *MerkleNode
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}
	mNode.Left = left
	mNode.Right = right
	return &mNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}
	mTree := MerkleTree{&nodes[0]}
	return &mTree
}

func showMerkleTree(root *MerkleNode) {
	if root == nil {
		return
	} else {
		PrintNode(root, 1)
	}
	showMerkleTree(root.Left)
	showMerkleTree(root.Right)
}

func PrintNode(node *MerkleNode, level int) {
	fmt.Printf("%p\n", node)
	if node != nil {
		fmt.Printf("left[%p], right[%p], data(%x)\n", node.Left, node.Right, node.Data)
	}
}

func main() {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
		[]byte("node4"),
	}
	tree := NewMerkleTree(data)
	showMerkleTree(tree.RootNode)
}
