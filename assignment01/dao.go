package dao

import (
	"database/sql"
	"errors"
	"log"
)

// sql.ErrNoRows
// 我认为不应该 Wrap这个error抛给上层
// 对于ErrNoRows 是一个查询结果为空的说明，它表示数据库成功执行了sql，返回为空是因为没有符合条件的行
// 在dao层可以直接返回nil，error，对ErrNoRows的处理交给业务层

var db *sql.DB

func init() {
	db, err := sql.Open("mysql", "url")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

func QueryRows(id int) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM table_name WHERE id = ?", id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, err
	case err != nil:
		// 权限、sql语法、网络等异常可以wrap更详细的信息
		return nil, errors.New("wrap an error with more detail")
	}
	return rows, nil
}
