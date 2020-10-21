package db

import "github.com/bhupeshpandey/employees/model"

type mockDB struct {
}

func newMockDB() *mockDB {
	return &mockDB {}
}

func (db *mockDB) AddEmployee(employee model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
} // POST

func (db *mockDB) UpdateEmployee(employee model.Employee) (bool, error) {
	return false, nil
} // PUT
func (db *mockDB) DeleteEmployee(employee model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
} // DELETE
func (db *mockDB) RestoreEmployee(employee model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
} // PUT
func (db *mockDB) SearchEmployees(search string) ([]model.Employee, error) {
	return nil, nil
} // GET
func (db *mockDB) ListEmployees(search string) ([]model.Employee, error) {
	return nil, nil
} // GET

func (db *mockDB) GenEmployeeNumber() int {
	return -777
}

//func (mockdb *mockDB) GetProducts() ([]model.Employee, error) {
//	return nil, nil
//}
//
//func (mockdb *mockDB) GetCarWithId(id string) (model.Employee, error) {
//	return model.Employee{}, nil
//}
//
//func (mockdb *mockDB) GetAccessoryWithId(id string) (model.Employee, error) {
//	return model.Employee{}, nil
//}
//func (mockdb *mockDB) GetAccessories() ([]model.Employee, error) {
//	return nil, nil
//}
//func (mockdb *mockDB) GetCars() ([]model.Employee, error) {
//	return nil, nil
//}
//func (mockdb *mockDB) GetRecentItems() ([]model.Employee, error) {
//	return nil, nil
//}
//func (mockdb *mockDB) SetSelectedItem(product model.Employee) (bool, error) {
//	return false, nil
//}
