package v1alpha1

type ApplicationPrototyper interface {
	Metadata() ApplicationMeta
	Workloads() []Workload
	Exports() map[string][]string
	Dependencies() []Dependency
}

type Metadata struct{
	Name string
	Annotations map[string]string
}

// ApplicationMeta is Here
type ApplicationMeta struct {
	Name         string
	Version      string
	Description  string
	Keywords     []string
	Author       string
	Maintainers  []Maintainer
	Repositories []string
	Bugs         string
	Licenses     []License
	Annotations  map[string]string
}

type Workload struct {
	Name       string
	Type       WorkloadTypePrototyper
	Vendor     WorkloadPrototyper
	Properties Properties
	Traits     []struct {
		Type       TraitTypePrototyper
		Properties Properties
	}
}

type Properties map[string]interface{}

type Maintainer struct {
	Email string
	Name  string
	Web   string
}

type License struct {
	Type string
	URL  string
}

type Dependency struct {
	Name     string
	Version  string
	Location string
	Items    map[string][]string
}
