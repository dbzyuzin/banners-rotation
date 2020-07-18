package bandit

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandsManagement(t *testing.T) {
	bandit := NewDefault()

	bandit.SetHands([]interface{}{1, 2, 3, 4, 5})

}

func TestSelect(t *testing.T) {
	t.Run("hands not set", func(t *testing.T) {
		bandit := NewDefault()

		hand, err := bandit.SelectHand()
		require.Error(t, err)
		require.Nil(t, hand)
	})

	t.Run("first time all hands selected one by one", func(t *testing.T) {
		bandit := NewDefault()

		hands := []interface{}{1, 2, 3, 4, 5}
		bandit.SetHands(hands)

		used := make(map[interface{}]struct{})
		for range hands {
			hand, err := bandit.SelectHand()
			require.NoError(t, err)
			used[hand] = struct{}{}
		}

		require.Len(t, used, len(hands))
	})

	t.Run("hands will selected without confirm", func(t *testing.T) {
		bandit := NewDefault()

		hands := []interface{}{1, 2, 3, 4, 5}
		bandit.SetHands(hands)

		used := make(map[interface{}]int)
		for range append(hands, hands...) {
			hand, err := bandit.SelectHand()
			require.NoError(t, err)
			used[hand]++
		}

		require.Len(t, used, len(hands))
		for _, count := range used {
			require.Equal(t, 2, count)
		}
	})
}

func TestConfirm(t *testing.T) {
	t.Run("can confirm added and selected hand", func(t *testing.T) {
		bandit := NewDefault()
		hand := 1
		bandit.AddHand(hand)

		_, err := bandit.SelectHand()
		require.NoError(t, err)

		err = bandit.ConfirmIncome(hand)
		require.NoError(t, err)
	})

	t.Run("confirm more then select", func(t *testing.T) {
		bandit := NewDefault()
		bandit.AddHand(1)

		err := bandit.ConfirmIncome(1)
		require.EqualError(t, err, ErrConfirmMoreThenSelect.Error())
	})

	t.Run("confirm hand not found", func(t *testing.T) {
		bandit := NewDefault()
		hand := 1
		bandit.AddHand(hand)

		handSelected, err := bandit.SelectHand()
		require.NoError(t, err)
		require.Equal(t, hand, handSelected)

		err = bandit.ConfirmIncome(2)
		require.EqualError(t, err, ErrHandNotFound.Error())
	})
}

func TestComplex(t *testing.T) {
	hand, hand2 := 1, 2
	initBandit := func() *Bandit {
		bandit := NewDefault()

		bandit.SetHands([]interface{}{hand, hand2})

		_, err := bandit.SelectHand()
		require.NoError(t, err)
		_, err = bandit.SelectHand()
		require.NoError(t, err)
		return bandit
	}

	t.Run("every time one has income -> all will selected in first time", func(t *testing.T) {
		bandit := NewDefault()
		bandit.SetHands([]interface{}{hand, hand2})

		h1, err := bandit.SelectHand()
		require.NoError(t, err)
		err = bandit.ConfirmIncome(hand)
		require.NoError(t, err)
		h2, err := bandit.SelectHand()
		require.NoError(t, err)

		require.Truef(t, h1 == hand2 || h2 == hand2, "need to get all hands ones at least")
	})

	t.Run("if have income then more often selected", func(t *testing.T) {
		bandit := initBandit()

		err := bandit.ConfirmIncome(hand)
		require.NoError(t, err)
		err = bandit.ConfirmIncome(hand)
		require.NoError(t, err)

		stats := make(map[interface{}]int)
		for i := 0; i < 50; i++ {
			h, err := bandit.SelectHand()
			require.NoError(t, err)
			stats[h]++
		}

		require.Less(t, stats[hand2], stats[hand])
		require.Less(t, 0, stats[hand2])
	})

	t.Run("only one have income, but another will selected", func(t *testing.T) {
		bandit := initBandit()

		err := bandit.ConfirmIncome(hand)
		require.NoError(t, err)
		err = bandit.ConfirmIncome(hand)
		require.NoError(t, err)

		stats := make(map[interface{}]int)
		for i := 0; i < 50; i++ {
			h, err := bandit.SelectHand()
			require.NoError(t, err)
			err = bandit.ConfirmIncome(hand)
			require.NoError(t, err)
			stats[h]++
		}

		require.Contains(t, stats, hand2)
	})
}
