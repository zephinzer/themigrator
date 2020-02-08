# `themigrator`

[![pipeline status](https://gitlab.com/zephinzer/themigrator/badges/master/pipeline.svg)](https://gitlab.com/zephinzer/themigrator/-/commits/master)
[![Build Status](https://travis-ci.org/zephinzer/themigrator.svg?branch=master)](https://travis-ci.org/zephinzer/themigrator)

Forward looking migrator that only goes one way - UP!

> Note: Still a WIP


- - -


## Usage

### Principles

#### Never look back in regret - move on to the next thing

This migrator was made with only forward migrations in mind.

### CLI

#### `themigrator -h`: Get help on the CLI

Use this when all seems lost and all you have is a terminal

#### `themigrator init -d db_name`: Initialise the migration table

Use this to initialise the database schema in the database

#### `themigrator verify connection`: Verifies database credentials

Use this to verify that the database credentials passed to `themigrator` works

#### `themigrator verify migrations -d db_name ./path/to/migrations`: Verifies migrations integrity

Use this to verify the integrity of the local and remote migrations

#### `themigrator apply ./path/to/migrations -d db_name`: Show the migrations to be applied

Use this to view the migrations planned

#### `themigrator apply ./path/to/migrations -d db_name --confirm`: Apply migrations

Use this to apply the migrations


### Flags

#### `--help`, `-h`

Displays the help message for the CLI

#### `--user`, `-u`

User to use to login to the database

#### `--password`, `-p`

Password of the user used to use to login to the database

#### `--host`, `-H`

Hostname where the database service is reachable at

#### `--port`, `-P`

Port which the database service is listening on

#### `--database`, `-d`

Name of the database to apply the migrations to

#### `--log-level`, `-l`

> This is a **global flag**

A number from 0 to 5 indicating the verbosity of logs.

- 0 silences the logs
- 1 outputs logs at only the `error` level
- 2 outputs logs at the `warning` level and above
- 3 outputs logs at the `info` level and above
- 4 outputs logs at the `debug` level and above
- 5 outputs logs at the `trace` level and above

#### `--log-format`, `-f`

> This is a **global flag**

One of `"json"` or `"text"`. Defaults to `"text"`.



- - -


## Development

### Setting up

THe `init` recipe in the `Makefile` will set up an environment with MySQL and MariaDB instances.

### Shell access to the database

To get a shell to the MySQL database, use `make check_mysql`.

To get a shell to the Maria database, use `make check_maria`.

### Manual Testing

- There are sample migrations located at `./test/migrations`

### Tearing down

The `denit` recipe in the `Makefile` should do the job.


- - -


## Licensing

Code is licensed under [the MIT license](./LICENSE).
