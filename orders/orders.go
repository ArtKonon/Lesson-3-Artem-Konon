package orders

// OrderStore is the minimal dependency needed by the service.
type OrderStore interface {
	Exec(query string, args ...any) error
}

// OrderService stores the injected dependency and performs order operations.
type OrderService struct {
	store OrderStore
}

// NewOrderService creates a service with an injected store.
func NewOrderService(store OrderStore) *OrderService {
	return &OrderService{store: store}
}

// PlaceOrder inserts a new order using the injected store.
func (s *OrderService) PlaceOrder(orderID string, amount float64) error {
	err := s.store.Exec("INSERT INTO orders (id, amount) VALUES (?, ?)", orderID, amount)
	if err != nil {
		return err
	}
	return nil
}
