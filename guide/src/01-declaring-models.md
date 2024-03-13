# Declaring Models

- [Declaring Models](#declaring-models)
  - [Declaring gorm model](#declaring-gorm-model)
  - [Declaring rapier model](#declaring-rapier-model)
  - [How to define model](#how-to-define-model)

## Declaring gorm model

[Declaring gorm Models](https://gorm.io/docs/models.html)

## Declaring rapier model

Supported field:

- bool: `Bool`
- []byte: `Bytes`
- float: `Float32`, `Float64`, `Decimal`
- integer: `Int`, `Int8`, `Int16`, `Int32`, `Int64`
- unsigned integer: `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
- string: `String`
- time.Time: `Time`
- any: `Field`
- raw filed: `Raw`

## How to define model

model defined for test [testdata](https://github.com/thinkgos/gorm-rapier/tree/main/testdata).

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
var ref_Dict_Native = New_Dict("dict")

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
