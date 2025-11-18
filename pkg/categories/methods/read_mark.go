package methods

type ReadMarkCategory struct {
	GreenAPI GreenAPIInterface
}

// ReadChat is designed to mark chat messages as read.
// https://green-api.com/en/docs/api/marks/ReadChat/
func (c ReadMarkCategory) ReadChat(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "readChat", parameters, "")
}
