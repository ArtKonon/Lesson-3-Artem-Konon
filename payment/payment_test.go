package payment

import (
	"strings"
	"testing"
	"time"
)

// TestTransactionLogInfo — базова перевірка того, що LogInfo вже
// працює правильно (він наданий готовим, але тест підтверджує, що
// ви не зламали його випадково).
func TestTransactionLogInfo(t *testing.T) {
	ts := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	tx := Transaction{ID: "tx-1", Timestamp: ts}

	got := tx.LogInfo()
	if got == "" {
		t.Fatal("LogInfo() повернув порожній рядок")
	}
	if !strings.Contains(got, "tx-1") {
		t.Errorf("LogInfo() = %q, має містити ID %q", got, "tx-1")
	}
}

// TestCreditCardFields перевіряє Завдання 1.3: наявність полів
// CardNumber і Balance, і що Transaction справді вбудована (а не
// іменована окремим полем).
func TestCreditCardFields(t *testing.T) {
	card := CreditCard{
		Transaction: Transaction{ID: "c-1", Timestamp: time.Now()},
		CardNumber:  "4242-4242-4242-4242",
		Balance:     100,
	}
	// Просунуте поле ID має бути доступне напряму, без card.Transaction.ID
	if card.ID != "c-1" {
		t.Errorf("card.ID = %q, want %q (перевірте, що Transaction вбудована, а не іменована)", card.ID, "c-1")
	}
	if card.CardNumber != "4242-4242-4242-4242" {
		t.Errorf("card.CardNumber = %q, want %q", card.CardNumber, "4242-4242-4242-4242")
	}
	if card.Balance != 100 {
		t.Errorf("card.Balance = %v, want 100", card.Balance)
	}
}

// TestCryptoWalletFields — те саме для CryptoWallet.
func TestCryptoWalletFields(t *testing.T) {
	wallet := CryptoWallet{
		Transaction:   Transaction{ID: "w-1", Timestamp: time.Now()},
		WalletAddress: "0xABCDEF",
		TokenBalance:  250,
	}
	if wallet.ID != "w-1" {
		t.Errorf("wallet.ID = %q, want %q", wallet.ID, "w-1")
	}
	if wallet.WalletAddress != "0xABCDEF" {
		t.Errorf("wallet.WalletAddress = %q, want %q", wallet.WalletAddress, "0xABCDEF")
	}
	if wallet.TokenBalance != 250 {
		t.Errorf("wallet.TokenBalance = %v, want 250", wallet.TokenBalance)
	}
}

// TestCreditCardPayMutatesBalance — ключовий тест Завдання 1.4.
//
// Якщо Pay() залишено з приймачем-значенням, ця мутація НЕ
// збережеться в оригінальній структурі — тест провалиться саме так,
// як провалився б реальний баг із заняття.
func TestCreditCardPayMutatesBalance(t *testing.T) {
	card := &CreditCard{Balance: 100}

	if err := card.Pay(30); err != nil {
		t.Fatalf("Pay(30) повернув неочікувану помилку: %v", err)
	}
	if card.Balance != 70 {
		t.Errorf("card.Balance = %v, want 70 (перевірте, що Pay має приймач-вказівник)", card.Balance)
	}
}

func TestCreditCardPayInsufficientBalance(t *testing.T) {
	card := &CreditCard{Balance: 10}

	if err := card.Pay(50); err == nil {
		t.Error("Pay(50) не повернув помилку при недостатньому балансі")
	}
	if card.Balance != 10 {
		t.Errorf("card.Balance = %v, want 10 (баланс не має змінюватися при помилці)", card.Balance)
	}
}

// TestCryptoWalletPayMutatesBalance — той самий ключовий тест для
// CryptoWallet.TokenBalance.
func TestCryptoWalletPayMutatesBalance(t *testing.T) {
	wallet := &CryptoWallet{TokenBalance: 500}

	if err := wallet.Pay(200); err != nil {
		t.Fatalf("Pay(200) повернув неочікувану помилку: %v", err)
	}
	if wallet.TokenBalance != 300 {
		t.Errorf("wallet.TokenBalance = %v, want 300 (перевірте, що Pay має приймач-вказівник)", wallet.TokenBalance)
	}
}

func TestCryptoWalletPayInsufficientBalance(t *testing.T) {
	wallet := &CryptoWallet{TokenBalance: 5}

	if err := wallet.Pay(999); err == nil {
		t.Error("Pay(999) не повернув помилку при недостатньому балансі")
	}
}

// TestProcessPaymentPolymorphic — Завдання 1.5: одна функція
// працює однаково і з CreditCard, і з CryptoWallet через інтерфейс
// PaymentMethod. Обидва передаються як ВКАЗІВНИКИ — саме так, як
// це буде в реальному коді з pointer receiver.
func TestProcessPaymentPolymorphic(t *testing.T) {
	methods := []PaymentMethod{
		&CreditCard{
			Transaction: Transaction{ID: "poly-c", Timestamp: time.Now()},
			Balance:     100,
		},
		&CryptoWallet{
			Transaction:  Transaction{ID: "poly-w", Timestamp: time.Now()},
			TokenBalance: 100,
		},
	}

	for _, m := range methods {
		if err := ProcessPayment(m, 20); err != nil {
			t.Errorf("ProcessPayment(%T, 20) повернув неочікувану помилку: %v", m, err)
		}
	}
}
