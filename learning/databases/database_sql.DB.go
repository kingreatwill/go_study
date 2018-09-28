// GO语言中的SQL数据库的泛用接口(database/sql包);

// 安装MySQL驱动:
// $ go get -u github.com/go-sql-driver/mysql

// 可参考:
// https://github.com/go-sql-driver/mysql
// https://github.com/go-sql-driver/mysql/wiki/Examples

// 关于Golang中database/sql包的学习笔记;
// https://segmentfault.com/a/1190000003036452
package databases

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/kingreatwill/go_study/learning/common"

// sqlDB01;
func sqlDB01() {
	common.Println("database_sql.DB.go")

	// (返回一个连接池对象不是单个连接)打开一个dirverName指定的数据库, 返回的DB可以安全的被多个go程同时使用, 并会维护自身的闲置连接池, 很少需要关闭DB;
	// func Open(driverName, dataSourceName string) (*DB, error);
	db, err := sql.Open("mysql", "DemoCloudUser:123456@lcb@tcp(192.168.1.50:3306)/DemoCloud?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	// 返回数据库下层驱动;
	// func (db *DB) Driver() driver.Driver;
	// 检查与数据库的连接是否仍有效, 如果需要将创建连接;
	// func (db *DB) Ping() error;
	// 关闭数据库, 释放任何打开的资源, 一般不会关闭DB, 因为DB句柄通常被多个go程共享, 并长期活跃;
	// func (db *DB) Close() error;
	defer db.Close()

	// 设置与数据库建立连接的最大数目(默认值为0, 表示不限制);
	// func (db *DB) SetMaxOpenConns(n int);
	// 设置连接池中的最大闲置连接数;
	// func (db *DB) SetMaxIdleConns(n int);
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)

	// (少用)执行一次命令(包括查询、删除、更新、插入等), 不返回任何执行结果;
	// func (db *DB) Exec(query string, args ...interface{}) (Result, error);
	// 执行一次查询, 返回多行结果(Rows), 一般用于执行select命令;
	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error);
	rows, err := db.Query("SELECT * FROM GoTest WHERE GoId > ?", 0)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()
	for rows.Next() {
		var goId int64
		var name string
		_ = rows.Scan(&goId, &name)
		fmt.Println("Query -- GoId:", goId, "Name:", name)
	}

	// 执行一次查询, 并期望返回最多一行结果(即Row);
	// func (db *DB) QueryRow(query string, args ...interface{}) *Row;
	var goId int64
	var name string
	_ = db.QueryRow("SELECT * FROM GoTest WHERE GoId = ?", 1).Scan(&goId, &name)
	fmt.Println("QueryRow -- GoId:", goId, "Name:", name)

	// 创建一个准备好的状态用于之后的查询和命令, 返回值可以同时执行多个查询和命令;
	// func (db *DB) Prepare(query string) (*Stmt, error);
	// 开始一个事务, 隔离水平由数据库驱动决定;
	// func (db *DB) Begin() (*Tx, error);
	tx, _ := db.Begin()
	stmt, _ := db.Prepare(`UPDATE GoTest SET Name = ? WHERE GoId = ?`)
	defer stmt.Close()
	result, _ := stmt.Exec("KY测试2", 2)
	n1, _ := result.LastInsertId()
	n2, _ := result.RowsAffected()
	defer tx.Rollback()
	_ = tx.Commit()
	fmt.Println("Begin & Prepare: ", n1, n2)
}
