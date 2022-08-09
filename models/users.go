package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName   string        `valid:"Required" json:"first_name" bson:"first_name"`
	LastName    string        `valid:"Required" json:"last_name" bson:"last_name"`
	Username    string        `valid:"Required" json:"username" bson:"username"`
	PhoneNumber string        `valid:"Required" json:"phone_number" bson:"phone_number"`
	Email       string        `valid:"Email" json:"email" bson:"email"`
	Address     string        `valid:"Required" json:"address" bson:"address"`
}

func (this *Users) Name() string {
	return "users"
}

func (this *Users) List(perPage int, page int) (result []Users, err error) {
	err = mHandler.C(this.Name()).
		Find(nil).
		Limit(perPage).
		Skip(page).
		All(&result)
	if err != nil {
		result = make([]Users, 0)
	}
	return
}

func (this *Users) Create() (result *Users, err error) {
	this.Id = bson.NewObjectId()
	err = mHandler.C(this.Name()).Insert(this)
	result = this
	return
}

func (this *Users) UpdateByID(id bson.ObjectId) error {
	return mHandler.C(this.Name()).UpdateId(id, this)
}

func (this *Users) DeleteByID(id bson.ObjectId) error {
	return mHandler.C(this.Name()).RemoveId(id)
}

func (this *Users) GetByID(id bson.ObjectId) (result Users, err error) {
	err = mHandler.C(this.Name()).FindId(id).One(&result)
	return
}
