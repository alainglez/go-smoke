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

func (r *UrlRepository) Create(url *models.TestUrl) error {
	obj_id := bson.NewObjectId()
	url.Id = obj_id
	url.CreatedOn = time.Now()
	err := r.C.Insert(&url)
	return err
}

func (r *UrlRepository) Update(url *models.TestUrl) error {
	// partial update on MongoDB
	err := r.C.Update(bson.M{"_id": url.Id},
		bson.M{"$set": bson.M{
			"siteid":	url.SiteId,
			"url": 		url.Url,
			"htmlfragment": url.HtmlFragment,
		}})
	return err
}
func (r *UrlRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *UrlRepository) GetBySite(id string) []models.TestUrl {
	var urls []models.TestUrl
	siteid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"siteid": siteid}).Iter()
	result := models.TestUrl{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}
func (r *UrlRepository) GetAll() []models.TestUrl {
	var urls []models.TestUrl
	iter := r.C.Find(nil).Iter()
	result := models.TestUrl{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}
func (r *UrlRepository) GetById(id string) (url models.TestUrl, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&url)
	return
}
