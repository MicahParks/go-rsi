package rsi_test

import (
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/MicahParks/go-rsi"
)

// func BenchmarkBigRSI_Calculate(b *testing.B) {
// 	avgGain := avg(avgGains[0:rsi.DefaultPeriods])
// 	avgLoss := avg(avgLosses[0:rsi.DefaultPeriods])
//
// 	r, _ := rsi.NewBig(0, rsi.BigInput{
// 		AverageGain: big.NewFloat(avgGain),
// 		AverageLoss: big.NewFloat(avgLoss),
// 	})
//
// 	for i := rsi.DefaultPeriods; i < len(bigAvgGains); i++ {
// 		avgGain, _ = bigAvgGains[i].Float64()
// 		avgLoss, _ = bigAvgLosses[i].Float64()
// 		_ = r.Calculate(rsi.BigInput{
// 			AverageGain: big.NewFloat(avgGain),
// 			AverageLoss: big.NewFloat(avgLoss),
// 		})
// 	}
// }
//
// func BenchmarkRSI_Calculate(b *testing.B) {
// 	avgGain := avg(avgGains[0:rsi.DefaultPeriods])
// 	avgLoss := avg(avgLosses[0:rsi.DefaultPeriods])
//
// 	r, _ := rsi.New(0, rsi.Input{
// 		AverageGain: avgGain,
// 		AverageLoss: avgLoss,
// 	})
//
// 	for i := rsi.DefaultPeriods; i < len(avgGains); i++ {
// 		avgGain = avgGains[i]
// 		avgLoss = avgLosses[i]
// 		_ = r.Calculate(rsi.Input{
// 			AverageGain: avgGain,
// 			AverageLoss: avgLoss,
// 		})
// 	}
// }
//
// func ExampleRSI_Calculate() {
// 	// Create a logger.
// 	logger := log.New(os.Stdout, "", 0)
//
// 	// Gather some data.
// 	//
// 	// For production systems, it'd be best to gather test data asynchronously.
// 	avgGains, avgLosses = avgGains, avgLosses
//
// 	// Determine the number of periods for the initial inputs. Defaults to 14.
// 	periods := rsi.DefaultPeriods
//
// 	// Average the gains and losses over the given period.
// 	avgGain := avg(avgGains[0:periods])
// 	avgLoss := avg(avgLosses[0:periods])
// 	initialInput := rsi.Input{
// 		AverageGain: avgGain,
// 		AverageLoss: avgLoss,
// 	}
//
// 	// Create the RSI data structure and get the first result.
// 	//
// 	// If the first argument, the initial periods is 0, the default value, 14, will be used.
// 	r, result := rsi.New(uint(periods), initialInput)
// 	logger.Printf("Period index: %d\nAverage gain: %.2f\nAverage loss: %.2f\nRSI: %.2f", periods-1, avgGain, avgLoss, result)
//
// 	// Use the remaining data to generate the RSI for each period.
// 	for i := periods; i < len(avgGains); i++ {
// 		avgGain = avgGains[i]
// 		avgLoss = avgLosses[i]
// 		result = r.Calculate(rsi.Input{
// 			AverageGain: avgGain,
// 			AverageLoss: avgLoss,
// 		})
// 	}
// 	logger.Printf("Period index: %d\nAverage gain: %.2f\nAverage loss: %.2f\nRSI: %.2f", len(avgGains)-1, avgGain, avgLoss, result)
// 	// Output: Period index: 13
// 	// Average gain: -0.11
// 	// Average loss: -0.11
// 	// RSI: 50.37
// 	// Period index: 99
// 	// Average gain: 0.35
// 	// Average loss: 0.00
// 	// RSI: 271.62
// }

func TestBigRSI_Calculate(t *testing.T) {
	logger := log.New(os.Stdout, "", 0)

	r, result := rsi.NewBig(bigPrices[:rsi.DefaultPeriods+1])
	logger.Printf("%.8f", result)

	for _, next := range bigPrices[rsi.DefaultPeriods+1:] {
		result = r.Calculate(next)
		logger.Printf("%.8f", result)
	}
}

func TestRSI_Calculate(t *testing.T) {
	logger := log.New(os.Stdout, "", 0)

	r, result := rsi.New(prices[0 : rsi.DefaultPeriods+1])
	logger.Printf("%.8f", result)

	for _, next := range prices[rsi.DefaultPeriods+1:] {
		result = r.Calculate(next)
		logger.Printf("%.8f", result)
	}
}

func avg(s []float64) (avg float64) {
	for _, v := range s {
		avg += v
	}
	avg /= float64(len(s))
	return avg
}

func floatToBig(s []float64) (b []*big.Float) {
	l := len(s)
	b = make([]*big.Float, l)
	for i := 0; i < l; i++ {
		b[i] = big.NewFloat(s[i])
	}
	return b
}

var (
	bigPrices = floatToBig(prices)
	prices    = []float64{
		88.55,
		88.96,
		87.93,
		88.48,
		88.14,
		89.1,
		89.61,
		89.58,
		89.53,
		89.47,
		90.2,
		90.8,
		90.92,
		90.7,
		91.35,
		91.73,
		91.11,
		91.76,
		90.67,
		89.97,
		90.04,
		89.15,
		89.74,
		89.79,
		88.25,
		88.61,
		88.28,
		88.46,
		88.67,
		88.52,
		88.36,
		89.32,
		88.97,
		88.59,
		89.2,
		87.7,
		88.51,
		89.53,
		89.55,
		89.43,
		88.95,
		89.61,
		89.56,
		88.71,
		88.97,
		89.16,
		89.16,
		89.26,
		89.83,
		89.37,
		88.77,
		88.58,
		88.63,
		88.17,
		89.29,
		89.02,
		90.13,
		90.58,
		90.11,
		90.77,
		90.58,
		89.93,
		89.77,
		89.01,
		88.6,
		88.77,
		88.68,
		89.22,
		88.99,
		86.49,
		86.61,
		86.08,
		86.62,
		87.45,
		88.09,
		89.3,
		89.43,
		88.07,
		88,
		88.21,
		88.99,
		87.51,
		88.63,
		90.02,
		90.58,
		91.15,
		91.7,
		91.03,
		90.14,
		90.1,
		91.11,
		91.24,
		90.7,
		91.76,
		92.52,
		92.8,
		92.19,
		91.79,
		91.53,
		92.38,
		92.08,
		92.09,
		90.47,
		91.46,
		90.79,
		89.65,
		90.23,
		91.45,
		91.12,
		90.46,
		89.74,
		91.33,
		90.58,
		91.51,
		91.93,
		91,
		89.05,
		89.48,
		88.48,
		89.11,
		89.8,
		90.04,
		90.35,
		90.3,
		90.64,
		90.69,
		91.57,
		92.25,
		93.34,
		92.82,
		93.27,
		93.4,
		93.25,
		93.3,
		93.46,
		94.38,
		95.17,
		94.83,
		95.2,
		94.99,
		94.97,
		94.79,
		94.03,
		93.81,
		93.75,
		92.73,
		91.95,
		91.59,
		91.47,
		91.34,
		91.78,
		90.72,
		89.89,
		89.4,
		88.95,
		88.16,
		87.91,
		86.92,
		86.83,
		86.42,
		86.98,
		87.52,
		87.09,
		86.82,
		88.84,
		87.27,
		87.16,
	}
)
