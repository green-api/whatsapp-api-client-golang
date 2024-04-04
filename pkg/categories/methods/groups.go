package methods

type GroupsCategory struct {
	GreenAPI GreenAPIInterface
}

// CreateGroup is designed to create a group chat.
// https://green-api.com/en/docs/api/groups/CreateGroup/
func (c GroupsCategory) CreateGroup(groupName string, chatIds []string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "createGroup", map[string]interface{}{
		"groupName": groupName,
		"chatIds":   chatIds,
	}, "")
}

// UpdateGroupName changes the name of the group chat.
// https://green-api.com/en/docs/api/groups/UpdateGroupName/
func (c GroupsCategory) UpdateGroupName(groupId, groupName string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "updateGroupName", map[string]interface{}{
		"groupId":   groupId,
		"groupName": groupName,
	}, "")
}

// GetGroupData gets group chat data.
// https://green-api.com/en/docs/api/groups/GetGroupData/
func (c GroupsCategory) GetGroupData(groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "getGroupData", map[string]interface{}{
		"groupId": groupId,
	}, "")
}

// AddGroupParticipant adds a participant to the group chat.
// https://green-api.com/en/docs/api/groups/AddGroupParticipant/
func (c GroupsCategory) AddGroupParticipant(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "addGroupParticipant", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// RemoveGroupParticipant removes the participant from the group chat.
// https://green-api.com/en/docs/api/groups/RemoveGroupParticipant/
func (c GroupsCategory) RemoveGroupParticipant(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "removeGroupParticipant", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// SetGroupAdmin designates a member of a group chat as an administrator.
// https://green-api.com/en/docs/api/groups/SetGroupAdmin/
func (c GroupsCategory) SetGroupAdmin(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "setGroupAdmin", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// RemoveAdmin deprives the participant of group chat administration rights.
// https://green-api.com/en/docs/api/groups/RemoveAdmin/
func (c GroupsCategory) RemoveAdmin(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "removeAdmin", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// SetGroupPicture sets the avatar of the group.
// https://green-api.com/en/docs/api/groups/SetGroupPicture/
func (c GroupsCategory) SetGroupPicture(filePath, groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "setGroupPicture", map[string]interface{}{
		"groupId": groupId,
	}, filePath)
}

// LeaveGroup logs the user of the current account out of the group chat.
// https://green-api.com/en/docs/api/groups/LeaveGroup/
func (c GroupsCategory) LeaveGroup(groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "leaveGroup", map[string]interface{}{
		"groupId": groupId,
	}, "")
}
