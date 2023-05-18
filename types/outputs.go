package types

import "time"

/* ***** USERS ***** */
type UserOutput struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

/* ***** EVENTS ***** */
type EventOutput struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Img         string    `json:"img"`
	Location    string    `json:"location"`
	Topic       string    `json:"topic"`
	UserID      int       `json:"user_id"`
}

type TopEventOutput struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	AVG           float64   `json:"avg"`
	Name          string    `json:"name"`
	Date          time.Time `json:"date"`
	Description   string    `json:"description"`
	Img           string    `json:"img"`
	Location      string    `json:"location"`
	Topic         string    `json:"topic"`
	UserID        int       `json:"user_id"`
	Accessibility string    `json:"accessibility"`
}

type SearchEventOutput struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Date          time.Time `json:"date"`
	Description   string    `json:"description"`
	Img           string    `json:"img"`
	Location      string    `json:"location"`
	Topic         string    `json:"topic"`
	UserID        int       `json:"user_id"`
	Accessibility string    `json:"accessibility"`
}

/* ***** GUESTS ***** */
type GuestsOutput struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Status      string `json:"status"`
}

type InvitationsOutput struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Img         string    `json:"img"`
}

/* ***** GRADES ***** */
type GradeOutput struct {
	ID      int    `json:"id"`
	Opinion string `json:"opinion"`
	Grade   int    `json:"grade"`
	User    string `json:"user"`
	EventID uint   `json:"eventID"`
	UserID  uint   `json:"userID"`
}

/* ***** MEMORIES ***** */
type MemoryOutput struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	EventID uint   `json:"eventID"`
}

/* ***** TODOS ***** */
type TodoOutput struct {
	ID       int       `json:"id"`
	Text     string    `json:"text"`
	Date     time.Time `json:"date"`
	Complete bool      `json:"complete"`
	EventID  uint      `json:"eventID"`
}

/* ***** PURCHASES ***** */
type PurchaseOutput struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	EventID uint    `json:"eventID"`
}
