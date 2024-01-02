package gorm

// Pagination is a struct for pagination setting.
type Pagination struct {
	cursor         string
	limit          int
	previousPaging bool
	nextPaging     bool
}

// PreviousPaging returns true if this paging is previous paging otherwise false.
func (p *Pagination) PreviousPaging() bool {
	return p.previousPaging
}

// NextPaging returns true if this paging is next paging otherwise false.
func (p *Pagination) NextPaging() bool {
	return p.nextPaging
}

// Cursor returns the cursor.
func (p *Pagination) Cursor() string {
	return p.cursor
}

// Limit returns the limit.
func (p *Pagination) Limit() int {
	return p.limit
}

type paginationOption func(pagination *Pagination)

// WithPreviousPaging returns a PaginationOption for previous paging.
func WithPreviousPaging(limit int, cursor *string) paginationOption {
	return func(pagination *Pagination) {
		if cursor != nil {
			pagination.cursor = *cursor
		}
		pagination.limit = limit
		pagination.previousPaging = true
		pagination.nextPaging = false
	}
}

// WithNextPaging returns a PaginationOption for next paging.
func WithNextPaging(limit int, cursor *string) paginationOption {
	return func(pagination *Pagination) {
		if cursor != nil {
			pagination.cursor = *cursor
		}
		pagination.limit = limit
		pagination.previousPaging = false
		pagination.nextPaging = true
	}
}
