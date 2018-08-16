package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
)

// Set the difficulty parameter
var difficulty = "00000"
// Enable or disable logging of mining
var logMining = true

func main() {
	// Create a new blockchain
	chain := newChain()

	// Create a genesis block with some data and mine it
	genesis := newGenesis(" -> Daniel: 25")
	genesis.mine(difficulty)
	chain.addGenesis(genesis)

	// Append new blocks with transaction data an mine them
	b1 := newBlock(genesis, "Daniel -> John: 1.5")
	b1.mine(difficulty)
	chain.addBlock(b1)

	b2 := newBlock(b1, "Daniel -> Alice: 2, John -> David 1")
	b2.mine(difficulty)
	chain.addBlock(b2)

	b3 := newBlock(b2, "David -> John .5")
	b3.mine(difficulty)
	chain.addBlock(b3)

	b4 := newBlock(b3, "Some data")
	b4.mine(difficulty)
	chain.addBlock(b4)

	// Print the final result with good visuals
	for _, b := range chain.blocks {
		b.print()
	}
}

// Block hash is sha256(nonce + data + prev)
type block struct {
	number int
	prev string
	hash string
	data string
	nonce *big.Int
}

func newBlock(prev *block, data string ) *block {
	return &block{
		nonce: new(big.Int),
		data: data,
		prev: prev.hash,
	}
}

func newGenesis(data string) *block {
	return &block{
		nonce: new(big.Int),
		data: data,
	}
}

func (b *block) calculateHash() string {
	sum := sha256.Sum256([]byte(b.nonce.String() + b.data + b.prev))
	return fmt.Sprintf("%x", sum)
}

func (b *block) upNonce() {
	b.nonce.Add(b.nonce, new(big.Int).SetInt64(1))
}

func (b *block) mine(diff string) {
	for {
		sum := sha256.Sum256([]byte(b.nonce.String() + b.data + b.prev))
		b.hash = fmt.Sprintf("%x", sum)
		if logMining {
			fmt.Printf("%s | %s | %d \n", b.hash, b.data, b.nonce) // Log the mining process
		}

		if b.hash[0:len(diff)] == diff {
			break
		}

		b.upNonce()
	}
}

func (b *block) print() {
	data := [][]string{
		{"Hash", b.hash},
		{"Previous block", b.prev},
		{"Nonce", fmt.Sprintf("%d", b.nonce)},
		{"Data", b.data},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", fmt.Sprintf("Block #%d", b.number)})

	for _, v := range data {
		table.Append(v)
	}

	table.Render() // Send output
}

type chain struct {
	blocks []*block
}

func newChain() *chain {
	return &chain{
		blocks: make([]*block, 0),
	}
}

func (c *chain) lastBlock() *block{
	return c.blocks[len(c.blocks) - 1]
}

func (c *chain) addGenesis(genesis *block) {
	c.blocks = append(c.blocks, genesis)
}

func (c *chain) addBlock(newBlock *block) error {
	last := c.blocks[len(c.blocks) - 1]

	if newBlock.prev == last.hash && newBlock.hash[0:len(difficulty)] == difficulty {
		newBlock.number = last.number + 1
		c.blocks = append(c.blocks, newBlock)
		return nil
	}

	fmt.Sprintf( "Not valid %s", newBlock)
	return errors.New("Block not valid!")
}
