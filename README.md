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

```
themigrator
        When the only way to go is up!

Usage:
  themigrator [flags]
  themigrator [command]

Available Commands:
  apply       Applies pending migrations
  help        Help about any command
  init        Initialises the migration table
  new         Creates a new migration
  verify      Verifies things

Flags:
  -h, --help   help for themigrator

Use "themigrator [command] --help" for more information about a command.
```


#### `themigrator init -d db_name`: Initialise the migration table

Example output:

```
DEBUG[20200206185107] [initialise] checking if migration table exists... 
INFO[20200206185107] [DB_CONNECTION_SUCCEEDED] connection succeeded 
DEBUG[20200206185107] [initialise] creating migration table...     
INFO[20200206185107] [initialise] migration table successfully created 
```

#### `themigrator verify connection`: Verifies database credentials

Example output:

```
DEBUG[20200206185131] [verify-connection] connecting as 'user' to 'localhost:3306/' with parameters map[] 
INFO[20200206185131] [verify-connection] connection credentials verified 
INFO[20200206185131] [verify-connection] DB_CONNECTION_SUCCEEDED: connection succeeded 
```


#### `themigrator verify migrations -d db_name ./path/to/migrations`: Verifies migrations integrity

When no migrations have been done:

```
INFO[20200206185226] [verify-migrations] validating provided input arguments ([./test/migrations/]) 
INFO[20200206185226] [verify-migrations] retrieving local migrations from '/media/z/z/documents/gitlab.com/zephinzer/themigrator/test/migrations'... 
INFO[20200206185226] [verify-migrations] found 3 local migrations 
DEBUG[20200206185226] [verify-migrations] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185226] [verify-migrations] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185226] [verify-migrations] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185226] [verify-migrations] retrieving remote migrations from the database... 
INFO[20200206185226] [verify-migrations] found 0 remote migrations
```

After all migrations have been done:

```
INFO[20200206185427] [verify-migrations] validating provided input arguments ([./test/migrations/]) 
INFO[20200206185427] [verify-migrations] retrieving local migrations from '/media/z/z/documents/gitlab.com/zephinzer/themigrator/test/migrations'... 
INFO[20200206185427] [verify-migrations] found 3 local migrations 
DEBUG[20200206185427] [verify-migrations] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185427] [verify-migrations] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185427] [verify-migrations] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185427] [verify-migrations] retrieving remote migrations from the database... 
INFO[20200206185427] [verify-migrations] found 3 remote migrations 
DEBUG[20200206185427] [verify-migrations] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185427] [verify-migrations] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185427] [verify-migrations] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185427] [verify-migrations] processed 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7 
INFO[20200206185427] [verify-migrations] processed 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3 
INFO[20200206185427] [verify-migrations] processed 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8 
```

#### `themigrator apply ./path/to/migrations -d db_name`: Show the migrations to be applied

When nothing has been applied:

```
INFO[20200206185300] [apply] validating provided input arguments ([./test/migrations/]) 
INFO[20200206185300] [apply] retrieving local migrations from '/media/z/z/documents/gitlab.com/zephinzer/themigrator/test/migrations'... 
INFO[20200206185300] [apply] found 3 local migrations             
DEBUG[20200206185300] [apply] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185300] [apply] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185300] [apply] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185300] [apply] retrieving remote migrations from the database... 
INFO[20200206185300] [apply] found 0 remote migrations            
INFO[20200206185300] [apply] 3 migrations have yet to be applied  
DEBUG[20200206185300] [apply] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185300] [apply] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185300] [apply] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
WARNING[20200206185300] [apply] refusing to apply because --confirm was not specified
```

#### `themigrator apply ./path/to/migrations -d db_name --confirm`: Apply migrations

When nothing has been applied:

```
INFO[20200206185333] [apply] validating provided input arguments ([./test/migrations/]) 
INFO[20200206185333] [apply] retrieving local migrations from '/media/z/z/documents/gitlab.com/zephinzer/themigrator/test/migrations'... 
INFO[20200206185333] [apply] found 3 local migrations             
DEBUG[20200206185333] [apply] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185333] [apply] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185333] [apply] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185333] [apply] retrieving remote migrations from the database... 
INFO[20200206185333] [apply] found 0 remote migrations            
INFO[20200206185333] [apply] 3 migrations have yet to be applied  
DEBUG[20200206185333] [apply] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185333] [apply] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185333] [apply] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185333] [apply] crossing fingers and applying the migrations now... 
INFO[20200206185333] [apply] successfully applied migration '20200201152310_hello_world.sql' 
INFO[20200206185333] [apply] successfully applied migration '20200201160939_add_text_column.sql' 
INFO[20200206185333] [apply] successfully applied migration '20200201161156_add_number_column.sql'
```

When nothing is left to apply:

```
INFO[20200206185355] [apply] validating provided input arguments ([./test/migrations/]) 
INFO[20200206185355] [apply] retrieving local migrations from '/media/z/z/documents/gitlab.com/zephinzer/themigrator/test/migrations'... 
INFO[20200206185355] [apply] found 3 local migrations             
DEBUG[20200206185355] [apply] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185355] [apply] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185355] [apply] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185355] [apply] retrieving remote migrations from the database... 
INFO[20200206185355] [apply] found 3 remote migrations            
DEBUG[20200206185355] [apply] 20200201152310_hello_world.sql/7db3db092838aa64471063337ec542645de4b4a2b71f9671af4984f4702208f7: 'CREATE TABLE hey ( id INTEGER PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT );' 
DEBUG[20200206185355] [apply] 20200201160939_add_text_column.sql/3126eaf6c51501c8f1431f904bf3a6d63ed42cc3ec603bcaefa4a4533b12fbb3: 'ALTER TABLE hey ADD COLUMN textval TEXT;' 
DEBUG[20200206185355] [apply] 20200201161156_add_number_column.sql/9c5927a4493348552e88463ab8d6908ac53d6ae6fd0a043e4238bfc0b41c70f8: 'ALTER TABLE hey ADD COLUMN intval INTEGER NOT NULL;' 
INFO[20200206185355] [apply] all migrations that could've been applied have been
```


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
