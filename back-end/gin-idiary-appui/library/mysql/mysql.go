package mysql

import (
	"context"
	"database/sql"

	"github.com/didi/gendry/scanner"
	_ "github.com/go-sql-driver/mysql"
)

func (dao *client) Query(ctx context.Context, tableName string, where map[string]interface{}, columns []string, data interface{}) error {
	builder := NewSelectBuilder(tableName, where, columns)
	err := QueryWithBuilder(ctx, dao, builder, data)
	if err != nil {
		// err := fmt.Errorf("Query Failed, err: %s", err.Error())
		// logit.Logger.Error(err.Error())
		return err
	}
	// logit.Logger.Info("Query Success")
	return nil
}

func (dao *client) Insert(ctx context.Context, tableName string, data []map[string]interface{}) (sql.Result, error) {
	builder := NewInsertBuilder(tableName, data, insertCommon)
	res, err := ExecWithBuilder(ctx, dao, builder)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (dao *client) InsertIgnore(ctx context.Context, tableName string, data []map[string]interface{}) (sql.Result, error) {
	builder := NewInsertBuilder(tableName, data, insertIgnore)
	res, err := ExecWithBuilder(ctx, dao, builder)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (dao *client) InsertReplace(ctx context.Context, tableName string, data []map[string]interface{}) (sql.Result, error) {
	builder := NewInsertBuilder(tableName, data, insertReplace)
	res, err := ExecWithBuilder(ctx, dao, builder)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (dao *client) Update(ctx context.Context, tableName string, where map[string]interface{}, update map[string]interface{}) (sql.Result, error) {
	builder := NewUpdateBuilder(tableName, where, update)
	res, err := ExecWithBuilder(ctx, dao, builder)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (dao *client) Delete(ctx context.Context, tableName string, where map[string]interface{}) (sql.Result, error) {
	builder := NewDeleteBuilder(tableName, where)
	res, err := ExecWithBuilder(ctx, dao, builder)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (dao *client) ExecRaw(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	builder := NewRawBuilder(sql, args)
	res, err := ExecWithBuilder(ctx, dao, builder)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryWithBuilder 传入一个 SQLBuilder 并执行 QueryContext
func QueryWithBuilder(ctx context.Context, client Client, builder Builder, data interface{}) error {
	db, err := client.connect(ctx)
	if err != nil {
		return err
	}
	cond, values, err := builder.CompileContext(ctx, client)
	if err != nil {
		return err
	}
	rows, err := db.QueryContext(ctx, cond, values...)
	if err != nil {
		return err
	}
	return scanner.ScanClose(rows, data)
}

func ExecWithBuilder(ctx context.Context, client Client, builder Builder) (sql.Result, error) {
	db, err := client.connect(ctx)
	if err != nil {
		return nil, err
	}
	cond, values, err := builder.CompileContext(ctx, client)
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, cond, values...)
}

func Execraw(ctx context.Context, client Client, builder Builder) (sql.Result, error) {
	db, err := client.connect(ctx)
	if err != nil {
		return nil, err
	}
	cond, values, err := builder.CompileContext(ctx, client)
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, cond, values...)
}

var _ Client = (*client)(nil)
