# 📋 Challenge Summary 
### The task is to build two Go services that communicate using NATS messaging.

## Services Overview
**Service A: RESTful Web Service**

    - Endpoints:
        POST /: Accepts a JSON object (2-3 fields of your choice), stores it in an SQLite database, and publishes it to NATS.
        GET /{id}: Retrieves and returns the stored JSON object based on an identifier.
    - Database: Uses SQLite to persist incoming objects.
    - Messaging: Publishes the stored object to NATS after saving.

**Service B: NATS Subscriber**

    - Purpose: Subscribes to messages from Service 1 via NATS.
    - Functionality: Receives messages and prints them to the console.

## Requirements

    - Code Quality: Clean, modular, and concise with good naming conventions.
    - Testing: At least one unit test for core functionality.
    - Docker Compose Setup:
        Both services and NATS should run via docker-compose up.
        The web service is accessible from the Docker host with exposed ports.


# ✅ Solution
The solution now realises 2 services in go in which the described requirements are implemented adequately.

## ℹ️ Notes / special features:
1. In service_a, the messaging service is called in the invoice service after the invoice has been saved. If an error occurs here, the invoice is only saved in the sqlite db but not published on the bus. As can be seen from the comment in the code, there is potential for improvement here in terms of transaction security. A rollback mechanism / transaction could be opened for the database, which only persists if the message is on the bus.
In my opinion, this is where the opinion of domain experts is needed. Is eventual consistency sufficient here?  Should we go the extra mile and implement the rollback / transaction mechanism or even the outbox pattern?

2. The implemented unit test currently mainly tests third-party code / functionality. I would not normally implement this as I am now dependent on technical aspects and components in the test. However, there is no domain logic in the services that can ideally be tested, for this reason the test should act as an example and test the post endpoint including the underlying logic

# 🚀 Run
To run the services with nats, simply use the following command  </br>
`docker compose up`