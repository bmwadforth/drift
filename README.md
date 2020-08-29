# WARNING!
Do not use this library in production/enterprise systems. This library is not tested and stable yet.  

# Installation

## Build from source
* Clone the repository
* Navigate to cloned repository folder
* go build main.go

The binary will now be available in the working directory

## Binaries

`Not implemented yet` 

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

`drift add my_migration` will add a new file under patch with the following template <unix_timestamp>_<migration_name>.sql

    .
    ├── migration                                             # The top level migration folder
        ├── patch                                             # When you run drift add <migration_name> - the migration is placed here
            ├── 1595849780_my_migration.sql                   # The migration file


### Remove

`drift remove`

This has not been implemented yet

### Up
`drift up`

Running this command will create a drift_migrations table in your database, and will apply any migrations that exist under patch.

### Down
`drift down` or `drift down <migration_name>`

This has not been implemented yet. 
