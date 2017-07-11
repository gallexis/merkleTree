package main

import (
	"fmt"
	"errors"
	"crypto/sha256"
)

type merkleNode struct {
	isLeaf     bool
	hash       string
	leftChild  *merkleNode
	rightChild *merkleNode
}

func CreateMerkleTree(hashes []string) (merkleNode, error) {
	nodes, err := getArrayOfNodes(hashes)
	if err != nil {
		return merkleNode{}, err
	}

	for len(nodes) > 1 {
		nodes = createFathers(nodes)
	}

	t := nodes[0]
	fmt.Println(t)

	return t, nil
}

func getArrayOfNodes(hashes []string) ([]merkleNode, error) {
	if 0 == len(hashes) {
		return nil, errors.New("Empty hashes array")
	}
	nodesArray := []merkleNode{}

	for _, hash := range hashes {
		node := merkleNode{isLeaf: true, hash: "", leftChild: nil, rightChild: nil}
		node.hashMe(hash)
		nodesArray = append(nodesArray, node)
	}
	return nodesArray, nil
}

func createFathers(nodes []merkleNode) []merkleNode {
	nodesArray := []merkleNode{}
	lengthNodes := len(nodes)
	i := 0

	for i = 0; lengthNodes-i > 1; i += 2 {
		lc := nodes[i]
		rc := nodes[i+1]
		node := merkleNode{isLeaf: false, hash: "", leftChild: &lc, rightChild: &rc}
		node.hashMe(lc.hash + rc.hash)
		nodesArray = append(nodesArray, node)
	}

	if 1 == lengthNodes-i {
		lc := nodes[lengthNodes-1]
		node := merkleNode{isLeaf: false, hash: "", leftChild: &lc, rightChild: nil}
		nodesArray = append(nodesArray, node)
	}
	return nodesArray
}

func (node *merkleNode) hashMe(data string) {
	sum := sha256.Sum256([]byte(data))
	node.hash = fmt.Sprintf("%x", sum)
}

func (node merkleNode) GetRoot() (string) {
	return node.hash
}

func (node *merkleNode) GetHeight() (i int) {
	for i = 0; !node.isLeaf; i++ {
		node = node.leftChild
	}
	return
}

func (node *merkleNode) getNodesByLevel(level int) ([]string) {
	if level <= 0 {
		return []string{node.hash}

	} else {
		if node.rightChild != nil {
			return append(node.leftChild.getNodesByLevel(level-1), node.rightChild.getNodesByLevel(level-1) ...)
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
	t, err := CreateMerkleTree([]string{"a","b"})
	if err != nil {
		return
	}
	fmt.Println(t.leftChild)
	fmt.Println(t.rightChild)


	fmt.Println(t.GetRoot())
	fmt.Println(t.GetHeight())
	fmt.Println(t.GetLevel(1))
}
