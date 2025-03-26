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

package v1beta2

import (
	http "net/http"

	resourcev1beta2 "k8s.io/api/resource/v1beta2"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type ResourceV1beta2Interface interface {
	RESTClient() rest.Interface
	DeviceClassesGetter
	ResourceClaimsGetter
	ResourceClaimTemplatesGetter
	ResourceSlicesGetter
}

// ResourceV1beta2Client is used to interact with features provided by the resource.k8s.io group.
type ResourceV1beta2Client struct {
	restClient rest.Interface
}

func (c *ResourceV1beta2Client) DeviceClasses() DeviceClassInterface {
	return newDeviceClasses(c)
}

func (c *ResourceV1beta2Client) ResourceClaims(namespace string) ResourceClaimInterface {
	return newResourceClaims(c, namespace)
}

func (c *ResourceV1beta2Client) ResourceClaimTemplates(namespace string) ResourceClaimTemplateInterface {
	return newResourceClaimTemplates(c, namespace)
}

func (c *ResourceV1beta2Client) ResourceSlices() ResourceSliceInterface {
	return newResourceSlices(c)
}

// NewForConfig creates a new ResourceV1beta2Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ResourceV1beta2Client, error) {
	config := *c
	setConfigDefaults(&config)
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ResourceV1beta2Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ResourceV1beta2Client, error) {
	config := *c
	setConfigDefaults(&config)
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ResourceV1beta2Client{client}, nil
}

// NewForConfigOrDie creates a new ResourceV1beta2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ResourceV1beta2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ResourceV1beta2Client for the given RESTClient.
func New(c rest.Interface) *ResourceV1beta2Client {
	return &ResourceV1beta2Client{c}
}

func setConfigDefaults(config *rest.Config) {
	gv := resourcev1beta2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ResourceV1beta2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
