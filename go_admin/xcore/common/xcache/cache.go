package xcache

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/lib/v4/cache"
	bigcacheStore "github.com/eko/gocache/store/bigcache/v4"
	"time"
)

type CacheStore[T interface{}] struct {
	context     context.Context
	prefix      string
	cacheManger *cache.Cache[[]byte]
}

func NewCacheStore[T interface{}](invalidDuration time.Duration, prefixKey string) *CacheStore[T] {
	var ctx = context.Background()
	bigCacheClient, _ := bigcache.New(ctx, bigcache.DefaultConfig(invalidDuration))
	bigCacheStore := bigcacheStore.NewBigcache(bigCacheClient)
	cacheManager := cache.New[[]byte](bigCacheStore)
	return &CacheStore[T]{context.Background(), prefixKey, cacheManager}
}

func (receiver CacheStore[T]) Set(key string, value T) (ok bool) {
	var bufferByte = new(bytes.Buffer)
	newEncoder := gob.NewEncoder(bufferByte)
	err := newEncoder.Encode(value)
	err = receiver.cacheManger.Set(receiver.context, receiver.prefix+key, bufferByte.Bytes())
	return err == nil
}
func (receiver CacheStore[T]) Get(key string) (result T) {
	byteInCatch, _ := receiver.cacheManger.Get(receiver.context, receiver.prefix+key)
	decoder := gob.NewDecoder(bytes.NewReader(byteInCatch))
	_ = decoder.Decode(&result)
	return
}

func (receiver CacheStore[T]) Delete(key string) (err error) {
	err = receiver.cacheManger.Delete(receiver.context, receiver.prefix+key)
	return
}

func (receiver CacheStore[T]) Clear() (err error) {
	err = receiver.cacheManger.Clear(receiver.context)
	return
}
func (receiver CacheStore[T]) Exist(key string) (exit bool) {
	return true
}
