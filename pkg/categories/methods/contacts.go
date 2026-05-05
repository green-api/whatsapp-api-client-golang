package methods

type ContactsCategory struct {
	GreenAPI GreenAPIInterface
}

// AddContact adds a new contact to the contact list.
// https://green-api.com/en/docs/api/contacts/AddContact/
func (c ContactsCategory) AddContact(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "addContact", parameters, "")
}

// EditContact edits an existing contact in the contact list.
// https://green-api.com/en/docs/api/contacts/EditContact/
func (c ContactsCategory) EditContact(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "editContact", parameters, "")
}

// DeleteContact deletes a contact from the contact list.
// https://green-api.com/en/docs/api/contacts/DeleteContact/
func (c ContactsCategory) DeleteContact(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "deleteContact", parameters, "")
}
