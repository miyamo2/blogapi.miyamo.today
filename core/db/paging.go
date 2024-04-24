package db

// Pagination is a struct for pagination setting.
type Pagination struct {
	cursor           string
	limit            int
	isPreviousPaging bool
	isNextPaging     bool
}

// IsPreviousPaging returns true if this paging is previous paging otherwise false.
func (p *Pagination) IsPreviousPaging() bool {
	return p.isPreviousPaging
}

// IsNextPaging returns true if this paging is next paging otherwise false.
func (p *Pagination) IsNextPaging() bool {
	return p.isNextPaging
}

// Cursor returns the cursor.
func (p *Pagination) Cursor() string {
	return p.cursor
}

// Limit returns the limit.
func (p *Pagination) Limit() int {
	return p.limit
}

type PaginationOption func(pagination *Pagination)

// WithPreviousPaging returns a PaginationOption for previous paging.
func WithPreviousPaging(limit int, cursor *string) PaginationOption {
	return func(pagination *Pagination) {
		if cursor != nil {
			pagination.cursor = *cursor
		}
		pagination.limit = limit
		pagination.isPreviousPaging = true
		pagination.isNextPaging = false
	}
}

// WithNextPaging returns a PaginationOption for next paging.
func WithNextPaging(limit int, cursor *string) PaginationOption {
	return func(pagination *Pagination) {
		if cursor != nil {
			pagination.cursor = *cursor
		}
		pagination.limit = limit
		pagination.isPreviousPaging = false
		pagination.isNextPaging = true
	}
}
