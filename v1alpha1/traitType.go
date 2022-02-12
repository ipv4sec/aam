package v1alpha1



type TraitTypePrototyper interface {
	Spec() string
	Metadata() Metadata
}

