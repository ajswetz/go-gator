# go-gator
Blog AggreGATOR written in Golang

## Requirements

The following software is required to install and use `go-gator`:

### Go toolchain
You need to install the Go toolchain. Installation instructions can be found on the official Go website here: [Go - Download and Install](https://go.dev/doc/install)

### Postgres database
`go-gator` utilizes a local PostgreSQL database. You can download it here: [PostgreSQL Downloads](https://www.postgresql.org/download/)

### Goose (_optional - quickly initialize the database_)
You will need to create the initial database schema before using `go-gator`. The fastest way to do this involves using the [`goose` database migration tool](https://github.com/pressly/goose). After you have installed Go, you can install `goose` using the following command:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Alternatively, you can setup the database manually by referencing the SQL files in the 'sql/schema' directory. **Do this at your own risk.** If you do not prepare the database exactly as specified in these schema files, you will likely run into problems using `go-gator`.

## Installation

Once you have installed the above mentioned requirements, you have two options for installing `go-gator`:

1. To install `go-gator` globally for repeated use, simply run this command:
```bash
go install github.com/ajswetz/go-gator
```

2. If you want to tinker with `go-gator` or run the binary locally without installing it, you will first need to download a copy of the raw Go files. Then, you can use the `go build` command to create the binary. Here's a full example:
```bash
git clone https://github.com/ajswetz/go-gator
cd go-gator
go build
```

## Setup

Before you can use `go-gator`, you need to complete some initial setup.

### PostgreSQL

1. Ensure the PostgreSQL installation worked. The `psql` command-line utility is the default client for Postgres. Use it to make sure you're on version 16+ of Postgres:

```bash
psql --version
```

2. (**Linux only**) Update postgres password:

```bash
sudo passwd postgres
```

Enter a password, and be sure you won't forget it. You can just use something easy like `postgres`.

3. Start the Postgres server in the background

- Mac: `brew services start postgresql`
- Linux: `sudo service postgresql start`

4. Connect to the server. You can use the default `psql` client included with the PostgreSQL install.

Enter the `psql` shell:

- Mac: `psql postgres`
- Linux: `sudo -u postgres psql`

You should see a new prompt that looks like this:

```bash
postgres=#
```

5. Create a new database. I called mine `gator`:

```sql
CREATE DATABASE gator;
```

6. Connect to the new database:

```bash
\c gator
```

You should see a new prompt that looks like this:

```bash
gator=#
```

7. (**Linux only**) Set the user password

```sql
ALTER USER postgres PASSWORD 'postgres';
```

For simplicity, I used `postgres` as the password. Before, we altered the _system_ user's password, now we're altering the _database_ user's password.

8. Exit the `psql` shell by running `exit`

9. Get your database connection string.

A connection string is just a URL with all of the information needed to connect to a database. The format is:
protocol://username:password@host:port/database

Here are examples:

* Mac OS (no password, your username): `postgres://myusername:@localhost:5432/gator`

* Linux (postgres user : password set in step #7 above): `postgres://postgres:postgres@localhost:5432/gator`

Test your connection string by running psql, for example:

```bash
psql "postgres://myusername:@localhost:5432/gator"
```

It should connect you to the gator database directly. If it's working, great. Exit out of psql and save the connection string for later.

10. Use `goose` to 'migrate' the database (_i.e.: create necessary tables and columns_)

In your terminal, `cd` into the sql/schema directory and run this command, replacing <my_connection_string> with the string you prepared in the previous step:

```
goose postgres <my_connection_string> up
```

To confirm whether the tables have been created successfully, run these commands:

```
psql <my_connection_string>
\dt
```

The output should look something like this:

```
gator=# \dt
              List of relations
 Schema |       Name       | Type  |  Owner   
--------+------------------+-------+----------
 public | feed_follows     | table | postgres
 public | feeds            | table | postgres
 public | goose_db_version | table | postgres
 public | posts            | table | postgres
 public | users            | table | postgres
(5 rows)
```

### go-gator Configuration File

Finally, you will need to create a JSON configuration file that will live in your profile's home directory.

Create a config file at the following location: `~/.gatorconfig.json`

The file should have the following contents:

```JSON
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}           
```

**IMPORTANT: Be sure to replace the contents of the `db_url` field with your postgres connection string. Also, make sure to append `?sslmode=disable` to the end of the string, as shown in the example.**

You should now be ready to use `go-gator`!

## Usage

`go-gator` currently supports the following commands:

### register
The `register` command takes one argument, a name, and creates a new user in the database. It also sets this user as the "current_user_name" in the `.gatorconfig.json` file.

Example:
```
go-gator register john
```
```
New user 'john' was successfully created
User details:
{
  "ID": "18e41bc1-5c35-4a89-8f7b-42f5ed6b1f17",
  "CreatedAt": "2024-11-15T16:40:32.953237Z",
  "UpdatedAt": "2024-11-15T16:40:32.953237Z",
  "Name": "john"
}
```

### login
The `login` command takes one argument, a name, and sets this user as the "current_user_name" in the `.gatorconfig.json` file.

Example:
```
go-gator login john
```
```
john is now logged in
```

If this user has not yet been created in the database using the `register` command, the `login` command will return an error.

Example:
```
go-gator login frank
```
```
Unable to get user 'frank' from the database: sql: no rows in result set
```


### users
The `users` command takes no arguments and returns a list of all users currently registered in the database. It will also note which user is the "currently logged in" user.

Example:
```
go-gator users
```
```
* john (current)
* tim
* mary
```

### addfeed
The `addfeed` command takes two arguments, a 'name' of a blog (_can be anything you want_) and a valid URL of that blog's RSS feed. It will add a new RSS feed to the database. It will also automatically 'follow' that feed for the current user.

Example:
```
go-gator addfeed "JetBrains Go Blog" "https://blog.jetbrains.com/go/feed/"
```
```
Successfully added new feed to the database
{
  "ID": "bcaf7137-de0a-48b0-930e-76d0d7ef5b4c",
  "CreatedAt": "2024-11-15T16:50:21.127404Z",
  "UpdatedAt": "2024-11-15T16:50:21.127404Z",
  "Name": "JetBrains Go Blog",
  "Url": "https://blog.jetbrains.com/go/feed/",
  "UserID": "18e41bc1-5c35-4a89-8f7b-42f5ed6b1f17",
  "LastFetchedAt": {
    "Time": "0001-01-01T00:00:00Z",
    "Valid": false
  }
}
john is now following feed 'JetBrains Go Blog'
```

### feeds
The `feeds` command takes no arguments and returns a list of all RSS feeds currently saved in the database. It also shows the name of the user who initially registered that feed.

Example:
```
go-gator feeds
```
```
[
  {
    "FeedName": "Hacker News",
    "Url": "https://news.ycombinator.com/rss",
    "UserName": "mary"
  },
  {
    "FeedName": "Boot.dev Blog",
    "Url": "https://blog.boot.dev/index.xml",
    "UserName": "tim"
  },
  {
    "FeedName": "Tech Crunch",
    "Url": "https://techcrunch.com/feed/",
    "UserName": "tim"
  },
  {
    "FeedName": "JetBrains Go Blog",
    "Url": "https://blog.jetbrains.com/go/feed/",
    "UserName": "john"
  }
]
```

### follow
The `follow` command takes one argument, the URL of a feed to follow. It establishes a relationship between the current user and the specified feed by utilizing the `feed_follows` table. 

Example:
```
go-gator follow "https://techcrunch.com/feed/"
```
```
john is now following feed 'Tech Crunch'
```

### unfollow
The `unfollow` command takes one argument, the URL of a feed to _un_follow. It removes the relationship between the current user and the specified feed by updating the `feed_follows` table.

Example:
```
go-gator unfollow "https://techcrunch.com/feed/"
```
```
john has successfully unfollowed feed 'https://techcrunch.com/feed/'
```

### following
The `following` command takes no arguments and returns a list of all RSS feeds currently followed by the logged in user.

Example:
```
go-gator following
```
```
john is currently following these feeds:
 - Tech Crunch
 - JetBrains Go Blog
 ```

### agg
The `agg` command takes one parameter, a stringified duration value (e.g.: 1s, 3m, 2h). The `agg` command is a never-ending loop that fetches feeds and saves posts to the database. The intended use case is to leave the agg command running in the background while you interact with the program in another terminal. The `agg` command will attempt to fetch new feeds once every `time.Duration` passed in as the parameter with the initial command. You can kill the `agg` loop with Ctrl+C.

Example:
```
go-gator agg 30s
```
```
Collecting feeds every 30s...
Running scrapeFeeds()...
```

### browse
The `browse` command takes no mandatory parameters and returns a list of recent posts stored in the database associated with the feeds followed by the logged in user. By default, the `browse` command returns the most recent two posts.

The `browse` command can take one option parameter, a 'limit' that will adjust the number of posts returned.

Example:
```
go-gator browse
```
```
Here are your most recent posts:

Title: ‘I went to Greenland to try to buy it’: Meet the founder who wants to recreate Mars on Earth
Description: <p>Praxis co-founder Dryden Brown wants to build a city in Greenland that emulates what a community on Mars could be like.</p>
<p>© 2024 TechCrunch. All rights reserved. For personal use only.</p>

Publication Date: 2024-11-15 19:01:40
Link: https://techcrunch.com/2024/11/15/i-went-to-greenland-to-try-to-buy-it-meet-the-founder-who-wants-to-re-create-mars-on-earth/
--------------------------------------------------------------------------------

Title: Think you need a VPN? Start here.
Description: <p>Not everyone actually needs to use a VPN. This simple guide will help you decide if you need a VPN for your situation. </p>
<p>© 2024 TechCrunch. All rights reserved. For personal use only.</p>

Publication Date: 2024-11-15 19:00:00
Link: https://techcrunch.com/2024/11/15/think-you-need-a-vpn-guide-start-here/
--------------------------------------------------------------------------------
```

```
go-gator browse 3
```
```
Here are your most recent posts:

Title: ‘I went to Greenland to try to buy it’: Meet the founder who wants to recreate Mars on Earth
Description: <p>Praxis co-founder Dryden Brown wants to build a city in Greenland that emulates what a community on Mars could be like.</p>
<p>© 2024 TechCrunch. All rights reserved. For personal use only.</p>

Publication Date: 2024-11-15 19:01:40
Link: https://techcrunch.com/2024/11/15/i-went-to-greenland-to-try-to-buy-it-meet-the-founder-who-wants-to-re-create-mars-on-earth/
--------------------------------------------------------------------------------

Title: Think you need a VPN? Start here.
Description: <p>Not everyone actually needs to use a VPN. This simple guide will help you decide if you need a VPN for your situation. </p>
<p>© 2024 TechCrunch. All rights reserved. For personal use only.</p>

Publication Date: 2024-11-15 19:00:00
Link: https://techcrunch.com/2024/11/15/think-you-need-a-vpn-guide-start-here/
--------------------------------------------------------------------------------

Title: This ‘AI Granny’ hack wastes telephone scammers’ time with boring chit-chat
Description: <p>Telephone scams are nothing new, but with the advent of AI, it has become harder than ever for people to know whether the person they’re speaking to is in fact who they say they are. But U.K. mobile network O2 is turning the tables, creating what it calls an “AI granny” to keep scammers on […]</p>
<p>© 2024 TechCrunch. All rights reserved. For personal use only.</p>

Publication Date: 2024-11-15 18:11:21
Link: https://techcrunch.com/2024/11/15/ai-granny-scambaiter-wastes-telephone-fraudsters-time-with-boring-chat/
--------------------------------------------------------------------------------
```

### reset
**CAUTION: DANGER ZONE**

The `reset` command will delete all data stored in the database, including users, feeds, user-feed follows, and posts. It is primarily meant as a development convenience. Only run the `reset` command if you want to clear the database and start from scratch.

Example:
```
go-gator reset
```
```
All users successfully deleted from the database
```