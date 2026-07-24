# Gator

Gator is a simple command-line interface tool that lets you read and track multiple RSS feeds right from your terminal.

## Setup

Before you install Gator, make sure you have these tools on your computer:

- Go (version 1.26 or higher) `brew install go`
- PostgreSQL (version 15) `brew install postgresql@15`

## Installation

1. Clone the repository:

```bash
git clone https://github.com
cd gator
```

2. Build and install:

```bash
go build -o gator
```

## Config

Create a configuration file `.gatorconfig.json` in your home directory. You will need to replace `your_username` and `your_password` with the appropriate values (password is optional for postgres on OSX).

```json
{
  "db_url": "postgres://your_username:your_password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "placeholder"
}
```

## Usage

### example usage

```bash
# register a user
gator register examplename

# add a feed
gator addfeed h4X0rn00z https://news.ycombinator.com/rss

# tell gator check for new content every 10 seconds
gator agg 10

# kill the process by holding ctrl-c
## TODO: make a better UX flow. one that doesn't require the user killing a process.

# finally, browse your content
gator browse 10s

# bask in a sense of pride an accomplishment. well done, you!

# finally finally, clean up after yourself
# this will delete all user and feed data
gator reset
```

### options

You can call `gator` with any of the following options:

```bash
gator register <username> # register a user, logged in automatically
gator login <username> # switch current user
gator reset # clears all user and feed data
gator users # lists all registered users
gator agg <time_between_reqs> # gator will loop, checking for new RSS content every <time>
gator feeds # list all saved RSS feeds
gator addfeed <name> <url> # add a feed, the current user will follow automatically
gator follow <url> # follow a feed that another user added
gator following # list all feeds the current user follows
gator unfollow <url> # unfollow a feed
gator browse <optional limit> # display 2 or <optional limit> new posts from the current user's feeds
```

## thanks!

gator is a guided project for the boot.dev educational platform. check them out!


