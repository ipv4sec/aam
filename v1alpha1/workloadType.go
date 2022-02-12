package v1alpha1

type WorkloadTypePrototyper interface {
	Parameter() string
	OptionalTraits() []TraitTypePrototyper
	Metadata() Metadata
}