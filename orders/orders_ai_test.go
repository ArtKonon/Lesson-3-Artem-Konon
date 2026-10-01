package orders

import (
	"errors"
	"reflect"
	"testing"
)

// ExtendedMockStore зберігає повну історію викликів для глибокої перевірки.
type ExtendedMockStore struct {
	ErrToReturn error
	CallsCount  int
	LastQuery   string
	LastArgs    []any
}

// Exec імітує роботу БД та записує всі вхідні дані для подальших перевірок у тестах.
func (m *ExtendedMockStore) Exec(query string, args ...any) error {
	m.CallsCount++
	m.LastQuery = query
	m.LastArgs = args
	return m.ErrToReturn
}

func TestOrderService_PlaceOrder_Extended(t *testing.T) {
	// Створюємо стандартну помилку для тестування збоїв
	dbErr := errors.New("timeout connection to database")

	// Таблиця тестових сценаріїв
	tests := []struct {
		name          string
		orderID       string
		amount        float64
		mockError     error
		expectedError error
	}{
		{
			name:          "Успішне збереження стандартного замовлення",
			orderID:       "ORD-12345",
			amount:        99.99,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Помилка на рівні бази даних",
			orderID:       "ORD-99999",
			amount:        150.00,
			mockError:     dbErr,
			expectedError: dbErr,
		},
		{
			name:          "Граничний випадок: нульова сума замовлення",
			orderID:       "ORD-00001",
			amount:        0.0,
			mockError:     nil,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange (Підготовка)
			mockStore := &ExtendedMockStore{
				ErrToReturn: tt.mockError,
			}
			service := NewOrderService(mockStore)

			// Act (Дія)
			err := service.PlaceOrder(tt.orderID, tt.amount)

			// Assert (Перевірки)

			// 1. Перевірка результату (помилки)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("Очікувалася помилка: %v, отримано: %v", tt.expectedError, err)
			}

			// 2. Перевірка кількості викликів (чи не викликає сервіс Exec двічі?)
			if mockStore.CallsCount != 1 {
				t.Errorf("Очікувався 1 виклик Exec, але отримано: %d", mockStore.CallsCount)
			}

			// 3. Перевірка правильності SQL-запиту
			expectedQuery := "INSERT INTO orders (id, amount) VALUES (?, ?)"
			if mockStore.LastQuery != expectedQuery {
				t.Errorf("Очікувався SQL-запит: %q\nОтримано: %q", expectedQuery, mockStore.LastQuery)
			}

			// 4. Глибока перевірка переданих аргументів
			expectedArgs := []any{tt.orderID, tt.amount}
			if !reflect.DeepEqual(mockStore.LastArgs, expectedArgs) {
				t.Errorf("Очікувалися аргументи: %v, отримано: %v", expectedArgs, mockStore.LastArgs)
			}
		})
	}
}
