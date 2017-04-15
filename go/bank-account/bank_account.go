package account

import "sync"

const testVersion = 1

type Account struct {
	balance int
	open    bool
	sync.Mutex
}

func Open(i int) *Account {
	if i < 0 {
		return nil
	}
	return &Account{balance: i, open: true}
}

func (a *Account) Close() (int, bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, false
	}
	a.open = false
	return a.balance, true
}

func (a *Account) Balance() (int, bool) {
	return a.balance, a.open
}

func (a *Account) Deposit(i int) (int, bool) {
	a.Lock()
	defer a.Unlock()
	if a.balance+i < 0 {
		return a.balance, false
	}
	a.balance += i
	return a.balance, a.open
}
