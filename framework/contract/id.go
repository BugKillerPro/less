package contract

const IDKey = "less:id"

type IDService interface {
	NewID() string
}
