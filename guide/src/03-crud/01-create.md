# Create

- [Create](#create)
  - [Empty record](#empty-record)
  - [Single record](#single-record)
  - [Multiple record](#multiple-record)
  - [Batch insert multiple record](#batch-insert-multiple-record)

## Empty record

```go
// empty record
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).Create()
_ = err          // return error
_ = rowsAffected // return row affected
```

## Single record

```go
// single record
newDict := testdata.Dict{
    Key:    "key1",
    Name:   "name1",
    IsPin:  true,
    Remark: "remark1",
}
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).Create(&newDict)
_ = err          // return error
_ = rowsAffected // return row affected
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.135","2024-02-20 07:18:42.135")
```

## Multiple record

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
}
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).Create(newDicts...)
_ = err          // return error
_ = rowsAffected // return row affected
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136"),("key2","name2",true,"remark2","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136")
```

## Batch insert multiple record

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
rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).CreateInBatches(newDicts, 2)
_ = err          // return error
_ = rowsAffected // return row affected
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key1","name1",true,"remark1","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136"),("key2","name2",true,"remark2","2024-02-20 07:18:42.136","2024-02-20 07:18:42.136")
// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("key3","name3",true,"remark3","2024-02-20 07:18:42.135","2024-02-20 07:18:42.135")
```

more information [gorm Create](https://gorm.io/docs/create.html)
