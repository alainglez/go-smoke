package data

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alainglez/go-smoke/models"
)

type SmokeTestRepository struct {
	C *mgo.Collection
}

func (r *SmokeTestRepository) Create(smoketest *models.SmokeTest) error {
	obj_id := bson.NewObjectId()
	smoketest.Id = obj_id
	smoketest.CreatedOn = time.Now()
	err := r.C.Insert(&smoketest)
	return err
}
func (r *SmokeTestRepository) Update(smoketest *models.SmokeTest) error {
	// partial update on MongoDB
	err := r.C.Update(bson.M{"_id": smoketest.Id},
		bson.M{"$set": bson.M{
			"passfail":    smoketest.PassFail,
			"urlresults":  smoketest.UrlResults,
		}})
	return err
}
func (r *SmokeTestRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *SmokeTestRepository) GetAll() []models.SmokeTest {
	var smoketests []models.SmokeTest
	iter := r.C.Find(nil).Iter()
	result := models.SmokeTest{}
	for iter.Next(&result) {
		smoketests = append(smoketests, result)
	}
	return smoketests
}
func (r *SmokeTestRepository) GetById(id string) (smoketest models.SmokeTest, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&smoketest)
	return
}
func (r *SmokeTestRepository) GetByUser(user string) []models.SmokeTest {
	var smoketests []models.SmokeTest
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.SmokeTest{}
	for iter.Next(&result) {
		smoketests = append(smoketests, result)
	}
	return smoketests
}
