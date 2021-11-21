# Tweets streamer service

## Backend

In order to use backend service that connects to Twitter API you need to provide
credentials. Check `docker/golang/secrets.template.env` for more details. File with secrets
should be located at `docker/golang/secrets.env`. When you run the service using `docker-compose up`
the file will be created as a volume in service root dir.

### Start the service

To start the service you have to have docker installed on your machine. If you have run 
command `docker-compose up`. That will start REST service on default `3001` port.

To make a request to the service use `$ curl localhost:3001/stream/:keyword`
Keyword is the word that has to be inside tweet, e.g. rice. 

#### TODO

This service is a showcase project so it's not fully developed.

- Write more test
- Create CI/CD pipeline
- Expand the service with new endpoints