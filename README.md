# gorm-rapier

[gorm-rapier](https://github.com/thinkgos/gorm-rapier) is an assist rapier for gorm.

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/gorm-rapier?tab=doc)
[![codecov](https://codecov.io/gh/thinkgos/gorm-rapier/graph/badge.svg?token=aHu5wq1m6i)](https://codecov.io/gh/thinkgos/gorm-rapier)
[![Tests](https://github.com/thinkgos/gorm-rapier/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/thinkgos/gorm-rapier/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/gorm-rapier)](https://goreportcard.com/report/github.com/thinkgos/gorm-rapier)
[![Licence](https://img.shields.io/github/license/thinkgos/gorm-rapier)](https://raw.githubusercontent.com/thinkgos/gorm-rapier/main/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/gorm-rapier)](https://github.com/thinkgos/gorm-rapier/tags)

- [gorm-rapier](#gorm-rapier)
  - [Overview](#overview)
  - [Usage](#usage)
    - [1. Installation](#1-installation)
    - [2. Getting Started](#2-getting-started)
    - [2.1 Declaring gorm and rapier Model](#21-declaring-gorm-and-rapier-model)
      - [2.1.1 Declaring gorm model](#211-declaring-gorm-model)
      - [2.1.2 Declaring rapier model](#212-declaring-rapier-model)
      - [2.1.3 How to define model](#213-how-to-define-model)
    - [2.2 Connecting to a Database](#22-connecting-to-a-database)
    - [2.3 CRUD interface](#23-crud-interface)
      - [2.3.1 Create](#231-create)
        - [Empty record](#empty-record)
        - [Single record](#single-record)
        - [Multiple record](#multiple-record)
        - [Batch insert multiple record](#batch-insert-multiple-record)
      - [2.3.2 Query](#232-query)
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
      - [2.3.3 Advanced Query](#233-advanced-query)
        - [Locking](#locking)
        - [SubQuery](#subquery)
        - [From SubQuery](#from-subquery)
        - [FirstOrInit](#firstorinit)
        - [FirstOrCreate](#firstorcreate)
        - [Pluck](#pluck)
        - [Count](#count)
        - [Exist](#exist)
      - [2.3.4 Update](#234-update)
        - [`Save` will save all fields](#save-will-save-all-fields)
        - [Update single column](#update-single-column)
        - [Updates multiple columns](#updates-multiple-columns)
        - [Update from SubQuery](#update-from-subquery)
        - [Without Hooks/Time Tracking](#without-hookstime-tracking)
      - [2.3.5 Delete](#235-delete)
    - [2.4 Original gorm db](#24-original-gorm-db)
    - [2.5 Transaction](#25-transaction)
    - [2.6 Associations](#26-associations)
    - [2.7 Example](#27-example)
  - [Reference](#reference)
  - [License](#license)

## Overview

- Idiomatic and Reusable API from Dynamic Raw SQL
- 100% Type-safe API without interface{}
- Almost supports all features, plugins, DBMS that GORM supports
- Almost same behavior as gorm you used.

## Usage

[返回顶部](#gorm-rapier)

### 1. Installation

Use go get.

```bash
go get github.com/thinkgos/gorm-rapier
```

Then import the package into your own code.

```go
import "github.com/thinkgos/gorm-rapier"
```

### 2. Getting Started

[GORM Guides](http://gorm.io/docs)

### 2.1 Declaring gorm and rapier Model

[返回顶部](#gorm-rapier)

model defined detail see [testdata](./testdata).

#### 2.1.1 Declaring gorm model

see [Declaring Models](https://gorm.io/docs/models.html)

#### 2.1.2 Declaring rapier model

Supported field:

- bool: `Bool`
- []byte: `Bytes`
- float: `Float32`, `float64`, `Decimal`
- integer: `Int`, `Int8`, `Int16`, `Int32`, `Int64`
- unsigned integer: `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
- string: `String`
- time.Time: `Time`
- any: `Field`
- raw filed: `Raw`

#### 2.1.3 How to define model

[返回顶部](#gorm-rapier)

if we have a gorm model follow:

```go
// Dict 字典
type Dict struct {
    Id        int64     `gorm:"column:id;autoIncrement:true;not null;primaryKey" json:"id,omitempty"`
    Key       string    `gorm:"column:key;type:varchar(64);not null;default:'';uniqueIndex:uk_key" json:"key,omitempty"`
    Name      string    `gorm:"column:name;type:varchar(64);not null;default:''" json:"name,omitempty"`
    IsPin     bool      `gorm:"column:is_pin;type:tinyint(1);not null;default:0" json:"is_pin,omitempty"`
    Remark    string    `gorm:"column:remark;type:varchar(128);not null;default:''" json:"remark,omitempty"`
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at,omitempty"`
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at,omitempty"`
}

// TableName implement schema.Tabler interface
func (*Dict) TableName() string {
    return "dict"
}
```

then we can define rapier model:

```go
var ref_Dict_Native = new_Dict("dict", "dict")

type Dict_Native struct {
    refAlias     string
    refTableName string
    ALL          rapier.Asterisk
    Id           rapier.Int64
    Key          rapier.String
    Name         rapier.String
    IsPin        rapier.Bool
    Remark       rapier.String
    CreatedAt    rapier.Time
    UpdatedAt    rapier.Time
}

func new_Dict(tableName, alias string) *Dict_Native {
    return &Dict_Native{
        refAlias:     alias,
        refTableName: tableName,
        ALL:          rapier.NewAsterisk(alias),
        Id:           rapier.NewInt64(alias, "id"),
        Key:          rapier.NewString(alias, "key"),
        Name:         rapier.NewString(alias, "name"),
        IsPin:        rapier.NewBool(alias, "is_pin"),
        Remark:       rapier.NewString(alias, "remark"),
        CreatedAt:    rapier.NewTime(alias, "created_at"),
        UpdatedAt:    rapier.NewTime(alias, "updated_at"),
    }
}

// Ref_Dict model with TableName `dict`.
func Ref_Dict() *Dict_Native { return ref_Dict_Native }

// New_Dict new instance.
func New_Dict(tableName string) *Dict_Native {
    return new_Dict(tableName, tableName)
}

// As alias
func (x *Dict_Native) As(alias string) *Dict_Native {
    return new_Dict(x.refTableName, alias)
}

// Ref_Alias hold alias name when call Dict_Active.As that you defined.
func (x *Dict_Native) Ref_Alias() string { return x.refAlias }

// TableName hold table name when call New_Dict that you defined.
func (x *Dict_Native) TableName() string { return x.refTableName }
```

### 2.2 Connecting to a Database

[返回顶部](#gorm-rapier)

see [gorm Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)

### 2.3 CRUD interface

`Executor[T]`'s `Where`, `Or`, `Not`, `Having` and the suffix with `Expr` method can use field which implement `Expr` interface.

#### 2.3.1 Create

##### Empty record

```go
// empty record
err := rapier.NewExecutor[testdata.Dict](db).Create()
_ = err // return error
// do nothing
```

##### Single record

[返回顶部](#gorm-rapier)

```go
newDict := testdata.Dict{
    Key:    "key1",
    Name:   "name1",
    IsPin:  true,
    Remark: "remark1",
}
err := rapier.NewExecutor[testdata.Dict](db).Create(newDict)
_ = err // return error
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.135","2024-02-20 07:18:42.135")
```

##### Multiple record

[返回顶部](#gorm-rapier)

```go
newDicts := []*testdata.Dict{
    {
        Key:    "key1",
        Name:   "name1",
        IsPin:  true,
        Remark: "remark1",
    },
    {
        Key:    "key2",
        Name:   "name2",
        IsPin:  true,
        Remark: "remark2",
    },
}
err = rapier.NewExecutor[testdata.Dict](db).Create(newDicts...)
_ = err // return error
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136"),("key2","name2",true,"remark2","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136")
```

##### Batch insert multiple record

[返回顶部](#gorm-rapier)

```go
// batch insert multiple record
newDicts := []*testdata.Dict{
    {
        Key:    "key1",
        Name:   "name1",
        IsPin:  true,
        Remark: "remark1",
    },
    {
        Key:    "key2",
        Name:   "name2",
        IsPin:  true,
        Remark: "remark2",
    },
    {
        Key:    "key3",
        Name:   "name3",
        IsPin:  true,
        Remark: "remark3",
    },
}
err := rapier.NewExecutor[testdata.Dict](db).CreateInBatches(newDicts, 2)
_ = err // return error
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136"),("key2","name2",true,"remark2","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136")
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key3","name3",true,"remark3","2024-02-20 07:18:42.135","2024-02-20 07:18:42.135")
```

more information see [gorm Create](https://gorm.io/docs/create.html)

#### 2.3.2 Query

[返回顶部](#gorm-rapier)

##### Retrieving a single object

```go
var dummy testdata.Dict

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
err = rapier.NewExecutor[testdata.Dict](db).First(&dummy)
_ = err     // return error
_ = record1 // return record
// Get one record, no specified order with original gorm api
err = rapier.NewExecutor[testdata.Dict](db).Take(&dummy)
_ = err     // return error
_ = record1 // return record
// Get one record, no specified order with original gorm api
err = rapier.NewExecutor[testdata.Dict](db).Last(&dummy)
_ = err     // return error
_ = record1 // return record
```

##### Retrieving a single field

[返回顶部](#gorm-rapier)

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

##### Retrieving multiple objects

[返回顶部](#gorm-rapier)

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

##### Condition

[返回顶部](#gorm-rapier)

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

##### Selecting Specific Fields

[返回顶部](#gorm-rapier)

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

##### Order

[返回顶部](#gorm-rapier)

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

##### Limit & Offset

[返回顶部](#gorm-rapier)

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

##### Group By & Having

[返回顶部](#gorm-rapier)

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

##### Distinct

[返回顶部](#gorm-rapier)

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

##### Joins

[返回顶部](#gorm-rapier)

TODO...

##### Scan

[返回顶部](#gorm-rapier)

Retrieving a single field, the api like `ScanXXX` return follow type: `bool`,`string`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`.

```go
var dummy testdata.Dict

// Get one record, no specified order
record1, err := rapier.NewExecutor[testdata.Dict](db).ScanOne()
_ = err     // return error
_ = record1 // return record
// SELECT * FROM `dict`
// Get one record, no specified order with original gorm api
err = rapier.NewExecutor[testdata.Dict](db).Scan(&dummy)
_ = err     // return error
_ = record1 // return record
// SELECT * FROM `dict`
```

more information see [gorm Query](https://gorm.io/docs/query.html)

#### 2.3.3 Advanced Query

[返回顶部](#gorm-rapier)

##### Locking

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

##### SubQuery

[返回顶部](#gorm-rapier)

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

##### From SubQuery

[返回顶部](#gorm-rapier)

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

##### FirstOrInit

[返回顶部](#gorm-rapier)

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

##### FirstOrCreate

[返回顶部](#gorm-rapier)

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

##### Pluck

[返回顶部](#gorm-rapier)

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

##### Count

[返回顶部](#gorm-rapier)

The `Count` method is used to retrieve the number of records that match a given query. It’s a useful feature for understanding the size of a dataset, particularly in scenarios involving conditional queries or data analysis.

```go
total, err := rapier.NewExecutor[testdata.Dict](db).Count()
_ = err
_ = total
// SELECT count(*) FROM `dict`
```

##### Exist

[返回顶部](#gorm-rapier)

The `Exist` method is used to check whether the exist record that match a given query.

```go
refDict := testdata.Ref_Dict()
b, err := rapier.NewExecutor[testdata.Dict](db).Where(refDict.Id.Eq(100)).Exist()
_ = err
_ = b
// SELECT 1 FROM `dict` WHERE `dict`.`id` = 100 LIMIT 1
```

more information see [gorm Advanced Query](https://gorm.io/docs/advanced_query.html)

#### 2.3.4 Update

[返回顶部](#gorm-rapier)

##### `Save` will save all fields

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

##### Update single column

[返回顶部](#gorm-rapier)

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

##### Updates multiple columns

[返回顶部](#gorm-rapier)

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

##### Update from SubQuery

[返回顶部](#gorm-rapier)

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

##### Without Hooks/Time Tracking

[返回顶部](#gorm-rapier)

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

#### 2.3.5 Delete

[返回顶部](#gorm-rapier)

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

more information see [gorm Delete](https://gorm.io/docs/delete.html)

### 2.4 Original gorm db

[返回顶部](#gorm-rapier)

`IntoDB`, `IntoRawDB` will get original gorm db.

- `IntoDB`: with model or table
- `IntoRawDB`: without model or table

### 2.5 Transaction

[gorm transaction](https://gorm.io/docs/transactions.html)

### 2.6 Associations

not supported yet, you can use gorm original api.

### 2.7 Example

[返回顶部](#gorm-rapier)

- [create](./example_create_test.go): example create
- [query](./example_query_test.go): example query
- [advance query](./example_advance_query_test.go): example advance query
- [update](./example_update_test.go): example update
- [delete](./example_delete_test.go): example delete

## Reference

- [gorm](https://github.com:go-gorm/gorm)
- [sea-orm](https://github.com/SeaQL/sea-orm)
- [ent](https://github.com/ent/ent)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.