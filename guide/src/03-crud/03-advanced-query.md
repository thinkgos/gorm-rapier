# Advanced Query

- [Advanced Query](#advanced-query)
  - [Locking](#locking)
  - [SubQuery](#subquery)
  - [From SubQuery](#from-subquery)
  - [IN with multiple columns](#in-with-multiple-columns)
  - [FirstOrInit](#firstorinit)
  - [FirstOrCreate](#firstorcreate)
  - [Pluck](#pluck)
  - [Count](#count)
  - [Exist](#exist)
  - [Function](#function)
    - [Case When](#case-when)
    - [Concat](#concat)
    - [IF](#if)

more information [gorm Advanced Query](https://gorm.io/docs/advanced_query.html)

## Locking

```go
// Basic FOR UPDATE lock
_, _ = rapier.NewExecutor[testdata.Dict](db).
    LockingUpdate().
    TakeOne()
// SELECT * FROM `dict` LIMIT 1 FOR UPDATE
// Basic FOR UPDATE lock with Clauses api
_, _ = rapier.NewExecutor[testdata.Dict](db).
    Clauses(clause.Locking{Strength: "UPDATE"}).
    TakeOne()
// SELECT * FROM `dict` LIMIT 1 FOR UPDATE

// Basic FOR SHARE lock
_, _ = rapier.NewExecutor[testdata.Dict](db).
    LockingShare().
    TakeOne()
// SELECT * FROM `dict` LIMIT 1 FOR SHARE
// Basic FOR SHARE lock with Clauses api
_, _ = rapier.NewExecutor[testdata.Dict](db).
    Clauses(clause.Locking{Strength: "SHARE"}).
    TakeOne()
// SELECT * FROM `dict` LIMIT 1 FOR SHARE

// Basic FOR UPDATE NOWAIT lock with Clauses api
_, _ = rapier.NewExecutor[testdata.Dict](db).
    Clauses(clause.Locking{Strength: "UPDATE",Options: "NOWAIT"}).
    TakeOne()
// SELECT * FROM `dict` LIMIT 1 FOR UPDATE NOWAIT
```

## SubQuery

```go
refDict := testdata.Ref_Dict()

_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Id,
        refDict.Key,
        rapier.NewExecutor[testdata.Dict](db).
            SelectExpr(rapier.Star.Count()).
            Where(
                refDict.Name.Eq("kkk"),
            ).
            IntoSubQueryExpr().As("total"),
    ).
    Where(refDict.Key.LeftLike("key")).
    Find(&struct{}{})
// SELECT `dict`.`id`,`dict`.`key`,(SELECT COUNT(*) FROM `dict` WHERE `dict`.`name` = "kkk") AS `total` FROM `dict` WHERE `dict`.`key` LIKE "key%"

_, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(refDict.Key.EqSubQuery(
        rapier.NewExecutor[testdata.Dict](db).
            SelectExpr(refDict.Key).
            Where(refDict.Id.Eq(1001)).
            IntoDB(),
    )).
    FindAll()
// SELECT * FROM `dict` WHERE `dict`.`key` = (SELECT `dict`.`key` FROM `dict` WHERE `dict`.`id` = 1001)
```

## From SubQuery

```go
refDict := testdata.Ref_Dict()
_, _ = rapier.NewExecutor[testdata.Dict](db).
    TableExpr(
        rapier.From{
            Alias: "u",
            SubQuery: rapier.NewExecutor[testdata.Dict](db).
                SelectExpr(refDict.Key).
                IntoDB(),
        },
        rapier.From{
            Alias: "p",
            SubQuery: rapier.NewExecutor[testdata.Dict](db).
                SelectExpr(refDict.Key).
                IntoDB(),
        },
    ).
    FindAll()
// SELECT * FROM (SELECT `dict`.`key` FROM `dict`) AS `u`, (SELECT `dict`.`key` FROM `dict`) AS `p`
```

## IN with multiple columns

supports the IN clause with multiple columns, allowing you to filter data based on multiple field values in a single query.

```go
refDict := testdata.Ref_Dict()

record1, _ := rapier.NewExecutor[testdata.Dict](db).
    Where(
        rapier.NewColumns(refDict.Name, refDict.IsPin).
            In([][]any{{"name1", true}, {"name2", false}}),
    ).
    FindAll()
_ = record1
// SELECT * FROM `dict` WHERE (`dict`.`name`, `dict`.`is_pin`) IN (("name1",true),("name2",false))
record2, _ := rapier.NewExecutor[testdata.Dict](db).
    Where(
        rapier.NewColumns(refDict.Name, refDict.IsPin).
            In(
                rapier.NewExecutor[testdata.Dict](db).
                    SelectExpr(refDict.Name, refDict.IsPin).
                    Where(refDict.Id.In(10, 11)).
                    IntoDB(),
            ),
    ).
    FindAll()
_ = record2
// SELECT * FROM `dict` WHERE (`dict`.`name`,`dict`.`is_pin`) IN (SELECT `dict`.`name`,`dict`.`is_pin` FROM `dict` WHERE `dict`.`id` IN (10,11))
```

## FirstOrInit

`FirstOrInit` method is utilized to fetch the first record that matches given conditions, or initialize a new instance if no matching record is found. This method allows additional flexibility with the `Attrs`, `Assign`, `AttrsExpr`, `AssignExpr` methods.

- `Attrs`, `AttrsExpr`: When no record is found, you can use `Attrs`,`AttrsExpr` to initialize a struct with additional attributes. These attributes are included in the new struct but are not used in the SQL query.
- `Assign`, `AssignExpr` method allows you to set attributes on the struct regardless of whether the record is found or not. These attributes are set on the struct but are not used to build the SQL query and the final data won’t be saved into the database.

***NOTE!!!***: if with expr condition will not initialize the field when initializing, so we should use `Attrs`, `AttrsExpr`, `Assign`, `AssignExpr` attributes to indicate these fields.

```go
refDict := testdata.Ref_Dict()
// NOTE!!!: if with expr condition will not initialize the field when initializing, so we should use
// `Attrs`, `AttrsExpr`, `Assign`, `AssignExpr` attributes to indicate these fields.

// `Attrs`, `AttrsExpr`
// with expr
newdict, _ := rapier.NewExecutor[testdata.Dict](db).
    Where(refDict.Name.Eq("myname")).
    AttrsExpr(refDict.Remark.Value("remark11")).
    FirstOrInit()
_ = newdict
// NOTE: Condition use expr. here will not initialize the field of the condition when initializing.
// if not found
// newdict -> Dict{ Remark: "remark11" }
//
// if found, `Attrs`, `AttrsExpr` are ignored
// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark" }

// with original gorm api
newdict, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(&testdata.Dict{
        Name: "non_existing",
    }).
    FirstOrInit()
_ = newdict
// NOTE: Condition not use expr, here will initialize the field of the condition when initializing.
// newdict -> Dict{ Name: "non_existing" } if not found
newdict, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(&testdata.Dict{
        Name: "myname",
    }).
    Attrs(&testdata.Dict{Remark: "remark11"}).
    FirstOrInit()
_ = newdict
// NOTE: Condition not use expr, here will initialize the field of the condition when initializing.
// if not found
// newdict -> Dict{ Name: "myname", Remark: "remark11" }
//
// if found, `Attrs`, `AttrsExpr` are ignored
// newdict -> Dict{ Id: 1, Name: "myname", Remark: "remark" }

// `Assign`, `AssignExpr`
// with expr
newdict, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(refDict.Name.Eq("myname")).
    AssignExpr(refDict.Remark.Value("remark11")).
    FirstOrInit()
_ = newdict
// NOTE: Where condition use expr, here will not initialize the field of the condition when initializing.
//  if not found
// newdict -> Dict{ Remark: "remark11" }
//
//  if not found
// newdict -> Dict{ Name: "non_existing" }
newdict, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(&testdata.Dict{
        Name: "myname",
    }).
    Assign(&testdata.Dict{Remark: "remark11"}).
    FirstOrInit()
_ = newdict
// NOTE: condition not use expr, here will initialize the field of the condition when initializing.
// if not found
// newdict -> Dict{ Name: "myname", Remark: "remark11" }
//
// if found, `Assign`, `AssignExpr` are set on the struct
// newdict -> Dict{ Id: 1, Name: "myname", Remark: "remark11" }
```

## FirstOrCreate

`FirstOrCreate` is used to fetch the first record that matches given conditions or create a new one if no matching record is found. This method is effective with both struct and map conditions. The RowsAffected property is useful to determine the number of records created or updated.

- `Attrs`, `AttrsExpr` can be used to specify additional attributes for the new record if it is not found. These attributes are used for creation but not in the initial search query.
- `Assign`, `AssignExpr` method sets attributes on the record regardless of whether it is found or not, and these attributes are saved back to the database.

***NOTE!!!***: if with expr condition will not initialize the field when creating, so we should use
`Attrs`, `AttrsExpr`, `Assign`, `AssignExpr` attributes to indicate these fields.

```go
refDict := testdata.Ref_Dict()
// NOTE!!!: if with expr condition will not initialize the field when creating, so we should use
// `Attrs`, `AttrsExpr`, `Assign`, `AssignExpr` attributes to indicate these fields.

// `Attrs`, `AttrsExpr`
// with expr
newdict, _ := rapier.NewExecutor[testdata.Dict](db).
    Where(refDict.Name.Eq("myname")).
    AttrsExpr(refDict.Remark.Value("remark11")).
    FirstOrCreate()
_ = newdict
// NOTE: Condition use expr. here will not initialize the field of the condition when creating.
// if not found. initialize with additional attributes
// SELECT * FROM `dict` WHERE `dict`.`name` = "myname" ORDER BY `dict`.`id` LIMIT 1;
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("","",false,"remark11","2024-03-08 02:20:10.853","2024-03-08 02:20:10.853");
// newdict -> Dict{ Id: 11, Name: "", Remark: "remark11" } if not found
//
// if found, `Attrs`, `AttrsExpr` are ignored.
// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark" }

// with original gorm api
newdict, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(&testdata.Dict{
        Name: "myname",
    }).
    Attrs(&testdata.Dict{Remark: "remark11"}).
    FirstOrCreate()
_ = newdict
// NOTE: Condition not use expr, here will initialize the field of the condition when creating.
// if not found, initialize with given conditions and additional attributes
// SELECT * FROM `dict` WHERE `dict`.`name` = "myname" ORDER BY `dict`.`id` LIMIT 1;
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("","myname",false,"remark11","2024-03-08 02:20:10.853","2024-03-08 02:20:10.853");
// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark11" }
//
// if found, `Attrs`, `AttrsExpr` are ignored
// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark" }

// `Assign`, `AssignExpr`
// with expr
newdict, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(refDict.Name.Eq("myname")).
    AssignExpr(refDict.Remark.Value("remark11")).
    FirstOrCreate()
_ = newdict
// NOTE: Where condition use expr, here will not initialize the field of the condition when creating.
// whether it is found or not, and `Assign`, `AssignExpr` attributes are saved back to the database.
// if no found
// SELECT * FROM `dict` WHERE `dict`.`name` = "myname" ORDER BY `dict`.`id` LIMIT 1
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("","",false,"remark11","2024-03-08 02:26:12.619","2024-03-08 02:26:12.619");
// newdict -> Dict{ Id: 11, Name: "", Remark: "remark11", ... }
//
// if found
// SELECT * FROM `dict` WHERE `dict`.`name` = "myname" ORDER BY `dict`.`id` LIMIT 1
// UPDATE `dict` SET `remark` = "remark11" WHERE id = "11"
// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark11", ... }

newdict, _ = rapier.NewExecutor[testdata.Dict](db).
    Where(&testdata.Dict{
        Name: "myname",
    }).
    Assign(&testdata.Dict{Remark: "remark11"}).
    FirstOrCreate()
_ = newdict
// NOTE: condition not use expr, here will initialize the field of the condition when creating.
// whether it is found or not, and `Assign`, `AssignExpr` attributes are saved back to the database.
// if no found
// SELECT * FROM `dict` WHERE `dict`.`name` = "myname" ORDER BY `dict`.`id` LIMIT 1
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("","myname",false,"remark11","2024-03-08 02:26:12.619","2024-03-08 02:26:12.619");
// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark11", ... }
//
// if found
// SELECT * FROM `dict` WHERE `dict`.`name` = "myname" ORDER BY `dict`.`id` LIMIT 1
// UPDATE `dict` SET `remark` = "remark11" WHERE id = "11"
// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark11", ... }
```

## Pluck

The `Pluck`, `PluckExpr` method is used to query a single column from the database and scan the result into a slice. This method is ideal for when you need to retrieve specific fields from a model.

If you need to query more than one column, you can use Select with `Scan` or `Find` instead.

```go
var ids []int64

refDict := testdata.Ref_Dict()
// with expr api
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprString(refDict.Name)
// SELECT `name` FROM `dict`
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprBool(refDict.IsPin)
// SELECT `is_pin` FROM `dict`
_ = rapier.NewExecutor[testdata.Dict](db).Pluck("id", &ids)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt8(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt16(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt32(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt64(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint8(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint16(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint32(refDict.Id)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint64(refDict.Id)
// SELECT `id` FROM `dict`

// with original gorm api
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckString("name")
// SELECT `name` FROM `dict`
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckBool("is_pin")
// SELECT `is_pin` FROM `dict`
_ = rapier.NewExecutor[testdata.Dict](db).Pluck("id", &ids)
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt8("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt16("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt32("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt64("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint8("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint16("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint32("id")
_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint64("id")
// SELECT `id` FROM `dict`
```

## Count

The `Count` method is used to retrieve the number of records that match a given query. It’s a useful feature for understanding the size of a dataset, particularly in scenarios involving conditional queries or data analysis.

```go
total, err := rapier.NewExecutor[testdata.Dict](db).Count()
_ = err
_ = total
// SELECT count(*) FROM `dict`
```

## Exist

The `Exist` method is used to check whether the exist record that match a given query.

```go
refDict := testdata.Ref_Dict()
b, err := rapier.NewExecutor[testdata.Dict](db).Where(refDict.Id.Eq(100)).Exist()
_ = err
_ = b
// SELECT 1 FROM `dict` WHERE `dict`.`id` = 100 LIMIT 1
```

## Function

### Case When

```go
NewCaseWhen().
WhenThen(NewField("", "id1").Gt(100), NewField("", "value1")).
WhenThen(NewField("", "id2").Gt(200), NewField("", "value2")).
Else(NewField("", "result")).
Build()
// (CASE WHEN `id1` > ? THEN `value1` WHEN `id2` > ? THEN `value2` ELSE `result` END)
```

### Concat

```go
ConcatCol(NewString("", "id"), NewString("", "new_id"), NewString("", "new_id2"))
// CONCAT(`id`,`new_id`,`new_id2`)

ConcatWsCol(NewRaw(`'-'`), NewString("", "id"), NewString("", "new_id"), NewString("", "new_id2"))
// CONCAT_WS('-',`id`,`new_id`,`new_id2`)
```

### IF

```go
IF(NewField("", "id1").Gt(100), NewRaw("t"), NewField("", "f").Sub(1))
// IF(`id1` > ?,t,`f`-?)
```
