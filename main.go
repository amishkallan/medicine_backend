// package main

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB
// var err error

// func main() {
// 	db, err = sql.Open("mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<dbname>")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	log.Print("started")
// }
package main

import (
	"github.com/gin-gonic/gin"

	"Documents/rest-api-mysql/controllers" //
	"Documents/rest-api-mysql/models"      // new
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()                      // new
	r.GET("/medicine", controllers.FindMedicines) // new
	r.POST("/medicine", controllers.CreateMedicine)
	r.GET("/medicine/:medname", controllers.FindMedicine)
	r.PATCH("/medicine/:medname", controllers.UpdateMedicine)
	r.DELETE("/medicine/:medname", controllers.DeleteMedicine)
	r.Run()
}
