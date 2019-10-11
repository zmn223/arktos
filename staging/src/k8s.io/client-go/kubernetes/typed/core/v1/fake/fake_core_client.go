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
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCoreV1 struct {
	*testing.Fake
}

func (c *FakeCoreV1) Actions(namespace string, optional_tenant ...string) v1.ActionInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeActions{c, namespace, tenant}
}

func (c *FakeCoreV1) ComponentStatuses() v1.ComponentStatusInterface {

	return &FakeComponentStatuses{c}
}

func (c *FakeCoreV1) ConfigMaps(namespace string, optional_tenant ...string) v1.ConfigMapInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeConfigMaps{c, namespace, tenant}
}

func (c *FakeCoreV1) ControllerInstances() v1.ControllerInstanceInterface {

	return &FakeControllerInstances{c}
}

func (c *FakeCoreV1) Endpoints(namespace string, optional_tenant ...string) v1.EndpointsInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeEndpoints{c, namespace, tenant}
}

func (c *FakeCoreV1) Events(namespace string, optional_tenant ...string) v1.EventInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeEvents{c, namespace, tenant}
}

func (c *FakeCoreV1) LimitRanges(namespace string, optional_tenant ...string) v1.LimitRangeInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeLimitRanges{c, namespace, tenant}
}

func (c *FakeCoreV1) Namespaces(optional_tenant ...string) v1.NamespaceInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeNamespaces{c, tenant}
}

func (c *FakeCoreV1) Nodes() v1.NodeInterface {

	return &FakeNodes{c}
}

func (c *FakeCoreV1) PersistentVolumes(optional_tenant ...string) v1.PersistentVolumeInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakePersistentVolumes{c, tenant}
}

func (c *FakeCoreV1) PersistentVolumeClaims(namespace string, optional_tenant ...string) v1.PersistentVolumeClaimInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakePersistentVolumeClaims{c, namespace, tenant}
}

func (c *FakeCoreV1) Pods(namespace string, optional_tenant ...string) v1.PodInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakePods{c, namespace, tenant}
}

func (c *FakeCoreV1) PodTemplates(namespace string, optional_tenant ...string) v1.PodTemplateInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakePodTemplates{c, namespace, tenant}
}

func (c *FakeCoreV1) ReplicationControllers(namespace string, optional_tenant ...string) v1.ReplicationControllerInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeReplicationControllers{c, namespace, tenant}
}

func (c *FakeCoreV1) ResourceQuotas(namespace string, optional_tenant ...string) v1.ResourceQuotaInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeResourceQuotas{c, namespace, tenant}
}

func (c *FakeCoreV1) Secrets(namespace string, optional_tenant ...string) v1.SecretInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeSecrets{c, namespace, tenant}
}

func (c *FakeCoreV1) Services(namespace string, optional_tenant ...string) v1.ServiceInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeServices{c, namespace, tenant}
}

func (c *FakeCoreV1) ServiceAccounts(namespace string, optional_tenant ...string) v1.ServiceAccountInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeServiceAccounts{c, namespace, tenant}
}

func (c *FakeCoreV1) Tenants() v1.TenantInterface {

	return &FakeTenants{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
