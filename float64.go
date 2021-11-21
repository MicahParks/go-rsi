package rsi

// DefaultPeriods is the default number of periods for the averages for the RSI algorithm.
const DefaultPeriods = 14

// RSI represents the state of a Relative Strength Index (RSI) algorithm.
type RSI struct {
	periods         float64
	periodsMinusOne float64
	prev            float64
	prevAvgGain     float64
	prevAvgLoss     float64
}

// New creates a new RSI data structure and returns the initial result.
//
// The length of the initial input slice should be 1 + `periods`. Where `periods` is the length of the lookback period.
func New(initial []float64) (r *RSI, result float64) {
	if len(initial) < 2 {
		return nil, 0
	}
	periods := len(initial) - 1

	r = &RSI{
		periods:         float64(periods),
		periodsMinusOne: float64(periods - 1),
		prev:            initial[periods],
	}

	// Calculate AvgGain and AvgLoss from initial slice.
	var prev float64
	for i := 0; i < len(initial); i++ {
		if i != 0 {
			diff := initial[i] - prev
			if diff > 0 {
				r.prevAvgGain += diff
			} else {
				r.prevAvgLoss -= diff
			}
		}
		prev = initial[i]
	}
	r.prevAvgGain /= r.periods
	r.prevAvgLoss /= r.periods

	result = 100 - (100 / (1 + r.prevAvgGain/r.prevAvgLoss))

	return r, result
}

// Calculate produces the next RSI result given the next input.
func (r *RSI) Calculate(next float64) (result float64) {
	gain := float64(0)
	loss := float64(0)
	if diff := next - r.prev; diff > 0 {
		gain = diff
	} else {
		loss = -diff
	}

	r.prev = next

	r.prevAvgGain = (r.prevAvgGain*r.periodsMinusOne + gain) / r.periods
	r.prevAvgLoss = (r.prevAvgLoss*r.periodsMinusOne + loss) / r.periods

	result = 100 - 100/(1+r.prevAvgGain/r.prevAvgLoss)

	return result
}
