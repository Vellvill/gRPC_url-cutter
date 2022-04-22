package repository

import (
	"context"
	"fmt"
	"gRPC_cutter/pkg/model"
	"gRPC_cutter/pkg/utils"
	"sync"
	"sync/atomic"
)

type hash struct {
	*sync.RWMutex
	inc     int64
	hashmap map[string]*model.Model
}

func NewHash() (*hash, error) {
	repo := &hash{
		RWMutex: new(sync.RWMutex),
		inc:     1,
		hashmap: make(map[string]*model.Model),
	}
	return repo, nil
}

func (h *hash) AddModel(ctx context.Context, wg *sync.WaitGroup, url string) (*model.Model, error) {
	defer wg.Done()
	h.RLock()
	if j, ok := h.hashmap[url]; ok {
		h.RUnlock()
		return nil, fmt.Errorf("URL aldready exists, short: %s\n", j.Shorturl)
	}
	h.RUnlock()

	m := model.NewModel(int(h.inc), url, utils.Encode())

	atomic.AddInt64(&h.inc, 1)

	h.Lock()
	h.hashmap[m.Shorturl] = m
	h.Unlock()

	return m, nil
}
func (h *hash) GetModel(ctx context.Context, wg *sync.WaitGroup, shortURL string) (string, error) {
	h.RLock()
	defer h.RUnlock()
	defer wg.Done()
	if j, ok := h.hashmap[shortURL]; ok {
		return j.Longurl, nil
	} else {
		return "", fmt.Errorf("Didnt find %s shorturl in hash\n", shortURL)
	}
}
