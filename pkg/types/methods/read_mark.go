package methods

type ReadMarkCategory struct {
	GreenAPI GreenAPIInterface
}

func (c ReadMarkCategory) ReadChat(parameters map[string]interface{}) map[string]interface{} {
	return c.GreenAPI.Request("POST", "ReadChat", parameters, "")
}
