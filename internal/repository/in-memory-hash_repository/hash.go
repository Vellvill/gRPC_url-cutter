package in_memory_hash_repository

import (
	"context"
	"fmt"
	"gRPC_cutter/internal/model"
	"gRPC_cutter/internal/utils"
	"sync"
	"sync/atomic"
)

type hash struct {
	*sync.RWMutex
	inc          int64
	hashmapShort map[string]*model.Model
	hashmapLong  map[string]*model.Model
}

func NewHash() (*hash, error) {
	repo := &hash{
		RWMutex:      new(sync.RWMutex),
		inc:          1,
		hashmapShort: make(map[string]*model.Model),
		hashmapLong:  make(map[string]*model.Model),
	}
	return repo, nil
}

func (h *hash) AddModel(ctx context.Context, url string) (*model.Model, error) {

	m, err := model.NewModel(int(h.inc), url, utils.Encode())
	if err != nil {
		return nil, err
	}
	atomic.AddInt64(&h.inc, 1)

	h.RLock()

	if j, ok := h.hashmapLong[m.Longurl]; ok {
		h.RUnlock()
		return j, nil
	} else {
		h.RUnlock()
		h.Lock()
		h.hashmapLong[m.Longurl] = m
	}

	h.hashmapShort[m.Shorturl] = m
	h.Unlock()

	return m, nil
}

func (h *hash) GetModel(ctx context.Context, shortURL string) (string, error) {
	h.RLock()
	defer h.RUnlock()
	if j, ok := h.hashmapShort[shortURL]; ok {
		return j.Longurl, nil
	} else {
		return "", fmt.Errorf("Didnt find %s shorturl in hash\n", shortURL)
	}
}
