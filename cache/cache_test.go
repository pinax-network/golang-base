package cache

import (
	"context"
	"errors"
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func init() {
	_ = log.InitializeGlobalLogger(true)
}

func TestCache_Get(t *testing.T) {

	updateCnt := 0

	testCache := New(5*time.Minute, func(ctx context.Context, key string) (EntryUpdate[string], error) {
		updateCnt++
		return EntryUpdate[string]{
			Value: "test_result",
			Error: nil,
		}, nil
	})

	// test we get a valid response
	res, err := testCache.Get(context.Background(), "test_key")
	assert.NoError(t, err)
	assert.Equal(t, "test_result", res)
	assert.Equal(t, 1, updateCnt)

	// test the same key to ensure it's cached and UpdateFunc isn't called again
	res, err = testCache.Get(context.Background(), "test_key")
	assert.NoError(t, err)
	assert.Equal(t, "test_result", res)
	assert.Equal(t, 1, updateCnt)
}

func TestCache_GetParallel(t *testing.T) {

	updateCnt := 0

	testCache := New(5*time.Minute, func(ctx context.Context, key string) (EntryUpdate[string], error) {
		updateCnt++
		time.Sleep(1 * time.Second)
		return EntryUpdate[string]{
			Value: "test_result",
			Error: nil,
		}, nil
	})

	// run 10 requests in goroutines, all should get a valid response, cache should only be updated once
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			res, err := testCache.Get(context.Background(), "test_key")
			assert.NoError(t, err)
			assert.Equal(t, "test_result", res)
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, 1, updateCnt)
}

func TestCache_GetExpires(t *testing.T) {

	updateCnt := 0

	testCache := New(1*time.Minute, func(ctx context.Context, key string) (EntryUpdate[string], error) {
		updateCnt++
		return EntryUpdate[string]{
			Value: "test_result",
			Error: nil,
		}, nil
	})
	testCache.entries["test_key"] = &Entry[string]{
		Value:     "test_result",
		Error:     nil,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	// test we get a cached response
	res, err := testCache.Get(context.Background(), "test_key")
	assert.NoError(t, err)
	assert.Equal(t, "test_result", res)
	assert.Equal(t, 0, updateCnt)

	// test we update the expired entry
	testCache.entries["test_key"].ExpiresAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	res, err = testCache.Get(context.Background(), "test_key")
	assert.NoError(t, err)
	assert.Equal(t, "test_result", res)
	assert.Equal(t, 1, updateCnt)
}

func TestCache_GetError(t *testing.T) {

	updateCnt := 0
	testError := errors.New("test_error")

	testCache := New(5*time.Minute, func(ctx context.Context, key string) (EntryUpdate[string], error) {
		updateCnt++
		return EntryUpdate[string]{}, testError
	})

	// we should get a test error
	_, err := testCache.Get(context.Background(), "test_key")
	assert.Equal(t, testError, err)
	assert.Equal(t, 1, updateCnt)

	// no entry should be cached, so calling Get() again should trigger the UpdateFunc
	_, err = testCache.Get(context.Background(), "test_key")
	assert.Equal(t, testError, err)
	assert.Equal(t, 2, updateCnt)
}

func TestCache_GetEmbeddedError(t *testing.T) {

	updateCnt := 0
	testError := errors.New("test_error")

	testCache := New(5*time.Minute, func(ctx context.Context, key string) (EntryUpdate[string], error) {
		updateCnt++
		return EntryUpdate[string]{Value: "", Error: testError}, nil
	})

	// we should get a test error
	_, err := testCache.Get(context.Background(), "test_key")
	assert.Equal(t, testError, err)
	assert.Equal(t, 1, updateCnt)

	// as the error is embedded, it should be cached
	_, err = testCache.Get(context.Background(), "test_key")
	assert.Equal(t, testError, err)
	assert.Equal(t, 1, updateCnt)
}

func TestCache_GetParallelErrors(t *testing.T) {

	updateCnt := 0
	testError := errors.New("test_error")

	testCache := New(5*time.Minute, func(ctx context.Context, key string) (EntryUpdate[string], error) {
		updateCnt++
		time.Sleep(1 * time.Second)

		// we error the first 3 calls and afterwards return a valid response
		if updateCnt <= 2 {
			return EntryUpdate[string]{}, testError
		} else {
			return EntryUpdate[string]{
				Value: "test_result",
				Error: nil,
			}, nil
		}
	})

	// run 10 requests in goroutines, all should get a valid response, cache should only be updated once
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			_, _ = testCache.Get(context.Background(), "test_key")
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, 3, updateCnt)
}
