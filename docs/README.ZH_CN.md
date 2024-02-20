# rapier

## 1 Getting Started

[gorm-rapier](https://github.com/thinkgos/gorm-rapier) is an assist rapier for gorm.

- [GORM Guides](http://gorm.io/docs)
- [Rapier Guide](https://github.com/thinkgos/gorm-rapier/tree/main/docs/README.ZH_CN.md)

### 1.1 Overview

- Idiomatic and Reusable API from Dynamic Raw SQL
- 100% Type-safe API without interface{}
- Almost supports all features, plugins, DBMS that GORM supports
- Almost same behavior as gorm you used.

### 1.2 Declaring gorm and Rapier Model

#### 1.2.1 Declaring gorm model

see [Declaring Models](https://gorm.io/docs/models.html)

#### 1.2.2 Declaring Rapier model

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

#### 1.2.3 How to define model

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
var ref_Dict_Model = new_Dict("dict")

type Dict_Native struct {
    refAlias  string
    ALL       rapier.Asterisk
    Id        rapier.Int64
    Key       rapier.String
    Name      rapier.String
    IsPin     rapier.Bool
    Remark    rapier.String
    CreatedAt rapier.Time
    UpdatedAt rapier.Time
}

// Ref_Dict model with TableName `dict`.
func Ref_Dict() Dict_Native { return ref_Dict_Model }

func new_Dict(alias string) Dict_Native {
    return Dict_Native{
        refAlias:  alias,
        ALL:       rapier.NewAsterisk(alias),
        Id:        rapier.NewInt64(alias, "id"),
        Key:       rapier.NewString(alias, "key"),
        Name:      rapier.NewString(alias, "name"),
        IsPin:     rapier.NewBool(alias, "is_pin"),
        Remark:    rapier.NewString(alias, "remark"),
        CreatedAt: rapier.NewTime(alias, "created_at"),
        UpdatedAt: rapier.NewTime(alias, "updated_at"),
    }
}

// New_Dict new instance.
func New_Dict(xAlias string) Dict_Native {
    if xAlias == "dict" {
        return ref_Dict_Model
    } else {
        return new_Dict(xAlias)
    }
}

// As alias
func (*Dict_Native) As(alias string) Dict_Native { return New_Dict(alias) }

// Ref_Alias hold table name when call New_Dict or Dict_Active.As that you defined.
func (x *Dict_Native) Ref_Alias() string { return x.refAlias }

// New_Executor new entity executor which suggest use only once.
func (*Dict_Native) New_Executor(db *gorm.DB) *rapier.Executor[Dict] {
    return rapier.NewExecutor[Dict](db)
}

// TableName hold model `Dict` table name returns `dict`.
func (*Dict_Native) TableName() string { return "dict" }

// Select_Expr select model fields
func (x *Dict_Native) Select_Expr() []rapier.Expr {
    return []rapier.Expr{
        x.Id,
        x.Key,
        x.Name,
        x.IsPin,
        x.Remark,
        x.CreatedAt,
        x.UpdatedAt,
    }
}

// Select_VariantExpr select model fields, but time.Time field convert to timestamp(int64).
func (x *Dict_Native) Select_VariantExpr(prefixes ...string) []rapier.Expr {
    if len(prefixes) > 0 && prefixes[0] != "" {
        return []rapier.Expr{
            x.Id.As(x.Id.FieldName(prefixes...)),
            x.Key.As(x.Key.FieldName(prefixes...)),
            x.Name.As(x.Name.FieldName(prefixes...)),
            x.IsPin.As(x.IsPin.FieldName(prefixes...)),
            x.Remark.As(x.Remark.FieldName(prefixes...)),
            x.CreatedAt.UnixTimestamp().As(x.CreatedAt.FieldName(prefixes...)),
            x.UpdatedAt.UnixTimestamp().As(x.UpdatedAt.FieldName(prefixes...)),
        }
    } else {
        return []rapier.Expr{
            x.Id,
            x.Key,
            x.Name,
            x.IsPin,
            x.Remark,
            x.CreatedAt.UnixTimestamp().As(x.CreatedAt.FieldName()),
            x.UpdatedAt.UnixTimestamp().As(x.UpdatedAt.FieldName()),
        }
    }
}
```

### 1.3 Connecting to a Database

see [Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)

## 2 CRUD interface

### 2.1 Create

```go
newDict := model.Dict{
    Key:    "key1",
    Name:   "name1",
    IsPin:  true,
    Remark: "remark1",
}
err := rapier.NewExecutor[model.Dict](db).Create(newDict)
_ = err // return error
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.135","2024-02-20 07:18:42.135")
```

```go
newDicts := []*model.Dict{
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
err = rapier.NewExecutor[model.Dict](db).Create(newDicts...)
_ = err // return error
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136"),("key2","name2",true,"remark2","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136")
```

### 2.2 Query

### 2.3 Advanced Query

### 2.4 Update

```go
refDict := model.Ref_Dict()
rowsAffected, err := rapier.NewExecutor[model.Dict](db).
    Where(refDict.Id.Eq(100)).
    UpdatesExpr(
        refDict.Key.Value("k1"),
    )
_ = err          // return error
_ = rowsAffected // return row affected
 // UPDATE `dict` SET `key`="k1",`updated_at`="2024-02-20 07:22:59.171" WHERE `dict`.`id` = 100
```

### 2.5 Delete

### Original gorm db

## Transaction

## Associations

not supported yet
