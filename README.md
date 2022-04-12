# icanhazquotes

## Todo

4. Docker File
5. CI/CD file
6. Helm Chart
9. Failure Code
11. Different Environment Code


Mongo setup

docker run --name icanhazquotesdb \
-p 27017:27017 \
-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
-e MONGO_INITDB_ROOT_PASSWORD=secret \
-e MONGO_INITDB_DATABASE=icanhazquotes \
-v quotes-data:/data/db \
-d mongo

mongoimport --uri mongodb://localhost:27017/icanhazquotes --username=mongoadmin --authenticationDatabase "admin" --collection quotes --type json --file quotes.json --jsonArray

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
| ICHO_DB_DBNAME | icanhazqutes | Database to access data from |