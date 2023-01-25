package types

type ParseResponse struct {
	Status  int
	Message string
	Payload Payload
}

type Payload struct {
	Body   any
	Header any
}
