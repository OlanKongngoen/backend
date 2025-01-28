package main

import (
  "fmt"
  EmployeeController "backend/api/controller/employee"
  AdminController "backend/api/controller/admin"
  AuthController "backend/api/controller/auth"
  "github.com/gin-gonic/gin"
  "backend/api/db"
  "github.com/joho/godotenv"
  "backend/api/middleware" //เรียกใช้ไฟล์ที่อยู่ในห้อง middleware
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

  authorized := router.Group("/api", middleware.JwtAuthen())  //ทำการจัดกลุ่ม path ที่ต้องการล๊อค api
  
  //Employee API Method
  router.GET("/employee", EmployeeController.GET) //GET
  router.GET("/employee/:id", EmployeeController.GETEmployeeByID) //GET BY ID
  authorized.GET("/employeedb", EmployeeController.GETDB) ///ล๊อค api โดยต้องแนบ token ก่อนถึงใช้งานได้
  router.GET("/register", AdminController.GetAdmin) //GET Admin
 
  router.POST("/employee", EmployeeController.POST) //POST
  router.POST("/employeedb", EmployeeController.POSTDB) //POST TO DB
  router.POST("/register", AdminController.PostAdmin) //POST admin
  router.POST("/login", AuthController.Login) //POST LOGIN

  router.PUT("/employee", EmployeeController.PUT) //PUT
  router.PUT("/employeedb", EmployeeController.PutEmployeeDB) //PUT DB
 
  router.DELETE("/employee", EmployeeController.DELETE) //DELETE
  router.DELETE("/employeedb/:id", EmployeeController.DELETEDB) //DELETE DB

  //Customer API Method

  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  
}