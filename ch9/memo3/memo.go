// Package memo provides a concurrency unsafe
// memoisation of a function of type Func.
package memo

import "sync"

// A Memo caches the result of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex // Guards cache
	cache map[string]result
}

// Func is the type of the function to memoise.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: Not concurrency safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
