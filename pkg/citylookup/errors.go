// Custom error types
package citylookup

type ErrorNotImplementedError struct{}

func (e ErrorNotImplementedError) Msg() string {
	return "Not Implemented"
}
func (e ErrorNotImplementedError) Error() string {
	return "Not Implemented"
}
