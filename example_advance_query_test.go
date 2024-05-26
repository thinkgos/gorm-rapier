package rapier_test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/testdata"
	"gorm.io/gorm/clause"
)

func Test_Example_AdvanceQuery_Locking(t *testing.T) {
	// Basic FOR UPDATE lock
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		LockingUpdate().
		TakeOne()
	// Basic FOR UPDATE lock with Clauses api
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		TakeOne()

	// Basic FOR SHARE lock
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		LockingShare().
		TakeOne()
	// Basic FOR SHARE lock with Clauses api
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Clauses(clause.Locking{Strength: "SHARE"}).
		TakeOne()

	// Basic FOR UPDATE NOWAIT lock with Clauses api
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).
		TakeOne()
}

func Test_Example_AdvanceQuery_SubQuery(t *testing.T) {
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

	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(refDict.Key.EqSubQuery(
			rapier.NewExecutor[testdata.Dict](db).
				SelectExpr(refDict.Key).
				Where(refDict.Id.Eq(1001)).
				IntoDB(),
		)).
		FindAll()
}

func Test_Example_AdvanceQuery_FromSubQuery(t *testing.T) {
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
}

func Test_Example_AdvanceQuery_IN_WithMultipleColumns(t *testing.T) {
	refDict := testdata.Ref_Dict()

	record1, _ := rapier.NewExecutor[testdata.Dict](db).
		Where(
			rapier.NewColumns(refDict.Name, refDict.IsPin).
				In([][]any{{"name1", true}, {"name2", false}}),
		).
		FindAll()
	_ = record1

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
}

func Test_Example_AdvanceQuery_FirstOrInit(t *testing.T) {
	refDict := testdata.Ref_Dict()
	// NOTE!!!: if with expr condition will not initialize the field when initializing, so we should use
	// `Attrs`, `AttrsExpr`, `Assign`, `AssignExpr` attributes to indicate these fields.

	// `Attrs`, `AttrsExpr`
	// with expr
	result, _ := rapier.NewExecutor[testdata.Dict](db).
		Where(refDict.Name.Eq("myname")).
		AttrsExpr(refDict.Remark.Value("remark11")).
		FirstOrInit()
	newdict := result.Data
	rowsAffected := result.RowsAffected
	_ = newdict
	_ = rowsAffected
	// NOTE: Condition use expr. here will not initialize the field of the condition when initializing.
	// if not found
	// newdict -> Dict{ Remark: "remark11" }
	//
	// if found, `Attrs`, `AttrsExpr` are ignored
	// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark" }

	// with original gorm api
	result, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(&testdata.Dict{
			Name: "non_existing",
		}).
		FirstOrInit()
	newdict = result.Data
	rowsAffected = result.RowsAffected
	_ = newdict
	_ = rowsAffected
	// NOTE: Condition not use expr, here will initialize the field of the condition when initializing.
	// newdict -> Dict{ Name: "non_existing" } if not found
	result, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(&testdata.Dict{
			Name: "myname",
		}).
		Attrs(&testdata.Dict{Remark: "remark11"}).
		FirstOrInit()
	newdict = result.Data
	rowsAffected = result.RowsAffected
	_ = newdict
	_ = rowsAffected
	// NOTE: Condition not use expr, here will initialize the field of the condition when initializing.
	// if not found
	// newdict -> Dict{ Name: "myname", Remark: "remark11" }
	//
	// if found, `Attrs`, `AttrsExpr` are ignored
	// newdict -> Dict{ Id: 1, Name: "myname", Remark: "remark" }

	// `Assign`, `AssignExpr`
	// with expr
	result, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(refDict.Name.Eq("myname")).
		AssignExpr(refDict.Remark.Value("remark11")).
		FirstOrInit()
	newdict = result.Data
	rowsAffected = result.RowsAffected
	_ = newdict
	_ = rowsAffected
	// NOTE: Where condition use expr, here will not initialize the field of the condition when initializing.
	//  if not found
	// newdict -> Dict{ Remark: "remark11" }
	//
	//  if not found
	// newdict -> Dict{ Name: "non_existing" }
	result, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(&testdata.Dict{
			Name: "myname",
		}).
		Assign(&testdata.Dict{Remark: "remark11"}).
		FirstOrInit()
	newdict = result.Data
	rowsAffected = result.RowsAffected
	_ = newdict
	_ = rowsAffected
	// NOTE: condition not use expr, here will initialize the field of the condition when initializing.
	// if not found
	// newdict -> Dict{ Name: "myname", Remark: "remark11" }
	//
	// if found, `Assign`, `AssignExpr` are set on the struct
	// newdict -> Dict{ Id: 1, Name: "myname", Remark: "remark11" }
}

