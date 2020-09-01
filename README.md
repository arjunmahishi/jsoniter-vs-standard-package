# encoding/json vs jsoniter

This is a benchmark in go to compare the performance of standard package `encoding/json` and `jsoniter`

## How to Run

```bash
go test -bench "<regex of benchmark name>"
```

## Result

```bash
 $ go test -bench ".*"
goos: linux
goarch: amd64
BenchmarkJSONMarshalSmall-4            30000     48131 ns/op
BenchmarkJSONMarshalMed-4               1000   1214124 ns/op
BenchmarkJSONMarshalLarge-4              200   8338271 ns/op
BenchmarkJSONUnmarshalSmall-4          30000     49919 ns/op
BenchmarkJSONUnmarshalMed-4             1000   1257659 ns/op
BenchmarkJSONUnmarshalLarge-4            200   8570177 ns/op
BenchmarkJSONIterMarshalSmall-4       500000      2127 ns/op
BenchmarkJSONIterMarshalMed-4          30000     42359 ns/op
BenchmarkJSONIterMarshalLarge-4         3000    519634 ns/op
BenchmarkJSONIterUnmarshalSmall-4     100000     23554 ns/op
BenchmarkJSONIterUnmarshalMed-4         1000   1239436 ns/op
BenchmarkJSONIterUnmarshalLarge-4        500   4094496 ns/op
PASS
```
