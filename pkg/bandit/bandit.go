package bandit

type HandStats struct {
	TimesIsSelected int64
	AverageIncome   float64
}

type Bandit struct {
	hands []interface{}
}

func (b *Bandit) SetHands(hands []interface{}) {
	b.hands = hands
}
