package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		amount := Bitcoin(10)
		wallet.Deposit(amount)

		assetBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		Wallet := Wallet{
			balance: 100,
		}
		amt := Bitcoin(50)
		err := Wallet.Withdraw(amt)
		assertNoError(t, err)
	})

	t.Run("Low Balance", func(t *testing.T) {
		Wallet := Wallet{
			balance: 50,
		}
		err := Wallet.Withdraw(Bitcoin(60))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assetBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
