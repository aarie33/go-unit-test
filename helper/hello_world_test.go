package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Test")
	m.Run()
	fmt.Println("After Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac OS")
	}

	result := HelloWorld("Skip")
	require.Equal(t, "Hello Skip", result)
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		t.Errorf("HelloWorld(\"World\") = %s; want \"Hello World\"", result)
		// t.Fail()
		// t.FailNow()
		// t.Error("HelloWorld(\"World\") = %s; want \"Hello World\"", result)
		// t.Fatal("HelloWorld(\"World\") = %s; want \"Hello World\"", result)
	}
}

func TestHelloJoko(t *testing.T) {
	result := HelloWorld("Joko")
	if result != "Hello Joko" {
		t.Errorf("HelloWorld(\"Joko\") = %s; want \"Hello Joko\"", result)
	}
}

func TestHelloAssertion(t *testing.T) {
	result := HelloWorld("Joko")
	assert.Equal(t, "Hello Joko", result)
	// fmt.Println("TestHelloAssertion")
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("Joko")
	require.Equal(t, "Hello Joko", result) // manggil fail now
	// fmt.Println("TestHelloRequire")
}

func TestSubTest(t *testing.T) {
	t.Run("Joko", func(t *testing.T) {
		result := HelloWorld("Joko")
		require.Equal(t, "Hello Joko", result)
	})
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result)
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		}, {
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mamet")
	}
}

func BenchmarkHelloWorldBudi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Budi")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Santoso", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Santoso")
		}
	})

	b.Run("Jono", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jono")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Santoso",
			request: "Santoso",
		},
		{
			name:    "Jono",
			request: "Jono",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
