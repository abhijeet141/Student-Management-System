package models

import (
	"log"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Course struct {
	Id    int    `json:"id" orm:"auto"`
	Title string `json:"title"`
}
type User struct {
	Id       int    `json:"id" orm:"auto"`
	UserName string `json:"username" orm:"size(255);unique"`
	Password string `json:"password" orm:"size(255)"`
	Role     string `json:"role"`
}
type Student struct {
	Id     int     `json:"id" orm:"size(255)"`
	Name   string  `json:"name" orm:"size(255)"`
	Age    int     `json:"age"`
	Course *Course `json:"course" orm:"rel(fk)"`
	User   *User   `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Course), new(User), new(Student))
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connectionstring := os.Getenv("CONNECTION_STRING")
	if connectionstring == "" {
		log.Fatal("Connection String is not set in the environment")
	}
	orm.RegisterDataBase("default", "mysql", connectionstring)
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Database connection is established")
}
