package main

// blockchain 是一个block 指针数组
type BlockChain struct{
  blocks [] *Block
}

// NewBlockChain 创建一个创世块的链
func NewBlockChain() * BlockChain{
  return &BlockChain{[]*Block{NewGenesisBlock()}}

}

//Add Block 
// data 在实际中就是交易
func (bc *BlockChain) AddBlock(data string){
  prevBlock := bc.blocks[len(bc.blocks)-1]
  newBlock := NewBlock(data,prevBlock.Hash)
  bc.blocks = append(bc.blocks, newBlock)
}



