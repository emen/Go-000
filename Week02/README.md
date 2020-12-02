# hmmm, errors...

## Example

```
api (./api/server.go) -> biz (./model/user.go) -> dao (./model/mysql/mysql.go)
```

1. The DAO layer specific errors (implementation details) should not be exposed to API. Otherwise, there is coupling
and makes it harder to switch DB (e.g. from MySQL -> MongoDB, etc) 

2. Two types of errors could happen at DAO layer.
    1. known errors that should map to a biz layer error, e.g. `sql.ErrNoRows`. And in api layer, it is checked by `errors.Is`
  and mapped to a specific http response (in this case, 404).
    2. unexpected errors. There is no mapped error in biz layer, and api layer should do a catch-all check `err != nil` and
  normally returns 503 server error.




