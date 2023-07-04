package memstore

import (
	"ethscanner/parser"
	"sync"
)

func init() {
	sd := &inMemoryStoreDelegate{sync.RWMutex{}, make(map[string][]parser.Transaction)}
	parser.SetStoreDelegate(sd)
}

type inMemoryStoreDelegate struct {
	lock sync.RWMutex
	data map[string][]parser.Transaction
}

func (s *inMemoryStoreDelegate) Insert(address string, transaction parser.Transaction) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[address] = append(s.data[address], transaction)
}

func (s *inMemoryStoreDelegate) Get(address string) []parser.Transaction {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.data[address]
}
