package memo_test

import (
	"testing"

	memo "gopl.io/ch9/memo5"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
