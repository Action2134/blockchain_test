package main

import (
  "bytes"
  "crypto/sha256"
  "strconv"
  "time"
)

//block 由区块头和交易两个部分组成
//timestamp, PrevBlockHash , Hash 属于头 block header
//timestamp  :当前时间戳
//PrevBlockHash : 前一个块的哈希
//Hash : 当前块的hash
//Data : 区块实际存储的信息

type Block struct{
  Timestamp  int64
  PrevBlockHash  []byte
  Hash           []byte
  Data           []byte
}

//NewBlock  用于产生新块  参数需要 Data 与 PrevBlockHash
//当前块的哈希会基于 Data 和 PrevBlackHash计算得到
func NewBlock(data string , prevBlockHash []byte) *Block{
  block:= &Block{
    Timestamp: time.Now().Unix(),
	PrevBlockHash: prevBlockHash,
	Hash:    []byte{},
	Data:    []byte(data)}
  block.SetHash()
  return block
}

//setHash 设置当前 哈希
func (b *Block) SetHash(){
  timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp},[]byte{})
  hash := sha256.Sum256(headers)
  b.Hash = hash[:]
}
//生成创世块
func NewGenesisBlock() *Block{
  return NewBlock("Genesis Block",[]byte{})
}




