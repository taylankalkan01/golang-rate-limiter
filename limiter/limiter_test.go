package limiter

import (
	"testing"
	"time"
)

// TestTokenBucket_TakeTokens tests the TakeTokens method of TokenBucket.
func TestTokenBucket_TakeTokens(t *testing.T) {
	tb := NewTokenBucket(5, 2)

	// Initial state check
	if tb.currentToken != tb.GetCapacity() {
		t.Errorf("Expected initial currentToken to be equal to capacity, got %d", tb.currentToken)
	}

	// Try to take 3 tokens (should succeed)
	if !tb.TakeTokens(3) {
		t.Error("Failed to take 3 tokens")
	}

	// Check currentToken after taking 3 tokens
	if tb.currentToken != tb.GetCapacity()-3 {
		t.Errorf("Expected currentToken to be %d after taking 3 tokens, got %d", tb.GetCapacity()-3, tb.currentToken)
	}

	// Try to take 5 tokens (should fail, not enough tokens)
	if tb.TakeTokens(5) {
		t.Error("Should not be able to take 5 tokens, but succeeded")
	}

	// Check currentToken after failing to take 5 tokens
	if tb.currentToken != tb.GetCapacity()-3 {
		t.Errorf("Expected currentToken to still be %d after failing to take 5 tokens, got %d", tb.GetCapacity()-3, tb.currentToken)
	}

	// Wait for 2 seconds to allow token refill
	time.Sleep(2 * time.Second)

	// Try to take 5 tokens again (should succeed after refill)
	if !tb.TakeTokens(5) {
		t.Error("Failed to take 5 tokens after refill")
	}

	// Check currentToken after taking 5 tokens after refill
	if tb.currentToken != tb.GetCapacity()-5 {
		t.Errorf("Expected currentToken to be %d after taking 5 tokens after refill, got %d", tb.GetCapacity()-5, tb.currentToken)
	}
}
