package main

import (
  "fmt"
  EmployeeController "backend/api/controller/employee"
  "github.com/gin-gonic/gin"
  "backend/api/db"
  "github.com/joho/godotenv"
)

func main() {
  //Get .env
  err := godotenv.Load(".env")
  if err != nil {
		fmt.Println("Error loading .env file")
  }
  //get InitDB fuction
  db.InitDB()

  router := gin.Default()
  //Employee API Method
  router.GET("/employee", EmployeeController.GET) //GET
  router.GET("/employeedb", EmployeeController.GETDB) //GET TO DB
  router.POST("/employee", EmployeeController.POST) //POST
  router.POST("/employeedb", EmployeeController.POSTDB) //POST TO DB
  router.PUT("/employee", EmployeeController.PUT) //PUT
  router.DELETE("/employee", EmployeeController.DELETE) //DELETE

  router.GET("/employee/:id", EmployeeController.GETEmployeeByID)  //GET By ID
  router.POST("/employee/:id", EmployeeController.POSTEmployeeByID)  //POST By ID
  router.PUT("/employee/:id", EmployeeController.PUTEmployeeByID)  //PUT By ID
  router.DELETE("/employee/:id", EmployeeController.DELETEEmployeeByID)  //DELETE By ID

  //Customer API Method
  
  router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}