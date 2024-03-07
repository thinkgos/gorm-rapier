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

// New_Executor new entity executor which suggest use only once.
func (*Dict_Native) New_Executor(db *gorm.DB) *rapier.Executor[Dict] {
    return rapier.NewExecutor[Dict](db)
}
```

### 2.2 Connecting to a Database

see [gorm Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)

### 2.3 CRUD interface

`Executor[T]`'s `Where` and the suffix with `Expr` method can use field which implement `Expr` interface.

#### 2.3.1 Create

single record

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

multiple record

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

batch insert multiple record

```go
    // multiple record
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

#### 2.3.3 Advanced Query

#### 2.3.4 Update

```go
refDict := testdata.Ref_Dict()
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
    Model().
    Where(refDict.Id.Eq(100)).
    UpdatesExpr(
        refDict.Key.Value("k1"),
    )
_ = err          // return error
_ = rowsAffected // return row affected
 // UPDATE `dict` SET `key`="k1",`updated_at`="2024-02-20 07:22:59.171" WHERE `dict`.`id` = 100
```

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
