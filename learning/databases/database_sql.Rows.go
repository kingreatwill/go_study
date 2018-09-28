// GO语言中的SQL数据库的泛用接口(database/sql包);
package databases

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/kingreatwill/go_study/learning/common"

// sqlRows01;
func sqlRows01() {
	common.Println("database_sql.Rows.go")

	// 连接;
	db, _ := sql.Open("mysql", "DemoCloudUser:123456@lcb@tcp(192.168.1.50:3306)/DemoCloud?charset=utf8")
	defer db.Close()
	// 检索;
	rows, _ := db.Query("SELECT * FROM GoTest WHERE GoId > ?", 0)

	// 返回列名;
	// func (rs *Rows) Columns() ([]string, error);
	cols, _ := rows.Columns()
	fmt.Println("Columns:", cols)

	// 下一行结果;
	// func (rs *Rows) Next() bool;
	// 将当前行各列结果填充进dest指定的各个值中;
	// func (rs *Rows) Scan(dest ...interface{}) error;
	for rows.Next() {
		var goId int64
		var name string
		_ = rows.Scan(&goId, &name)
		fmt.Println("Next+Scan -- GoId:", goId, "Name:", name)
	}

	// 关闭Rows, 阻止对其更多的列举, 如果Next方法返回假, Rows会自动关闭;
	// func (rs *Rows) Close() error;
	// 返回可能的、在迭代时出现的错误, Err需在显式或隐式调用Close方法后调用;
	// func (rs *Rows) Err() error;
}
