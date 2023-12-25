package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 50}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet, 40)

	})

	t.Run("insufficient funds", func(t *testing.T) {
		startingBal := Bitcoin(20)
		wallet := Wallet{balance: startingBal}
		err := wallet.Withdraw(100)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBal)

	})

}

func assertBalance(t testing.TB, wallet Wallet, balance Bitcoin) {
	t.Helper()
	got := wallet.GetBalance()
	if got != balance {
		t.Errorf("got %s want %s", got, balance)
	}
}

func assertError(t testing.TB, err, errMessage error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if err != errMessage {
		t.Errorf("got %q, expected %q", err, errMessage)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but was not expecting one", err)
	}
}
