package blockchain

type block struct {
	Header blockHeader
	Data   []transaction
}

type blockHeader struct {
	PrevBlockHash []byte
	Signature     []byte
}
