package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/*category crud begin*/
func AddLog(level int64, content, logfrom string) error {
	o := orm.NewOrm()
	log := &Log{Level: level, Content: content, Logfrom: logfrom, Created: time.Now()}
	_, err := o.Insert(log)
	return err
}
func AllLog() ([]*Log, error) {
	var logs = make([]*Log, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Log").All(&logs)
	return logs, err
}
func GetLogById(id int64) (*Log, error) {
	o := orm.NewOrm()
	log := &Log{Id: id}
	err := o.Read(log)
	return log, err
}

/*--------category crud end*/
