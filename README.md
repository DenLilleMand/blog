#Creating migrations
goose -dir="migrations" postgres "user=<dbusrname> dbname=<dbname> sslmode=disable"  create <migrationname> <sql|go>
