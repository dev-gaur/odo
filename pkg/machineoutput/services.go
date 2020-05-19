package machineoutput

import (
	"github.com/openshift/odo/pkg/catalog"
	olm "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type CatalogListOutput struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Services  metav1.List `json:"services,omitempty"`
	// list of clusterserviceversions (installed by Operators)
	Operators *olm.ClusterServiceVersionList `json:"operators,omitempty"`
}

func NewCatalogListOutput(services *[]catalog.ServiceType, operators *olm.ClusterServiceVersionList) CatalogListOutput {
	var objects []runtime.RawExtension
	for _, service := range *services {
		objects = append(objects, runtime.RawExtension{Object: &service})
	}

	return CatalogListOutput{
		TypeMeta: metav1.TypeMeta{
			Kind: "CatalogListOutput",
		},
		Services:  metav1.List{
			TypeMeta : metav1.TypeMeta{
				Kind:       "List",
				APIVersion: catalog.ApiVersion,
			},
			Items: objects,
		},
		Operators: operators,
	}
}