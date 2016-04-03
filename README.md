# go-smoke
go microservice to smoke test a website. 

- RESTful JSON-based API Server. 
- Supports OAuth for authentication and JWT for authorization.
- Third-party packages: (no web development framework such as NET MVC, Ruby on Rails or Beego for Go.)
    - mgo.v2 for MongoDB driver for Go and implementation of BSON spec for Go.
    - gorilla/mux for request router and dispatcher.
    - dgrijalva/jwt-go helper functions for working with JSON Web Tokens (JWT).
    - codegangsta/negroni for idiomatic approach to HTTP middleware.
- Deployed with Docker.

# USAGE
runs a smoke test of multiple URLs on a website using concurrent go routines to prime and check the status of a host recently deployed to via a CI job, returns PASS or FAIL along with the status code, duration and size of http responses to each of the URLs. Normally called fro every host in a target environment during a rolling deployment by taking a host down behind the load balancer (reverse proxy), deploying a new build and smoke testing it. 

Happy flow

First time
- register user
- login user
- add site e.g. www.carnival.com
- add URLs
- run a smoke test against a site on a host
Going forward 
- login user
- run a smoke test against a site on a host

#URI                    HTTP Verb           Functionality
/users/register         Post                Creates a new user, e.g. chefdeliveryusr
/users/login            Post                User logs in to the system, which returns a JWT if loggin is successful. 
/sites                  Post                Creates a new site. E.g. www.carnival.com, www.carnival.co.uk, ww4.uatcarnival.com
/sites                  Get                 Gets all sites
/urls/sites/{id}        Get                 Gets all URLs for a given site ID. The value of the ID comes from the route parameter
/urls                   Post                Creates a new URL against an existing site

E.g. Prime and check status of .. 
 Core           "http://www.carnival.com"
 BookingEngine  "https://secure.carnival.com/BookingEngine/Booking/Book?durDays=4&embkCode=MIA&isMilitary=N&isOver55=N&isPastGuest=N&itinCode=ZW0&numGuests=2&sailDate=03062017&sailingID=77874&shipCode=VI&showDbl=False&subRegionCode=CW&be_version=22#/number-of-staterooms"
 Login          "https://secure.carnival.com/BookedGuest/guestmanagement/mycarnival/logon?returnUrl=http%3A%2F%2Fwww.carnival.com%2F" 
 BookedGuest    "https://secure.carnival.com/BookedGuest/"
 ShoreEx        "https://https://secure.carnival.com/shore-excursions"
 Funshops       "http://www.carnival.com/Funshops/"                
 OnlineCheckIn  "TO DO"
 
/smoketests             Post                Runs a smoke test. Host IP, site name from JSON data on Http request Body. JWT on Header.

TO DO: Add URIs for reporting, updating and deleting previouly created users, sites, URLs and smoke tests.
