# Description: **Service A: RESTful Web Service**

    - Endpoints:
        POST /: Accepts a JSON object (2-3 fields of your choice), stores it in an SQLite database, and publishes it to NATS.
        GET /{id}: Retrieves and returns the stored JSON object based on an identifier.
    - Database: Uses SQLite to persist incoming objects.
    - Messaging: Publishes the stored object to NATS after saving.

## Env vars
Env var | Description
-------- | -------- 
`SQLITE_DB_PATH`   | SQL database path    
`NATS_URL`   | NATS url for messaging   
`APP_PORT`   | App port where to start the http server   