# AggreGATOR

This is a project developed within [BootDev](http://boot.dev) backend track.

## Requirements 

- PostgreSQL >= 15
- Go >= 1.23

## Installation

To install `gator`, ensure you have go installed and run:

```bash
go install github.com/araujoarthur/aggregator
```

## Configuration

1. Create a file called `.gatorconfig.json` at your home directory.
2. Fill the contents

```json
{"db_url":"POSTGRES CONNECTION STRING"}
```

The connection string should look like something as `postgresql://USERNAME@HOST/gator?sslmode=disable`

## Usage

After installation and configuration, you can register a user into the tool:

```bash
aggregator register USERNAME
```

Then you can use any of the following commands

```bash
reset # Resets the users table
login # changes the current user
addfeed [url] # adds a feed
(un)follow [url] # follows/unfollows a feed
users # list users
feeds # list available feeds
following # lists following feeds
agg [time-between-reqs] # fetch posts and save them to database once each "time-between-reqs" duration.
browse [limit] # limit is optional, defaults to 2. Lists already downloaded posts for feeds you follow.
```