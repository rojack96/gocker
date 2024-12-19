package common

import "github.com/rojack96/gocker/helpers"

const (
	format = "--format"
	index  = "--index"
)

// Index - index of the container if service has multiple replicas
func Index(indexOfContainer int) string {
	return helpers.Int(index, indexOfContainer)
}

func Format(value string) string {
	return helpers.String(format, value)
}
