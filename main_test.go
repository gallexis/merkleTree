package main

import (
	"testing"
)

func TestTree_getHashRoot(t *testing.T) {
	hashRootExpected := "62af5c3cb8da3e4f25061e829ebeea5c7513c54949115b1acc225930a90154da"

	tree, err := CreateMerkleTree([]string{"a", "b"})
	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	if tree.GetRoot() != hashRootExpected {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, hashRootExpected)
	}
}

func TestTree_GetLevel(t *testing.T) {
	tree, err := CreateMerkleTree([]string{"a", "b"})
	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	_, err = tree.GetLevel(2)
	if err == nil {
		t.Errorf("Exception should have been raised while getting level 2, on a 1 level tree.")
	}
}

func TestTree_getLevel2(t *testing.T) {
	a := "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"
	b := "3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d"
	c := "2e7d2c03a9507ae265ecf5b5356885a53393a2029d241394997265a1a25aefc6"

	tree, err := CreateMerkleTree([]string{"a", "b", "c"})
	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	nodes, err := tree.GetLevel(2)
	if err != nil {
		t.Errorf("Error while getting level 0.")
		return
	}

	if len(nodes) != 3 {
		t.Errorf("Node must be of length 1")
	}

	if a != nodes[0] {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, nodes[0])
	}

	if b != nodes[1] {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, nodes[0])
	}

	if c != nodes[2] {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, nodes[0])
	}
}

func TestTree_0node(t *testing.T) {
	_, err := CreateMerkleTree([]string{})
	if err == nil {
		t.Errorf("Exception should have been raised while creating tree.")
		return
	}
}

func TestTree_4nodes(t *testing.T) {
	hashRootExpected := "58c89d709329eb37285837b042ab6ff72c7c8f74de0446b091b6a0131c102cfd"
	tree, err := CreateMerkleTree([]string{"a", "b", "c", "d"})

	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	if tree.hash != hashRootExpected {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, hashRootExpected)
	}
}

func TestTree_5nodes(t *testing.T) {
	hashRootExpected := "58c89d709329eb37285837b042ab6ff72c7c8f74de0446b091b6a0131c102cfd"
	tree, err := CreateMerkleTree([]string{"a", "b", "c", "d"})

	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	if tree.hash != hashRootExpected {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, hashRootExpected)
	}
}

func TestTree_9nodes(t *testing.T) {
	hashRootExpected := "95f57bf74ca2317b9056ad66ecd7582e6e470511d960be8bb594e51b8c8a6498"
	tree, err := CreateMerkleTree([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})

	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	if tree.hash != hashRootExpected {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, hashRootExpected)
	}
}
