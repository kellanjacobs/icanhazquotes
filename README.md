# icanhazquotes

## Todo
* CI/CD file
* Helm Chart
* Failure Code
* Different Environment Code

## Local Install

commands must be run from root directory of this repo

```
docker-compose up -d
docker cp quotes.json db:/
docker exec -it db bash
mongoimport --uri mongodb://localhost:27017/icanhazquotes --username=mongouser --authenticationDatabase "admin" --collection quotes --type json --file quotes.json --jsonArray
```

| ENV Variables | Default Value | Description |
|---------------|---------------|-------------|
| ICHQ_APP_PORT | 9999 | HTTP port |
| ICHQ_APP_IP | 0.0.0.0 | IP Listening Address |
| ICHQ_DB_HOST | localhost | Databse Host |
| ICHQ_DB_USER | mongouser | Database User |
| ICHQ_DB_PASSWORD | changeme | Database Password |
| ICHO_DB_AUTHSOURCE | auth | Authorization database for mongo user |
| ICHO_DB_PORT | 27017 | Database Port | 
| ICHO_DB_TIMEOUT | 2 | Database timeout |
| ICHO_DB_DBNAME | icanhazquotes | Database to access data from |
