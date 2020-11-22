# bcnchess

Barcelona Chess (bcnchess) is a free open source website that is aimed at making meet ups of tournaments within Barcelona/Spain easier. Currently the situation faced is that all the tournaments
are shared through whatsapp which makes it easy to miss one or not know about the others.

### Getting Started
- *IMPORTANT* WHILE RUNNING THE SET UP SCRIPT A NEW USER OF `'chess-web'@'localhost'` *AND* `'test_bcn'@'localhost'` IS GOING TO BE CREATED BOTH WITH PASSWORD `password`.
THIS IS DONE FOR BOTH SETTING UP A LOCAL ENVIRONMENT AND FOR RUNNING INTEGRATION TESTS
- Make sure that when starting the server you are passing the flag for `dsn` as this will connect to your local version of the database. The user will need permission to SELECT, INSERT.
- For testing make sure that you pass in the flag `test-dsn` to set for a user that has all priveledges, as we will need to CREATE, DROP etc.

- new with [#16](https://github.com/ctfrancia/bcnchess/issues/16) is a WIP. While the server can be built with the docker file by running: `$ docker build -t chess-server .` while in the root
folder the issue is currently connecting to the mysql database

#### Steps with running locally (Docker)
*note* need to have docker installed locally sorry for the long steps this will be automated better.

run these commands in order from the root of this folder:

- `$ docker network create bcnchess`
- `$ docker build -t chess-db ./dev/`
- `$ docker run -d --net bcnchess --name chess-db -p 3306:3306  chess-db`
- `$ docker build -t chess-server .`
- `$ docker run -d --net bcnchess --name go-server -p 4000:4000 chess-server`

## Troubleshooting
- if you don't see the two running then check the logs of the container with `$ docker container logs <CONTAINER_NAME>`
- adjust the error, if it's something related to the build then please make an issue.

### Goals of this site:
- see list of tournaments
- list a live or online tournament
- shared tournaments through links
- chat with people through the website

### Future goals
*see issues to see what it currently there*
- registered members can see who will be attending
- create chess clubs
- chess club managers can manage their clubs
- verification of chess club members (that "John Doe" is the manager or has permission to host chess games at said club)
- After a tournament is complete be able to update metadata related to the tournament for people to later see
- create `match` database where users can save their games, storing PEN so they can see their matches later for viewing(NOSQL database will likely be best option here)
- More will be added as the site comes online

## Getting Started
- The first thing that you will need to do is go to their [official install website](https://golang.org/doc/install) and download the source code for your OS. You can check the installation was correct by doing `go version`
- Note: that in your rc file (`.bachrc` or `.zshrc` you will need to export the path. if you folowed the instructions on the official website then you should be able to export the gopath with: `export GOPATH=$HOME/Documents/goworkspace`). The `$HOME` variable is equvilent to `~` command in the linux shell.
- after following the instructions create a folder somewhere on your machine where you will have all the Go code. For example on my machine(MacOS) all my Go code is located at `/Users/<USER_NAME>/Documents/goworkspace` when there I have three three folders `bin` `pkg` `src` so run the command `mkdir pkg src bin` this will create the three folders on linux/mac. 
- `bin` holds the binaries, `pkg` holds the packages, and `src` holds all the source files(go code)
- cd into `src` and create a folder called `github.com`, this is a convension the idea is that there needs to be a unique name to prevent any name collision within packages. After that, cd into `github.com` and create a folder with your username on github. Once inside there you can clone the repo into that file path.
- for example the full path to the code on my machine is: `/Users/<USER_NAME>/Documents/goworkspace/src/github.com/ctfrancia/bcnchess`
- After all that is done. make sure you have MySql downloaded and installed with all the correct permissions. (the script isn't done yet to automate the db set up and seed, however, you will see the code in the path `/pkg/models/mysql/db.setup.sql`) which you can copy and paste into your terminal as needed.
- to launch the application. from the root directory (`/github.com/ctfrancia/bcnchess`) run the command `go run ./cmd/web` and if all is working you will see the port that it is listening on (default is 4000)


## Flags
### Reminder
all cli commands can be seen by running `$ go run ./cmd/web -help`

*note* as this server is moving to an api many of these won't be needed anyone
1. `-addr=":<NUMBER>"` - this is used to set the address, default is :4000
2. `-static-dir="<PATH>"` -relative path to your static files directory default: "./ui/static"
3. `-dsn="<user>:<password>"` - usern and password of the user who will be writing to the db. default is "chess-web:password"
4. `-secret="<VALUE>"` - secret key used for the sessions token default: s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge *note* this is for dev, and will not work in the production :P
5. `-secretLifetime="<NUMBER>"` - secret key's lifetime before the key is invalid, default is: 12 hours.
6. `-debug=<BOOL>` - debug mode (errors with stack stracing shows up in the browser instead of the terminal)
