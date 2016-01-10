package models

import (
	"github.com/astaxie/beego/orm"
)

/*category crud begin*/
func AddCategory(title string) error {
	o := orm.NewOrm()
	category := &Category{Title: title}
	if o.Read(category) != nil { //已经存在
		return nil
	}
	_, err := o.Insert(category)
	return err
}
func DelCategory(id int64) error {
	o := orm.NewOrm()
	category := &Category{Id: id}

	if o.Read(category) == nil {
		_, err := o.Delete(category)
		return err
	}
	return nil
}
func AllCategory() ([]*Category, error) {
	var categorys = make([]*Category, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Category").All(&categorys)
	return categorys, err
}

/*--------category crud end*/
