# WARNING!
Do not use this library in production/enterprise systems. This library is not tested and stable yet.  

# Installation

## Build from source
* Clone the repository
* Navigate to cloned repository folder
* go build main.go

The binary will now be available in the working directory

## Binaries **(NOT IMPLEMENTED YET)**

## Supported Providers

When using drift you can specify a provider - which is the type of database management system that you are using. The table below shows currently supported database management systems.

| Name        | Config Name | Supported |
| ----------- | ----------- | --------- |
| PostgreSQL  | POSTGRES    |  ✅        |
| MySQL       | MYSQL       |  ❌        |
| SQL Server  | SQLSERVER   |  ❌        |


# Commands

### Init
`drift init`

Running this command will create the necessary files/folders in the current working directory.

For example, if I navigate under /home/users/drift and run this command, it will create 2 folders under a root migration folder and 1 config file. 

The folder hierarchy looks like so

    .
    ├── migration                   # The top level migration folder
        ├── patch                   # When you run drift add <migration_name> - the migration is placed here
        ├── seed                    # This is where seed files will live to seed your database migrations (not implemented yet)
        └── config.json             # This is the configuration file - which you will need to fill out to tell drift what provider you are using (database management system), as well as other information
        
        
Currently, the only supported 'provider' is postgres. SQL Server and MySQL are coming.

The configuration file looks like so:
```json
{
  "provider": "",
  "databaseName": "",
  "host": "",
  "username": "",
  "password": "",
  "port": 0
}
```

An example configuration file for POSTGRES would look like so

```json
{
  "provider": "POSTGRES",
  "databaseName": "postgres",
  "host": "127.0.0.1",
  "username": "postgres",
  "password": "password",
  "port": 5432
}
```

### Add

`drift add <migration_name>`

This command adds a new migration under the patch folder. For example

`drift add my_migration` will add two new files under patch with the following template <unix_timestamp>_<migration_name>_up.sql and <unix_timestamp>_<migration_name>_down.sql.

    .
    ├── migration                                             # The top level migration folder
        ├── patch                                             # When you run drift add <migration_name> - the migration is placed here
            ├── 1595849780_my_migration_up.sql                # The migration file used to apply the migration to a database schema
            ├── 1595849780_my_migration_down.sql              # The migration file used to remove the migration from a database schema


### Remove **(NOT IMPLEMENTED YET)**

`drift remove`


### Up
`drift up`

Running this command will create a drift_migrations table in your database, and will apply any migrations that exist under patch that have 'up' suffixed to the migration name.

### Down **(NOT IMPLEMENTED YET)**
`drift down` or `drift down <migration_name>`

Running this command apply any migrations that exist under patch that have 'down' suffixed to the migration name.
