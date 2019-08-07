// Package provides a concurrency-safe bank with one account.
package bank

type withdrawl struct {
	amount int
	result chan bool
}

var deposits = make(chan int)         // send amount to deposit
var balances = make(chan int)         // receive balance
var withdrawls = make(chan withdrawl) // withdrawls

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	result := make(chan bool)
	withdrawls <- withdrawl{amount, result}
	return <-result
}

func teller() {
	var balance int // balance is confined to the teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawls:
			if w.amount > balance {
				w.result <- false
				continue
			}
			balance -= w.amount
			w.result <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
