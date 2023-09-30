package service

type InventoryByIdResponse struct {
	ID            int    `json:"id"`
	Code          string `json:"code"`
	InventoryName string `json:"inventory_name"`
	Name          string `json:"name"`
	Tel           string `json:"tel"`
	Address_1     string `json:"address_1"`
	Address_2     string `json:"address_2"`
}

type InventoryRequest struct {
	Code          string `json:"code"`
	InventoryName string `json:"inventory_name"`
	Name          string `json:"name"`
	Tel           string `json:"tel"`
	Address_1     string `json:"address_1"`
	Address_2     string `json:"address_2"`
}

type InventoryResponse struct {
	ID            int    `json:"id"`
	Code          string `json:"code"`
	InventoryName string `json:"inventory_name"`
	Name          string `json:"name"`
	Tel           string `json:"tel"`
	Address       string `json:"address"`
	UpdateLast    string `json:"update_last"`
}

type DataInventoryResponse struct {
	Pagination PaginationResponse  `json:"pagination"`
	Data       []InventoryResponse `json:"inventory"`
}

type ParamsFilterInventory struct {
	Q       string
	Limit   int
	Page    int
	OwnerId int
	TeamId  int
}

type PaginationResponse struct {
	Page      int   `json:"page"`
	TotalRow  int64 `json:"total_row"`
	TotalPage int   `json:"total_page"`
}

type OrderService interface {
	//Inventory
	GetListAllInventory(ParamsFilterInventory) (*DataInventoryResponse, error)
	GetListInventory(int, int, int) (*InventoryByIdResponse, error)
	CreateInventory(InventoryRequest, int, int) (string, error)
	DeleteInventory(int, int, int) (string, error)
	UpdateInventory(InventoryRequest, int, int, int) (string, error)
}
