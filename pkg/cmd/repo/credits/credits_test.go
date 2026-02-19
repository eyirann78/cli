package credits

import (
	"math/rand"
	"testing"
)

// Benchmark for starLine function
func BenchmarkStarLine(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	widths := []int{50, 100, 200}
	
	for _, width := range widths {
		b.Run(string(rune(width)), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				starLine(r, width)
			}
		})
	}
}

// Benchmark for twinkle function
func BenchmarkTwinkle(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	testLine := starLine(r, 100)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twinkle(testLine)
	}
}

// Test to ensure starLine produces correct length output
func TestStarLine(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	widths := []int{0, 1, 10, 50, 100}
	
	for _, width := range widths {
		result := starLine(r, width)
		if len(result) != width {
			t.Errorf("starLine(%d) = len %d, want %d", width, len(result), width)
		}
	}
}

// Test to ensure twinkle performs the correct character swaps
func TestTwinkle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{".", "+"},
		{"+", "*"},
		{"*", "."},
		{" ", " "},
		{".+*", "+*."},
		{"..++**", "++**.."},
		{" . + * ", " + * . "},
	}
	
	for _, tt := range tests {
		result := twinkle(tt.input)
		if result != tt.expected {
			t.Errorf("twinkle(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
