package crud

type Table struct {
	TableName string // database table name
	ModelName string // go model file struct name
	ModelPkg  string // go model package module
}
