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