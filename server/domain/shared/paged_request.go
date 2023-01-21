package shared

type PagedRequest struct {
	Page   uint
	Offset uint
	Limit  uint
}

func NewPagedRequest(limit uint) *PagedRequest {
	return &PagedRequest{Page: 1, Offset: 0, Limit: limit}
}

func (p *PagedRequest) NextPage() {
	p.Page++
	p.Offset += p.Limit
}
