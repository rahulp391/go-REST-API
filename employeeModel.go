package model
type Employee struct {
 Name    string `json:"name"`
 Email   string `json:"email"`
 Phone   string `json:"phone"`
}
func (b *Employee) TableName() string {
 return "employee"
}