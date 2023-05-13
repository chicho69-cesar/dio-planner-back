package types

/* ***** USERS ***** */
type UserQuery struct {
	Name        string 		 	`json:"name"`
}

/* ***** EVENTS ***** */
type EventQuery struct {
	Name        string 		 	`json:"name"`
	Description string 		 	`json:"description"`
	Location 		string 		 	`json:"location"`
}
