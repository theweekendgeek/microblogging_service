# Twitter Service

## Description
This service fetches tweets for a predefined list of users and stores them in a database.
It is deployed automatically as a Google Cloud Function, which is triggered at regular intervals.

### Current limitations
On every call the Tweet table is wiped and only the last 20 tweets for a given user are stored.

### Todos

- [ ] add more tests
- [ ] get only new tweets since the last call
- [ ] change the way the user ids are passed
- [ ] add service account creation + function authorization to terraform
- [ ] simplify handling of .env files
