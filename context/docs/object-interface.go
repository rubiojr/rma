// Object is the interface that all object types in Risor must implement.
type Object interface {
	// Type of the object.
	Type() Type

	// Inspect returns a string representation of the given object.
	Inspect() string

	// Interface converts the given object to a native Go value.
	Interface() interface{}

	// Returns True if the given object is equal to this object.
	Equals(other Object) Object

	// GetAttr returns the attribute with the given name from this object.
	GetAttr(name string) (Object, bool)

	// SetAttr sets the attribute with the given name on this object.
	SetAttr(name string, value Object) error

	// IsTruthy returns true if the object is considered "truthy".
	IsTruthy() bool

	// RunOperation runs an operation on this object with the given
	// right-hand side object.
	RunOperation(opType op.BinaryOpType, right Object) Object

	// Cost returns the incremental processing cost of this object.
	Cost() int
}
