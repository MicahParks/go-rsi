package rsi

import (
	"math/big"
)

var (
	big0   = big.NewFloat(0)
	big1   = big.NewFloat(1)
	big100 = big.NewFloat(100)
)

// BigRSI represents the state of a Relative Strength Index (RSI) algorithm.
type BigRSI struct {
	periods       *big.Float
	periodsMinus1 *big.Float
	prev          *big.Float
	prevAvgGain   *big.Float
	prevAvgLoss   *big.Float
}

// NewBig creates a new RSI data structure and returns the initial result.
//
// The length of the initial input slice should be 1 + `periods`. Where `periods` is the length of the lookback period.
func NewBig(initial []*big.Float) (r *BigRSI, result *big.Float) {
	if len(initial) < 2 {
		return nil, nil
	}
	periods := len(initial) - 1

	r = &BigRSI{
		periods:     big.NewFloat(float64(periods)),
		prev:        new(big.Float).Copy(initial[periods]),
		prevAvgGain: big.NewFloat(0),
		prevAvgLoss: big.NewFloat(0),
	}
	r.periodsMinus1 = new(big.Float).Sub(r.periods, big1)

	// Calculate AvgGain and AvgLoss from initial slice.
	var prev *big.Float
	for i := 0; i < len(initial); i++ {
		if i != 0 {
			diff := new(big.Float).Sub(initial[i], prev)
			if diff.Cmp(big0) == 1 {
				r.prevAvgGain.Add(r.prevAvgGain, diff)
			} else {
				r.prevAvgLoss.Sub(r.prevAvgLoss, diff)
			}
		}
		prev = initial[i]
	}
	r.prevAvgGain.Quo(r.prevAvgGain, r.periods)
	r.prevAvgLoss.Quo(r.prevAvgLoss, r.periods)

	result = new(big.Float).Sub(big100, new(big.Float).Quo(big100, new(big.Float).Add(big1, new(big.Float).Quo(r.prevAvgGain, r.prevAvgLoss))))

	return r, result
}

// Calculate produces the next RSI result given the next input.
func (r *BigRSI) Calculate(next *big.Float) (result *big.Float) {
	gain := big.NewFloat(0)
	loss := big.NewFloat(0)
	if diff := new(big.Float).Sub(next, r.prev); diff.Cmp(big0) == 1 {
		gain.Add(gain, diff)
	} else {
		loss.Sub(loss, diff)
	}

	r.prev.Copy(next)

	r.prevAvgGain = new(big.Float).Quo(new(big.Float).Add(new(big.Float).Mul(r.prevAvgGain, r.periodsMinus1), gain), r.periods)
	r.prevAvgLoss = new(big.Float).Quo(new(big.Float).Add(new(big.Float).Mul(r.prevAvgLoss, r.periodsMinus1), loss), r.periods)

	result = new(big.Float).Sub(big100, new(big.Float).Quo(big100, new(big.Float).Add(big1, new(big.Float).Quo(r.prevAvgGain, r.prevAvgLoss))))

	return result
}
