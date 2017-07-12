package main

import (
	"testing"
)

func TestTree_1node(t *testing.T) {
	hashRootExpected := "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"
	tree, err := CreateMerkleTree([]string{"a"})
	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	if tree.hash != hashRootExpected {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, hashRootExpected)
	}
}

func TestTree_2nodes(t *testing.T) {
	hashRootExpected := "62af5c3cb8da3e4f25061e829ebeea5c7513c54949115b1acc225930a90154da"
	tree, err := CreateMerkleTree([]string{"a", "b"})
	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	if tree.hash != hashRootExpected {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, hashRootExpected)
	}
}

func TestTree_3nodes(t *testing.T) {
	hashRootExpected := "d71dc32fa2cd95be60b32dbb3e63009fa8064407ee19f457c92a09a5ff841a8a"
	tree, err := CreateMerkleTree([]string{"a", "b", "c"})

	if err != nil {
		t.Errorf("Exception raised while creating tree.")
		return
	}

	if tree.hash != hashRootExpected {
		t.Errorf("Hash was incorrect, got: %d, want: %d.", tree.hash, hashRootExpected)
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
