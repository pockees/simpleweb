package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/*category crud begin*/
func AddInquery(title, description, inqueryer, inquerycorp, email, phonenumber string, productid int64) error {
	o := orm.NewOrm()
	inquery := &Inquery{Title: title, Description: description, ProductId: productid, Inqueryer: inqueryer, InqueryerCorp: inquerycorp, EMail: email, PhoneNumber: phonenumber, Status: 0, Created: time.Now()}
	_, err := o.Insert(inquery)
	return err
}
func DelInquery(id int64) error {
	o := orm.NewOrm()
	inquery := &Inquery{Id: id}

	if o.Read(inquery) == nil {
		_, err := o.Delete(inquery)
		return err
	}
	return nil
}
func AllInquery() ([]*Inquery, error) {
	var inquerys = make([]*Inquery, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Inquery").All(&inquerys)
	return inquerys, err
}
func GetInqueryById(id int64) (*Inquery, error) {
	o := orm.NewOrm()
	inquery := &Inquery{Id: id}
	err := o.Read(inquery)
	return inquery, err
}

/*--------category crud end*/
