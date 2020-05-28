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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/k8snetworkplumbingwg/macvlan-networkpolicy/pkg/apis/k8s.cni.cncf.io/v1"
	scheme "github.com/k8snetworkplumbingwg/macvlan-networkpolicy/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MacvlanNetworkPoliciesGetter has a method to return a MacvlanNetworkPolicyInterface.
// A group's client should implement this interface.
type MacvlanNetworkPoliciesGetter interface {
	MacvlanNetworkPolicies(namespace string) MacvlanNetworkPolicyInterface
}

// MacvlanNetworkPolicyInterface has methods to work with MacvlanNetworkPolicy resources.
type MacvlanNetworkPolicyInterface interface {
	Create(*v1.MacvlanNetworkPolicy) (*v1.MacvlanNetworkPolicy, error)
	Update(*v1.MacvlanNetworkPolicy) (*v1.MacvlanNetworkPolicy, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MacvlanNetworkPolicy, error)
	List(opts metav1.ListOptions) (*v1.MacvlanNetworkPolicyList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MacvlanNetworkPolicy, err error)
	MacvlanNetworkPolicyExpansion
}

// macvlanNetworkPolicies implements MacvlanNetworkPolicyInterface
type macvlanNetworkPolicies struct {
	client rest.Interface
	ns     string
}

// newMacvlanNetworkPolicies returns a MacvlanNetworkPolicies
func newMacvlanNetworkPolicies(c *K8sCniCncfIoV1Client, namespace string) *macvlanNetworkPolicies {
	return &macvlanNetworkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the macvlanNetworkPolicy, and returns the corresponding macvlanNetworkPolicy object, and an error if there is any.
func (c *macvlanNetworkPolicies) Get(name string, options metav1.GetOptions) (result *v1.MacvlanNetworkPolicy, err error) {
	result = &v1.MacvlanNetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MacvlanNetworkPolicies that match those selectors.
func (c *macvlanNetworkPolicies) List(opts metav1.ListOptions) (result *v1.MacvlanNetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MacvlanNetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested macvlanNetworkPolicies.
func (c *macvlanNetworkPolicies) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a macvlanNetworkPolicy and creates it.  Returns the server's representation of the macvlanNetworkPolicy, and an error, if there is any.
func (c *macvlanNetworkPolicies) Create(macvlanNetworkPolicy *v1.MacvlanNetworkPolicy) (result *v1.MacvlanNetworkPolicy, err error) {
	result = &v1.MacvlanNetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		Body(macvlanNetworkPolicy).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a macvlanNetworkPolicy and updates it. Returns the server's representation of the macvlanNetworkPolicy, and an error, if there is any.
func (c *macvlanNetworkPolicies) Update(macvlanNetworkPolicy *v1.MacvlanNetworkPolicy) (result *v1.MacvlanNetworkPolicy, err error) {
	result = &v1.MacvlanNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		Name(macvlanNetworkPolicy.Name).
		Body(macvlanNetworkPolicy).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the macvlanNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *macvlanNetworkPolicies) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *macvlanNetworkPolicies) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched macvlanNetworkPolicy.
func (c *macvlanNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MacvlanNetworkPolicy, err error) {
	result = &v1.MacvlanNetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("macvlan-networkpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}