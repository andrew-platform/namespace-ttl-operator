/*
Copyright 2026.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// TemporaryNamespaceSpec defines the desired state of TemporaryNamespace
type TemporaryNamespaceSpec struct {
	TTL string `json:"ttl"`
	// +optional
	NamespaceName string `json:"namespaceName,omitempty"`
}

// TemporaryNamespaceStatus defines the observed state of TemporaryNamespace.
type TemporaryNamespaceStatus struct {
	// +optional
	CreatedAt metav1.Time `json:"createdAt,omitempty"`
	// +optional
	ExpiresAt metav1.Time `json:"expiresAt,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TemporaryNamespace is the Schema for the temporarynamespaces API
type TemporaryNamespace struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of TemporaryNamespace
	// +required
	Spec TemporaryNamespaceSpec `json:"spec"`

	// status defines the observed state of TemporaryNamespace
	// +optional
	Status TemporaryNamespaceStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// TemporaryNamespaceList contains a list of TemporaryNamespace
type TemporaryNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []TemporaryNamespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &TemporaryNamespace{}, &TemporaryNamespaceList{})
		return nil
	})
}
