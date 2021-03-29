package controller
import (
 "go-rest/model"
 "fmt"
 "net/http"
"github.com/gin-gonic/gin"
)
//GetUsers ... Get all users
func GetEmployees(c *gin.Context) {
 var employee []model.Employee
 err := model.GetAllEmployees(&employee)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, employee)
 }
}
//CreateUser ... Create User
func CreateEmployee(c *gin.Context) {
 var employee model.Employee
 c.BindJSON(&employee)
 err := model.CreateEmployee(&employee)
 if err != nil {
  fmt.Println(err.Error())
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, employee)
 }
}
//GetUserByID ... Get the user by id
func GetEmployeeByID(c *gin.Context) {
 id := c.Params.ByName("id")
 var employee model.Employee
 err := model.GetEmployeeByID(&employee, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, employee)
 }
}
//UpdateUser ... Update the user information
func UpdateEmployee(c *gin.Context) {
 var employee model.Employee
 id := c.Params.ByName("id")
 err := model.GetEmployeeByID(&employee, id)
 if err != nil {
  c.JSON(http.StatusNotFound, employee)
 }
 c.BindJSON(&employee)
 err = model.UpdateEmployee(&employee, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, employee)
 }
}
//DeleteUser ... Delete the user
func DeleteEmployee(c *gin.Context) {
 var employee model.Employee
 id := c.Params.ByName("id")
 err := model.DeleteEmployee(&employee, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
 }
}