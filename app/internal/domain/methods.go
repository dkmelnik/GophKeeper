package domain

// Method represents a string identifier for handler methods.
type Method string

// String returns the string representation of Method.
func (m Method) String() string {
	return string(m)
}

// Constants representing specific handler methods.
var (
	UserRegisterMethod Method = "UserHandler.Register"
	UserLoginMethod    Method = "UserHandler.Login"

	TextCreateMethod Method = "TextHandler.Create"
	TextListMethod   Method = "TextHandler.List"

	CardCreateMethod Method = "CardHandler.Create"
	CardListMethod   Method = "CardHandler.List"

	BinaryCreateMethod Method = "BinaryHandler.Create"
	BinaryListMethod   Method = "BinaryHandler.List"
)
