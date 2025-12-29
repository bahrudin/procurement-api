package requests

type SupplierCreateRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type SupplierUpdateRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
