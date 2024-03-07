# gorm-rapier

gorm rapier is an assist rapier for gorm.

[![GoDoc](https://godoc.org/github.com/thinkgos/gorm-rapier?status.svg)](https://godoc.org/github.com/thinkgos/gorm-rapier)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/gorm-rapier?tab=doc)
[![codecov](https://codecov.io/gh/thinkgos/gorm-rapier/graph/badge.svg?token=aHu5wq1m6i)](https://codecov.io/gh/thinkgos/gorm-rapier)
[![Tests](https://github.com/thinkgos/gorm-rapier/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/thinkgos/gorm-rapier/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/gorm-rapier)](https://goreportcard.com/report/github.com/thinkgos/gorm-rapier)
[![Licence](https://img.shields.io/github/license/thinkgos/gorm-rapier)](https://raw.githubusercontent.com/thinkgos/gorm-rapier/main/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/gorm-rapier)](https://github.com/thinkgos/gorm-rapier/tags)

## Overview

- Idiomatic and Reusable API from Dynamic Raw SQL
- 100% Type-safe API without interface{}
- Almost supports all features, plugins, DBMS that GORM supports
- Almost same behavior as gorm you used.

## Usage

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

see [gorm Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)

### 2.3 CRUD interface

`Executor[T]`'s `Where`, `Or`, `Not`, `Having` and the suffix with `Expr` method can use field which implement `Expr` interface.

#### 2.3.1 Create

##### empty record

```go
    // empty record
    err := rapier.NewExecutor[testdata.Dict](db).Create()
    _ = err // return error
    // do nothing
```

##### single record

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

##### multiple record

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

##### batch insert multiple record

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

##### Retrieving a single field, the api like `FirstXXX` or `TakeXXX`, return follow type: `bool`,`string`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`.

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

Specify order when retrieving records from the database

```go
    refDict := testdata.Ref_Dict()

    // with expr
    _, _ = rapier.NewExecutor[testdata.Dict](db).OrderExpr(refDict.Key.Desc(), refDict.Name).FindAll()
    _, _ = rapier.NewExecutor[testdata.Dict](db).OrderExpr(refDict.Key.Desc()).OrderExpr(refDict.Name).FindAll()

    // with original gorm api
    _, _ = rapier.NewExecutor[testdata.Dict](db).Order("`key` DESC,name").FindAll()
    _, _ = rapier.NewExecutor[testdata.Dict](db).Order("`key` DESC").Order("name").FindAll()
```

##### Limit & Offset

more information see [gorm Update](https://gorm.io/docs/query.html)

#### 2.3.3 Advanced Query

#### 2.3.4 Update

##### `Save` will save all fields when performing the Updating SQL

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

    // TODO: update with exprs

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

### 2.5 Transaction

### 2.6 Associations

not supported yet

### 2.7 Example

## Reference

- [gorm](https://github.com:go-gorm/gorm)
- [sea-orm](https://github.com/SeaQL/sea-orm)
- [ent](https://github.com/ent/ent)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
