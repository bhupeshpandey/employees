package model

type AppConfig struct {
	ServerConfig *ServerConfig `json:"serverConfig"`
	CacheConfig  *CacheConfig `json:"cacheConfig"`
	DBConfig     *DBConfig `json:"dbConfig"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
}

type DBConfig struct {
	Env string `json:"env"`// defines the env for DB (Testing, Development, Production)
	Host string `json:"host"`
	Port int `json:"port"`
	// TODO Add user name and password going forward.
}

type CacheConfig struct {
	Size int `json:"capacity"`
}

type Cache interface {
	Get(id string) interface{}
	Put(id string, value interface{})
}

//type SearchOption func() (string, interface{})

type DB interface {
	AddEmployee(employee Employee) (Employee, error)     // POST
	UpdateEmployee(employee Employee) (bool, error)      // PUT
	DeleteEmployee(employee Employee) (Employee, error)  // DELETE
	RestoreEmployee(employee Employee) (Employee, error) // PUT
    SearchEmployees(search string) ([]Employee, error) // GET
    ListEmployees(search string) ([]Employee, error)   // GET

	//GetCarWithId(id string) (Employee, error)
	//GetAccessoryWithId(id string) (Employee, error)
	//GetAccessories()([]Employee, error)
	//GetCars()([]Employee, error)
	//GetRecentItems()([]Employee, error)
	//SetSelectedItem(product Employee) (bool, error)
}


type Employee struct {
	ID         string `json:"empId,omitempty" bson:"empId,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Department string `json:"department" bson:"department,omitempty`
	Address    string `json:"address,omitempty" bson:"address,omitempty`
	Skills     string `json:"skills" bson:"skills,omitempty"`
	Status     string `json:"status" bson:"status"`
}

type EmployeeCount struct {
	Count int `json:"employeeCount" bson:"employeeCount"`
}
