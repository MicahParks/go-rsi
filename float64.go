package rsi

// TODO
type Input struct {
	AverageGain float64
	AverageLoss float64
}

// TODO
type RSI struct {
	periods  float64
	previous Input
}

// TODO
func New(periods uint, initial Input) (initialValue float64, r *RSI) {
	if periods == 0 {
		periods = 14
	}

	r = &RSI{
		periods:  float64(periods),
		previous: initial,
	}

	initialValue = 100 - (100 / (1 + ((r.previous.AverageGain / r.periods) / (r.previous.AverageLoss / r.periods))))

	return initialValue, r
}

// TODO
func (r *RSI) Calculate(i Input) (value float64) {
	r.previous.AverageGain = (r.previous.AverageGain*(r.periods-1) + i.AverageGain) / r.periods
	r.previous.AverageLoss = (r.previous.AverageLoss*(r.periods-1) + i.AverageLoss) / r.periods

	value = 100 - 100/(1+r.previous.AverageGain/r.previous.AverageLoss)

	return value
}
