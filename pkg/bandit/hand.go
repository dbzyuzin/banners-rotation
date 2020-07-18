package bandit

import "math"

type Hand struct {
	TimesIsSelected int64
	TimesSucceed    int64
	Data            interface{}
}

func (h Hand) MyScore(n int64) float64 {
	averageIncome := float64(h.TimesSucceed) / float64(h.TimesIsSelected)
	score := averageIncome + math.Sqrt(2*math.Log(float64(n))/float64(h.TimesIsSelected))
	return score
}
