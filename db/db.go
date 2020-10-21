package db

import (
	"context"
	model "github.com/bhupeshpandey/employees/model"
	"strings"
)

type Count interface {
	GenEmployeeNumber() int
}

type DB struct {
	db model.DB
}

func New(config *model.DBConfig) (model.DB, error) {

	var db model.DB
	var err error
	if strings.EqualFold(config.Env, "Testing") {
		db = newMockDB()
	} else {
		db, err = newMongoDB(config, context.Background())

		if err != nil {
			return nil, err
		}
	}
	return &DB{
		db,
	}, nil


}

func (db *DB) AddEmployee(employee model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
} // POST

func (db *DB) UpdateEmployee(employee model.Employee) (bool, error) {
	return false, nil
} // PUT
func (db *DB) DeleteEmployee(employee model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
} // DELETE
func (db *DB) RestoreEmployee(employee model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
} // PUT
func (db *DB) SearchEmployees(search string) ([]model.Employee, error) {
	return nil, nil
} // GET
func (db *DB) ListEmployees(search string) ([]model.Employee, error) {
	return nil, nil
} // GET

func (db *DB) GenEmployeeNumber() int {
	return db.db.(Count).GenEmployeeNumber()
}

//func (db DB) GetProducts() ([]model.Employee, error) {
//	return db.db.GetProducts()
//}
//
//func (db DB) GetCarWithId(id string) (model.Employee, error) {
//	return db.db.GetCarWithId(id)
//}
//
//func (db DB) GetAccessoryWithId(id string) (model.Employee, error) {
//	return db.db.GetAccessoryWithId(id)
//}
//
//func (db DB) GetAccessories() ([]model.Employee, error) {
//	return db.db.GetAccessories()
//}
//func (db DB) GetCars() ([]model.Employee, error) {
//	return db.db.GetCars()
//}
//func (db DB) GetRecentItems() ([]model.Employee, error) {
//	return db.db.GetRecentItems()
//}
//func (db DB) SetSelectedItem(product model.Employee) (bool, error) {
//	return db.db.SetSelectedItem(product)
//}
