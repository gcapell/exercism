package account

import (
	"sync"
)

type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// If Open is given a negative initial deposit, it must return nil.
func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	return &Account{balance: initalDeposit, closed: false}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	// a.RLock()
	// defer a.RUnlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
