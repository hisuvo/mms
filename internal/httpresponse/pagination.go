package httpresponse

type PaginatedResponse[T any] struct {
	Meta Meta `json:"meta"`
	Data []T  `json:"data"`
}