package product

type ProductRespository interface {
	Save(product *Product) error
	FindAll() ([]Product, error)
}