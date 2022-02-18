// +build !ignore_autogenerated

/*
Copyright 2020 The Kubernetes Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
	if in.Except != nil {
		in, out := &in.Except, &out.Except
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiNetworkPolicy) DeepCopyInto(out *MultiNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiNetworkPolicy.
func (in *MultiNetworkPolicy) DeepCopy() *MultiNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(MultiNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiNetworkPolicyEgressRule) DeepCopyInto(out *MultiNetworkPolicyEgressRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]MultiNetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]MultiNetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiNetworkPolicyEgressRule.
func (in *MultiNetworkPolicyEgressRule) DeepCopy() *MultiNetworkPolicyEgressRule {
	if in == nil {
		return nil
	}
	out := new(MultiNetworkPolicyEgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiNetworkPolicyIngressRule) DeepCopyInto(out *MultiNetworkPolicyIngressRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]MultiNetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]MultiNetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiNetworkPolicyIngressRule.
func (in *MultiNetworkPolicyIngressRule) DeepCopy() *MultiNetworkPolicyIngressRule {
	if in == nil {
		return nil
	}
	out := new(MultiNetworkPolicyIngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiNetworkPolicyList) DeepCopyInto(out *MultiNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiNetworkPolicyList.
func (in *MultiNetworkPolicyList) DeepCopy() *MultiNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(MultiNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiNetworkPolicyPeer) DeepCopyInto(out *MultiNetworkPolicyPeer) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IPBlock != nil {
		in, out := &in.IPBlock, &out.IPBlock
		*out = new(IPBlock)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiNetworkPolicyPeer.
func (in *MultiNetworkPolicyPeer) DeepCopy() *MultiNetworkPolicyPeer {
	if in == nil {
		return nil
	}
	out := new(MultiNetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiNetworkPolicyPort) DeepCopyInto(out *MultiNetworkPolicyPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(corev1.Protocol)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiNetworkPolicyPort.
func (in *MultiNetworkPolicyPort) DeepCopy() *MultiNetworkPolicyPort {
	if in == nil {
		return nil
	}
	out := new(MultiNetworkPolicyPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiNetworkPolicySpec) DeepCopyInto(out *MultiNetworkPolicySpec) {
	*out = *in
	in.PodSelector.DeepCopyInto(&out.PodSelector)
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]MultiNetworkPolicyIngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]MultiNetworkPolicyEgressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyTypes != nil {
		in, out := &in.PolicyTypes, &out.PolicyTypes
		*out = make([]MultiPolicyType, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiNetworkPolicySpec.
func (in *MultiNetworkPolicySpec) DeepCopy() *MultiNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(MultiNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}
