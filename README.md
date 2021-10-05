# go-rsi
The Relative Strength Index (RSI) algorithm implemented in Golang.

# Usage
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
logger.Printf("Period index: %d\n  Average gain: %.2f\n  Average loss: %.2f\n RSI: %.2f", periods-1, avgGain, avgLoss, result)
```

## Step 3
Use the remaining data per period to calculate the RSI value at that period.
```go
// Use the remaining data to generate the RSI for each period.
for i := periods; i < len(avgGains); i++ {
	avgGain = avgGains[i]
	avgLoss = avgLosses[i]
	result = r.Calculate(rsi.Input{
		AverageGain: avgGain,
		AverageLoss: avgLoss,
	})
	logger.Printf("Period index: %d\n  Average gain: %.2f\n  Average loss: %.2f\n RSI: %.2f", i, avgGain, avgLoss, result)
}
```

# Full example
See the `examples` directory for a full example.

# Testing
There is 100% test coverage and benchmarks for this project. Here is an example benchmark result:
```
$ go test -bench . goos: linux goarch: amd64 pkg: github.com/MicahParks/go-rsi cpu: Intel(R) Core(TM) i5-9600K CPU @
3.70GHz BenchmarkBigRSI_Calculate-6 1000000000 0.0001172 ns/op BenchmarkRSI_Calculate-6 1000000000 0.0000058 ns/op PASS
ok github.com/MicahParks/go-rsi 0.004s
```