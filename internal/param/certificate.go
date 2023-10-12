package param

type Certificate struct {
	Skip    bool                 `json:"skip,omitempty"`
	Root    string               `json:"root,omitempty"`
	Clients []*ClientCertificate `json:"clients,omitempty"`
}

func NewCertificate() *Certificate {
	return new(Certificate)
}
