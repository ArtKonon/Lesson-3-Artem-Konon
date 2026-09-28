package wallet

import "testing"

// TestApplyDeposits — ключовий тест Кроку 1. Якщо реалізація
// ApplyDeposits наївно використовує `for _, w := range wallets`,
// зміни губляться (мутують лише копії циклу), і цей тест провалиться
// — точно так, як провалився б реальний код із заняття.
func TestApplyDeposits(t *testing.T) {
	wallets := []SecureWallet{{balance: 100}, {balance: 200}}

	ApplyDeposits(wallets, 50)

	if got := wallets[0].Balance(); got != 150 {
		t.Errorf("wallets[0].Balance() = %v, want 150", got)
	}
	if got := wallets[1].Balance(); got != 250 {
		t.Errorf("wallets[1].Balance() = %v, want 250", got)
	}
}

func TestApplyDeposits_EmptySlice(t *testing.T) {
	wallets := []SecureWallet{}
	ApplyDeposits(wallets, 50) // не повинно панікувати
	if len(wallets) != 0 {
		t.Errorf("len(wallets) = %d, want 0", len(wallets))
	}
}

func TestApplyDeposits_MultipleCalls(t *testing.T) {
	wallets := []SecureWallet{{balance: 0}}

	ApplyDeposits(wallets, 10)
	ApplyDeposits(wallets, 15)

	if got := wallets[0].Balance(); got != 25 {
		t.Errorf("wallets[0].Balance() = %v, want 25", got)
	}
}
