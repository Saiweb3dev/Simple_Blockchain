package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"log"

	"Simple_Blockchain/models"
)

type Block struct {
	Index    int
	Data     models.TransactionCheckout
	Timestamp string
	Hash     string
	PrevHash string
}

type Blockchain struct {
	Blocks []*Block
}

func GetBlockChain(blockchain *Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			jbytes, err := json.MarshalIndent(&blockchain.Blocks, "", "  ")
			if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("Could not marshal the blockchain : %v", err)
					w.Write([]byte("Could not marshal the blockchain"))
					return
			}
			io.WriteString(w, string(jbytes))
	}
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, models.TransactionCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func CreateBlock(prevBlock *Block, data models.TransactionCheckout) *Block {
	block := &Block{}
	block.Index = prevBlock.Index + 1
	block.Timestamp = time.Now().String()
	block.Data = data
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Index+1 != block.Index {
		return false
	}
	return true
}

func (bc *Block) validateHash(hash string) bool {
	bc.generateHash()
	return bc.Hash == hash
}

func (bc *Block) generateHash() {
	bytes, _ := json.Marshal(bc.Data)
	data := string(bc.Index) + bc.Timestamp + string(bytes) + bc.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	bc.Hash = hex.EncodeToString(hash.Sum(nil))
}


func (bc *Blockchain) AddBlock(data models.TransactionCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := CreateBlock(prevBlock,data)

	if validBlock(block,prevBlock) {
		bc.Blocks = append(bc.Blocks,block)
	}
}