package server

import (
	"encoding/json"
	"fmt"
	"github.com/bhupeshpandey/employees/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	Config *model.ServerConfig
	Router *mux.Router
	Cache  model.Cache
	DB     model.DB
}

func New(serverConfig *model.ServerConfig, cache model.Cache, db model.DB) *Server {
	router := mux.NewRouter()

	server := &Server{
		Config: serverConfig,
		Cache:  cache,
		Router: router,
		DB: db,
	}

	// List All Employees
	router.HandleFunc("/api/list", server.ListEmployees).Methods("GET")
	router.HandleFunc("/api/list/{employeeId}", server.ListEmployees).Methods("GET")
	router.HandleFunc("/api/list/{department}", server.ListEmployees).Methods("GET")
	router.HandleFunc("/api/list/{name}", server.ListEmployees).Methods("GET")

	// Search All Employees
	router.HandleFunc("/api/list", server.SearchEmployees).Methods("GET")
	router.HandleFunc("/api/list/{address}", server.SearchEmployees).Methods("GET")
	router.HandleFunc("/api/list/{skills}", server.SearchEmployees).Methods("GET")
	router.HandleFunc("/api/list/{department}", server.SearchEmployees).Methods("GET")
	router.HandleFunc("/api/list/{name}", server.SearchEmployees).Methods("GET")

	// Add Employee
	router.HandleFunc("/api/add", server.AddEmployee).Methods("POST")

	// Update Employee
	router.HandleFunc("/api/update", server.UpdateEmployee).Methods("PUT")

	// Delete Employee
	router.HandleFunc("/api/delete", server.DeleteEmployee).Methods("DELETE")

	// Restore Employee
	router.HandleFunc("/api/restore", server.RestoreEmployee).Methods("PUT")

	//router.HandleFunc("/api/accessories", server.GetAccessories).Methods("GET")
	//router.HandleFunc("/api/books/{id}", server.GetAccessoryWithId).Methods("GET")
	//router.HandleFunc("/api/recent/items}", server.GetRecentItems).Methods("GET")
	//router.HandleFunc("/api/recent/items}", server.SetSelectedItem).Methods("PUT")
	// set our port address

	return server
}

func (Server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", Server.Config.Host, Server.Config.Port), Server.Router))
}

//func (ServerConfig *ServerConfig) GetEmployees()

//func (Server *Server) GetCars(w http.ResponseWriter, r *http.Request) {
//
//	//b, err := ioutil.ReadAll(r.Body)
//	//defer r.Body.Close()
//	//if err != nil {
//	//	http.Error(w, err.Error(), 500)
//	//	return
//	//}
//	fmt.Println(r.RequestURI)
//	uri := r.RequestURI
//	str := Server.Cache.Get(uri)
//
//	var cars interface{}
//	var err error
//	if str == "" {
//		cars, err = Server.DB.GetCars()
//
//		if err != nil {
//			w.WriteHeader(500)
//			w.Write([]byte(err.Error()))
//			return
//		}
//
//		if cars == nil {
//			w.WriteHeader(500)
//			w.Write([]byte("Nil entries in car"))
//			return
//		}
//		Server.Cache.Put(r.RequestURI, cars)
//	}
//	//Server.Cache.Get()
//	data, err := json.Marshal(cars.([]model.Employee))
//
//	if err != nil {
//		w.WriteHeader(500)
//		w.Write([]byte(err.Error()))
//	    return
//	}
//	w.Write(data)
//}

func (Server *Server) ListEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	//employeeID, exists := params["employeeId"]
	//
	//if exists {
	//
	//}
}

func (server *Server) SearchEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (server *Server) AddEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employeeData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	defer r.Body.Close()

	var emp model.Employee
	err = json.Unmarshal(employeeData, &emp)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	emp, err  = server.DB.AddEmployee(emp)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}



}

func (server *Server) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (Server *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (Server *Server) RestoreEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//func (Server *Server) GetAccessories(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func (Server *Server) GetAccessoryWithId(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func (Server *Server) GetRecentItems(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func (Server *Server) SetSelectedItem(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func (Server *Server) SetSelectedCar(w http.ResponseWriter, r *http.Request) {
//
//}
