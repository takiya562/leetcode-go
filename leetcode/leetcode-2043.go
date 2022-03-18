package leetcode

type Bank struct {
	blance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{blance: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	n := len(this.blance)
	if account1 < 1 || account1 > n || account2 < 1 || account2 > n || money > this.blance[account1-1] {
		return false
	}
	this.blance[account1-1] -= money
	this.blance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account < 1 || account > len(this.blance) {
		return false
	}
	this.blance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account < 1 || account > len(this.blance) || money > this.blance[account-1] {
		return false
	}
	this.blance[account-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
