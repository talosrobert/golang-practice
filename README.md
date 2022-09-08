# golang-practice

## Project directory structure
- The __bin__ directory will contain our compiled application binaries, ready for deployment to a production server.
- The __cmd/api__ directory will contain the application-specific code for our API. This will include the code for running the server, reading and writing HTTP requests, and managing authentication.
- The __internal__ directory will contain various ancillary packages used by our API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn’t application-specific and can potentially be reused will live in here. Our Go code under cmd/api will import the packages in the internal directory (but never the other way around).
- The __migrations__ directory will contain the SQL migration files for our database.
- The __remote__ directory will contain the configuration files and setup scripts for our production server.
- The __go.mod__ file will declare our project dependencies, versions and module path.
- The __Makefile__ will contain recipesfor automating common administrative tasks — like auditing our Go code, building binaries, and executing database migrations.

## further page 18
