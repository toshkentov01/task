package models

// Response ...
type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// Error ...
type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Success ...
type Success struct {
	Success bool `json:"success"`
}
