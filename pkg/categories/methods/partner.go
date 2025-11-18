package methods

type PartnerCategory struct {
	GreenAPI GreenAPIInterface
}

// CreateInstance is aimed to create an instace using partner account.
// https://green-api.com/en/docs/partners/createInstance/
func (c PartnerCategory) CreateInstance(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.PartnerRequest("POST", "createInstance", parameters, "")
}

// DeleteInstanceAccount is aimed to delete an instance using partner account.
// https://green-api.com/en/docs/partners/deleteInstanceAccount/
func (c PartnerCategory) DeleteInstanceAccount(idInstance int) (map[string]any, error) {
	return c.GreenAPI.PartnerRequest("POST", "deleteInstanceAccount", map[string]any{"idInstance": idInstance}, "")
}

// GetInstances is aimed to get all instances on a partner account.
// https://green-api.com/en/docs/partners/getInstances/
func (c PartnerCategory) GetInstances() ([]any, error) {
	return c.GreenAPI.ArrayPartnerRequest("GET", "getInstances", nil, "")
}
