[![Go Reference](https://pkg.go.dev/badge/github.com/MicahParks/go-rsi/v2.svg)](https://pkg.go.dev/github.com/MicahParks/go-rsi/v2) [![Go Report Card](https://goreportcard.com/badge/github.com/MicahParks/go-rsi/v2)](https://goreportcard.com/report/github.com/MicahParks/go-rsi/v2)
# go-rsi
The Relative Strength Index (RSI) technical analysis algorithm implemented in Golang.

```
import "github.com/MicahParks/go-rsi/v2"
```

# Usage
For full examples, please see the `examples` directory.

## Preconditions
1. Gather test data.
2. Decide on the number of periods, `p`, for the RSI algorithm. Populate a slice of prices whose length is `p + 1`.

```go
// Determine the number of periods for the initial inputs. Defaults to 14.
const initialLength = rsi.DefaultPeriods + 1
initial := prices[:initialLength]
```

## Step 1
Create the RSI data structure and get the first result.
```go
r, result := rsi.New(initial)
```

## Step 2
Use the remaining data to calculate the RSI value for that period.
```go
remaining := prices[initialLength:]
for i, next := range remaining {
	result = r.Calculate(next)
}
```

# Testing
There is 100% test coverage and benchmarks for this project. Here is an example benchmark result:
```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/MicahParks/go-rsi/v2
cpu: Intel(R) Core(TM) i5-9600K CPU @ 3.70GHz
BenchmarkBigRSI_Calculate-6     1000000000               0.0001744 ns/op
BenchmarkRSI_Calculate-6        1000000000               0.0000017 ns/op
PASS
ok      github.com/MicahParks/go-rsi/v2    0.005s
```

# Resources
I built and tested this package using these resources:
* [Investopedia](https://www.investopedia.com/terms/r/rsi.asp)
* [Invest Excel](https://investexcel.net/relative-strength-index-spreadsheet/)
