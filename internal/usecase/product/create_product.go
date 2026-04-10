package product

import (
	domain "go-pos/internal/domain/product"
)

type CreateProductUsecase struct {
	repo domain.Repository
}

func NewCreateProductUsecase(repo domain.Repository) *CreateProductUsecase {
	return &CreateProductUsecase{repo: repo}
}

func (uc *CreateProductUsecase) Execute (id, name string, stock, price int) error {
	product, err := domain.NewProduct(id, name, price, stock)
	if err != nil {
		return err
	}

	return uc.repo.Save(product)
}