package models

import (
	"os"
	"path"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	"github.com/mattn/"
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
	Id int64
	Type                  VARCHAR2(8),
  TITLE                 VARCHAR2(200) not null,
  CONTENT               VARCHAR2(4000),
  AUTHOR                NUMBER(38),
  COMMENT_IND           CHAR(1) default 'Y',
  REPRINT_IND           CHAR(1) default 'Y',
  DEFUNCT_IND           CHAR(1) default 'N' not null,
  CREATED_BY            NUMBER(38),
  CREATED_DATETIME      DATE default SYSDATE not null,
  LAST_UPDATED_BY       NUMBER(38),
  LAST_UPDATED_DATETIME DATE default SYSDATE not null
}

func RegisterDB() {

}