func Test_Example_AdvanceQuery_FirstOrCreate(t *testing.T) {
	refDict := testdata.Ref_Dict()
	// NOTE!!!: if with expr condition will not initialize the field when creating, so we should use
	// `Attrs`, `AttrsExpr`, `Assign`, `AssignExpr` attributes to indicate these fields.

	// `Attrs`, `AttrsExpr`
	// with expr
	result, _ := rapier.NewExecutor[testdata.Dict](db).
		Where(refDict.Name.Eq("myname")).
		AttrsExpr(refDict.Remark.Value("remark11")).
		FirstOrCreate()
	newdict := result.Data
	rowsAffected := result.RowsAffected
	_ = newdict
	_ = rowsAffected
	// NOTE: Condition use expr. here will not initialize the field of the condition when creating.
	// if not found. initialize with additional attributes
	// SELECT * FROM `dict` WHERE `dict`.`name` = "myname" ORDER BY `dict`.`id` LIMIT 1;
	// INSERT INTO `dict` (`key`,`name`,`is_pin`,`remark`,`created_at`,`updated_at`) VALUES ("","",false,"remark11","2024-03-08 02:20:10.853","2024-03-08 02:20:10.853");
	// newdict -> Dict{ Id: 11, Name: "", Remark: "remark11" } if not found
	//
	// if found, `Attrs`, `AttrsExpr` are ignored.
	// newdict -> Dict{ Id: 11, Name: "myname", Remark: "remark" }

	// with original gorm api
	result, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(&testdata.Dict{
			Name: "myname",
		}).
		Attrs(&testdata.Dict{Remark: "remark11"}).
		FirstOrCreate()
	newdict = result.Data
	rowsAffected = result.RowsAffected
	_ = newdict
	_ = rowsAffected
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
	result, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(refDict.Name.Eq("myname")).
		AssignExpr(refDict.Remark.Value("remark11")).
		FirstOrCreate()
	newdict = result.Data
	rowsAffected = result.RowsAffected
	_ = newdict
	_ = rowsAffected
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

	result, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(&testdata.Dict{
			Name: "myname",
		}).
		Assign(&testdata.Dict{Remark: "remark11"}).
		FirstOrCreate()
	newdict = result.Data
	rowsAffected = result.RowsAffected
	_ = newdict
	_ = rowsAffected
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
}

func Test_Example_AdvanceQuery_Pluck(t *testing.T) {
	var ids []int64

	refDict := testdata.Ref_Dict()
	// with expr api
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprString(refDict.Name)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprBool(refDict.IsPin)
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

	// with original gorm api
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckString("name")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckBool("is_pin")
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
}
func Test_Example_AdvanceQuery_Count(t *testing.T) {
	total, err := rapier.NewExecutor[testdata.Dict](db).Count()
	_ = err
	_ = total
}

func Test_Example_AdvanceQuery_Exist(t *testing.T) {
	refDict := testdata.Ref_Dict()
	b, err := rapier.NewExecutor[testdata.Dict](db).Where(refDict.Id.Eq(100)).Exist()
	_ = err
	_ = b
}

func Test_Example_AdvanceQuery_FindByPage(t *testing.T) {
	rows, total, err := rapier.NewExecutor[testdata.Dict](db).
		FindAllByPage(10, 10)
	_ = err
	_ = rows
	_ = total

	rows, total, err = rapier.NewExecutor[testdata.Dict](db).
		FindAllPaginate(2, 10)
	_ = err
	_ = rows
	_ = total
}
