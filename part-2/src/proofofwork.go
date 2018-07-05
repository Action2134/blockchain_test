package main
import(
  "bytes"
  "crypto/sha256"
  "fmt"
  "math"
  "math/big"
)

//难度值， 这里表示哈希的钱 24 位必须是 0
const targetBits = 24
const maxNonce = math.MaxInt64

//每个块的工作量必须嗷要证明  所有 有个指向Block的指针
// target 是目标， 我们最终找到的哈希必须小于目标
type ProofOfWork struct{
  block *Block
  target *big.Int
}

//target 等于 1  左移 257 -targetBits 位
func NewProofOfWork(b *Block) *ProofOfWork{
  target:=big.NewInt(1)
  target.Lsh(target, uint(256-targetBits))
  pow := &ProofOfWork{b, target}
  return pow
}

//工作量证明用到的数据 有 PrevBlockHash, Data, Timestampe , targetBits, nonce
func (pow *ProofOfWork) prepareData(nonce int)[]byte{
  data:= bytes.Join(
    [][]byte{
	  pow.block.PrevBlockHash,
	  pow.block.Data,
	  IntToHex(pow.block.Timestamp),
	  IntToHex(int64(targetBits)),
	  IntToHex(int64(nonce)),
	},
	[]byte{},
  )
  return data
}
//工作量证明的核心是寻找 有效哈希
func (pow *ProofOfWork)Run()(int, []byte){
  var hasInt big.Int
  var  hash [32]byte
  nonce :=0

  //fmt.Printf("mining the bloick containing \"%s \"\n",pow.block.Data);
  for nonce < maxNonce{
    data := pow.prepareDate(nonce)
	hash = sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.target) == -1{
	  fmt.Printf("%x \n",hash);
	  break
	}else{
	  nonce ++
	}
  }
  return nonce,hash[:]
}



//检验工作量 ， 只要哈希小雨目标就ok
func (pow *ProofOfWork) Validate() bool{
  var hashInt big.Int
  data:=pow.prepareData(pow.block.Nonce)
  hash := sha256.Sum256(data)
  hashInt.SetBytes(hash[:])
  isValid:=hashInt.Cmp(pow.target)==-1
  return isValid
}


