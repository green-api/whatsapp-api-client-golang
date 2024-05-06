package greenapi

type GreenAPIInterface interface {
	Request(httpMethod, APImethod string, data map[string]interface{}) (any, error)
}
