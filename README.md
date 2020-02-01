# `themigrator`

# Usage



```sh
# create the migrations table
themigrator init -d table_name

# verify your database connection
themigrator verify -d table_name

# create new migration at ./path/to/migrations
themigrator new ./path/to/migrations

# view pending migrations
themigrator plan ./path/to/migrations
```