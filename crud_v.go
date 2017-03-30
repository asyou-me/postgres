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
	q.Engine.Info(`INSERT INTO "` + q.table + `" (` + fields + `) VALUES (` + inserts + `)`)
	commandTag, err := q.Engine.Exec(`INSERT INTO "`+q.table+`" (`+fields+`) VALUES (`+inserts+`)`, values...)
	fmt.Println("err:", err)
	return string(commandTag), err
}

// UpdateV 更新数据到数据库
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
	q.Engine.Info(`UPDATE "` + q.table + `" SET ` + sets + q.whereStr())
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
	query := `SELECT ` + gets + ` FROM "` + q.table + `"` + q.whereStr()
	q.Engine.Info(query)
	err = q.Engine.QueryRow(query).Scan(values...)
	if err != nil {
		return
	}
	return nil
}

// ListV 获取字段数据
func (q *QueryBuilder) ListV(files map[string]uint8, limit int, offset int) ([]*map[string]*V, error) {
	var gets string
	var indexData = len(files) - 1
	var index = 0
	var fs = make([]string, len(files))
	for k := range files {
		if index == indexData {
			gets = gets + "\"" + k + "\""
		} else {
			gets = gets + "\"" + k + "\"" + `,`
		}
		fs[index] = k
		index = index + 1
	}
	outs := make([]*map[string]*V, 0, limit)

	query := `SELECT ` + gets + ` FROM "` + q.table + `"` + q.whereStr() + "" + ` LIMIT ` + fmt.Sprint(limit) + ` OFFSET ` + fmt.Sprint(offset)
	q.Engine.Info(query)
	rows, err := q.Engine.Query(query)
	defer rows.Close()

	if err != nil {
		q.Engine.Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		item := map[string]*V{}
		values := make([]interface{}, len(files))
		i := 0
		for _, k := range fs {
			v := &V{
				T: files[k],
			}
			item[k] = v
			values[i] = v
			i = i + 1
		}
		err = rows.Scan(values...)
		if err != nil {
			q.Engine.Warn(err.Error())
			continue
		}
		outs = append(outs, &item)
	}
	return outs, nil
}

// Search 获取字段数据
func (q *QueryBuilder) Search(files map[string]uint8) ([]*map[string]*V, error) {
	var gets string
	var indexData = len(files) - 1
	var index = 0
	var fs = make([]string, len(files))
	for k := range files {
		if index == indexData {
			gets = gets + "`" + k + "`"
		} else {
			gets = gets + "`" + k + "`,"
		}
		fs[index] = k
		index = index + 1
	}
	outs := make([]*map[string]*V, 0, 10)
	query := `SELECT ` + gets + ` FROM "` + q.table + `"` + q.whereStr() + q.OrderStr()
	q.Engine.Info(query)
	rows, err := q.Engine.Query(query)
	defer rows.Close()

	if err != nil {
		q.Engine.Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		item := map[string]*V{}
		values := make([]interface{}, len(files))
		i := 0
		for _, k := range fs {
			v := &V{
				T: files[k],
			}
			item[k] = v
			values[i] = v
			i = i + 1
		}
		err = rows.Scan(values...)
		if err != nil {
			continue
		}
		outs = append(outs, &item)
	}
	return outs, nil
}
