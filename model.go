//go:generate msgp
package main

const (
	// Standard primitive types.
	String  = "string"
	Boolean = "boolean"
	Integer = "integer"
	Float   = "float"
)

// Version represents a semantic version.
type Version [3]uint8

// Domain represents a group of concepts that apply to a use case.
type Domain struct {
	Name string
	Doc  string
}

// Concept is a description of a concept that may be applicable to
// multiple domains. The fields provide non-functional information, but
// also fields that can provide functional utility for query and selection.
type Concept struct {
	// Stable identifier of the concept.
	ID string

	// The category the concept is in.
	Category string

	// The free-text name of the entity.
	Name string

	// The version of the concept.
	Version Version

	// Documentation for the concept.
	Doc string

	// Arbitrary set of keywords for free-text indexing.
	Keywords string

	// Status of the concept. TBD.
	Status string

	// If true, denotes the concept is published.
	Published bool

	// The type of values that are associated to this concept. If omitted,
	// the name denotes the type. Unlike
	Type string

	// Parameters are the parameters of the concept.
	Parameters map[string]*Parameter
}

// Parameter represents a query parameter for a concept. An argument for the
// parameter is composed of the
type Parameter struct {
	// Name of the parameter
	Name string

	// Document of the parameter.
	Doc string

	// Type is the default type for the parameter. Operators defined for this
	// parameter can override the type.
	Type string

	// The set of operators that apply to the parameter.
	Operators map[string]*Operator

	// If true, an argument for this parameter is required. That is, the concept
	// cannot be queried without supplying an argument for this parameter.
	Required bool
}

// Operator is essentially a user-defined function (UDF) in the domain.
// In this form, it is generalized for a set of types if defined.
type Operator struct {
	// Name is the unique name of the operation across all uses of it.
	Name string

	// Doc is documentation of the operation.
	Doc string

	// Type overrides the type defined by the parameter the operator is bound to.
	Type string

	// If true, the operator can be applied to multiple values.
	Multiple bool
}
