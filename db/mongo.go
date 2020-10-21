package db

import (
	"context"
	"fmt"
	"github.com/bhupeshpandey/employees/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoDB struct {
	host                  string
	port                  int
	ctx                   context.Context
	client                *mongo.Client
	employees             *mongo.Collection
	employeeCount         *mongo.Collection
}

func newMongoDB(db *model.DBConfig, ctx context.Context) (*mongoDB, error) {
	ctxx, _ := context.WithCancel(ctx)
	mongoInst := &mongoDB{
		db.Host,
		db.Port,
		ctxx,
		nil,
		nil,
		nil,
	}

	if err := mongoInst.Init(); err != nil {
		return nil, err
	}
	return mongoInst, nil

}

func (mongoDB mongoDB) Init() error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%v/", mongoDB.host, mongoDB.port))
	client, err := mongo.Connect(mongoDB.ctx, clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(mongoDB.ctx, nil)
	if err != nil {
		return err
	}

	mongoDB.client = client

	database := client.Database("Company")

	mongoDB.employees = database.Collection("Employees")
	//mongoDB.accessoriesCollection = database.Collection("Accessories")

	return nil
}

func (mongoDB *mongoDB) AddEmployee(employee model.Employee) (model.Employee, error) {
	number := mongoDB.GenEmployeeNumber() + 1

	empID := fmt.Sprintf("EMP%v", number)

	employee.ID = empID

	_, err := mongoDB.employees.InsertOne(mongoDB.ctx, employee)

	if err != nil {
		return model.Employee{}, err
	}

	return employee, nil
} // POST

func (mongoDB *mongoDB) UpdateEmployee(employee model.Employee) (bool, error) {

	_, err := mongoDB.employees.UpdateOne(mongoDB.ctx, bson.M{"employeeId": employee.ID}, bson.D{
		{"$set", bson.D{{"name", employee.Name}, {"department", employee.Department}, {"address", employee.Address}, {"skills", employee.Skills}, {"status", employee.Status}}},
	})

	if err != nil {
		return false, err
	}
	return true, nil
} // PUT

func (mongoDB *mongoDB) DeleteEmployee(employee model.Employee) (model.Employee, error) {

	// Check for the flags to delete it permanently
	var err error
	// if perm flag true
	if true {
	    _, err = mongoDB.employees.DeleteOne(mongoDB.ctx, bson.M{"employeeId": employee.ID})
	} else {
		// Get the employee from db and set status to inactive
		employee.Status = "Inactive"
		_, err = mongoDB.employees.UpdateOne(mongoDB.ctx, bson.M{"employeeId": employee.ID}, bson.D{
			{"$set", bson.D{{"status", employee.Status}}},
		})
	}

	if err != nil {
		return model.Employee{}, err
	}

	return model.Employee{}, nil
} // DELETE

func (mongoDB *mongoDB) RestoreEmployee(employee model.Employee) (model.Employee, error) {
	// Get the employee from db and set status to inactive
	employee.Status = "Active"
	_, err := mongoDB.employees.UpdateOne(mongoDB.ctx, bson.M{"employeeId": employee.ID}, bson.D{
		{"$set", bson.D{{"status", employee.Status}}},
	})

	return employee, err
} // PUT

func (mongoDB *mongoDB) SearchEmployees(search string) ([]model.Employee, error) {

	query := bson.M{
		"$text": bson.M{
			"$search": search,
		},
	}

	cur, err := mongoDB.employees.Find(mongoDB.ctx, query)

	if err  != nil {
		return nil, err
	}

	var employees []model.Employee
	for cur.Next(mongoDB.ctx) {
		var e model.Employee
		err := cur.Decode(&e)
		if err != nil {
			fmt.Println("Could not decode Employee")
			return []model.Employee{}, err
		}
		employees = append(employees, e)
	}
	return employees, nil
} // GET

func (mongoDB *mongoDB) ListEmployees(search string) ([]model.Employee, error) {
	query := bson.M{
		"$text": bson.M{
			"$search": search,
		},
	}

	cur, err := mongoDB.employees.Find(mongoDB.ctx, query)

	if err  != nil {
		return nil, err
	}

	var employees []model.Employee
	for cur.Next(mongoDB.ctx) {
		var e model.Employee
		err := cur.Decode(&e)
		if err != nil {
			fmt.Println("Could not decode Employee")
			return []model.Employee{}, err
		}
		employees = append(employees, e)
	}
	return employees, nil
} // GET

//func (mongoDB *mongoDB) GetProducts() ([]model.Employee, error) {
//	return nil, nil
//}
//
//func (mongoDB *mongoDB) GetCarWithId(id string) (model.Employee, error) {
//	return model.Employee{}, nil
//}
//
//func (mongoDB *mongoDB) GetAccessoryWithId(id string) (model.Employee, error) {
//	return model.Employee{}, nil
//}
//func (mongoDB *mongoDB) GetAccessories() ([]model.Employee, error) {
//	return nil, nil
//}
func (mongoDB *mongoDB) GetCars() ([]model.Employee, error) {


	cur, err := mongoDB.employees.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	//var episodes []bson.M
	//if err = cur.All(mongoDB.ctx, &episodes); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Episodes ", episodes)

	defer cur.Close(context.Background())

	var products []model.Employee

	for cur.Next(mongoDB.ctx) {
		var episode bson.M
		var product = model.Employee{}
		if err = cur.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		//episode.
		//fmt.Println(episode)

		//marshal, _ := bson.Marshal(episode)
		r, _ := episode["_id"]
		product.ID = r.(string)
		//fmt.Println(r)
		r, _ = episode["name"]
		product.Name = r.(string)
		//fmt.Println(r)
		//r, _ = episode["type"]
		//product.Type = r.(string)
		////fmt.Println(r)
		//r, _ = episode["price"]
		//product.Department = r.(float64)
		//fmt.Println(r)


		//bson.Unmarshal(marshal, &product)

		fmt.Println(product)
		products = append(products, product)
	}



	//cur.Decode(&products)
	//for cur.Next(context.Background()) {
	//	product := model.Employee{
	//
	//	}
	//	err := cur.Decode(&product)
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//	products = append(products, product)
	//}


	fmt.Println(products)
	return products, nil
}

func (mongoDb *mongoDB) GenEmployeeNumber() int {
	ctx := context.Background()
	cursor, err := mongoDb.employeeCount.Find(ctx, bson.M{})


	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
	count := -888
	var exists bool
	count, exists = episodes[0]["employeeCount"].(int)

	if !exists {
		return count
	}

	return count
}

//func (mongoDB *mongoDB) GetRecentItems() ([]model.Employee, error) {
//	return nil, nil
//}
//func (mongoDB *mongoDB) SetSelectedItem(product model.Employee) (bool, error) {
//	return false, nil
//}
