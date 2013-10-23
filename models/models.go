package models

import (
	"os"
	"path"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqllite3"
)

const (
	DB_NAME        = "data/tomzblog.db"
	SQLITE3_DRIVER = "sqlite3"
)

type User struct {
	Id        int64
	Name      string    //姓名
	UserCode  string    //用户编码
	PassWord  string    //密码
	BirthDate time.Time `orm:"index"` // 出生日期
	Sex       string    //性别
	QQ        int
	Mobile    string
	Phone     string
	Email     string
	HomePage  string
}

type Article struct {
	Id          int64
	Type        string
	Title       string
	Content     string `orm:"size(5000)"`
	Author      int64
	CommentInd  string
	ReprintInd  string
	DefunctInd  string
	CreatedBy   int64
	CreatedTime time.Time `orm:"index"`
	UpdatedBy   int64
	UpdatedTime time.Time `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(DB_NAME) {
		os.MkdirAll(path.Dir(DB_NAME), os.ModePerm)
		os.Create(DB_NAME)
	}

	orm.RegisterModel(new(User), new(Article))
	orm.RegisterDriver(SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", SQLITE3_DRIVER, DB_NAME, 10)
}
