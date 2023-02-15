package structs

type AnyMap []interface{}

type Storage struct {
	Data []Image
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    AnyMap `json:"data,omitempty"`
}

type Image struct {
	TargetId int    `json:"targetId"`
	State    string `json:"state"`
	ImageUrl string `json:"imageUrl"`
}

type RobloxResponse struct {
	Data []Image
}
