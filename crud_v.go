package postgres

import (
	"errors"
	"fmt"
)

// InsertV 插入数据到数据库
func (q *QueryBuilder) InsertV(data map[string]*V) (string, error) {
	var inserts = ""
	var fields = ""
	var lenData = len(data)
	var indexData = lenData - 1
	var values = make([]interface{}, lenData)
	var index int
	for k, v := range data {
		if index == indexData {
			fields = fields + "\"" + k + "\""
			inserts = inserts + `$` + fmt.Sprint(index+1)
		} else {
			fields = fields + "\"" + k + "\"" + `,`
			inserts = inserts + `$` + fmt.Sprint(index+1) + `,`
		}
		values[index] = v
		index = index + 1
	}
	commandTag, err := q.Engine.Exec(`INSERT INTO "`+q.table+`" (`+fields+`) VALUES (`+inserts+`)`, values...)
	return string(commandTag), err
}

// UpdateV 更新数据到数据库
//
// req:查询条件 sql写法 where xxx
//
// data:需要更新的数据的对象
//
// column:需要更新的字段
func (q *QueryBuilder) UpdateV(data map[string]*V) (string, error) {
	if q.where == "" {
		return "", errors.New("更新条件不能为空")
	}
	var sets = ""
	var lenData = len(data)
	var indexData = lenData - 1
	var index int
	var values = make([]interface{}, lenData)
	for k, v := range data {
		if index == indexData {
			sets = sets + "\"" + k + "\"" + `=$` + fmt.Sprint(index+1)
		} else {
			sets = sets + "\"" + k + "\"" + `=$` + fmt.Sprint(index+1) + `,`
		}
		values[index] = v
		index = index + 1
	}
	commandTag, err := q.Engine.Exec(`UPDATE "`+q.table+`" SET `+sets+q.whereStr(), values...)
	return string(commandTag), err
}

// GetV 获取字段数据
func (q *QueryBuilder) GetV(data map[string]*V) (err error) {
	var gets = ""
	var values = make([]interface{}, len(data))
	var indexData = len(data) - 1
	var index = 0
	for k, v := range data {
		if index == indexData {
			gets = gets + "\"" + k + "\""
		} else {
			gets = gets + "\"" + k + "\"" + `,`
		}
		values[index] = v
		index = index + 1
	}
	fmt.Println(`SELECT ` + gets + ` FROM "` + q.table + `"` + q.whereStr())
	err = q.Engine.QueryRow(`SELECT ` + gets + ` FROM "` + q.table + `"` + q.whereStr()).Scan(values...)
	if err != nil {
		return
	}
	return nil
}
