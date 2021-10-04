package rsi

// TODO
type Input struct {
	AverageGain float64
	AverageLoss float64
}

// TODO
type RSI struct {
	periods  uint
	previous Input
}

// TODO
func New(periods uint, initial Input) (initialValue float64, r *RSI) {
	if periods == 0 {
		periods = 14
	}

	r = &RSI{
		periods:  periods,
		previous: initial,
	}

	initialValue = 100 - (100 / (1 + ((r.previous.AverageGain / float64(r.periods)) / (r.previous.AverageLoss / float64(r.periods)))))

	return initialValue, r
}

// TODO
func (r *RSI) Calculate(i Input) (value float64) {
	r.previous.AverageGain = (r.previous.AverageGain*float64(r.periods-1) + i.AverageGain) / float64(r.periods)
	r.previous.AverageLoss = (r.previous.AverageLoss*float64(r.periods-1) + i.AverageLoss) / float64(r.periods)

	value = 100 - 100/(1+r.previous.AverageGain/r.previous.AverageLoss)

	return value
}
