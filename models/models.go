package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	Site struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name        string        `json:"name,omitempty"`
		Description string        `json:"description,omitempty"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
	SiteUrl struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		SiteId      bson.ObjectId `json:"siteid,omitempty"`
		Url	    string        `json:"url,omitempty"`
		HtmlFragment string	  `json:"htmlfragment"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
	SmokeTest struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		SiteName    string	  `json:"sitename"`
		CreatedBy   string        `json:"createdby"`
		Host        string        `json:"host,omitempty"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		PassFail    string        `json:"passfail,omitempty"`
		UrlResults  []UrlCodes    `json:"urlresults,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	UrlCodes struct {
		Url	    string	  `json:"url,omitempty"` 
		StatusCode  string	  `json:"statuscode,omitempty"`
		PassFail    string	  'json:"passfail,omitempty"`
	}
	
)
