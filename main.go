package main
import (
 "go-rest/config"
 "go-rest/model"
 "go-rest/service"
 "fmt"
"github.com/jinzhu/gorm"
)
var err error
func main() {
 config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
if err != nil {
  fmt.Println("Status:", err)
 }
defer config.DB.Close()
 config.DB.AutoMigrate(&model.Employee{})
r := service.SetupRouter()
 //running
 r.Run()
}