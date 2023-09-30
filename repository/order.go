package repository

import (
	"gorm.io/gorm"
)

type Inventory struct {
	ID            uint   `gorm:"primaryKey;comment:PK"`
	Code          string `gorm:"size:100;not null;comment:รหัสคลังสินค้า"`
	InventoryName string `gorm:"size:100;not null;comment:ชื่อคลังสินค้า"`
	Name          string `gorm:"size:100;not null;comment:ชื่อผู้ส่ง"`
	Tel           string `gorm:"size:100;not null;comment:เบอร์โทรผู้ส่ง"`
	Address_1     string `gorm:"size:100;not null;comment:ที่อยู่ผู้ส่ง"`
	SubDistrict   string `gorm:"size:50;not null;comment:ตำบล/แขวง"`
	District      string `gorm:"size:50;not null;comment:อำเภอ/เขต"`
	Province      string `gorm:"size:50;not null;comment:จังหวัด"`
	Zipcode       string `gorm:"size:20;not null;comment:รหัสไปรษณีย์"`
	OwnerTeam     int64  `gorm:"not null;comment:id ของ owner team นั้นๆ"`
	TeamID        int64  `gorm:"comment:id ของ  team นั้นๆ"`
	gorm.Model
}

type FilterGetAllInventory struct {
	Q       string
	Limit   int
	Page    int
	OwnerId int
	TeamId  int
}

type Pagination struct {
	Page      int
	TotalRow  int64
	TotalPage int
}

type OrderRepository interface {

	//Inventory
	CreateInventoryRepo(Inventory) (string, error)
	CheckInventory(string, int, int) (*Inventory, error)
	CheckInventoryCode(string, int, int) (*Inventory, error)
	GetInventory(FilterGetAllInventory) ([]Inventory, *Pagination, error)
	GetByIdInventory(int, int, int) (*Inventory, error)
	UpdateInventoryRepo(Inventory, int, int, int) (string, error)
	DeleteInventoryRepo(int, int, int) (string, error)
}
