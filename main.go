package main

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// Blockchain is a simple blockchain with a slice of blocks.
type Blockchain struct {
	Blocks []Block
}

// NewBlock creates a new block and adds it to the blockchain.
func (bc *Blockchain) NewBlock(transaction string, nonce int) {
	// Get the previous block's information
	previousBlock := bc.GetLatestBlock()
	previousHash := previousBlock.CurrentHash

	// Calculate the current block's hash
	currentHash := CreateHash(transaction, nonce, previousHash)

	// Create a new block
	block := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  currentHash,
	}

	// Append the new block to the blockchain
	bc.Blocks = append(bc.Blocks, block)
}

// DisplayBlocks prints all the blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	// Iterate through each block and display its details
	for i, block := range bc.Blocks {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println("------------------------------------------------------------")
	}
}

// GetLatestBlock returns the latest block in the blockchain.
func (bc *Blockchain) GetLatestBlock() Block {
	// Check if there are any blocks in the blockchain
	if len(bc.Blocks) == 0 {
		// Return an empty block for the genesis block
		return Block{}
	}

	// Return the last block in the blockchain
	return bc.Blocks[len(bc.Blocks)-1]
}

// ChangeBlock changes the transaction of a given block.
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
	// Check if the provided index is valid
	if index >= 0 && index < len(bc.Blocks) {
		// Update the block's transaction
		bc.Blocks[index].Transaction = newTransaction

		// Recalculate the block's current hash
		bc.Blocks[index].CurrentHash = CreateHash(newTransaction, bc.Blocks[index].Nonce, bc.Blocks[index].PreviousHash)
	}
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	// Iterate through each block starting from the second block
	for i := 1; i < len(bc.Blocks); i++ {
		// Check if the previous block's hash matches the current block's previous hash
		if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].CurrentHash {
			return false
		}
	}
	return true
}

// CreateHash calculates the hash of a block.
func CreateHash(transaction string, nonce int, previousHash string) string {
	// Concatenate the block data
	data := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)

	// Calculate the SHA-256 hash
	hash := sha256.Sum256([]byte(data))

	// Convert the hash to a hexadecimal string
	return fmt.Sprintf("%x", hash)
}

func main() {
	// Initialize the blockchain variable
	blockchain := &Blockchain{
		Blocks: []Block{}, // Initialize the Blocks slice
	}

	// Add some blocks to the blockchain
	blockchain.NewBlock("Alice to Bob", 123)
	blockchain.NewBlock("Eeman to Qadeer", 456)
	blockchain.NewBlock("Baheej to Nimra", 789)
	blockchain.NewBlock("Amna to Kiran", 345)
	blockchain.NewBlock("Shifa to Ayesha", 234)
	blockchain.NewBlock("Mama to Baba", 986)
	blockchain.NewBlock("Zain to Ahtisham", 567)

	// Display all blocks
	blockchain.DisplayBlocks()

	// Change the transaction of the second block
	//blockchain.ChangeBlock(1, "New transaction")

	// Verify the blockchain
	if blockchain.VerifyChain() {
		fmt.Println("---------Blockchain is valid--------------")
	} else {
		fmt.Println("-----------Blockchain is not valid------------")
	}
}
