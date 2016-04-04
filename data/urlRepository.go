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

func (r *UrlRepository) Create(url *models.SiteUrl) error {
	obj_id := bson.NewObjectId()
	url.Id = obj_id
	url.CreatedOn = time.Now()
	err := r.C.Insert(&url)
	return err
}

func (r *UrlRepository) Update(url *models.SiteUrl) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": url.Id},
		bson.M{"$set": bson.M{
			"Url": url.Url,
		}})
	return err
}
func (r *UrlRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *UrlRepository) GetBySite(id string) []models.SiteUrl {
	var urls []models.SiteUrl
	siteid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"siteid": siteid}).Iter()
	result := models.SiteUrl{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}
func (r *UrlRepository) GetAll() []models.SiteUrl {
	var urls []models.SiteUrl
	iter := r.C.Find(nil).Iter()
	result := models.SiteUrl{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}
func (r *UrlRepository) GetById(id string) (url models.SiteUrl, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&url)
	return
}
