// Package orders реалізує Розділ 2 домашньої роботи: рефакторинг
// OrderService через виокремлення мінімального інтерфейсу OrderStore
// та впровадження залежності через конструктор.
//
// Саме цей пакет ви й попросите ШІ-асистента рефакторити (Розділ 2)
// та покрити тестами (Розділ 3, Крок 2). Файл orders_test.go вже
// містить автотести з фейковим store — використайте їх, щоб
// перевірити свою реалізацію локально, а власний промпт і згенерований
// ШІ тест додайте окремо (наприклад, у orders_ai_test.go) для звіту.
package orders

// OrderStore — мінімальний інтерфейс для залежності від БД
// (Завдання 2.1). Замінює прямий *sql.DB, дотримуючись приказки
// Роба Пайка: "чим більший інтерфейс, тим слабша абстракція".
type OrderStore interface {
	Exec(query string, args ...any) error
}

// OrderService — Завдання 2.2: сервіс приймає залежність через
// конструктор (dependency injection), а не створює її сам.
type OrderService struct {
	store OrderStore
}

// NewOrderService — конструктор із впровадженням залежності.
func NewOrderService(store OrderStore) *OrderService {
	return &OrderService{store: store}
}

// PlaceOrder виконує вставку замовлення через store.
//
// TODO (Завдання 2.2): реалізуйте цей метод.
//   - викличте s.store.Exec(...) із SQL-запитом і аргументами
//     orderID та amount;
//   - поверніть помилку, якщо Exec її повернув;
//   - інакше поверніть nil.
func (s *OrderService) PlaceOrder(orderID string, amount float64) error {
	// TODO: ваш код тут
	return nil
}
