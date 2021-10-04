package rsi

import (
	"math/big"
)

// TODO
type BigInput struct {
	AverageGain *big.Float
	AverageLoss *big.Float
}

// TODO
type BigRSI struct {
	periods  *big.Float
	previous BigInput
}

// TODO
func NewBig(periods uint, initial BigInput) (initialValue *big.Float, r *BigRSI) {
	if periods == 0 {
		periods = 14
	}

	r = &BigRSI{
		periods:  big.NewFloat(float64(periods)),
		previous: initial,
	}

	initialValue = new(big.Float).Sub(big.NewFloat(100), new(big.Float).Quo(big.NewFloat(100), new(big.Float).Add(big.NewFloat(1), new(big.Float).Quo(new(big.Float).Quo(r.previous.AverageGain, r.periods), new(big.Float).Quo(r.previous.AverageLoss, r.periods)))))

	return initialValue, r
}

// TODO
func (r *BigRSI) Calculate(i BigInput) (value float64) {
	new(big.Float).Mul(new(big.Float), r.previous.AverageGain)
	r.previous.AverageGain = (r.previous.AverageGain*float64(r.periods-1) + i.AverageGain) / float64(r.periods)
	r.previous.AverageLoss = (r.previous.AverageLoss*float64(r.periods-1) + i.AverageLoss) / float64(r.periods)

	value = 100 - 100/(1+r.previous.AverageGain/r.previous.AverageLoss)

	return value
}
