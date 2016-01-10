package models

import (
	"github.com/astaxie/beego/orm"
)

/*Product crud begin*/
func AddProduct(categoryId int64, name, item, picture, tag string) error {
	o := orm.NewOrm()
	product := &Product{CategoryId: categoryId, Name: name, Item: item, Picture: picture, Tag: tag}
	_, err := o.Insert(product)
	return err
}
func DelProduct(id int64) error {
	o := orm.NewOrm()
	product := &Product{Id: id}

	if o.Read(product) == nil {
		_, err := o.Delete(product)
		return err
	}
	return nil
}
func GetProductById(id int64) (*Product, error) {
	o := orm.NewOrm()
	product := &Product{Id: id}
	err := o.Read(product)
	return product, err
}
func GetProducts() ([]*Product, error) {
	var products = make([]*Product, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Product").All(&products)
	return products, err
}
func GetProductsByCategoryId(categoryId int64) ([]*Product, error) {
	var products = make([]*Product, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Product").Filter("CategoryId", categoryId).All(&products)
	return products, err
}

/*--------product crud end*/
