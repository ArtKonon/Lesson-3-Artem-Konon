package orders

import (
	"errors"
	"testing"
)

// fakeStore — просунутий фейк (fake), що записує кожен викликання
// Exec. Це саме той тип мок-об'єкта, який мав би згенерувати ШІ-
// асистент у Розділі 2.3 — тут він наданий готовим, щоб автотести
// працювали незалежно від вашого власного промпту.
type fakeStore struct {
	calls     int
	lastQuery string
	lastArgs  []any
	err       error
}

func (f *fakeStore) Exec(query string, args ...any) error {
	f.calls++
	f.lastQuery = query
	f.lastArgs = args
	return f.err
}

func TestPlaceOrder_Success(t *testing.T) {
	store := &fakeStore{}
	svc := NewOrderService(store)

	if err := svc.PlaceOrder("order-1", 42.5); err != nil {
		t.Fatalf("PlaceOrder повернув неочікувану помилку: %v", err)
	}

	if store.calls != 1 {
		t.Fatalf("Exec викликано %d разів, очікувалось 1", store.calls)
	}
	if store.lastQuery == "" {
		t.Error("Exec отримав порожній запит (query)")
	}
	if len(store.lastArgs) < 2 {
		t.Fatalf("Exec отримав %d аргументів, очікувалось щонайменше 2 (orderID, amount)", len(store.lastArgs))
	}
}

func TestPlaceOrder_PropagatesError(t *testing.T) {
	store := &fakeStore{err: errors.New("db is down")}
	svc := NewOrderService(store)

	if err := svc.PlaceOrder("order-2", 10); err == nil {
		t.Fatal("PlaceOrder мав повернути помилку, коли store.Exec її повертає")
	}
}

// TestOrderServiceUsesInjectedStore перевіряє, що OrderService
// дійсно використовує ІНТЕРФЕЙС, а не конкретний тип — головна
// мета рефакторингу з Розділу 2.
func TestOrderServiceUsesInjectedStore(t *testing.T) {
	var _ OrderStore = &fakeStore{} // fakeStore має задовольняти OrderStore

	store := &fakeStore{}
	svc := NewOrderService(store)
	if svc == nil {
		t.Fatal("NewOrderService повернув nil")
	}

	_ = svc.PlaceOrder("order-3", 5)
	if store.calls == 0 {
		t.Error("PlaceOrder не викликав store.Exec — перевірте, що сервіс справді використовує injected store")
	}
}
