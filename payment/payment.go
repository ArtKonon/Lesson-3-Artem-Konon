// Package payment реалізує Розділ 1 домашньої роботи: платіжну
// систему з малим інтерфейсом PaymentMethod, композицією через
// вбудовування Transaction, і поліморфну обробку платежів.
//
// Примітка щодо дизайну: оригінальне завдання просить ProcessPayment
// виводити "просунуті" (promoted) поля ID і Timestamp. Оскільки
// інтерфейс PaymentMethod не може напряму "бачити" поля конкретного
// типу, ми додали до інтерфейсу другий метод LogInfo() — і ось
// цікавий момент: вам НЕ потрібно писати LogInfo() для CreditCard чи
// CryptoWallet окремо. Він уже реалізований один раз на Transaction
// (нижче) і автоматично "просувається" на обидва типи завдяки
// вбудовуванню. Це ілюструє саме те просування методів, яке ми
// проходили на занятті.
package payment

import (
	"errors"
	"fmt"
	"time"
)

// Transaction містить спільні поля для будь-якого платежу.
type Transaction struct {
	ID        string
	Timestamp time.Time
}

// LogInfo повертає відформатований рядок з даними транзакції.
// Уже реалізовано — просувається на CreditCard і CryptoWallet
// автоматично через вбудовування.
func (t Transaction) LogInfo() string {
	return fmt.Sprintf("ID: %s, Timestamp: %s", t.ID, t.Timestamp.Format(time.RFC3339))
}

// PaymentMethod — малий, сфокусований інтерфейс (Завдання 1.1).
type PaymentMethod interface {
	Pay(amount float64) error
	LogInfo() string
}

// CreditCard — Завдання 1.3.
//
// TODO: додайте поля:
//   - CardNumber string
//   - Balance    float64
type CreditCard struct {
	Transaction
	// TODO: ваші поля тут
}

// CryptoWallet — Завдання 1.3.
//
// TODO: додайте поля:
//   - WalletAddress string
//   - TokenBalance  float64
type CryptoWallet struct {
	Transaction
	// TODO: ваші поля тут
}

// Pay реалізує PaymentMethod для CreditCard (Завдання 1.4).
//
// Правила:
//   - якщо amount > c.Balance — поверніть помилку з текстом
//     "insufficient balance"
//   - інакше відніміть amount від c.Balance і поверніть nil
//
// TODO (Завдання 1.4): приймач тут має бути ВКАЗІВНИКОМ (*CreditCard).
// Метод змінює баланс — приймач-значення призведе до тієї самої
// "мовчазної" помилки, яку ми розбирали на занятті.
func (c CreditCard) Pay(amount float64) error {
	// TODO: ваш код тут
	return errors.New("not implemented")
}

// Pay реалізує PaymentMethod для CryptoWallet (Завдання 1.4).
// Правила ті самі, що і для CreditCard, але для TokenBalance.
//
// TODO (Завдання 1.4): приймач має бути ВКАЗІВНИКОМ (*CryptoWallet).
func (w CryptoWallet) Pay(amount float64) error {
	// TODO: ваш код тут
	return errors.New("not implemented")
}

// ProcessPayment — поліморфна функція (Завдання 1.5).
// Вона працює з БУДЬ-ЯКИМ типом, що задовольняє PaymentMethod, —
// не важливо, CreditCard це чи CryptoWallet.
func ProcessPayment(method PaymentMethod, amount float64) error {
	if err := method.Pay(amount); err != nil {
		return err
	}
	fmt.Printf("Payment successful. %s\n", method.LogInfo())
	return nil
}
