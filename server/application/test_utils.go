package application

import "testing"

type RepositorySpy[T any] struct {
	Calls   []interface{}
	Results []T
	Errors  []error
}

func (spy *RepositorySpy[T]) Error() error {
	if len(spy.Errors) > 0 {
		err := spy.Errors[0]
		spy.Errors = spy.Errors[1:]
		return err
	}
	return nil
}

func (spy *RepositorySpy[T]) Result() T {
	if len(spy.Results) > 0 {
		res := spy.Results[0]
		spy.Results = spy.Results[1:]
		return res
	}

	var res T
	return res
}

func (spy *RepositorySpy[T]) CalledOnce(t *testing.T) {
	if !spy.CallTimes(1) {
		t.Errorf("repository expected %d calls received %d\n", 1, len(spy.Calls))
	}
}

func (spy *RepositorySpy[T]) CallTimes(times int) bool {
	return len(spy.Calls) == times
}

func NewRepositoryWithResults[T any](results []T) RepositorySpy[T] {
	return RepositorySpy[T]{Results: results}
}

func NewRepositoryWithErrors[T any](errs []error) RepositorySpy[T] {
	return RepositorySpy[T]{Errors: errs}
}

func NewRepositoryWithResultsAndErrors[T any](results []T, errs []error) RepositorySpy[T] {
	return RepositorySpy[T]{Results: results, Errors: errs}
}

// Helper functions

func CheckPopertyEquality(
	t *testing.T, prop string, entProp, reqProp interface{},
) {
	if entProp != reqProp {
		t.Errorf("%s differs, wanted %v, got %v\n", prop, entProp, reqProp)
	}
}
