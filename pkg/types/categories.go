package types

import "github.com/green-api/whatsapp-api-client-golang/pkg/types/methods"

type GreenAPICategories struct {
	GreenAPI methods.GreenAPIInterface
}

func (c GreenAPICategories) Account() methods.AccountCategory {
	return methods.AccountCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) Device() methods.DeviceCategory {
	return methods.DeviceCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) Groups() methods.GroupsCategory {
	return methods.GroupsCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) Journals() methods.JournalsCategory {
	return methods.JournalsCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) Queues() methods.QueuesCategory {
	return methods.QueuesCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) ReadMark() methods.ReadMarkCategory {
	return methods.ReadMarkCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) Receiving() methods.ReceivingCategory {
	return methods.ReceivingCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) Sending() methods.SendingCategory {
	return methods.SendingCategory{GreenAPI: c.GreenAPI}
}

func (c GreenAPICategories) Service() methods.ServiceCategory {
	return methods.ServiceCategory{GreenAPI: c.GreenAPI}
}
