package methods

type PartnerCategory struct {
	GreenAPI GreenAPIInterface
}

// This method is aimed to create an instace using partner account.
// https://green-api.com/en/docs/partners/createInstance/
func (c PartnerCategory) CreateInstance(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.PartnerRequest("POST", "createInstance", parameters, "")
}

// This method is aimed to delete an instance using partner account.
// https://green-api.com/en/docs/partners/deleteInstanceAccount/
func (c PartnerCategory) DeleteInstanceAccout(idInstance int) (map[string]interface{}, error) {

	return c.GreenAPI.PartnerRequest("POST", "deleteInstanceAccount", map[string]interface{}{
		"idInstance": idInstance,
	}, "")
}

// This method is aimed to get all instances on a partner account.
// https://green-api.com/en/docs/partners/getInstances/
func (c PartnerCategory) GetInstances() ([]interface{}, error) {

	return c.GreenAPI.ArrayPartnerRequest("GET", "getInstances", nil, "")
}
