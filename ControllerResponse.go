package main

type ResponseController struct {
	ID    string `json:"ID"`
	STATE string `json:"STATE"`
	ERROR string `json:"ERROR"`
	CODE  string `json:"CODE"`
}

func NewResponseControllerEmpty() *ResponseController {
	return &ResponseController{
		ID:    "0",
		STATE: "0",
		ERROR: "0",
		CODE:  "0",
	}
}

func NewResponseController(id string, state string, _error string, code string) *ResponseController {
	return &ResponseController{
		ID:    id,
		STATE: state,
		ERROR: _error,
		CODE:  code,
	}
}
