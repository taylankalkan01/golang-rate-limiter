// Package limiter provides a rate limiter.
package limiter

import (
	"sync"
	"time"
)

// TokenBucket struct represents the token bucket structure.
type TokenBucket struct {
	capacity     int        // Maximum tokens the bucket can hold
	currentToken int        // Current number of tokens in the bucket
	rate         int        // Rate at which tokens are added to the bucket per second
	last         time.Time  // Timestamp of the last token addition
	mu           sync.Mutex // Mutex for synchronization
	refillAmount int        // Number of tokens added per second
}

// NewTokenBucket creates a new TokenBucket.
func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		currentToken: capacity,
		rate:         rate,
		last:         time.Now(),
		refillAmount: rate,
	}
}

// TakeTokens checks if tokens are available and consumes them if so.
func (tb *TokenBucket) TakeTokens(tokens int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.refillTokens()
	if tb.currentToken >= tokens {
		tb.currentToken -= tokens
		return true
	}
	return false
}

// refillTokens refills the token bucket based on the time passed since the last refill.
func (tb *TokenBucket) refillTokens() {
	now := time.Now()
	diff := now.Sub(tb.last)
	tb.last = now

	// Calculate the number of tokens that should be added based on the time elapsed
	addedTokens := int(diff.Seconds()) * tb.refillAmount

	// Refill the bucket up to its capacity
	tb.currentToken += addedTokens
	if tb.currentToken > tb.capacity {
		tb.currentToken = tb.capacity
	}
}

func (tb *TokenBucket) GetCapacity() int {
	return tb.capacity
}
func (tb *TokenBucket) GetRefillRate() int {
	return tb.rate
}
