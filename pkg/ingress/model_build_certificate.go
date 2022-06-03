package ingress

import (
	"context"

	"sigs.k8s.io/aws-load-balancer-controller/pkg/annotations"
	acm "sigs.k8s.io/aws-load-balancer-controller/pkg/model/acm"
)

const (
	resourceIDCertificate = "Certificate"
)

func (t *defaultModelBuildTask) buildCertificate(ctx context.Context) (*acm.Certificate, error) {
	certSpec, err := t.buildCertificateSpec(ctx)
	if err != nil {
		return nil, err
	}
	cert := acm.NewCertificate(t.stack, resourceIDCertificate, *certSpec)
	return cert, nil
}

func (t *defaultModelBuildTask) buildCertificateSpec(ctx context.Context) (*acm.CertificateSpec, error) {
	certHostName := t.computeIngressHostedZoneCertificate(ctx)
	if certHostName == "" {
		return nil, nil
	}
	return &acm.CertificateSpec{
		Name: certHostName,
		Tags: map[string]string{"ManagedBy": "Colum-SMWYG"},
	}, nil

}

func (t *defaultModelBuildTask) computeIngressHostedZoneCertificate(ctx context.Context) string {
	// if exists := t.annotationParser.ParseStringAnnotation(annotations.IngressHostedZoneCertificate, &certHostName, ing.Annotations); !exists {
	// 	return nil, errors.Errorf("empty ingress-hosted-zone-certificate configuration: `%s`", certHostName)
	// }

	certHostName := ""
	for _, member := range t.ingGroup.Members {
		if exists := t.annotationParser.ParseStringAnnotation(annotations.IngressHostedZoneCertificate, &certHostName, member.Ing.Annotations); !exists {
			continue
		}
	}
	return certHostName
}
