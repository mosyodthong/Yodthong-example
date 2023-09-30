package service

import (
	"errors"
	"order/repository"
	"strings"
)

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return orderService{orderRepo: orderRepo}
}

func (s orderService) CreateInventory(request InventoryRequest, OwnerId, TeamId int) (string, error) {

	qName := strings.ReplaceAll(request.InventoryName, " ", "") //ตัดช่องว่างออกทั้งหมด
	qCode := strings.ReplaceAll(request.Code, " ", "")          //ตัดช่องว่างออกทั้งหมด

	checkInventory, err := s.orderRepo.CheckInventory(qName, OwnerId, TeamId)

	if err != nil {
		return "", err
	}

	if checkInventory != nil { //เช็คชื่อ Inventory ซ้ำ
		return "", errors.New("มีชื่อคลังสินค้านี้แล้ว")
	}

	checkInventoryCode, err := s.orderRepo.CheckInventoryCode(qCode, OwnerId, TeamId)

	if err != nil {
		return "", err
	}

	if checkInventoryCode != nil { //เช็คชื่อ InventoryCode ซ้ำ
		return "", errors.New("มีโค้ดคลังสินค้านี้แล้ว")
	}

	masterAddress := strings.Split(request.Address_2, ">>")

	if len(masterAddress) < 4 {
		return "", errors.New("master_address is invalid")
	}

	// ตรวจสอบว่ามี " แขวง" นำหน้าหรือไม่
	if strings.HasPrefix(masterAddress[0], "แขวง") && masterAddress[2] == "กรุงเทพมหานคร" {
		// ถ้ามีให้ตัดออก
		masterAddress[0] = strings.TrimPrefix(masterAddress[0], "แขวง")
	}

	// ตรวจสอบว่ามี " เขต" นำหน้าหรือไม่
	if strings.HasPrefix(masterAddress[1], "เขต") && masterAddress[2] == "กรุงเทพมหานคร" {
		// ถ้ามีให้ตัดออก
		masterAddress[1] = strings.TrimPrefix(masterAddress[1], "เขต")
	}

	inventory := repository.Inventory{
		Code:          request.Code,
		InventoryName: request.InventoryName,
		Name:          request.Name,
		Tel:           request.Tel,
		Address_1:     request.Address_1,
		SubDistrict:   masterAddress[0],
		District:      masterAddress[1],
		Province:      masterAddress[2],
		Zipcode:       masterAddress[3],
		OwnerTeam:     int64(OwnerId),
		TeamID:        int64(TeamId),
	}

	message, err := s.orderRepo.CreateInventoryRepo(inventory)
	if err != nil {
		return "", err
	}

	return message, nil

}

func (s orderService) GetListInventory(id, OwnerId, TeamId int) (*InventoryByIdResponse, error) {

	response, err := s.orderRepo.GetByIdInventory(id, OwnerId, TeamId)
	if err != nil {
		return nil, err
	}

	inventoryResponse := InventoryByIdResponse{
		ID:            int(response.ID),
		Code:          response.Code,
		InventoryName: response.InventoryName,
		Name:          response.Name,
		Tel:           response.Tel,
		Address_1:     response.Address_1,
		Address_2:     response.SubDistrict + ">>" + response.District + ">>" + response.Province + ">>" + response.Zipcode,
	}

	return &inventoryResponse, nil

}

func (s orderService) DeleteInventory(id, OwnerId, TeamId int) (string, error) {

	message, err := s.orderRepo.DeleteInventoryRepo(id, OwnerId, TeamId)
	if err != nil {
		return "", err
	}

	return message, nil

}

func (s orderService) UpdateInventory(request InventoryRequest, id, OwnerId, TeamId int) (string, error) {

	qName := strings.ReplaceAll(request.InventoryName, " ", "") //ตัดช่องว่างออกทั้งหมด
	qCode := strings.ReplaceAll(request.Code, " ", "")          //ตัดช่องว่างออกทั้งหมด

	checkInventory, err := s.orderRepo.CheckInventory(qName, OwnerId, TeamId)

	if err != nil {
		return "", err
	}

	if checkInventory != nil && checkInventory.ID != uint(id) { //เช็คชื่อ Inventory ซ้ำ
		return "", errors.New("มีชื่อคลังสินค้านี้แล้ว")
	}

	checkInventoryCode, err := s.orderRepo.CheckInventoryCode(qCode, OwnerId, TeamId)

	if err != nil {
		return "", err
	}

	if checkInventoryCode != nil && checkInventoryCode.ID != uint(id) { //เช็คชื่อ InventoryCode ซ้ำ
		return "", errors.New("มีโค้ดคลังสินค้านี้แล้ว")
	}

	masterAddress := strings.Split(request.Address_2, ">>")

	if len(masterAddress) < 4 {
		return "", errors.New("master_address is invalid")
	}

	Inventory := repository.Inventory{
		Code:          request.Code,
		InventoryName: request.InventoryName,
		Name:          request.Name,
		Tel:           request.Tel,
		Address_1:     request.Address_1,
		SubDistrict:   masterAddress[0],
		District:      masterAddress[1],
		Province:      masterAddress[2],
		Zipcode:       masterAddress[3],
		OwnerTeam:     int64(OwnerId),
		TeamID:        int64(TeamId),
	}

	message, err := s.orderRepo.UpdateInventoryRepo(Inventory, id, OwnerId, TeamId)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (s orderService) GetListAllInventory(params ParamsFilterInventory) (*DataInventoryResponse, error) {

	qName := strings.ReplaceAll(params.Q, " ", "") //ตัดช่องว่างออกทั้งหมด

	filter := repository.FilterGetAllInventory{
		Q:       qName,
		Limit:   params.Limit,
		Page:    params.Page,
		OwnerId: params.OwnerId,
		TeamId:  params.TeamId,
	}

	all_status, pagination, err := s.orderRepo.GetInventory(filter)
	if err != nil {
		return nil, err
	}

	inventoryResponses := []InventoryResponse{}
	for _, status := range all_status {

		subdistrict := status.SubDistrict
		district := status.District
		if status.Province == "กรุงเทพมหานคร" {
			subdistrict = "แขวง" + subdistrict
			district = "เขต" + district
		} else {
			subdistrict = "ต." + subdistrict
			district = "อ." + district
		}

		inventoryResponse := InventoryResponse{
			ID:            int(status.ID),
			Code:          status.Code,
			InventoryName: status.InventoryName,
			Name:          status.Name,
			Tel:           status.Tel,
			Address: status.Address_1 + " " +
				subdistrict + " " +
				district + " จ." +
				status.Province + " " +
				status.Zipcode,
			UpdateLast: status.UpdatedAt.Format("02/01/2006 15:04"),
		}
		inventoryResponses = append(inventoryResponses, inventoryResponse)

	}

	paginationResponse := PaginationResponse{
		Page:      pagination.Page,
		TotalRow:  pagination.TotalRow,
		TotalPage: pagination.TotalPage,
	}

	invenResponse := DataInventoryResponse{
		Data:       inventoryResponses,
		Pagination: paginationResponse,
	}

	return &invenResponse, nil
}
