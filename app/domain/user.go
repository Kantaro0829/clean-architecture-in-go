package domain

//domain 層
//Entity の定義を行う
type User struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
