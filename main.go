package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type merkleNode struct {
	isLeaf     bool
	hash       string
	leftChild  *merkleNode
	rightChild *merkleNode
}

func toSHA256(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}

func CreateMerkleTree(hashes []string) (merkleNode, error) {
	lengthHashed := len(hashes)
	t := merkleNode{isLeaf: false, hash: "", leftChild: nil, rightChild: nil}

	if lengthHashed <= 0 {
		return merkleNode{}, errors.New("Empty array of hashes")
	}

	if lengthHashed > 0 {
		for _, elt := range hashes {
			t.AddNode(elt)
		}
	}

	return t, nil
}

func (node *merkleNode) AddNode(data string) {

	if node.leftChild == nil {
		node.leftChild = &merkleNode{isLeaf: true, hash: toSHA256(data), leftChild: nil, rightChild: nil}

	} else if node.rightChild == nil {
		/*
			If the left part of the current tree is complete
		 */
		if node.leftChild.isCompleteTree() {
			node.rightChild = &merkleNode{isLeaf: false, hash: "", leftChild: nil, rightChild: nil}
			height := node.GetHeight()

			node.rightChild.insertLeft(height-1, data)

			/*
				Otherwise, just insert the node where we find a nil right child from this current part of the tree.
			 */
		} else {
			node.leftChild.AddNode(data)
		}

		/*
			If it's a complete tree, we need to add a new layer at the top of the tree,
			then insert the node at the bottom right of this tree.
		 */
	} else if node.isCompleteTree() {
		root := merkleNode{isLeaf: node.isLeaf, hash: node.hash, leftChild: node.leftChild, rightChild: node.rightChild}
		node.leftChild = &root
		node.rightChild = &merkleNode{isLeaf: false, hash: "", leftChild: nil, rightChild: nil}
		height := node.GetHeight()

		node.rightChild.insertLeft(height-1, data)

		/*
			Otherwise, just insert the node where we find a nil right child from this current part of the tree.
		 */
	} else {
		node.rightChild.AddNode(data)
	}

	// Once the node is inserted, we need to recalculate the current node's hash
	node.hashMe()
}

/*
	Insert the node at a given level of a tree.
 */
func (node *merkleNode) insertLeft(level int, data string) {
	if level == 0 {
		node.isLeaf = true
		node.hash = toSHA256(data)
		return
	}

	node.leftChild = &merkleNode{isLeaf: false, hash: "", leftChild: nil, rightChild: nil}
	node.leftChild.insertLeft(level-1, data)
	node.hashMe()
}

func (node *merkleNode) isCompleteTree() bool {
	i := 0
	height := node.GetHeight()

	for i = 0; i < height && node.rightChild != nil; i++ {
		node = node.rightChild
	}

	return height == i
}

func (node *merkleNode) hashMe() {
	if node.leftChild != nil && node.rightChild != nil {
		node.hash = toSHA256(node.leftChild.hash + node.rightChild.hash)

	} else if node.leftChild != nil && node.rightChild == nil {
		node.hash = node.leftChild.hash
	}
}

func (node merkleNode) GetRoot() string {
	return node.hash
}

func (node *merkleNode) GetHeight() (height int) {
	for height = 0; !node.isLeaf && node.leftChild != nil; height++ {
		node = node.leftChild
	}
	return
}

func (node merkleNode) getNodesByLevel(level int) []string {
	if level <= 0 {
		return []string{node.hash}

	} else {
		if node.rightChild != nil {
			return append(node.leftChild.getNodesByLevel(level-1), node.rightChild.getNodesByLevel(level-1)...)
		}
		return node.leftChild.getNodesByLevel(level - 1)
	}
}

func (node *merkleNode) GetLevel(level int) ([]string, error) {
	height := node.GetHeight()

	if level > height {
		return nil, errors.New("Level too deep")
	}
	return node.getNodesByLevel(level), nil
}

func main() {

}
