package freebsdrandom

import (
	"github.com/seehuhn/mt19937"
	_ "github.com/seehuhn/mt19937"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkFreebsdrandomUint64n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64n(1e6)
	}
}

func BenchmarkMathRandUnit64n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Uint64()
	}
}

func BenchmarkMT19937Unit64n(b *testing.B) {
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		rng.Uint64()
	}
}

func BenchmarkFreebsdrandomIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intn(i)
	}
}

func BenchmarkMathRandIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1e6)
	}
}

func BenchmarkMT19937Intn(b *testing.B) {
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		rng.Intn(1e6)
	}
}