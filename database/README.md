# Database Migration

### Using golang migrate

- creating migration

```
$ migrate create -ext sql -dir database/migrations -seq [name]
```

- running migration

```
$ migrate -database "postgres://[uname]:[pass]@tcp([host]:[port])/[db_name]?sslmode=disable" -path database/migrations up
```

- undone all migrations

```
$ migrate -database "postgres://[uname]:[pass]@tcp([host]:[port])/[db_name]?sslmode=disable" -path database/migrations down
```

- rollback migration to specified version

```
$ migrate -database "postgres://[uname]:[pass]@tcp([host]:[port])/[db_name]?sslmode=disable" -path database/migrations goto [version]
```
