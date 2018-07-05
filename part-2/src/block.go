package main

import(
  "time"
)

//nonce 在对工作量验证世用
type Block struct{
  Timestamp  int64
  PrevBlockHash []byte
  Hash          []byte
  Data          []byte
  Nonce         int
}

//创建新块石 需要运行工作量证明找到有效的哈希
func NewBlock(data string, prevBlockHash []byte) *Block{
  block := &Block{
    Timestamp:     time.Now().Unix(),
	PrevBlockHash: prevBlockHash,
	Hash:          []byte{},
	Data:          []byte(data),
	Nonce:         0}
  pow := NewProofOfWork(block)
  nonce, hash := pow.Run()
  block.Hash = hash[:]
  block.Nonce = nonce
  
  return block
}

func NewGenesisBlock() *Block{
  return NewBlock("Genesis Block", []byte{})
}
