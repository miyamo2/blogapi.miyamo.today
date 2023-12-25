package infra

// ExecuteOption is an option for statement execution.
type ExecuteOption func(Statement)

// Statement is a database statement.
//
// e.g. Select, Insert, Update, Delete
type Statement interface {
	// Execute executes statement.
	Execute(opts ...ExecuteOption) error
	Result() StatementResult
}

// StatementResult is result of statement.
type StatementResult interface {
	// Get returns a result of statement.
	Get() interface{}
	// Set sets a result of statement.
	Set(v interface{})
}

// SingleStatementResult is a single result of statement.
type SingleStatementResult[T any] struct {
	value      T
	hasBeenSet bool
}

func (r *SingleStatementResult[T]) Get() interface{} {
	return r.value
}

// StrictGet returns a result of statement.
// The return value is the type specified in the generics.
func (r *SingleStatementResult[T]) StrictGet() T {
	return r.value
}

func (r *SingleStatementResult[T]) Set(v interface{}) {
	if r.hasBeenSet {
		return
	}
	r.value = v.(T)
}

func NewSingleStatementResult[T any]() *SingleStatementResult[T] {
	return &SingleStatementResult[T]{}
}

// MultipleStatementResult is a results iterator of statement.
type MultipleStatementResult[T any] struct {
	value      []T
	hasBeenSet bool
	next       int
}

func (r *MultipleStatementResult[T]) Get() interface{} {
	return r.value
}

// StrictGet returns a results of statement.
// The return value is a slice of the type specified in the generics.
func (r *MultipleStatementResult[T]) StrictGet() []T {
	return r.value
}

func (r *MultipleStatementResult[T]) Set(v interface{}) {
	if r.hasBeenSet {
		return
	}
	r.value = v.([]T)
}

func (r *MultipleStatementResult[T]) Next() T {
	defer func() {
		r.next++
	}()
	return r.value[r.next]
}

func (r *MultipleStatementResult[T]) HasNext() bool {
	return r.next < len(r.value)-1
}

func NewMultipleStatementResult[T any]() *MultipleStatementResult[T] {
	return &MultipleStatementResult[T]{}
}
