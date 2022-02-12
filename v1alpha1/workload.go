package v1alpha1

type WorkloadPrototyper interface {
	Metadata() Metadata
	Spec() string
	//GetType() WorkloadTypePrototyper // 继承的约束
}