package methods

type GroupsCategory struct {
	GreenAPI GreenAPIInterface
}

// CreateGroup is designed to create a group chat
func (c GroupsCategory) CreateGroup(groupName string, chatIds []string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "CreateGroup", map[string]interface{}{
		"groupName": groupName,
		"chatIds":   chatIds,
	}, "")
}

// UpdateGroupName changes the name of the group chat
func (c GroupsCategory) UpdateGroupName(groupId, groupName string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "UpdateGroupName", map[string]interface{}{
		"groupId":   groupId,
		"groupName": groupName,
	}, "")
}

// GetGroupData gets group chat data
func (c GroupsCategory) GetGroupData(groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "GetGroupData", map[string]interface{}{
		"groupId": groupId,
	}, "")
}

// AddGroupParticipant adds a participant to the group chat
func (c GroupsCategory) AddGroupParticipant(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "AddGroupParticipant", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// RemoveGroupParticipant removes the participant from the group chat
func (c GroupsCategory) RemoveGroupParticipant(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "RemoveGroupParticipant", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// SetGroupAdmin designates a member of a group chat as an administrator
func (c GroupsCategory) SetGroupAdmin(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "SetGroupAdmin", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// RemoveAdmin deprives the participant of group chat administration rights
func (c GroupsCategory) RemoveAdmin(groupId, participantChatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "RemoveAdmin", map[string]interface{}{
		"groupId":           groupId,
		"participantChatId": participantChatId,
	}, "")
}

// SetGroupPicture sets the avatar of the group
func (c GroupsCategory) SetGroupPicture(filePath, groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "setGroupPicture", map[string]interface{}{
		"groupId": groupId,
	}, filePath)
}

// LeaveGroup logs the user of the current account out of the group chat
func (c GroupsCategory) LeaveGroup(groupId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "LeaveGroup", map[string]interface{}{
		"groupId": groupId,
	}, "")
}
