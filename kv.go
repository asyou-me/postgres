package postgres

func (this *DB) Get(table, path, req string, out ReflectTableInterface) (err error) {
	var relSlice *[]interface{}
	var reStr *string

	query := `SELECT ` + *reStr + ` FROM "` + table + `" ` + req
	err = this.Pool.QueryRow(query).Scan(*relSlice...)

	this.Info(query)

	return err
}
