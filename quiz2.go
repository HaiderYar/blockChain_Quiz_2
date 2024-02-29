package main

import "fmt"

// This structure will be used for managing blocks in a blockchain
type Block_Manager struct {
	blocks []string
}

func (bm *Block_Manager) displayAllBlocks() {
	fmt.Println("Blocks in the blockchain: ")
	for _, block := range bm.blocks {
		fmt.Println(block)
	}
}

func (bm *Block_Manager) newBlock(bData string) {
	bm.blocks = append(bm.blocks, bData)
}

func (bm *Block_Manager) ModifyBlock(bIndex int, newData string) {
	if bIndex >= 0 && bIndex < len(bm.blocks) {
		bm.blocks[bIndex] = newData
	} else {
		fmt.Println("The block index is out of range")
	}
}

func main() {
	// Making a manager obj through which we are going to manage the blockchain
	manage := Block_Manager{}

	// The blocks which we are going to add
	block1 := "Block 1"
	block2 := "Block 2"
	block3 := "Block 3"

	// Adding new blocks to blockchain
	manage.newBlock(block1)
	manage.newBlock(block2)
	manage.newBlock(block3)

	// Displaying all the blocks that we added
	manage.displayAllBlocks()

	// Modifying a block to change its data
	manage.ModifyBlock(1, "Modified Block 2")

	// Again displaying all the blocks
	manage.displayAllBlocks()

}
