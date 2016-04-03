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
