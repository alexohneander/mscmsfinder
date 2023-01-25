package types

type ParseResponse struct {
	Status  int
	Message string
	Payload Payload
	Found   bool
}

type Payload struct {
	Body   string
	Header any
}
