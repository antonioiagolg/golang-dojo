package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	checkBalance := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}
	checkError := func(t testing.TB, got, want error) {

		if want != nil {
			if want.Error() != got.Error() {
				t.Errorf("Got %v, want wantor to be %v", got, want)
			}
		} else {
			if got != nil {
				t.Errorf("Got %v, want wantor to be %v", got, want)
			}
		}
	}
	t.Run("Test deposity", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)
		checkBalance(t, got, want)

	})

	t.Run("Test withdraw", func(t *testing.T) {

		tests := []struct {
			title    string
			wallet   Wallet
			withdraw Bitcoin
			want     Bitcoin
			err      error
		}{
			{
				title:    "Has Balance",
				wallet:   Wallet{balance: 20},
				withdraw: Bitcoin(10),
				want:     Bitcoin(10),
				err:      nil,
			},
			{
				title:    "No Balance",
				wallet:   Wallet{balance: 20},
				withdraw: Bitcoin(50),
				want:     Bitcoin(20),
				err:      ErrInsufficientFunds,
			},
		}

		for _, test := range tests {
			t.Run(test.title, func(t *testing.T) {
				err := test.wallet.Withdraw(test.withdraw)
				checkBalance(t, test.wallet.Balance(), test.want)
				checkError(t, err, test.err)
			})
		}

	})
}
