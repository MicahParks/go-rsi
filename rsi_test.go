package rsi_test

import (
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/MicahParks/go-rsi"
)

func BenchmarkBigRSI_Calculate(b *testing.B) {
	gains := []float64{
		0.00,
		0.06,
		0.00,
		0.72,
		0.50,
		0.27,
		0.33,
		0.42,
		0.24,
		0.00,
		0.14,
		0.00,
		0.67,
		0.00,
		0.00,
		0.03,
		0.38,
		0.00,
		0.00,
		0.57,
		0.04,
		0.00,
		0.74,
		0.00,
		0.00,
		0.00,
		0.15,
		0.04,
		0.35,
		0.00,
		0.00,
		0.47,
	}

	losses := []float64{
		0.25,
		0.00,
		0.54,
		0.00,
		0.00,
		0.00,
		0.00,
		0.00,
		0.00,
		0.19,
		0.00,
		0.42,
		0.00,
		0.00,
		0.28,
		0.00,
		0.00,
		0.19,
		0.58,
		0.00,
		0.00,
		0.54,
		0.00,
		0.67,
		0.43,
		1.33,
		0.00,
		0.00,
		0.00,
		1.15,
		0.76,
		0.00,
	}

	avgGain := avg(gains[0:rsi.DefaultPeriods])
	avgLoss := avg(losses[0:rsi.DefaultPeriods])

	result, r := rsi.NewBig(rsi.DefaultPeriods, rsi.BigInput{
		AverageGain: big.NewFloat(avgGain),
		AverageLoss: big.NewFloat(avgLoss),
	})

	res, _ := result.Float64()
	log.Println(fmt.Sprintf("Starting:\n  Gains: %.2f\n  Losses: %.2f\n  RSI: %.2f", avgGain, avgLoss, res))

	for i := rsi.DefaultPeriods; i < len(gains); i++ {
		avgGain = gains[i]
		avgLoss = losses[i]
		result = r.Calculate(rsi.BigInput{
			AverageGain: big.NewFloat(avgGain),
			AverageLoss: big.NewFloat(avgLoss),
		})
		res, _ = result.Float64()
		log.Println(fmt.Sprintf("Starting:\n  Gains: %.2f\n  Losses: %.2f\n  RSI: %.2f", avgGain, avgLoss, res))
	}
}

func BenchmarkRSI_Calculate(b *testing.B) {
	gains := []float64{
		0.00,
		0.06,
		0.00,
		0.72,
		0.50,
		0.27,
		0.33,
		0.42,
		0.24,
		0.00,
		0.14,
		0.00,
		0.67,
		0.00,
		0.00,
		0.03,
		0.38,
		0.00,
		0.00,
		0.57,
		0.04,
		0.00,
		0.74,
		0.00,
		0.00,
		0.00,
		0.15,
		0.04,
		0.35,
		0.00,
		0.00,
		0.47,
	}

	losses := []float64{
		0.25,
		0.00,
		0.54,
		0.00,
		0.00,
		0.00,
		0.00,
		0.00,
		0.00,
		0.19,
		0.00,
		0.42,
		0.00,
		0.00,
		0.28,
		0.00,
		0.00,
		0.19,
		0.58,
		0.00,
		0.00,
		0.54,
		0.00,
		0.67,
		0.43,
		1.33,
		0.00,
		0.00,
		0.00,
		1.15,
		0.76,
		0.00,
	}

	avgGain := avg(gains[0:rsi.DefaultPeriods])
	avgLoss := avg(losses[0:rsi.DefaultPeriods])

	result, r := rsi.New(rsi.DefaultPeriods, rsi.Input{
		AverageGain: avgGain,
		AverageLoss: avgLoss,
	})

	log.Println(fmt.Sprintf("Starting:\n  Gains: %.2f\n  Losses: %.2f\n  RSI: %.2f", avgGain, avgLoss, result))

	for i := rsi.DefaultPeriods; i < len(gains); i++ {
		avgGain = gains[i]
		avgLoss = losses[i]
		result = r.Calculate(rsi.Input{
			AverageGain: avgGain,
			AverageLoss: avgLoss,
		})
		log.Println(fmt.Sprintf("Starting:\n  Gains: %.2f\n  Losses: %.2f\n  RSI: %.2f", avgGain, avgLoss, result))
	}
}

func avg(s []float64) (a float64) {
	for _, v := range s {
		a += v
	}
	a /= float64(len(s))
	return a
}

// func testData() (gains, losses, results []float64) {
//
// }
