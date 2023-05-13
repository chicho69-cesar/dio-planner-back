package types

/* ***** USER ***** */
type FacebookUserRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GoogleUserRes struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}

/* ***** GUESTS ***** */
type GuestResponse struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	UserID  uint   `json:"userID"`
	EventID uint   `json:"eventID"`
	Status  string `json:"status"`
}
