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


func CreateMerkleTree(values []string)  merkleNode {


	length := len(values)
	heightTree := math.Ceil(math.Log2(float64(length)))
	fmt.Println(heightTree)

	t := createTree(int(heightTree), values)

	fmt.Println(t)
	fmt.Println(t.leftChild)
	fmt.Println(t.leftChild.leftChild)
	fmt.Println(t.leftChild.rightChild)
	fmt.Println(t.rightChild)
	fmt.Println(t.rightChild.leftChild)
	fmt.Println(t.rightChild.rightChild)

	return t
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
		node := createNode("")
		h := lenghtHashes % 2

		if h == 1{
			h = (lenghtHashes / 2) + 1
		}else{
			h = (lenghtHashes / 2)
		}

		lc := createTree(level,hashes[:h])
		rc := createTree(level,hashes[h:])
		node.leftChild  = &lc
		node.rightChild = &rc
		//node.hash = hash256(node.leftChild.hash + node.rightChild.hash)
		return node

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

func (node *merkleNode) addData(str string){

}

func (node merkleNode) root() (string){
	return node.hash
}

func (node *merkleNode) height() (i int){
	for i=1; !node.leftChild.isLeaf ;i++{
		node=node.leftChild
	}
	return
}

func (node *merkleNode) level(index int){

}

func main(){
	t := CreateMerkleTree([]string{"aa","bb","cc","ccd"})
	fmt.Println(t.height())
}
