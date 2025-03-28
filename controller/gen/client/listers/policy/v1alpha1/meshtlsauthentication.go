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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	policyv1alpha1 "github.com/linkerd/linkerd2/controller/gen/apis/policy/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// MeshTLSAuthenticationLister helps list MeshTLSAuthentications.
// All objects returned here must be treated as read-only.
type MeshTLSAuthenticationLister interface {
	// List lists all MeshTLSAuthentications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*policyv1alpha1.MeshTLSAuthentication, err error)
	// MeshTLSAuthentications returns an object that can list and get MeshTLSAuthentications.
	MeshTLSAuthentications(namespace string) MeshTLSAuthenticationNamespaceLister
	MeshTLSAuthenticationListerExpansion
}

// meshTLSAuthenticationLister implements the MeshTLSAuthenticationLister interface.
type meshTLSAuthenticationLister struct {
	listers.ResourceIndexer[*policyv1alpha1.MeshTLSAuthentication]
}

// NewMeshTLSAuthenticationLister returns a new MeshTLSAuthenticationLister.
func NewMeshTLSAuthenticationLister(indexer cache.Indexer) MeshTLSAuthenticationLister {
	return &meshTLSAuthenticationLister{listers.New[*policyv1alpha1.MeshTLSAuthentication](indexer, policyv1alpha1.Resource("meshtlsauthentication"))}
}

// MeshTLSAuthentications returns an object that can list and get MeshTLSAuthentications.
func (s *meshTLSAuthenticationLister) MeshTLSAuthentications(namespace string) MeshTLSAuthenticationNamespaceLister {
	return meshTLSAuthenticationNamespaceLister{listers.NewNamespaced[*policyv1alpha1.MeshTLSAuthentication](s.ResourceIndexer, namespace)}
}

// MeshTLSAuthenticationNamespaceLister helps list and get MeshTLSAuthentications.
// All objects returned here must be treated as read-only.
type MeshTLSAuthenticationNamespaceLister interface {
	// List lists all MeshTLSAuthentications in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*policyv1alpha1.MeshTLSAuthentication, err error)
	// Get retrieves the MeshTLSAuthentication from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*policyv1alpha1.MeshTLSAuthentication, error)
	MeshTLSAuthenticationNamespaceListerExpansion
}

// meshTLSAuthenticationNamespaceLister implements the MeshTLSAuthenticationNamespaceLister
// interface.
type meshTLSAuthenticationNamespaceLister struct {
	listers.ResourceIndexer[*policyv1alpha1.MeshTLSAuthentication]
}
