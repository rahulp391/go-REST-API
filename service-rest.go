package service
import (
 "go-rest/controller"
"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
 r := gin.Default()
 grp1 := r.Group("/employee-api")
 {
  grp1.GET("employee", controller.GetEmployees)
  grp1.POST("employee", controller.CreateEmployee)
  grp1.GET("employee/:id", controller.GetEmployeeByID)
  grp1.PUT("employee/:id", controller.UpdateEmployee)
  grp1.DELETE("employee/:id", controller.DeleteEmployee)
 }
 return r
}