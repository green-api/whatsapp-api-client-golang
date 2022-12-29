package methods

type GroupsCategory struct {
	GreenAPI GreenAPIInterface
}

func (c GroupsCategory) CreateGroup(groupName string, chatIds []string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "CreateGroup", map[string]interface{}{
		"groupName": groupName,
		"chatIds":   chatIds,
	}, "")
}

func (c GroupsCategory) UpdateGroupName(groupId, groupName string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "UpdateGroupName", map[string]interface{}{
		"groupId":   groupId,
		"groupName": groupName,
	}, "")
}

func (c GroupsCategory) GetGroupData(groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "GetGroupData", map[string]interface{}{
		"groupId": groupId,
	}, "")
}

func (c GroupsCategory) AddGroupParticipant(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "AddGroupParticipant", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

func (c GroupsCategory) RemoveGroupParticipant(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "RemoveGroupParticipant", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

func (c GroupsCategory) SetGroupAdmin(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "SetGroupAdmin", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

func (c GroupsCategory) RemoveAdmin(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "RemoveAdmin", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

func (c GroupsCategory) SetGroupPicture(filePath, groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "setGroupPicture", map[string]interface{}{
		"groupId": groupId,
	}, filePath)
}

func (c GroupsCategory) LeaveGroup(groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "LeaveGroup", map[string]interface{}{
		"groupId": groupId,
	}, "")
}
