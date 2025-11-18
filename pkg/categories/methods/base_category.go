package methods

type GreenAPIInterface interface {
	Request(method, APIMethod string, data map[string]any, filePath string) (map[string]any, error)
	RawRequest(method, APIMethod string, data map[string]any, filePath string) (any, error)
	ArrayRequest(method, APIMethod string, data map[string]any, filePath string) ([]any, error)
	PartnerRequest(method, APIMethod string, data map[string]any, filePath string) (map[string]any, error)
	ArrayPartnerRequest(method, APIMethod string, data map[string]any, filePath string) ([]any, error)
}
