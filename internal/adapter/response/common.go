package response

import "github.com/arvinpaundra/ngekost-api/pkg/helper/format"

type (
	WithPagination struct {
		Results    any
		Pagination format.Pagination
	}
)
