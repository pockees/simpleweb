package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/*category crud begin*/
func AddFeedback(title, description, poster, postercorp, email, phonenumber string) error {
	o := orm.NewOrm()
	feedback := &Feedback{Title: title, Description: description, Poster: poster, PosterCorp: postercorp, EMail: email, PhoneNumber: phonenumber, Status: 0, Created: time.Now()}
	_, err := o.Insert(feedback)
	return err
}
func DelFeedback(id int64) error {
	o := orm.NewOrm()
	feedback := &Feedback{Id: id}

	if o.Read(feedback) == nil {
		_, err := o.Delete(feedback)
		return err
	}
	return nil
}
func AllFeedback() ([]*Feedback, error) {
	var feedbacks = make([]*Feedback, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Feedback").All(&feedbacks)
	return feedbacks, err
}
func GetFeedbackById(id int64) (*Feedback, error) {
	o := orm.NewOrm()
	feedback := &Feedback{Id: id}
	err := o.Read(feedback)
	return feedback, err
}

/*--------category crud end*/
