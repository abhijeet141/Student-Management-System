package config

// import (
// 	"log"
// 	"os"

// 	"github.com/beego/beego/v2/client/orm"
// 	"github.com/joho/godotenv"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}
// 	connectionstring := os.Getenv("CONNECTION_STRING")
// 	if connectionstring == "" {
// 		log.Fatal("Connection String is not set in the environment")
// 	}
// 	orm.RegisterDataBase("default", "mysql", connectionstring)
// 	err = orm.RunSyncdb("default", false, true)
// 	if err != nil {
// 		log.Fatal("Error connecting to the database: ", err)
// 	}
// 	log.Println("Database connection is established")
// }
