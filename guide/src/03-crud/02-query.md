# Query

- [Query](#query)
  - [Retrieving a single object](#retrieving-a-single-object)
  - [Retrieving a single field](#retrieving-a-single-field)
  - [Retrieving multiple objects](#retrieving-multiple-objects)
  - [Condition](#condition)
  - [Selecting Specific Fields](#selecting-specific-fields)
  - [Order](#order)
  - [Limit \& Offset](#limit--offset)
  - [Group By \& Having](#group-by--having)
  - [Distinct](#distinct)
  - [Joins](#joins)
  - [Scan](#scan)

## Retrieving a single object

```go
var record2 testdata.Dict

// Get the first record ordered by primary key
record1, err := rapier.NewExecutor[testdata.Dict](db).FirstOne()
_ = err     // return error
_ = record1 // return record
// Get one record, no specified order
record1, err = rapier.NewExecutor[testdata.Dict](db).TakeOne()
_ = err     // return error
_ = record1 // return record
// Get one record, no specified order
record1, err = rapier.NewExecutor[testdata.Dict](db).LastOne()
_ = err     // return error
_ = record1 // return record

// Get the first record ordered by primary key  with original gorm api
err = rapier.NewExecutor[testdata.Dict](db).First(&record2)
_ = err     // return error
_ = record2 // return record
// Get one record, no specified order with original gorm api
err = rapier.NewExecutor[testdata.Dict](db).Take(&record2)
_ = err     // return error
_ = record2 // return record
// Get one record, no specified order with original gorm api
err = rapier.NewExecutor[testdata.Dict](db).Last(&record2)
_ = err     // return error
_ = record2 // return record
```

## Retrieving a single field

the api like `FirstXXX` or `TakeXXX`, return follow type: `bool`,`string`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`

```go
refDict := testdata.Ref_Dict()
// Get the first record ordered returned single field.
_ = err // return error
_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstString()
// SELECT `dict`.`key` FROM `dict` ORDER BY `dict`.`id` LIMIT 1

// Get one record, no specified order returned single field.
_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeString()
// SELECT `dict`.`key` FROM `dict` LIMIT 1
```

## Retrieving multiple objects

```go
// Get the multiple record.
records1, err := rapier.NewExecutor[testdata.Dict](db).
    FindAll()
_ = err      // return error
_ = records1 // return records
// SELECT * FROM `dict`

var records2 []*testdata.Dict
// Get the multiple record.
err = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(rapier.All).
    Find(&records2)
_ = err      // return error
_ = records1 // return records
// SELECT * FROM `dict`
```

## Condition

In addition to [gorm Conditions](https://gorm.io/docs/query.html#Conditions) usages, there are usable usage related to `rapier`

```go
refDict := testdata.Ref_Dict()

// =
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` = "key1" LIMIT 1
// <>
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Neq("key1")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` <> "key1" LIMIT 1
// IN
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.In("key1", "key2")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` IN ("key1","key2") LIMIT 1
// NOT IN
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.NotIn("key1", "key2")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` NOT IN ("key1","key2") LIMIT 1
// Fuzzy LIKE
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.FuzzyLike("key1")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` LIKE "%key1%" LIMIT 1
// Left LIKE
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.LeftLike("key1")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` LIKE "key1%" LIMIT 1
// LIKE
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Like("%key1%")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` LIKE "%key1%" LIMIT 1
// NOT LIKE
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.NotLike("%key1%")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` NOT LIKE "%key1%" LIMIT 1
// AND
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1"), refDict.IsPin.Eq(true)).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` = "key1" AND `dict`.`is_pin` = true LIMIT 1
// >
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Gt(time.Now())).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`created_at` > "2024-03-07 06:20:47.057" LIMIT 1
// >=
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Gte(time.Now())).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`created_at` >= "2024-03-07 06:20:47.057" LIMIT 1
// <
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Lt(time.Now())).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`created_at` < "2024-03-07 06:20:47.057" LIMIT 1
// <=
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Lte(time.Now())).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`created_at` <= "2024-03-07 06:20:47.057" LIMIT 1
// BETWEEN
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Between(time.Now().Add(time.Hour), time.Now())).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`created_at` BETWEEN "2024-03-07 07:20:47.057" AND "2024-03-07 06:20:47.057" LIMIT 1
// NOT BETWEEN
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.NotBetween(time.Now().Add(time.Hour), time.Now())).TakeOne()
// SELECT * FROM `dict` WHERE NOT (`dict`.`created_at` BETWEEN "2024-03-07 07:20:47.057" AND "2024-03-07 06:20:47.057") LIMIT 1
 
// not condition
_, _ = rapier.NewExecutor[testdata.Dict](db).Not(refDict.Key.Eq("key1")).TakeOne()
// SELECT * FROM `dict` WHERE NOT `dict`.`key` = "key1" LIMIT 1
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Not(refDict.Key.Eq("key1"))).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` <> "key1" LIMIT 1
_, _ = rapier.NewExecutor[testdata.Dict](db).Not(refDict.Key.In("key1", "key2")).TakeOne()
// SELECT * FROM `dict` WHERE NOT `dict`.`key` IN ("key1","key2") LIMIT 1
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Not(refDict.Key.In("key1", "key2"))).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` NOT IN ("key1","key2") LIMIT 1

// Or condition
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1")).Or(refDict.Key.Eq("key2")).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` = "key1" OR `dict`.`key` = "key2" LIMIT 1
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Or(refDict.Key.Eq("key1"), refDict.Key.Eq("key2"))).TakeOne()
// SELECT * FROM `dict` WHERE (`dict`.`key` = "key1" OR `dict`.`key` = "key2") LIMIT 1
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1")).Or(refDict.Key.Eq("key2"), refDict.IsPin.Eq(true)).TakeOne()
// SELECT * FROM `dict` WHERE `dict`.`key` = "key1" OR (`dict`.`key` = "key2" AND `dict`.`is_pin` = true) LIMIT 1
_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Or(refDict.Key.Eq("key1"), rapier.And(refDict.Key.Eq("key2"), refDict.IsPin.Eq(true)))).TakeOne()
// SELECT * FROM `dict` WHERE (`dict`.`key` = "key1" OR (`dict`.`key` = "key2" AND `dict`.`is_pin` = true))
```

## Selecting Specific Fields

`Select`, `SelectExpr` allows you to specify the fields that you want to retrieve from database.

```go
var records []*struct {
    Key   string
    IsPin bool
}
refDict := testdata.Ref_Dict()

// with expr
_ = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key, refDict.IsPin).Find(&records)
// SELECT `dict`.`key`,`dict`.`is_pin` FROM `dict`
_ = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key.Trim("1").As(refDict.Key.ColumnName()), refDict.IsPin).Find(&records)
// SELECT TRIM(BOTH "1" FROM `dict`.`key`) AS `key`,`dict`.`is_pin` FROM `dict`

// with original gorm api
_ = rapier.NewExecutor[testdata.Dict](db).Select("key", "is_pin").Find(&records)
// SELECT `key`,`is_pin` FROM `dict`
```

## Order

Specify order when retrieving records from the database

```go
refDict := testdata.Ref_Dict()

// with expr
_, _ = rapier.NewExecutor[testdata.Dict](db).OrderExpr(refDict.Key.Desc(), refDict.Name).FindAll()
// SELECT * FROM `dict` ORDER BY `dict`.`key` DESC,`dict`.`name`
_, _ = rapier.NewExecutor[testdata.Dict](db).OrderExpr(refDict.Key.Desc()).OrderExpr(refDict.Name).FindAll()
// SELECT * FROM `dict` ORDER BY `dict`.`key` DESC,`dict`.`name`

// with original gorm api
_, _ = rapier.NewExecutor[testdata.Dict](db).Order("`key` DESC,name").FindAll()
// SELECT * FROM `dict` ORDER BY `key` DESC,name
_, _ = rapier.NewExecutor[testdata.Dict](db).Order("`key` DESC").Order("name").FindAll()
// SELECT * FROM `dict` ORDER BY `key` DESC,name
```

## Limit & Offset

`Pagination`:

- `page`: page index
- `perPage`: per page size (default size is 50, default max size is 500)
- `maxPerPages`: override default max size.

`Limit`: specify the max number of records to retrieve.  
`Offset`: specify the number of records to skip before starting to return the records

```go
// with Pagination
_, _ = rapier.NewExecutor[testdata.Dict](db).Pagination(3, 5).FindAll()
// SELECT * FROM `dict` LIMIT 5 OFFSET 10

// with original gorm api
_, _ = rapier.NewExecutor[testdata.Dict](db).Limit(3).FindAll()
// SELECT * FROM `dict` LIMIT 3
_, _ = rapier.NewExecutor[testdata.Dict](db).Offset(3).FindAll()
// SELECT * FROM `dict` OFFSET 3
_, _ = rapier.NewExecutor[testdata.Dict](db).Limit(10).Offset(5).FindAll()
// SELECT * FROM `dict` LIMIT 10 OFFSET 5
```

## Group By & Having

```go
var result struct {
    Name  string
    Total int
}

refDict := testdata.Ref_Dict()
// with expr
_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Name,
        rapier.Star.Count().As("total"),
    ).
    Where(refDict.Name.LeftLike("group")).
    GroupExpr(refDict.Name).
    Take(&result)
