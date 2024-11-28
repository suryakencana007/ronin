package ronin

// Meta holds the response definition for the Meta entity.
type Meta struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

// Version holds the response definition for the Version entity.
type Version struct {
	Label  string `json:"label,omitempty"`
	Number string `json:"number,omitempty"`
}

// Pagination holds the response definition for the Pagination entity.
type Pagination struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"per_page,omitempty"`
	Size  int `json:"page_count,omitempty"`
	Total int `json:"total_count,omitempty"`
}

// Response represents response body of this API.
type Response struct {
	Meta       `json:"meta"`
	Version    `json:"version"`
	Pagination `json:"pagination,omitempty"`
	Data       any `json:"data,omitempty"`
}
