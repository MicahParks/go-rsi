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
	periods  uint
	previous BigInput
}

// TODO
func NewBig(periods uint, initial BigInput) (initialValue float64, r *BigRSI) {
	if periods == 0 {
		periods = 14
	}

	r = &BigRSI{
		periods:  periods,
		previous: initial,
	}

	initialValue = 100 - (100 / (1 + ((r.previous.AverageGain / float64(r.periods)) / (r.previous.AverageLoss / float64(r.periods)))))

	return initialValue, r
}

// TODO
func (r *BigRSI) Calculate(i Input) (value float64) {
	r.previous.AverageGain = (r.previous.AverageGain*float64(r.periods-1) + i.AverageGain) / float64(r.periods)
	r.previous.AverageLoss = (r.previous.AverageLoss*float64(r.periods-1) + i.AverageLoss) / float64(r.periods)

	value = 100 - 100/(1+r.previous.AverageGain/r.previous.AverageLoss)

	return value
}