// SELECT `dict`.`name`,COUNT(*) AS `total` FROM `dict` WHERE `dict`.`name` LIKE "group%" GROUP BY `dict`.`name`

_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Name,
        rapier.Star.Count().As("total"),
    ).
    GroupExpr(refDict.Name).
    Having(refDict.Name.Eq("group")).
    Take(&result)
// SELECT `dict`.`name`,COUNT(*) AS `total` FROM `dict` GROUP BY `dict`.`name` HAVING `dict`.`name` = "group"

// with original gorm api
_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Name,
        rapier.Star.Count().As("total"),
    ).
    Where(refDict.Name.LeftLike("group")).
    Group("name").
    Take(&result)
// SELECT `dict`.`name`,COUNT(*) AS `total` FROM `dict` WHERE `dict`.`name` LIKE "group%" GROUP BY `name` LIMIT 1

_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Name,
        rapier.Star.Count().As("total"),
    ).
    Group("name").
    Having("name = ?", "group").
    Take(&result)
// SELECT `dict`.`name`,COUNT(*) AS `total` FROM `dict` GROUP BY `name` HAVING name = "group" LIMIT 1
```

## Distinct

```go
refDict := testdata.Ref_Dict()
// with expr
_, _ = rapier.NewExecutor[testdata.Dict](db).
    DistinctExpr(
        refDict.Name,
        refDict.IsPin,
    ).
    FindAll()
