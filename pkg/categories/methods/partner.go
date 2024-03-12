package methods

type PartnerCategory struct {
	GreenAPI GreenAPIInterface
}

func (c PartnerCategory) CreateInstance(parameters map[string]interface{}) (map[string]interface{}, error) {

	return c.GreenAPI.PartnerRequest("POST", "createInstance", parameters, "")
}

func (c PartnerCategory) DeleteInstanceAccout(idInstance int) (map[string]interface{}, error) {

	return c.GreenAPI.PartnerRequest("POST", "deleteInstanceAccount", map[string]interface{}{
		"idInstance": idInstance,
	}, "")
}

func (c PartnerCategory) GetInstances() ([]interface{}, error) {

	return c.GreenAPI.ArrayPartnerRequest("GET", "getInstances", nil, "")
}
