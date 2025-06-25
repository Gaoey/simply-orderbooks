package orderbooks

type Order struct {
	Side     string // "BUY" or "SELL"
	Price    int    // Price of the order
	Quantity int    // Quantity of the order
}

type OrderBook struct {
	orders []Order
}

func NewOrderBook() OrderBook {
	return OrderBook{
		orders: []Order{},
	}
}

func (ob *OrderBook) AddOrder(side string, price int, quantity int) {
	order := Order{
		Side:     side,
		Price:    price,
		Quantity: quantity,
	}
	ob.orders = append(ob.orders, order)
}

func (ob *OrderBook) GetAll() []Order {
	return ob.orders
}
