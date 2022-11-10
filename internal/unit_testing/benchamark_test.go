package unittesting

import (
	"testing"

	"github.com/GabrielEdwinSP/go-csv/internal/services"
)

func BenchmarkProcessFileTest(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		services.ProcessFile()
	}
}
