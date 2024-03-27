package greenapi

type GreenAPIInterface interface {
	Request(httpMethod, APImethod string, data map[string]interface{}) (interface{}, error)
	ArrayRequest(url, method string, data map[string]interface{}, filePath string) ([]interface{}, error)
	GetURL(method, APIMethod string, data map[string]interface{}) string
	GetPartnerURL(APIMethod string) (string, error)
}
