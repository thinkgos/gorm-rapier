# Update

- [Update](#update)
  - [`Save` will save all fields](#save-will-save-all-fields)
  - [Update single column](#update-single-column)
  - [Updates multiple columns](#updates-multiple-columns)
  - [Update from SubQuery](#update-from-subquery)
  - [Without Hooks/Time Tracking](#without-hookstime-tracking)

## `Save` will save all fields

`Save` will save all fields when performing the Updating SQL

```go
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
    Model().
    Save(&testdata.Dict{
        Id:     100,
        Key:    "k1",
        Remark: "remark1",
    })
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1",`name`="",`is_pin`=false,`remark`="remark1",`created_at`="2024-03-07 01:53:14.633",`updated_at`="2024-03-07 01:53:14.633" WHERE `id` = 100
```

## Update single column

```go
refDict := testdata.Ref_Dict()
// update with expr
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateExpr(refDict.Key, "k1")
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1",`updated_at`="2024-03-07 02:10:44.258" WHERE `dict`.`id` = 100

// update SetExpr with expr
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateExpr(refDict.UpdatedAt, refDict.CreatedAt.Add(time.Second))
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `updated_at`=DATE_ADD(`dict`.`created_at`, INTERVAL 1000000 MICROSECOND) WHERE `dict`.`id`

// update with original gorm api
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    Update("key", "k1")
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1",`updated_at`="2024-03-07 02:10:44.258" WHERE `dict`.`id` = 100
```

## Updates multiple columns

```go
refDict := testdata.Ref_Dict()
// update with expr
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdatesExpr(
        refDict.Key.Value("k1"),
        refDict.Remark.Value(""),
    )
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1",`remark`="",`updated_at`="2024-03-07 02:19:10.144" WHERE `dict`.`id` = 100

// update use `struct` with original gorm api
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    Updates(&testdata.Dict{
        Key:    "k1",
        Remark: "remark1",
    })
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1",`remark`="remark1",`updated_at`="2024-03-07 02:19:10.144" WHERE `dict`.`id` = 100

// update use map with original gorm api
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdatesMap(map[string]any{
        "key":    "k1",
        "remark": "remark1",
    })
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1",`remark`="remark1",`updated_at`="2024-03-07 02:19:10.144" WHERE `dict`.`id` = 100
```

## Update from SubQuery

```go
refDict := testdata.Ref_Dict()
// update with expr
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateExpr(
        refDict.Key,
        rapier.NewExecutor[testdata.Dict](db).Model().
            SelectExpr(refDict.Key).
            Where(refDict.Id.Eq(101)).
            IntoDB(),
    )
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`=(SELECT `dict`.`key` FROM `dict` WHERE `dict`.`id` = 101),`updated_at`="2024-03-07 02:41:40.548" WHERE `dict`.`id` = 100

// update with exprs
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdatesExpr(
        refDict.Key.ValueSubQuery(
            rapier.NewExecutor[testdata.Dict](db).Model().
                SelectExpr(refDict.Key).
                Where(refDict.Id.Eq(101)).
                IntoDB(),
        ),
    )
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`=(SELECT `dict`.`key` FROM `dict` WHERE `dict`.`id` = 101),`updated_at`="2024-03-07 02:41:40.548" WHERE `dict`.`id` = 100

// update use map with original gorm api
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdatesMap(map[string]any{
        "key": rapier.NewExecutor[testdata.Dict](db).Model().
            SelectExpr(refDict.Key).
            Where(refDict.Id.Eq(101)).
            IntoDB(),
    })
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`=(SELECT `dict`.`key` FROM `dict` WHERE `dict`.`id` = 101),`updated_at`="2024-03-07 02:41:40.548" WHERE `dict`.`id` = 100
```

## Without Hooks/Time Tracking

```go
refDict := testdata.Ref_Dict()
// update with expr
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateColumnsExpr(
        refDict.Key.Value("k1"),
        refDict.Remark.Value(""),
    )
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1",`remark`="" WHERE `dict`.`id` = 100

// update with expr
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateColumnExpr(refDict.Key, "k1")
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1" WHERE `dict`.`id` = 100

// update with original gorm api
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateColumn("key", "k1")
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1" WHERE `dict`.`id` = 100

// update with original gorm api
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateColumns(&testdata.Dict{
        Key: "k1",
    })
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1" WHERE `dict`.`id` = 100

// update with original gorm api
rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdateColumnsMap(map[string]any{
        "key": "k1",
    })
_ = err          // return error
_ = rowsAffected // return row affected
// UPDATE `dict` SET `key`="k1" WHERE `dict`.`id` = 100
```

more information see [gorm Update](https://gorm.io/docs/update.html)
