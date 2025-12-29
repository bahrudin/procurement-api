package requests

type ItemCreateRequest struct {
	SKU  string `json:"sku" validate:"required,max=50"`
	Name string `json:"name" validate:"required"`
	//Price float64 `json:"price" validate:"required,gt=0"`
	Price string `json:"price" validate:"required"`
	Stock int    `json:"stock" validate:"gte=0"`
	Unit  string `json:"unit" validate:"omitempty,max=20"`
}

type ItemUpdateRequest struct {
	SKU  *string `json:"sku" validate:"omitempty,max=50"`
	Name string  `json:"name"`
	//Price float64 `json:"price" validate:"omitempty,gt=0"`
	Price *string `json:"price" validate:"omitempty"`
	Stock int     `json:"stock" validate:"omitempty,gte=0"`
	Unit  *string `json:"unit" validate:"omitempty,max=20"`
}
