package bc

// BlockHeaderEntry contains the header information for a blockchain
// block. It satisfies the Entry interface.
type BlockHeaderEntry struct {
	body struct {
		Version              uint64
		Height               uint64
		PreviousBlockID      Hash
		TimestampMS          uint64
		TransactionsRoot     Hash
		AssetsRoot           Hash
		NextConsensusProgram []byte
		ExtHash              Hash
	}
	witness struct {
		Arguments [][]byte
	}
}

func (BlockHeaderEntry) Type() string          { return "blockheader" }
func (bh *BlockHeaderEntry) Body() interface{} { return bh.body }

func (BlockHeaderEntry) Ordinal() int { return -1 }

func (bh *BlockHeaderEntry) Version() uint64 {
	return bh.body.Version
}

func (bh *BlockHeaderEntry) Height() uint64 {
	return bh.body.Height
}

func (bh *BlockHeaderEntry) PreviousBlockID() Hash {
	return bh.body.PreviousBlockID
}

func (bh *BlockHeaderEntry) TimestampMS() uint64 {
	return bh.body.TimestampMS
}

func (bh *BlockHeaderEntry) TransactionsRoot() Hash {
	return bh.body.TransactionsRoot
}

func (bh *BlockHeaderEntry) AssetsRoot() Hash {
	return bh.body.AssetsRoot
}

func (bh *BlockHeaderEntry) NextConsensusProgram() []byte {
	return bh.body.NextConsensusProgram
}

func (bh *BlockHeaderEntry) Arguments() [][]byte {
	return bh.witness.Arguments
}

func (bh *BlockHeaderEntry) SetArguments(args [][]byte) {
	bh.witness.Arguments = args
}

// NewBlockHeaderEntry creates a new BlockHeaderEntry and populates
// its body.
func NewBlockHeaderEntry(version, height uint64, previousBlockID Hash, timestampMS uint64, transactionsRoot, assetsRoot Hash, nextConsensusProgram []byte) *BlockHeaderEntry {
	bh := new(BlockHeaderEntry)
	bh.body.Version = version
	bh.body.Height = height
	bh.body.PreviousBlockID = previousBlockID
	bh.body.TimestampMS = timestampMS
	bh.body.TransactionsRoot = transactionsRoot
	bh.body.AssetsRoot = assetsRoot
	bh.body.NextConsensusProgram = nextConsensusProgram
	return bh
}

func (bh *BlockHeaderEntry) CheckValid(prev *BlockHeaderEntry, txs []*TxEntries) error {
	if prev == nil {
		if bh.body.Height != 1 {
			// xxx error
		}
	} else {
		if bh.body.Version < prev.body.Version {
			// xxx error
		}

		if bh.body.Height != prev.body.Height + 1 {
			// xxx error
		}

		// xxx check EntryID(prev) == bh.body.PreviousBlockID

		if bh.body.TimestampMS <= prev.body.TimestampMS {
			// xxx error
		}
	}

	// xxx eval NextConsensusProgram

	for i, tx := range txs {
		err := tx.CheckValid(bh.body.TimestampMS, bh.body.Version)
		if err != nil {
			return errors.Wrapf(err, "checking validity of transaction %d of %d", i, len(txs))
		}
	}

	// xxx check bh.body.TransactionsRoot == computeMerkleRoot(txs)

	if bh.body.Version == 1 {
		if (bh.body.ExtHash != bh.Hash{}) {
			// xxx error
		}
	}

	return nil
}
