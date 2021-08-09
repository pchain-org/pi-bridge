package merkle

import (
	"errors"
	"io"
	"os"

	"github.com/pchain-org/pi-bridge/tools/types"
)

// HashStore is an interface for persist hash
type HashStore interface {
	Append(hash []types.Uint256) error
	Flush() error
	Close()
	GetHash(pos uint32) (types.Uint256, error)
}

type fileHashStore struct {
	file_name string
	file      *os.File
}

// NewFileHashStore returns a HashStore implement in file
func NewFileHashStore(name string, tree_size uint32) (HashStore, error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	store := &fileHashStore{
		file_name: name,
		file:      f,
	}

	err = store.checkConsistence(tree_size)
	if err != nil {
		return nil, err
	}

	num_hashes := getStoredHashNum(tree_size)
	size := int64(num_hashes) * int64(types.UINT256_SIZE)

	_, err = store.file.Seek(size, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func getStoredHashNum(tree_size uint32) int64 {
	subtreesize := getSubTreeSize(tree_size)
	sum := int64(0)
	for _, v := range subtreesize {
		sum += int64(v)
	}

	return sum
}

func (self *fileHashStore) checkConsistence(tree_size uint32) error {
	num_hashes := getStoredHashNum(tree_size)

	stat, err := self.file.Stat()
	if err != nil {
		return err
	} else if stat.Size() < int64(num_hashes)*int64(types.UINT256_SIZE) {
		return errors.New("stored hashes are less than expected")
	}

	return nil
}

func (self *fileHashStore) Append(hash []types.Uint256) error {
	if self == nil {
		return nil
	}
	buf := make([]byte, 0, len(hash)*types.UINT256_SIZE)
	for _, h := range hash {
		buf = append(buf, h[:]...)
	}
	_, err := self.file.Write(buf)
	return err
}

func (self *fileHashStore) Flush() error {
	if self == nil {
		return nil
	}
	return self.file.Sync()
}

func (self *fileHashStore) Close() {
	if self == nil {
		return
	}
	self.file.Close()
}

func (self *fileHashStore) GetHash(pos uint32) (types.Uint256, error) {
	if self == nil {
		return EMPTY_HASH, errors.New("FileHashstore is nil")
	}
	hash := EMPTY_HASH
	_, err := self.file.ReadAt(hash[:], int64(pos)*int64(types.UINT256_SIZE))
	if err != nil {
		return EMPTY_HASH, err
	}

	return hash, nil
}

type memHashStore struct {
	hashes []types.Uint256
}

// NewMemHashStore returns a HashStore implement in memory
func NewMemHashStore() HashStore {
	return &memHashStore{}
}

func (self *memHashStore) Append(hash []types.Uint256) error {
	self.hashes = append(self.hashes, hash...)
	return nil
}

func (self *memHashStore) GetHash(pos uint32) (types.Uint256, error) {
	return self.hashes[pos], nil
}

func (self *memHashStore) Flush() error {
	return nil
}

func (self *memHashStore) Close() {}
