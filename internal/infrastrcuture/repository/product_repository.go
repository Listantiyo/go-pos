package repository

type ProductModel struct {
	ID 	  string `gorm:"primaryKey"`
	Name  string
	Price int
	Stock int
}

func toDomain()