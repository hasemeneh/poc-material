### Driver DB
Use MariaDB

### Folder Structure

#### migration
List of DDL that have been happening so far.


#### codegen
Inside this folder, there are other folders which is a representation of repository.
Inside those folder, there are 2 important files:
1. query.sql; List of queries what used on the repository
2. schema.sql; Partial DDL for tables, must be updated every time there is a schema change.
Note: Since using soft delete method, the implication is all queries need adding clause
```
deleted_at IS NULL
```
For retrieving deleted data, will be use other function, e.g `GetDeletedByID`, etc.


### Code Generator

`codegen` folder will be used for code generator from query to Go which is type-safe code.
The generator is [sqlc](https://github.com/kyleconroy/sqlc), using `sqlc.yaml` as sqlc configuration.
Please check sqlc docs [here](https://docs.sqlc.dev/en/stable/) for more details.

### How to generate code
After install `sqlc`, just use CLI `sqlc generate`

### Notes for codegen
There is limitation of current build `sqlc` for MySQL driver, such as
- `IN` clause not working
- `LIKE` clause not working 
        workaround: just edit the codegen-ed code.
- queries can not be exported