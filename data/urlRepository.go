package data

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alainglez/go-smoke/models"
)

type UrlRepository struct {
	C *mgo.Collection
}

func (r *UrlRepository) Create(url *models.TaskUrl) error {
	obj_id := bson.NewObjectId()
	url.Id = obj_id
	url.CreatedOn = time.Now()
	err := r.C.Insert(&url)
	return err
}

func (r *UrlRepository) Update(url *models.TaskUrl) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": url.Id},
		bson.M{"$set": bson.M{
			"description": url.Description,
		}})
	return err
}
func (r *UrlRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *UrlRepository) GetByTask(id string) []models.TaskUrl {
	var urls []models.TaskUrl
	taskid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.TaskUrl{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}
func (r *UrlRepository) GetAll() []models.TaskUrl {
	var urls []models.TaskUrl
	iter := r.C.Find(nil).Iter()
	result := models.TaskUrl{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}
func (r *UrlRepository) GetById(id string) (url models.TaskUrl, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&url)
	return
}
