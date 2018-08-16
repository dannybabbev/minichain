# Minichain: A minimalistic blockchain for teaching purposes

If you want to teach a technical audience about the very basics of a blockchain: showcasing minichain is a great start

## Usage

Run with:
```
go run main.go
```

You append blocks like:
```
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
```

The resulting blockchain will be:
```
+----------------+------------------------------------------------------------------+
|      NAME      |                             BLOCK #0                             |
+----------------+------------------------------------------------------------------+
| Hash           | 0000099c52e91e15209b41b101586cf21e4e0a28ba08468869584da3771c4498 |
| Previous block |                                                                  |
| Nonce          |                                                           210485 |
| Data           |  -> Daniel: 25                                                   |
+----------------+------------------------------------------------------------------+
+----------------+------------------------------------------------------------------+
|      NAME      |                             BLOCK #1                             |
+----------------+------------------------------------------------------------------+
| Hash           | 0000002b4448df516ce79f8a1715cb36248dc3878ed222da8b12985516667a15 |
| Previous block | 0000099c52e91e15209b41b101586cf21e4e0a28ba08468869584da3771c4498 |
| Nonce          |                                                           602786 |
| Data           | Daniel -> John: 1.5                                              |
+----------------+------------------------------------------------------------------+
+----------------+------------------------------------------------------------------+
|      NAME      |                             BLOCK #2                             |
+----------------+------------------------------------------------------------------+
| Hash           | 00000b7fa45d2192920e0640f6dd80f8f3be85333c1d1ae4aa66d218f389f636 |
| Previous block | 0000002b4448df516ce79f8a1715cb36248dc3878ed222da8b12985516667a15 |
| Nonce          |                                                          1809075 |
| Data           | Daniel -> Alice: 2, John ->                                      |
|                | David 1                                                          |
+----------------+------------------------------------------------------------------+
+----------------+------------------------------------------------------------------+
|      NAME      |                             BLOCK #3                             |
+----------------+------------------------------------------------------------------+
| Hash           | 000009269347fa33dc0e22fc51e2c3116ed5cb58c07dce9eb4e180fe05bc890f |
| Previous block | 00000b7fa45d2192920e0640f6dd80f8f3be85333c1d1ae4aa66d218f389f636 |
| Nonce          |                                                            51538 |
| Data           | David -> John .5                                                 |
+----------------+------------------------------------------------------------------+
+----------------+------------------------------------------------------------------+
|      NAME      |                             BLOCK #4                             |
+----------------+------------------------------------------------------------------+
| Hash           | 000005d2c5bd4c126f9f671a90e89827b8fa3147b93c64d0839c7a0b1f25ed9a |
| Previous block | 000009269347fa33dc0e22fc51e2c3116ed5cb58c07dce9eb4e180fe05bc890f |
| Nonce          |                                                          3519468 |
| Data           | Some data                                                        |
+----------------+------------------------------------------------------------------+
```

You can adjust the difficulty and logging:
```
// Set the difficulty parameter
var difficulty = "00000"
// Enable or disable logging of mining
var logMining = true
```