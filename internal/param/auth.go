package param

type Auth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
}

func NewAuth() *Auth {
	return new(Auth)
}
