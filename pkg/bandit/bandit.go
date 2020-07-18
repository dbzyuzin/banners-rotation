package bandit

type Bandit struct {
	hands          []*Hand
	TimesIsUsed    int64
	TimesIsSucceed int64
}

func NewDefault() *Bandit {
	return &Bandit{}
}

func (b *Bandit) SetHands(hands []interface{}) {
	b.hands = make([]*Hand, 0, len(hands))
	for _, hand := range hands {
		b.hands = append(b.hands, &Hand{Data: hand})
	}
}

func (b *Bandit) AddHand(hand interface{}) {
	b.hands = append(b.hands, &Hand{Data: hand})
}

func (b *Bandit) RemoveHand(hand interface{}) {
	index := -1
	for i, h := range b.hands {
		if h.Data == hand {
			index = i
			break
		}
	}
	if index != -1 {
		b.hands = append(b.hands[:index], b.hands[index+1:]...)
	}
}

func (b *Bandit) DropHands() {
	b.hands = make([]*Hand, 0)
}

func (b *Bandit) DropStats() {
	b.TimesIsUsed = 0
	b.TimesIsSucceed = 0
	for _, hand := range b.hands {
		hand.TimesIsSelected = 0
		hand.TimesSucceed = 0
	}
}

func (b *Bandit) SelectHand() (interface{}, error) {
	if len(b.hands) == 0 {
		return nil, ErrHandsNotSet
	}

	for _, hand := range b.hands {
		if hand.TimesIsSelected == 0 {
			b.TimesIsUsed++
			hand.TimesIsSelected++
			return hand.Data, nil
		}
	}

	selectedHand := b.hands[0]
	maxScore := selectedHand.MyScore(b.TimesIsUsed)
	for _, hand := range b.hands {
		score := hand.MyScore(b.TimesIsUsed)
		if score > maxScore || score == maxScore && hand.TimesIsSelected < selectedHand.TimesIsSelected {
			maxScore = score
			selectedHand = hand
		}
	}

	b.TimesIsUsed++
	selectedHand.TimesIsSelected++
	return selectedHand.Data, nil
}

func (b *Bandit) ConfirmIncome(handData interface{}) error {
	if b.TimesIsUsed < b.TimesIsSucceed+1 {
		return ErrConfirmMoreThenSelect
	}
	for _, hand := range b.hands {
		if handData == hand.Data {
			b.TimesIsSucceed++
			hand.TimesSucceed++
			return nil
		}
	}
	return ErrHandNotFound
}
