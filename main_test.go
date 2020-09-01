package main

import "testing"

var (
	smallFile = readFile(small)
	medFile   = readFile(med)
	largeFile = readFile(large)
)

func BenchmarkJSONMarshalSmall(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		marshalJSON(smallFile)
	}
}

func BenchmarkJSONMarshalMed(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		marshalJSON(medFile)
	}
}

func BenchmarkJSONMarshalLarge(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		marshalJSON(largeFile)
	}
}

func BenchmarkJSONUnmarshalSmall(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseJSON(smallFile)
	}
}

func BenchmarkJSONUnmarshalMed(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseJSON(medFile)
	}
}

func BenchmarkJSONUnmarshalLarge(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseJSON(largeFile)
	}
}

// JSONIter

func BenchmarkJSONIterMarshalSmall(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		marshalJSONIter(smallFile)
	}
}

func BenchmarkJSONIterMarshalMed(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		marshalJSONIter(medFile)
	}
}

func BenchmarkJSONIterMarshalLarge(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		marshalJSONIter(largeFile)
	}
}

func BenchmarkJSONIterUnmarshalSmall(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseJSONIter(smallFile)
	}
}

func BenchmarkJSONIterUnmarshalMed(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseJSONIter(medFile)
	}
}

func BenchmarkJSONIterUnmarshalLarge(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseJSONIter(largeFile)
	}
}
