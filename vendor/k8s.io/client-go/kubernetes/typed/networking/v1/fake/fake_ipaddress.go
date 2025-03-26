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
	v1 "k8s.io/api/networking/v1"
	networkingv1 "k8s.io/client-go/applyconfigurations/networking/v1"
	gentype "k8s.io/client-go/gentype"
	typednetworkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

// fakeIPAddresses implements IPAddressInterface
type fakeIPAddresses struct {
	*gentype.FakeClientWithListAndApply[*v1.IPAddress, *v1.IPAddressList, *networkingv1.IPAddressApplyConfiguration]
	Fake *FakeNetworkingV1
}

func newFakeIPAddresses(fake *FakeNetworkingV1) typednetworkingv1.IPAddressInterface {
	return &fakeIPAddresses{
		gentype.NewFakeClientWithListAndApply[*v1.IPAddress, *v1.IPAddressList, *networkingv1.IPAddressApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("ipaddresses"),
			v1.SchemeGroupVersion.WithKind("IPAddress"),
			func() *v1.IPAddress { return &v1.IPAddress{} },
			func() *v1.IPAddressList { return &v1.IPAddressList{} },
			func(dst, src *v1.IPAddressList) { dst.ListMeta = src.ListMeta },
			func(list *v1.IPAddressList) []*v1.IPAddress { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.IPAddressList, items []*v1.IPAddress) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
