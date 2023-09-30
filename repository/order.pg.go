package repository

import (
	"errors"
	"math"
	"strings"

	"gorm.io/gorm"
)

type orderRepositoryPG struct {
	db *gorm.DB
}

func NewOrderRepositoryPG(db *gorm.DB) OrderRepository {

	db.AutoMigrate(Inventory{})

	return orderRepositoryPG{db: db}
}

func (r orderRepositoryPG) GetInventory(filter FilterGetAllInventory) ([]Inventory, *Pagination, error) {
	inventory := []Inventory{}

	if filter.Page == 0 {
		filter.Page = 1
	}

	//จำนวนข้อมูล
	var count int64

	//แสดงข้อมูลในหนึ่งหน้า
	if filter.Limit == 0 {
		r.db.Model(&inventory).Where("owner_team = ? AND team_id = ? ", filter.OwnerId, filter.TeamId).Count(&count)
		filter.Limit = int(count)
	}

	if filter.Q != "" {
		tx := r.db.Where("owner_team = ? AND team_id = ? AND LOWER(REPLACE(inventory_name, ' ', '')) LIKE ?   ", filter.OwnerId, filter.TeamId, "%"+strings.ToLower(filter.Q)+"%").
			Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).
			Order("id asc").Find(&inventory)

		if tx.Error != nil {
			return nil, nil, tx.Error
		}

		//หาจำนวนข้อมูล
		tc := r.db.Model(&Inventory{}).Where("owner_team = ? AND team_id = ? AND LOWER(REPLACE(inventory_name, ' ', '')) LIKE ?   ", filter.OwnerId, filter.TeamId, "%"+strings.ToLower(filter.Q)+"%").Order("id asc").Count(&count)
		if tc.Error != nil {
			return nil, nil, tc.Error
		}

	} else {

		tx := r.db.Where("owner_team = ? AND team_id = ?", filter.OwnerId, filter.TeamId).
			Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).
			Order("id asc").Find(&inventory)

		if tx.Error != nil {
			return nil, nil, tx.Error
		}

		//หาจำนวนข้อมูล
		tc := r.db.Model(&Inventory{}).
			Where("owner_team = ? AND team_id = ?", filter.OwnerId, filter.TeamId).
			Count(&count)
		if tc.Error != nil {
			return nil, nil, tc.Error
		}
	}

	count_page := math.Ceil(float64(count) / float64(filter.Limit))
	if int(count_page) < 0 {
		count_page = 0
	}
	pagination := Pagination{
		Page:      filter.Page,
		TotalRow:  count,
		TotalPage: int(count_page),
	}

	return inventory, &pagination, nil
}

func (r orderRepositoryPG) GetByIdInventory(id, OwnerId, TeamId int) (*Inventory, error) {
	inventory := Inventory{}

	tx := r.db.Where("owner_team = ? AND team_id = ?  AND id = ?", OwnerId, TeamId, id).First(&inventory, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &inventory, nil
}

func (r orderRepositoryPG) CheckInventory(nameInventory string, OwnerId, TeamId int) (*Inventory, error) {
	inventory := Inventory{}

	tx := r.db.Where("(owner_team = ? AND team_id = ?  AND REPLACE(inventory_name, ' ', '') = ?)", OwnerId, TeamId, nameInventory).First(&inventory)

	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &inventory, nil
}

func (r orderRepositoryPG) CheckInventoryCode(inventoryCode string, OwnerId, TeamId int) (*Inventory, error) {
	inventory := Inventory{}

	tx := r.db.Where("(owner_team = ? AND team_id = ?  AND REPLACE(code, ' ', '') = ?)", OwnerId, TeamId, inventoryCode).First(&inventory)

	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &inventory, nil
}

func (r orderRepositoryPG) CreateInventoryRepo(inventory Inventory) (string, error) {

	tx := r.db.Create(&inventory)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("data is not creating")
	}
	return "Created.", nil

}

func (r orderRepositoryPG) DeleteInventoryRepo(id, OwnerId, TeamId int) (string, error) {

	//ลบข้อมูลตาราง Order
	tx := r.db.Where("owner_team = ? AND team_id = ? ", OwnerId, TeamId).Delete(&Inventory{}, id)
	if tx.Error != nil {
		return "", tx.Error
	}
	if tx.RowsAffected != 1 {
		return "", errors.New("data is not deleting")
	}

	return "Deleted.", nil

}

func (r orderRepositoryPG) UpdateInventoryRepo(inventory Inventory, id, OwnerId, TeamId int) (string, error) {

	tx := r.db.Model(&Inventory{}).Where("owner_team = ? AND team_id = ? AND id = ?", OwnerId, TeamId, id).Updates(&inventory)

	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("data is not update")
	}

	return "success.", nil

}
