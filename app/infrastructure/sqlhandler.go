package infrastructure

import (
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "root:ecc@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

//データベースが変わった場合や使用しているフレームワークが
//変更された場合などはここを変更する
//interface層内の./database配下にinterfaceを定義する
func (handler *SqlHandler) Create(obj interface{}) {
	//Gorm.Createメソッド
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	//Gorm.Findメソッド
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	//Gorm.Deleteメソッド
	handler.db.Delete(obj, id)
}

func (handler *SqlHandler) UpdateById(obj domain.User, name string) {
	//Gorm.Updateメソッド
	handler.db.First(&obj)
	fmt.Println("objの中身")
	fmt.Println(obj.ID)
	fmt.Println(obj.Name)
	obj.Name = name
	handler.db.Save(&obj)

}
