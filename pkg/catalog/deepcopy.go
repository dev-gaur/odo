package catalog

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec)  {
	out.Hidden = in.Hidden
	out.PlanList = make([]string, len(in.PlanList))
	copy(out.PlanList, in.PlanList)
}

func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *ServiceType) DeepCopyInto(out *ServiceType) {
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.TypeMeta = metav1.TypeMeta{
		in.TypeMeta.APIVersion,
		in.TypeMeta.Kind,
	}
}

func (in *ServiceType) DeepCopy() *ServiceType {
	if in == nil {
		return nil
	}
	out := new(ServiceType)
	in.DeepCopyInto(out)
	return out
}

func (in ServiceType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec)  {
	out.ImageStreamRef = in.ImageStreamRef

	out.AllTags = make([]string, len(in.AllTags))
	copy(out.AllTags, in.AllTags)

	out.NonHiddenTags = make([]string, len(in.NonHiddenTags))
	copy(out.NonHiddenTags, in.NonHiddenTags)

	out.SupportedTags = make([]string, len(in.SupportedTags))
	copy(out.SupportedTags, in.SupportedTags)
}

func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *ComponentType) DeepCopyInto(out *ComponentType) {
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.TypeMeta = metav1.TypeMeta{
		in.TypeMeta.APIVersion,
		in.TypeMeta.Kind,
	}
}

func (in *ComponentType) DeepCopy() *ComponentType {
	if in == nil {
		return nil
	}
	out := new(ComponentType)
	in.DeepCopyInto(out)
	return out
}

func (in ComponentType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}