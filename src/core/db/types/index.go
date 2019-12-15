package types

type (
	Where      map[string]interface{}
	Attributes []string
	GroupBy    []string
	Include    struct {
		TableName  string
		FkTableId  string
		RefTableID string
		JoinType   string
		AS         string
	}
)

type QueryOptions struct {
	Attributes Attributes
	Where      Where
	Includes   []Include
	GroupBy    GroupBy
	AS         string
}
