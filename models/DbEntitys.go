package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"

	"os"
	"path"
)

const (
	_DB_NAME       = "data/myapp.db"
	_SQLITE3_DRIVE = "sqlite3"
)

type Category struct {
	Id      int64
	Title   string
	Created time.Time `orm:"index"`
}

type User struct {
	Id      int64
	Account string
	Name    string
	Pass    string
	Created time.Time `orm:"index"`
}

type Companyinfo struct {
	Id           int64
	Name         string
	Name4en      string
	Address      string
	Telephone    string
	Fax          string
	Email        string
	Cellphone    string
	OtherContact string
	IM           string
	Created      time.Time `orm:"index"`
}
type Product struct {
	Id         int64
	CategoryId int64
	Name       string
	Item       string
	Picture    string
	Tag        string
	Status     int64     /*0 active 1 close*/
	Created    time.Time `orm:"index"`
}

type Inquery struct {
	Id            int64
	Title         string
	Description   string
	ProductId     int64
	Inqueryer     string
	InqueryerCorp string
	EMail         string
	PhoneNumber   string
	Status        int64 /*0 active 1 replyed 2 close*/
	Note          string
	Created       time.Time `orm:"index"`
}

type Feedback struct {
	Id          int64
	Title       string
	Description string
	Poster      string
	PosterCorp  string
	EMail       string
	PhoneNumber string
	Status      int64 /*0 active 1 replyed 2 close*/
	Note        string
	Created     time.Time `orm:"index"`
}

type Log struct {
	Id      int64
	Level   int64
	Content string
	Logfrom string
	Created time.Time `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.Mkdir(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(User), new(Companyinfo), new(Inquery), new(Product), new(Feedback), new(Log))
	orm.RegisterDriver(_SQLITE3_DRIVE, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVE, _DB_NAME, 10)
}

/*增加用户*/
func AddUser(account, name, pass string) error {
	o := orm.NewOrm()
	user := new(User)
	user.Account = account
	user.Name = name
	user.Pass = pass
	user.Created = time.Now()

	//q := o.QueryTable("User")
	_, err := o.Insert(&user)

	return err
}
