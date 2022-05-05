package users

type User struct {
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	PersonalCode uint64 `json:"personalCode,omitempty"`
}
