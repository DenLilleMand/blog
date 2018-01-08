#Setting up database

1.CREATE DATABASE blog;
2.CREATE USER denlillemand;
3.GRANT ALL PRIVILEGES ON DATABASE blog TO denlillemand;
4.CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


#Creating migrations
$ goose -dir="migrations" postgres "user=<dbusrname> dbname=<dbname> sslmode=disable"  create <migrationname> <sql|go>
