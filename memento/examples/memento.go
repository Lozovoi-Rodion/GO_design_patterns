package memento

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{balance: balance}
	b.changes = append(b.changes, &Memento{balance})
	return b
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = %d$, current = %d", b.balance, b.current)
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	b.current++
	m := &Memento{b.balance}
	b.changes = append(b.changes, m)
	fmt.Printf("Deposited %d, balance now is %d \n", amount, b.balance)
	return m
}

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 2
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		fmt.Println("MEMENTO before:", b.changes[b.current])
		b.current--
		m := b.changes[b.current]
		fmt.Println("MEMENTO after:", m)
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}
