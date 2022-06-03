package acm

import "sigs.k8s.io/aws-load-balancer-controller/pkg/model/core"

type Certificate struct {
	core.ResourceMeta `json:"-"`

	// desired state of ListenerRule
	Spec CertificateSpec `json:"spec"`

	// observed state of ListenerRule
	// +optional
	Status *CertificateStatus `json:"status,omitempty"`
}

// NewCertificate constructs new Certificate resource.
func NewCertificate(stack core.Stack, id string, spec CertificateSpec) *Certificate {
	lr := &Certificate{
		ResourceMeta: core.NewResourceMeta(stack, "AWS::ElasticLoadBalancingV2::ListenerCertificate", id),
		Spec:         spec,
		Status:       nil,
	}
	stack.AddResource(lr)
	lr.registerDependencies(stack)
	return lr
}

// register dependencies for Certificate.
func (ls *Certificate) registerDependencies(stack core.Stack) {
	// for _, sgToken := range lb.Spec.SecurityGroups {
	// 	for _, dep := range sgToken.Dependencies() {
	// 		stack.AddDependency(dep, lb)
	// 	}
	// }
	// todo: possible dependencies when doing cert verification
}

// CertificateSpec defines the desired state of LoadBalancer
type CertificateSpec struct {
	// The name of the certificate.
	Name string `json:"name"`

	// The status of the certificate (issue/ pending verification)
	Status string `json:"status"`

	// The tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// CertificateStatus defines the observed state of Certificate
type CertificateStatus struct {
	// The Amazon Resource Name (ARN) of the certificate.
	CertificateARN string `json:"certificateARN"`

	// The hostname on the certificate
	Hostname string `json:"hostname"`
}
