package param

type Pool struct {
	All  int `json:"all,omitempty"`
	Host int `json:"host,omitempty"`
}

func NewPool() *Pool {
	return &Pool{
		All:  100,
		Host: 10,
	}
}
