package product

import (
	domain "go-pos/internal/domain/product"
)

type GetProductUsecase struct {
	repo domain.Repository
}

func NewGetProductUsecase(repo domain.Repository) *GetProductUsecase {
	return &GetProductUsecase{repo: repo}
}

func (uc *GetProductUsecase) Execute() ([]domain.Product, error) {
	return uc.repo.FindAll()
}