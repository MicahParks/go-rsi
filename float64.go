package rsi

// DefaultPeriods is the default number of periods for the averages for the RSI algorithm.
const DefaultPeriods = 14

// Input represents the averages for a period that are inputted into the RSI algorithm.
type Input struct {
	AverageGain float64
	AverageLoss float64
}

// RSI represents the state of a Relative Strength Index (RSI) algorithm.
type RSI struct {
	periods         float64
	periodsMinusOne float64
	previous        Input
}

// New creates a new RSI data structure and returns the initial result.
func New(periods uint, initial Input) (r *RSI, result float64) {
	if periods == 0 {
		periods = DefaultPeriods
	}

	r = &RSI{
		periods:         float64(periods),
		periodsMinusOne: float64(periods - 1),
		previous:        initial,
	}

	result = 100 - (100 / (1 + ((r.previous.AverageGain / r.periods) / (r.previous.AverageLoss / r.periods))))

	return r, result
}

// Calculate produces the next RSI result given the next input.
func (r *RSI) Calculate(next Input) (result float64) {
	r.previous.AverageGain = (r.previous.AverageGain*(r.periodsMinusOne) + next.AverageGain) / r.periods
	r.previous.AverageLoss = (r.previous.AverageLoss*(r.periodsMinusOne) + next.AverageLoss) / r.periods

	result = 100 - 100/(1+r.previous.AverageGain/r.previous.AverageLoss)

	return result
}
