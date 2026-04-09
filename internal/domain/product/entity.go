package product

import "errors"

type Product struct {
	id	 string
	name string
	price int
	stock int
}

func NewProduct(id, name string, price, stock int) (*Product, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	if stock < 0  {
		return nil, errors.New("stock cannot be negative")
	}

	return &Product{
		id: id,
		name: name,
		stock: stock,
		price: price,
	}, nil
}

func (p *Product) ReduceStock(qty int) error {
	if qty <= 0 {
		return errors.New("qty must be greater than 0")
	}
	if qty > p.stock {
		return errors.New("stock not enough")
	}

	p.stock -= qty
	return nil
}