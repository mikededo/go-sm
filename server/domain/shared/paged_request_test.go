package shared_test

import (
	"testing"

	"github.com/mddg/go-sm/server/domain/shared"
)

func validatePagedRequestProperties(t *testing.T, got *shared.PagedRequest, page, offset, limit uint) {
	if got.Page != page {
		t.Errorf("wanted %d page, got %d\n", got.Page, page)
	}
	if got.Offset != offset {
		t.Errorf("wanted %d offset, got %d\n", got.Offset, offset)
	}
	if got.Limit != limit {
		t.Errorf("wanted %d limit, got %d\n", got.Limit, limit)
	}
}

func TestPageRequest(t *testing.T) {
	t.Run("create new page", func(t *testing.T) {
		validatePagedRequestProperties(t, shared.NewPagedRequest(25), 1, 0, 25)
	})
}

func TestPageRequest_NextPage(t *testing.T) {
	t.Run("increase to next page", func(t *testing.T) {
		p := shared.NewPagedRequest(25)
		p.NextPage()
		validatePagedRequestProperties(t, p, 2, 25, 25)
		p.NextPage()
		validatePagedRequestProperties(t, p, 3, 50, 25)
	})
}
