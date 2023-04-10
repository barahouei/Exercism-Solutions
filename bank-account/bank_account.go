// Package account manages bank account.
package account

import (
	"sync"
)

// A customer account.
type Account struct {
	mu     sync.Mutex
	amount int64
	status bool // status false means account is closed.
}

// Open returns a new account with the given amount.
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{amount: amount, status: true}
}

// Balance returns balance of the account,
// if account is closed it returns false.
func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.status {
		return 0, false
	}

	return a.amount, true
}

// Deposit is used for both Deposits and Withdrawals, it returns true if operation succeeds.
//
// positive number = Deposit
//
// Negative Number = Withdrawal
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.status {
		return 0, false

	}

	if amount < 0 {
		amount *= -1 // Convert amount to a positive number.

		if a.amount-amount < 0 {
			return a.amount, false
		}

		a.amount -= amount

		return a.amount, true
	}

	a.amount += amount

	return a.amount, true
}

// Close closes the account. it returns account's balance,
// and returns true if account is successfully closed.
func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.status {
		return 0, false
	}

	b := a.amount

	a.amount = 0
	a.status = false

	return b, true
}
