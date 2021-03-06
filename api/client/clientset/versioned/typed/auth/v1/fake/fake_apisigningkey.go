/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	authv1 "tkestack.io/tke/api/auth/v1"
)

// FakeAPISigningKeys implements APISigningKeyInterface
type FakeAPISigningKeys struct {
	Fake *FakeAuthV1
}

var apisigningkeysResource = schema.GroupVersionResource{Group: "auth.tkestack.io", Version: "v1", Resource: "apisigningkeys"}

var apisigningkeysKind = schema.GroupVersionKind{Group: "auth.tkestack.io", Version: "v1", Kind: "APISigningKey"}

// Get takes name of the aPISigningKey, and returns the corresponding aPISigningKey object, and an error if there is any.
func (c *FakeAPISigningKeys) Get(name string, options v1.GetOptions) (result *authv1.APISigningKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apisigningkeysResource, name), &authv1.APISigningKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.APISigningKey), err
}

// List takes label and field selectors, and returns the list of APISigningKeys that match those selectors.
func (c *FakeAPISigningKeys) List(opts v1.ListOptions) (result *authv1.APISigningKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apisigningkeysResource, apisigningkeysKind, opts), &authv1.APISigningKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authv1.APISigningKeyList{ListMeta: obj.(*authv1.APISigningKeyList).ListMeta}
	for _, item := range obj.(*authv1.APISigningKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPISigningKeys.
func (c *FakeAPISigningKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apisigningkeysResource, opts))
}

// Create takes the representation of a aPISigningKey and creates it.  Returns the server's representation of the aPISigningKey, and an error, if there is any.
func (c *FakeAPISigningKeys) Create(aPISigningKey *authv1.APISigningKey) (result *authv1.APISigningKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apisigningkeysResource, aPISigningKey), &authv1.APISigningKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.APISigningKey), err
}

// Update takes the representation of a aPISigningKey and updates it. Returns the server's representation of the aPISigningKey, and an error, if there is any.
func (c *FakeAPISigningKeys) Update(aPISigningKey *authv1.APISigningKey) (result *authv1.APISigningKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apisigningkeysResource, aPISigningKey), &authv1.APISigningKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.APISigningKey), err
}

// Delete takes name of the aPISigningKey and deletes it. Returns an error if one occurs.
func (c *FakeAPISigningKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apisigningkeysResource, name), &authv1.APISigningKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPISigningKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apisigningkeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &authv1.APISigningKeyList{})
	return err
}

// Patch applies the patch and returns the patched aPISigningKey.
func (c *FakeAPISigningKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.APISigningKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apisigningkeysResource, name, pt, data, subresources...), &authv1.APISigningKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.APISigningKey), err
}
