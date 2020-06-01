# bcnchess

Barcelona Chess (bcnchess) is a free open source website that is aimed at making meet ups of tournaments within Barcelona/Spain easier. Currently the situation faced is that all the tournaments
are shared through whatsapp which makes it easy to miss one or not know about the others.

### Getting Started
- *IMPORTANT* WHILE RUNNING THE SET UP SCRIPT A NEW USER OF `'chess-web'@'localhost'` *AND* `'test_bcn'@'localhost'` IS GOING TO BE CREATED BOTH WITH PASSWORD `password`.
THIS IS DONE FOR BOTH SETTING UP A LOCAL ENVIRONMENT AND FOR RUNNING INTEGRATION TESTS
- Make sure that when starting the server you are passing the flag for `dsn` as this will connect to your local version of the database. The user will need permission to SELECT, INSERT.
- For testing make sure that you pass in the flag `test-dsn` to set for a user that has all priveledges, as we will need to CREATE, DROP etc.

### Goals of this site:
- see list of tournaments
- list a live or online tournament
- shared tournaments through links


### Future goals
- registered members can see who will be attending
- create chess clubs
- chess club managers can manage their clubs
- verification of chess club members (that "John Doe" is the manager or has permission to host chess games at said club)
- After a tournament is complete be able to update metadata related to the tournament for people to later see
- create `match` database where users can save their games, storing PEN so they can see their matches later for viewing(NOSQL database will likely be best option here)
- More will be added as the site comes online
## flags
### reminder
all cli commands can be seen by running `$ go run ./cmd/web -help`

1. `-addr=":<NUMBER>"` - this is used to set the address, default is :4000
2. `-static-dir="<PATH>"` -relative path to your static files directory default: "./ui/static"
3. `-dsn="<user>:<password>"` - usern and password of the user who will be writing to the db. default is "chess-web:password"
4. `-secret="<VALUE>"` - secret key used for the sessions token default: s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge
5. `-secretLifetime="<NUMBER>"` - secret key's lifetime before the key is invalid, default is: 12 hours.