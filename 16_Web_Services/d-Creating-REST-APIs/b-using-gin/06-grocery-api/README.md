# Grocery API


## Project Setup

    $ go mod init grocery-api
    $ go mod tidy


├───.idea
├───handlers
├───middlewares
├───models
└───routes


MVC 
    Models      -- data 
    Views       -- business logic/ presentation layer
    Controllers -- routing




Assignment
===========


1) Add Rate Limiting

   Use middleware like github.com/ulule/limiter/v3 to prevent abuse of your API.

2) Secure the Application

   Use HTTPS by configuring TLS certificates.
