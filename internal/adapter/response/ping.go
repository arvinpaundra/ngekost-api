package response

type (
	Ping struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Author  string `json:"author"`
		Github  string `json:"github"`
	}
)
