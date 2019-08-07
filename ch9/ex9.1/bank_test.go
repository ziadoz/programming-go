package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/ex9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		if !bank.Withdraw(50) {
			t.Errorf("Could not withdraw 50")
		}
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		if !bank.Withdraw(25) {
			t.Errorf("Could not withdraw 25")
		}
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank.Balance(), 225; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
