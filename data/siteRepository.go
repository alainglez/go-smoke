package data

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alainglez/go-smoke/models"
)

type SiteRepository struct {
	C *mgo.Collection
}

func (r *SiteRepository) Create(site *models.Site) error {
	obj_id := bson.NewObjectId()
	site.Id = obj_id
	site.CreatedOn = time.Now()
	err := r.C.Insert(&site)
	return err
}

func (r *SiteRepository) Update(site *models.Site) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": site.Id},
		bson.M{"$set": bson.M{
			"sitedns":     site.SiteDns,
			"description": site.Description,
		}})
	return err
}
func (r *SiteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *SiteRepository) GetAll() []models.Site {
	var sites []models.Site
	iter := r.C.Find(nil).Iter()
	result := models.Site{}
	for iter.Next(&result) {
		sites = append(sites, result)
	}
	return sites
}
func (r *SiteRepository) GetById(id string) (site models.Site, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&site)
	return
}
func (r *SiteRepository) GetByUser(user string) []models.Site {
	var sites []models.Site
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.Site{}
	for iter.Next(&result) {
		sites = append(sites, result)
	}
	return sites
}
