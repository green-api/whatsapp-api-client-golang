package categories

import "github.com/green-api/whatsapp-api-client-golang/pkg/categories/methods"

type GreenAPICategories struct {
	GreenAPI methods.GreenAPIInterface
}

// Account category presents methods for working with the account.
// https://green-api.com/en/docs/api/account/
func (c GreenAPICategories) Account() methods.AccountCategory {
	return methods.AccountCategory{GreenAPI: c.GreenAPI}
}

// Device category presents methods for working with the device (phone).
// https://green-api.com/en/docs/api/phone/
func (c GreenAPICategories) Device() methods.DeviceCategory {
	return methods.DeviceCategory{GreenAPI: c.GreenAPI}
}

// Groups category presents methods for working with group chats.
// https://green-api.com/en/docs/api/groups/
func (c GreenAPICategories) Groups() methods.GroupsCategory {
	return methods.GroupsCategory{GreenAPI: c.GreenAPI}
}

// Journals present methods for working with incoming and outgoing messages.
// https://green-api.com/en/docs/api/journals/
func (c GreenAPICategories) Journals() methods.JournalsCategory {
	return methods.JournalsCategory{GreenAPI: c.GreenAPI}
}

// Queues category presents methods for working with a messages queue.
// https://green-api.com/en/docs/api/queues/
func (c GreenAPICategories) Queues() methods.QueuesCategory {
	return methods.QueuesCategory{GreenAPI: c.GreenAPI}
}

// ReadMark category presents methods for working with chat read mark.
// https://green-api.com/en/docs/api/marks/
func (c GreenAPICategories) ReadMark() methods.ReadMarkCategory {
	return methods.ReadMarkCategory{GreenAPI: c.GreenAPI}
}

// Receiving category presents methods for working with receiving events.
// https://green-api.com/en/docs/api/receiving/
func (c GreenAPICategories) Receiving() methods.ReceivingCategory {
	return methods.ReceivingCategory{GreenAPI: c.GreenAPI}
}

// Sending category presents methods for sending different messages.
// https://green-api.com/en/docs/api/sending/
func (c GreenAPICategories) Sending() methods.SendingCategory {
	return methods.SendingCategory{GreenAPI: c.GreenAPI}
}

// Service category presents different service methods.
// https://green-api.com/en/docs/api/service/
func (c GreenAPICategories) Service() methods.ServiceCategory {
	return methods.ServiceCategory{GreenAPI: c.GreenAPI}
}

// Partner category presents exclusive methods for partners.
// The partnership scheme involves deeper integration with the service
// and working with a larger number of instances on your side:
//
// * Instance management via API
// * Postpaid billing system (starting from the second month of operation)
// * Daily billing (for created and not deleted instances)
// * Dedicated support line
// For questions regarding connection to the partnership scheme
// and additional conditions, please contact us via email
// at support@green-api.com or via chat on the website.
// https://green-api.com/en/docs/partners/
func (c GreenAPICategories) Partner() methods.PartnerCategory {
	return methods.PartnerCategory{GreenAPI: c.GreenAPI}
}
