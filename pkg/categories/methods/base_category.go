package methods

type GreenAPIInterface interface {
	Request(method, APIMethod string, data map[string]interface{}, filePath string) (map[string]interface{}, error)
	RawRequest(method, APIMethod string, data map[string]interface{}, filePath string) (interface{}, error)
	ArrayRequest(method, APIMethod string, data map[string]interface{}, filePath string) ([]interface{}, error)
	PartnerRequest(method, APIMethod string, data map[string]interface{}, filePath string) (map[string]interface{}, error)
}
