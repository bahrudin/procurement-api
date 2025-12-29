package requests

type PurchaseItemRequest struct {
	ItemID uint `json:"item_id" validate:"required"`
	Qty    int  `json:"qty" validate:"required,gt=0"`
}

type PurchaseCreateRequest struct {
	SupplierID uint                  `json:"supplier_id" validate:"required"`
	Items      []PurchaseItemRequest `json:"items" validate:"required,min=1,dive"`
}
