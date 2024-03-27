package greenapi

type AccountCategory struct {
	GreenAPI GreenAPIInterface
}

func (c AccountCategory) GetSettings() (interface{}, error) {
	return c.GreenAPI.Request("GET", "getSettings", nil)
}
