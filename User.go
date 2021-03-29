package model
import (
 "go-rest/config"
 "fmt"
_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Fetch all user data
func GetAllEmployees(employee *[]Employee) (err error) {
 if err = config.DB.Find(employee).Error; err != nil {
  return err
 }
 return nil
}
//CreateUser ... Insert New data
func CreateEmployee(employee *Employee) (err error) {
 if err = config.DB.Create(employee).Error; err != nil {
  return err
 }
 return nil
}
//GetUserByID ... Fetch only one user by Id
func GetEmployeeByID(employee *Employee, id string) (err error) {
 if err = config.DB.Where("id = ?", id).First(employee).Error; err != nil {
  return err
 }
 return nil
}
//UpdateUser ... Update user
func UpdateEmployee(employee *Employee, id string) (err error) {
 fmt.Println(employee)
 config.DB.Save(employee)
 return nil
}
//DeleteUser ... Delete user
func DeleteEmployee(employee *Employee, id string) (err error) {
 config.DB.Where("id = ?", id).Delete(employee)
 return nil
}