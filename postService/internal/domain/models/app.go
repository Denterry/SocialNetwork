package models

// App represent the structure of the app component of the service
type App struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}
