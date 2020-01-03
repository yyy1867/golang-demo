package deploy

type ColumnInfo struct {
	Id           string `orm:"id" json:"id"`
	Code         string `orm:"code" json:"code"`
	Name         string `orm:"name" json:"name"`
	TableCode    string `orm:"table_code" json:"tableCode"`
	ColumnName   string `orm:"column_name" json:"columnName"`
	ColumnType   string `orm:"column_type" json:"columnType"`
	DefaultValue string `orm:"default_value" json:"defaultvalue"`
	Length       string `orm:"length" json:"length"`
}
