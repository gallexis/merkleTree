package main

import (
	"fmt"
	"math"
	"crypto/sha256"
)

type merkleNode struct {
	isLeaf bool
	hash string
	leftChild  *merkleNode
	rightChild *merkleNode
}

func CreateMerkleTree(values []string)  merkleNode {
	length := len(values)
	heightTree := math.Ceil(math.Log2(float64(length)))

	return *createTree(int(heightTree), values)
}


/*
	1.				[1,2,3,4,5]


	2.				     []
			[1,2,3,4]		   [5]


	3.						        []
					[]		                      []
			[1,2]        [3,4]                   [5]


	4.							[]
				   []		                 []
			[]	          []                 []
		[1]	  [2]     [3]   [4]             [5]
 */
func createTree(level int, hashes []string) *merkleNode{
	lengthHashes := len(hashes)
	halfLengthHashes := 0

	if level == 0{
		if lengthHashes == 1{
			leaf := merkleNode{isLeaf:true, hash: "", leftChild: nil, rightChild: nil}
			leaf.hashMe(hashes[0])
			return &leaf
		}else {
			leaf :=  merkleNode{isLeaf:true, hash: "", leftChild: nil, rightChild: nil}
			return &leaf
		}

	} else if level > 0{
		level -= 1
		isNumberOfHashesEven := (lengthHashes % 2) == 0
		node := merkleNode{isLeaf:false, hash: "", leftChild: nil, rightChild: nil}

		if isNumberOfHashesEven{
			halfLengthHashes = (lengthHashes / 2)
		}else{
			halfLengthHashes = (lengthHashes / 2) + 1
		}

		node.leftChild  = createTree(level,hashes[:halfLengthHashes])
		node.rightChild = createTree(level,hashes[halfLengthHashes:])

		if node.rightChild.hash != ""{
			node.hashMe(node.leftChild.hash + node.rightChild.hash)
		}else{
			node.hashMe(node.leftChild.hash)
		}

		return &node

	} else {
		fmt.Println("erreur")
		return &merkleNode{}
	}
}

func (node *merkleNode) hashMe(data string){
	sum := sha256.Sum256([]byte(data))
	node.hash = fmt.Sprintf("%x", sum)
}

func (node merkleNode) GetRoot() (string){
	return node.hash
}

func (node *merkleNode) GetHeight() (i int){
	for i=0; !node.isLeaf ;i++{
		node=node.leftChild
	}
	return
}

func (node *merkleNode) GetLevel(level int) []string{

	if(level <= 0){
		return []string{node.hash}
	} else{
		return append( node.leftChild.GetLevel(level-1), node.rightChild.GetLevel(level-1) ... )
	}

}


func main(){
	t := CreateMerkleTree([]string{"a", "b","c","d","e"})

	fmt.Println(t.GetRoot())
	fmt.Println(t.GetHeight())
	fmt.Println(t.GetLevel(3))
}
