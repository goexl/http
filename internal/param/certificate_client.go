package param

type ClientCertificate struct {
	Public  string `json:"public,omitempty"`
	Private string `json:"private,omitempty"`
}
