package rsi

import (
	"math/big"
)

var (
	big1   = big.NewFloat(1)
	big100 = big.NewFloat(100)
)

// BigInput represents the averages for a period that are inputted into the RSI algorithm.
type BigInput struct {
	AverageGain *big.Float
	AverageLoss *big.Float
}

// BigRSI represents the state of a Relative Strength Index (RSI) algorithm.
type BigRSI struct {
	periods       *big.Float
	periodsMinus1 *big.Float
	previous      BigInput
}

// NewBig creates a new RSI data structure and returns the initial value.
func NewBig(periods uint, initial BigInput) (r *BigRSI, result *big.Float) {
	if periods == 0 {
		periods = DefaultPeriods
	}

	r = &BigRSI{
		periods:  big.NewFloat(float64(periods)),
		previous: initial,
	}
	r.periodsMinus1 = new(big.Float).Sub(r.periods, big1)

	result = new(big.Float).Sub(big100, new(big.Float).Quo(big100, new(big.Float).Add(big1, new(big.Float).Quo(new(big.Float).Quo(r.previous.AverageGain, r.periods), new(big.Float).Quo(r.previous.AverageLoss, r.periods)))))

	return r, result
}

// Calculate produces the next RSI value given the next input.
func (r *BigRSI) Calculate(next BigInput) (result *big.Float) {
	r.previous.AverageGain = new(big.Float).Quo(new(big.Float).Add(new(big.Float).Mul(r.previous.AverageGain, r.periodsMinus1), next.AverageGain), r.periods)
	r.previous.AverageLoss = new(big.Float).Quo(new(big.Float).Add(new(big.Float).Mul(r.previous.AverageLoss, r.periodsMinus1), next.AverageLoss), r.periods)

	result = new(big.Float).Sub(big100, new(big.Float).Quo(big100, new(big.Float).Add(big1, new(big.Float).Quo(r.previous.AverageGain, r.previous.AverageLoss))))

	return result
}
