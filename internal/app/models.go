package app

type options struct {
	Addr    string `json:"addr"`
	Reflect bool   `json:"reflect"`

	Insecure   bool   `json:"insecure"`
	Plaintext  bool   `json:"plaintext"`
	Rootca     string `json:"rootca"`
	Clientcert string `json:"clientcert"`
	Clientkey  string `json:"clientkey"`
}

type methodSelect struct {
	FullName string `json:"full_name"`
	Name     string `json:"name"`
}

type methodsSelect []methodSelect

type serviceSelect struct {
	FullName string        `json:"full_name"`
	Methods  methodsSelect `json:"methods"`
}

type servicesSelect []serviceSelect
