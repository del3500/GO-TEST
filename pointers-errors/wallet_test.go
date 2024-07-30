package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{balance: 11}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(1))
	})

	t.Run("withdraw w/ insufficient BTC", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(21))

		assertError(t, err, ErrInsufficientBTC)
		assertBalance(t, wallet, startingBalance)

	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("want an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
