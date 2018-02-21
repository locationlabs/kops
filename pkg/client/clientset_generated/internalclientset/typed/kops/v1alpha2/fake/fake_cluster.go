/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
)

// FakeClusters implements ClusterInterface
type FakeClusters struct {
	Fake *FakeKopsV1alpha2
	ns   string
}

var clustersResource = schema.GroupVersionResource{Group: "kops", Version: "v1alpha2", Resource: "clusters"}

var clustersKind = schema.GroupVersionKind{Group: "kops", Version: "v1alpha2", Kind: "Cluster"}

// Get takes name of the cluster, and returns the corresponding cluster object, and an error if there is any.
func (c *FakeClusters) Get(name string, options v1.GetOptions) (result *v1alpha2.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustersResource, c.ns, name), &v1alpha2.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Cluster), err
}

// List takes label and field selectors, and returns the list of Clusters that match those selectors.
func (c *FakeClusters) List(opts v1.ListOptions) (result *v1alpha2.ClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustersResource, clustersKind, c.ns, opts), &v1alpha2.ClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ClusterList{}
	for _, item := range obj.(*v1alpha2.ClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusters.
func (c *FakeClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustersResource, c.ns, opts))

}

// Create takes the representation of a cluster and creates it.  Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Create(cluster *v1alpha2.Cluster) (result *v1alpha2.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustersResource, c.ns, cluster), &v1alpha2.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Cluster), err
}

// Update takes the representation of a cluster and updates it. Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Update(cluster *v1alpha2.Cluster) (result *v1alpha2.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustersResource, c.ns, cluster), &v1alpha2.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Cluster), err
}

// Delete takes name of the cluster and deletes it. Returns an error if one occurs.
func (c *FakeClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustersResource, c.ns, name), &v1alpha2.Cluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.ClusterList{})
	return err
}

// Patch applies the patch and returns the patched cluster.
func (c *FakeClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustersResource, c.ns, name, data, subresources...), &v1alpha2.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Cluster), err
}