// SELECT DISTINCT `dict`.`name`,`dict`.`is_pin` FROM `dict`

// with original gorm api
_, _ = rapier.NewExecutor[testdata.Dict](db).
    Distinct("name", "is_pin").
    FindAll()
// SELECT DISTINCT `name`,`is_pin` FROM `dict`
```

## Joins

```go
refDict := testdata.Ref_Dict()
refDictItem := testdata.Ref_DictItem()
d := refDict.As("d")
di := refDictItem.As("di")

// join
_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Id.As(refDict.Id.FieldName(refDict.TableName())),
        refDict.Key.As(refDict.Key.FieldName(refDict.TableName())),
        refDictItem.Name.As(refDictItem.Name.FieldName(refDictItem.TableName())),
    ).
    InnerJoinsExpr(refDictItem, refDictItem.DictId.EqCol(refDict.Id), refDictItem.IsEnabled.Eq(true)).
    Take(&struct{}{})
// SELECT `dict`.`id` AS `dict_id`,`dict`.`key` AS `dict_key`,`dict_item`.`name` AS `dict_item_name` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = true LIMIT 1

// join with alias
_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Id.As(refDict.Id.FieldName(refDict.TableName())),
        refDict.Key.As(refDict.Key.FieldName(refDict.TableName())),
        d.Id.As(d.Id.FieldName(d.Alias())),
    ).
    InnerJoinsExpr(rapier.NewJoinTable(d, d.Alias()), d.Name.EqCol(refDict.Name), d.IsPin.Eq(true)).
    Take(&struct{}{})
// SELECT `dict`.`id` AS `dict_id`,`dict`.`key` AS `dict_key`,`d`.`id` AS `d_id` FROM `dict` INNER JOIN `dict` `d` ON `d`.`name` = `dict`.`name` AND `d`.`is_pin` = true LIMIT 1

// join with alias which table implements Alias interface.
// we can directly use it. no need `NewJoinTable`.
_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Id.As(refDict.Id.FieldName(refDict.TableName())),
        refDict.Key.As(refDict.Key.FieldName(refDict.TableName())),
        d.Id.As(d.Id.FieldName(d.Alias())),
    ).
    InnerJoinsExpr(d, d.Name.EqCol(refDict.Name), d.IsPin.Eq(true)).
    Take(&struct{}{})
// SELECT `dict`.`id` AS `dict_id`,`dict`.`key` AS `dict_key`,`d`.`id` AS `d_id` FROM `dict` INNER JOIN `dict` `d` ON `d`.`name` = `dict`.`name` AND `d`.`is_pin` = true LIMIT 1

// join with SubQuery
_ = rapier.NewExecutor[testdata.Dict](db).
    SelectExpr(
        refDict.Id.As(refDict.Id.FieldName(refDict.TableName())),
        refDict.Key.As(refDict.Key.FieldName(refDict.TableName())),
        di.Sort.As(di.Sort.FieldName(di.Alias())),
    ).
    InnerJoinsExpr(
        rapier.NewJoinTableSubQuery(
            rapier.NewExecutor[testdata.DictItem](db).
                Where(refDictItem.IsEnabled.Eq(true)).
                IntoDB(),
            "di",
        ),
        di.Name.EqCol(refDict.Name),
        di.Sort.Gt(10),
    ).
    Take(&struct{}{})
// SELECT `dict`.`id` AS `dict_id`,`dict`.`key` AS `dict_key`,`di`.`sort` AS `di_sort` FROM `dict` INNER JOIN (SELECT * FROM `dict_item` WHERE `dict_item`.`is_enabled` = true) AS `di` ON `di`.`name` = `dict`.`name` AND `di`.`sort` > 10 LIMIT 1
```

## Scan

Retrieving a single field, the api like `ScanXXX` return follow type: `bool`,`string`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`.

```go
var record2 testdata.Dict

// Get one record, no specified order
record1, err := rapier.NewExecutor[testdata.Dict](db).ScanOne()
_ = err     // return error
_ = record1 // return record
// SELECT * FROM `dict`
// Get one record, no specified order with original gorm api
err = rapier.NewExecutor[testdata.Dict](db).Scan(&record2)
_ = err     // return error
_ = record2 // return record
// SELECT * FROM `dict`
```

more information [gorm Query](https://gorm.io/docs/query.html)
