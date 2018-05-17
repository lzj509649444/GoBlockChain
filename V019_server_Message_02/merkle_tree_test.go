package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -cover=true merkle_tree.go merkle_tree_test.go
func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
		[]byte("node4"),
		[]byte("node5"),
	}
	// Level 1
	n11 := NewMerkleNode(nil, nil, data[0])
	n12 := NewMerkleNode(nil, nil, data[1])
	n13 := NewMerkleNode(nil, nil, data[2])
	n14 := NewMerkleNode(nil, nil, data[3])
	n15 := NewMerkleNode(nil, nil, data[4])

	// Level 2
	n21 := NewMerkleNode(n11, n12, nil)
	n22 := NewMerkleNode(n13, n14, nil)
	n23 := NewMerkleNode(n15, n15, nil)

	// Level 3
	n31 := NewMerkleNode(n21, n22, nil)
	n32 := NewMerkleNode(n23, n23, nil)

	// Level 4
	n41 := NewMerkleNode(n31, n32, nil)

	rootHash := fmt.Sprintf("%x", n41.Data)
	mTree := NewMerkleTree(data)

	assert.Equal(t, rootHash, fmt.Sprintf("%x", mTree.RootNode.Data), "Merkle tree root hash is correct")
}
