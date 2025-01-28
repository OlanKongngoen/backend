package employee

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "backend/api/db"
)

type Tbl_employee struct {
    Emp_id int`gorm:"primaryKey"`
    Emp_firstname string
    Emp_lastname string
    Emp_department string
    Emp_salary float64
}

// Biding from Employee JSON
type EmployeeBody struct {
    Emp_id int`json:"emp_id" binding:"required"`
    Emp_firstname   string `son:"emp_firstname" binding:"required"`
    Emp_lastname   string `son:"emp_lastname" binding:"required"`
    Emp_department   string `son:"emp_department" binding:"required"`
    Emp_salary   float64 `son:"emp_salary" binding:"required"`
}

//GET Employee
func GET(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

//GET Employee to Database
func GETDB(c *gin.Context) {
 var employees []Tbl_employee
 db.Db.Find(&employees)
 c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Read Success", "employees": employees})
}

//POST Employee
func POST(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee POST Method!",
    })
}

//POST Employee to Database
func POSTDB(c *gin.Context) {
    var json EmployeeBody
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tbl_employee := Tbl_employee{Emp_id: json.Emp_id, Emp_firstname: json.Emp_firstname, Emp_lastname: json.Emp_lastname, Emp_department: json.Emp_department, Emp_salary: json.Emp_salary}
    db.Db.Create(&tbl_employee)
    if tbl_employee.Emp_firstname != "" {
        c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Created", "tbl_fund": tbl_employee})
    } else {
        c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Failed", "tbl_fund": tbl_employee})
    }

}

//PUT Employee
func PUT(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee PUT Method!",
    })
}

//PUT Employee to Database
func PutEmployeeDB(c *gin.Context) {
    var json EmployeeBody

    var UpdateEmployees Tbl_employee
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Db.First(&UpdateEmployees, json.Emp_id)
    UpdateEmployees.Emp_firstname = json.Emp_firstname
    UpdateEmployees.Emp_lastname = json.Emp_lastname
    UpdateEmployees.Emp_department = json.Emp_department
    UpdateEmployees.Emp_salary = json.Emp_salary

    db.Db.Save(&UpdateEmployees)

    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Updated"})

}

//DELETE Employee
func DELETE(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee DELETE Method!",
    })
}

//DELETE Employee to Database
func DELETEDB(c *gin.Context) {
    id := c.Param("id")
    var employees []Tbl_employee
    db.Db.Delete(&employees, "emp_id = ?", id)
    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Delete Profit Success"})
}



//GET By ID ALL
func GETEmployeeByID(c *gin.Context) {
    id := c.Param("id")
        var employees []Tbl_employee
 db.Db.First(&employees, id)
 c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Read Success", "employees": employees})
}

func POSTEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "POST By ID": id,
    })
}
func PUTEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "PUT By ID": id,
    })
}
func DELETEEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    var employees []Tbl_employee
    db.Db.Delete(&employees, "emp_id = ?", id)
    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Delete Profit Success"})
}