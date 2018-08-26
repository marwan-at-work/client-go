/*
Copyright The Kubernetes Authors.

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

package fake

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/v8/testing"
)

// FakeNamespaces implements NamespaceInterface
type FakeNamespaces struct {
	Fake *FakeCoreV1
}

var namespacesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "namespaces"}

var namespacesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}

// Get takes name of the namespace, and returns the corresponding namespace object, and an error if there is any.
func (c *FakeNamespaces) Get(name string, options v1.GetOptions) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(namespacesResource, name), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors.
func (c *FakeNamespaces) List(opts v1.ListOptions) (result *corev1.NamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(namespacesResource, namespacesKind, opts), &corev1.NamespaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.NamespaceList{ListMeta: obj.(*corev1.NamespaceList).ListMeta}
	for _, item := range obj.(*corev1.NamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaces.
func (c *FakeNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(namespacesResource, opts))
}

// Create takes the representation of a namespace and creates it.  Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Create(namespace *corev1.Namespace) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(namespacesResource, namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// Update takes the representation of a namespace and updates it. Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Update(namespace *corev1.Namespace) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(namespacesResource, namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNamespaces) UpdateStatus(namespace *corev1.Namespace) (*corev1.Namespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(namespacesResource, "status", namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// Delete takes name of the namespace and deletes it. Returns an error if one occurs.
func (c *FakeNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(namespacesResource, name), &corev1.Namespace{})
	return err
}

// Patch applies the patch and returns the patched namespace.
func (c *FakeNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(namespacesResource, name, data, subresources...), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}
