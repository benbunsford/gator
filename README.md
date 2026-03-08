# gator
## Requirements
- [Go](https://go.dev/dl/) (version 1.25+)
- [Postgres](https://www.postgresql.org/download/) (version 15+)
- [goose](https://github.com/pressly/goose) (to run database migration after install) Install with:
    ```
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```
## Configuration
### Setting up Postgres
1. Set up a database named `gator` with Postgres. Instructions vary by OS. On Linux:
    1. If you haven't already, set a password for the `postgres` user, which was created when you installed postgres:
    ```
    sudo passwd postgres <desired_password>
    ```
    2. Start the postgres service:
    ```
    sudo service postgresql start
    ```
    3. Connect to the server:
    ```
    sudo -u postgres psql
    ```
    4. Create the database:
    ```
    postgres=# CREATE DATABASE gator;
    ```
    5. Connect to the new database:
    ```
    postgres=# \c gator
    ```
    6. Set the database user password(we set the system user before):
    ```
    gator=# ALTER USER postgres PASSWORD 'postgres';
    ```

1. Create a config file in your home directory `~/.gatorconfig.json` with this content:
```
{
  "db_url": "<your_db_string>"
}
```
Format your db string like this:
```
protocol://username:password@host:port/database?sslmode=disable

examples:
postgres://ben:password@localhost:5432/gator?sslmode=disable

postgres://postgres:postgres@localhost:5432/gator?sslmode=disable
```
### Installing gator
You'll use your db string again here, but without the parameter `?sslmode=disable` at the end:
```
git clone github.com/benbunsford/gator@latest
cd gator
goose postgres <your_db_string_no_param> up
go install
```
## Usage
- Register a new user: `gator register <username>`
- Login to a different user account: `gator login <username>`
- List all users: `gator users`
- Add a new RSS feed to the database: `gator addfeed <feed_name> <feed_url>`
    - The current user automatically follows the added feed.
- Follow an existing feed: `gator follow <feed_url>`
- List all feeds: `gator feeds`
- List feeds followed by current user: `gator following`
- Unfollow a feed: `gator unfollow`
- Aggregate posts from feeds: `gator agg <duration>`
    - duration string example formats: 1s, 1m, 1h, 1m1s, 1h1m1s
    - don't make requests to quickly!
    - to stop aggregating, exit the program with ctrl+c
- Browse posts from followed feeds: `gator browse <limit>'
    - limit determines how many posts to display
    - default value if left blank is 2
