# `themigrator`

[![pipeline status](https://gitlab.com/zephinzer/themigrator/badges/master/pipeline.svg)](https://gitlab.com/zephinzer/themigrator/-/commits/master)
[![Build Status](https://travis-ci.org/zephinzer/themigrator.svg?branch=master)](https://travis-ci.org/zephinzer/themigrator)

Forward looking migrator that only goes one way - UP!

> Note: Still a WIP

## Usage

### Principles

#### Never look back in regret - move on to the next thing

This migrator was made with only forward migrations in mind.

### CLI

```sh
# display the help message on the cli
themigrator -h

# create the migrations table
themigrator init -d table_name

# verify your database connection
themigrator verify connection

# create new migration at ./path/to/migrations
themigrator new ./path/to/migrations

# view pending migrations
themigrator apply ./path/to/migrations -d database_name ./path/to/migrations

# apply pending migrations
themigrator apply ./path/to/migrations -d database_name ./path/to/migrations --confirm
```

### Flags

| Long | Short | Description |
| --- | --- | --- |
| `--help` | `-h` | Displays the help message for the CLI |
| `--user` | `-u` | User to use to login to the database |
| `--password` | `-p` | Password of the user used to use to login to the database |
| `--host` | `-H` | Hostname where the database service is reachable at |
| `--port` | `-P` | Port which the database service is listening on |
| `--database` | `-d` | Name of the database to apply the migrations to |
| `--driver` | `-D` | Type of database you're applying the migration on |

## Development

### Setting up

THe `init` recipe in the `Makefile` will set up an environment with MySQL, MariaDB, and PostgreSQL instances.

### Tearing down

The `denit` recipe in the `Makefile` should do the job.

## Licensing

Code is licensed under [the MIT license](./LICENSE).
