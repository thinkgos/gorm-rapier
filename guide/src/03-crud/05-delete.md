# Delete

```go
refDict := testdata.Ref_Dict()
    rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
        Model().
        Where(refDict.Id.Eq(100)).
        Delete()
    _ = err          // return error
    _ = rowsAffected // return row affected
 // DELETE FROM `dict` WHERE `dict`.`id` = 100
```

more information [gorm Delete](https://gorm.io/docs/delete.html)
