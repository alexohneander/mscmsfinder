package types

type ParseResponse struct {
	Status  int
	Message string
	Payload Payload
	Found   bool
	CMS     string
}

type Payload struct {
	Body   string
	Header any
}
