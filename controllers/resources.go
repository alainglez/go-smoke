package controllers

import (
	"github.com/alainglez/go-smoke/models"
)

//Models for JSON resources
type (
	//For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	//For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	//Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// For Post/Put - /sites
	// For Get - /sites/id
	SiteResource struct {
		Data models.Site `json:"data"`
	}
	// For Get - /sites
	SitesResource struct {
		Data []models.Site `json:"data"`
	}
	// For Post/Put - /smoketests
	// For Get - /smoketests/id
	SmokeTestResource struct {
		Data models.SmokeTest `json:"data"`
	}
	// For Get - /smoketests
	SmokeTestsResource struct {
		Data []models.SmokeTest `json:"data"`
	}
	// For Post/Put - /urls
	UrlResource struct {
		Data UrlModel `json:"data"`
	}
	// For Get - /urls
	// For /urls/sites/id
	UrlsResource struct {
		Data []models.SiteUrl `json:"data"`
	}
	//Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
	//Model for a SiteUrl
	UrlModel struct {
		SiteId      string `json:"siteid"`
		Description string `json:"description"`
	}
)
