package main

import (
	"fmt"
	"math"
)

type merkleNode struct {
	isLeaf bool
	hash string
	leftChild  *merkleNode
	rightChild *merkleNode
}


func CreateMerkleTree(values []string) /* merkleNode */ {


	length := len(values)
	heightTree := math.Ceil(math.Log2(float64(length)))
	fmt.Println(heightTree)

	t := createTree(int(heightTree), values)
	fmt.Println(t.leftChild)
	fmt.Println(t.leftChild.leftChild)
	fmt.Println(t.leftChild.rightChild)
	fmt.Println(t.rightChild)
	fmt.Println(t.rightChild.leftChild)
	fmt.Println(t.rightChild.rightChild)

}

func  createTree(level int, hashes []string) merkleNode{
	lenghtHashes := len(hashes)

	if level == 0{
		if lenghtHashes == 1{
			return createLeaf(hashes[0])
		}else {
			return createLeaf("emptyleaf")
		}

	} else if level > 0{
		level -= 1
		m := lenghtHashes % 2

		if m == 1{
			m = (lenghtHashes / 2) + 1
		}else{
			m = (lenghtHashes / 2)
		}
		t := createNode("")


		lc := createTree(level,hashes[:m])
		rc := createTree(level,hashes[m:])
		t.leftChild  = &lc
		t.rightChild = &rc
		//t.hash = hash256(t.leftChild.hash + t.rightChild.hash)
		return t

	} else {
		fmt.Println("erreur")
		return merkleNode{}
	}
}


func  createLeaf(hsh string) merkleNode{
	return merkleNode{isLeaf:true, hash: hsh, leftChild: nil, rightChild: nil}
}

func  createNode(hsh string) merkleNode{
	return merkleNode{isLeaf:false, hash: hsh, leftChild: nil, rightChild: nil}
}

func  createParent(leftChild *merkleNode, rightChild *merkleNode) merkleNode{
	n := createNode("")
	n.leftChild  = leftChild
	n.rightChild = rightChild
	// n.hash = hash256(n.leftChild.hash + n.rightChild.hash)
	return n
}

func (node *merkleNode) addData(str string){

}

func (node merkleNode) root() (string){
	return node.hash
}

func (node *merkleNode) height(){

}

func (node *merkleNode) level(index int){

}

func main(){
	CreateMerkleTree([]string{"aa","bb","cc","dd","ee"})
}
