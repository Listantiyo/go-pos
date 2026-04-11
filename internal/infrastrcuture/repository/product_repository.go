package repository

import (
	domain "go-pos/internal/domain/product"

	"gorm.io/gorm"
)

type ProductModel struct {
	ID 	  string `gorm:"primaryKey"`
	Name  string
	Price int
	Stock int
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.Repository {
	return &productRepository{db: db}
}

func (r *productRepository) Save(p *domain.Product) error {
	model := ProductModel{
		ID: p.ID(),
		Name: p.Name(),
		Stock: p.Stock(),
		Price: p.Price(),
	}

	return r.db.Create(&model).Error
}

func (r *productRepository) FindAll() ([]domain.Product, error) {
	var models []ProductModel

	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	var products []domain.Product

	for _, m := range models {
		p, err := toDomain(m)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func toDomain(m ProductModel) (domain.Product, error) {
	p, err := domain.NewProduct(
		m.ID,
		m.Name,
		m.Price,
		m.Stock,
	)

	if err != nil {
		return domain.Product{}, err
	}

	return *p, nil
}