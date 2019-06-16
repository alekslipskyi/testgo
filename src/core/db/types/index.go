package types

type (
	Where map[string]interface{}
)

type QueryOptions struct {
	Attributes []string
	Where      Where
}
