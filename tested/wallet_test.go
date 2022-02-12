package tested

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet.Balance(), wallet.Balance())
	})

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{100}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet.Balance(), Bitcoin(90))
	})

	t.Run("Withdraw with insufficient balance", func(t *testing.T) {
		wallet := Wallet{1}
		err := wallet.Withdraw(Bitcoin(10))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet.Balance(), Bitcoin(1))
	})
}
