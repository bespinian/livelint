package livelint

import (
	"testing"

	"github.com/matryer/is"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	existingIngressClassName    = "name2"
	nonExistingIngressClassName = "name3"
)

func TestCheckHasValidIngressClass(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		ingress         netv1.Ingress
		ingressClasses  map[string]netv1.IngressClass
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds, if there is only one ingress class and the ingress does not specify an ingress class",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{},
				},
			},
			ingressClasses:  map[string]netv1.IngressClass{"name1": {ObjectMeta: metav1.ObjectMeta{Name: "name1"}}},
			expectedToFail:  false,
			expectedMessage: "Ingress INGRESSNAME has a valid ingress class",
		},
		{
			it: "succeeds, if there is an explicit default ingress class and the ingress does not specify an ingress class",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{},
				},
			},
			ingressClasses:  map[string]netv1.IngressClass{"name1": {ObjectMeta: metav1.ObjectMeta{Name: "name1"}}, "name2": {ObjectMeta: metav1.ObjectMeta{Name: "name2", Annotations: map[string]string{"ingressclass.kubernetes.io/is-default-class": "true"}}}},
			expectedToFail:  false,
			expectedMessage: "Ingress INGRESSNAME has a valid ingress class",
		},
		{
			it: "fails, if there is no default ingress class and the ingress does not specify an ingress class",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{},
				},
			},
			ingressClasses:  map[string]netv1.IngressClass{"name1": {ObjectMeta: metav1.ObjectMeta{Name: "name1"}}, "name2": {ObjectMeta: metav1.ObjectMeta{Name: "name2", Annotations: map[string]string{}}}},
			expectedToFail:  true,
			expectedMessage: "Ingress INGRESSNAME does not specify an ingress class and there is no default ingress class in the cluster",
		},
		{
			it: "succeeds, if the ingress specifies an existing ingress class",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					IngressClassName: &existingIngressClassName,
					Rules:            []netv1.IngressRule{},
				},
			},
			ingressClasses:  map[string]netv1.IngressClass{"name1": {ObjectMeta: metav1.ObjectMeta{Name: "name1"}}, "name2": {ObjectMeta: metav1.ObjectMeta{Name: "name2", Annotations: map[string]string{}}}},
			expectedToFail:  false,
			expectedMessage: "Ingress INGRESSNAME has a valid ingress class",
		},
		{
			it: "fails, if the ingress specifies an non-existing ingress class",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					IngressClassName: &nonExistingIngressClassName,
					Rules:            []netv1.IngressRule{},
				},
			},
			ingressClasses:  map[string]netv1.IngressClass{"name1": {ObjectMeta: metav1.ObjectMeta{Name: "name1"}}, "name2": {ObjectMeta: metav1.ObjectMeta{Name: "name2", Annotations: map[string]string{}}}},
			expectedToFail:  true,
			expectedMessage: "Ingress INGRESSNAME declares ingress class name name3 but there is no such ingress class in the cluster",
		},
		{
			it: "succeeds, if the ingress specifies an existing ingress class the legacy way",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME", Annotations: map[string]string{"kubernetes.io/ingress.class": existingIngressClassName}},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{},
				},
			},
			ingressClasses:  map[string]netv1.IngressClass{"name1": {ObjectMeta: metav1.ObjectMeta{Name: "name1"}}, "name2": {ObjectMeta: metav1.ObjectMeta{Name: "name2", Annotations: map[string]string{}}}},
			expectedToFail:  false,
			expectedMessage: "Ingress INGRESSNAME has a valid ingress class",
		},
		{
			it: "fails, if the ingress specifies an non-existing ingress class the legacy way",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME", Annotations: map[string]string{"kubernetes.io/ingress.class": nonExistingIngressClassName}},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{},
				},
			},
			ingressClasses:  map[string]netv1.IngressClass{"name1": {ObjectMeta: metav1.ObjectMeta{Name: "name1"}}, "name2": {ObjectMeta: metav1.ObjectMeta{Name: "name2", Annotations: map[string]string{}}}},
			expectedToFail:  true,
			expectedMessage: "Ingress INGRESSNAME declares ingress class name name3 but there is no such ingress class in the cluster",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkHasValidIngressClass(tc.ingress, tc.ingressClasses)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
