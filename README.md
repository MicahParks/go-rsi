[![Go Reference](https://pkg.go.dev/badge/github.com/MicahParks/go-rsi.svg)](https://pkg.go.dev/github.com/MicahParks/go-rsi) [![Go Report Card](https://goreportcard.com/badge/github.com/MicahParks/go-rsi)](https://goreportcard.com/report/github.com/MicahParks/go-rsi)
# go-rsi
The Relative Strength Index (RSI) algorithm implemented in Golang.

# Usage
For full examples, please see the `examples` directory.

## Preconditions
1. Gather test data.
2. Decide on the number of periods for the averages for the RSI algorithm.

```go
// Gather some test data.
//
// For production systems, it'd be best to gather test data asynchronously.
avgGains, avgLosses, _ := testData()

// Determine the number of periods for the initial inputs. Defaults to 14.
periods := rsi.DefaultPeriods
```

## Step 1
Create the initial input. This is the average of the gains and losses over a given number of periods.
```go
// Average the gains and losses over the given period.
avgGain := avg(avgGains[0:periods])
avgLoss := avg(avgLosses[0:periods])
initialInput := rsi.Input{
	AverageGain: avgGain,
	AverageLoss: avgLoss,
}
```

## Step 2
Create the RSI data structure and get the first result.
```go
// Create the RSI data structure and get the first result.
//
// If the first argument, the initial periods is 0, the default value, 14, will be used.
r, result := rsi.New(uint(periods), initialInput)
```

## Step 3
Use the remaining data per period to calculate the RSI result at that period.
```go
// Use the remaining data to generate the RSI for each period.
for i := periods; i < len(avgGains); i++ {
	avgGain = avgGains[i]
	avgLoss = avgLosses[i]
	result = r.Calculate(rsi.Input{
		AverageGain: avgGain,
		AverageLoss: avgLoss,
	})
}
```

# Testing
There is 100% test coverage and benchmarks for this project. Here is an example benchmark result:
```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/MicahParks/go-rsi
cpu: Intel(R) Core(TM) i5-9600K CPU @ 3.70GHz
BenchmarkBigRSI_Calculate-6     1000000000               0.0001274 ns/op
BenchmarkRSI_Calculate-6        1000000000               0.0000007 ns/op
PASS
ok      github.com/MicahParks/go-rsi    0.006s
```

# Resources
I built and tested this package using these resources:
* [Investopedia](https://www.investopedia.com/terms/r/rsi.asp)
* [Invest Excel](https://investexcel.net/relative-strength-index-spreadsheet/)
