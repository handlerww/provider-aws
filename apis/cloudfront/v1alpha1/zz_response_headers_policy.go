/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ResponseHeadersPolicyParameters defines the desired state of ResponseHeadersPolicy
type ResponseHeadersPolicyParameters struct {
	// Region is which region the ResponseHeadersPolicy will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Contains metadata about the response headers policy, and a set of configurations
	// that specify the HTTP headers.
	// +kubebuilder:validation:Required
	ResponseHeadersPolicyConfig           *ResponseHeadersPolicyConfig `json:"responseHeadersPolicyConfig"`
	CustomResponseHeadersPolicyParameters `json:",inline"`
}

// ResponseHeadersPolicySpec defines the desired state of ResponseHeadersPolicy
type ResponseHeadersPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ResponseHeadersPolicyParameters `json:"forProvider"`
}

// ResponseHeadersPolicyObservation defines the observed state of ResponseHeadersPolicy
type ResponseHeadersPolicyObservation struct {
	// The version identifier for the current version of the response headers policy.
	ETag *string `json:"eTag,omitempty"`
	// The URL of the response headers policy.
	Location *string `json:"location,omitempty"`
	// Contains a response headers policy.
	ResponseHeadersPolicy *ResponseHeadersPolicy_SDK `json:"responseHeadersPolicy,omitempty"`
}

// ResponseHeadersPolicyStatus defines the observed state of ResponseHeadersPolicy.
type ResponseHeadersPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ResponseHeadersPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicies API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResponseHeadersPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResponseHeadersPolicySpec   `json:"spec"`
	Status            ResponseHeadersPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResponseHeadersPolicyList contains a list of ResponseHeadersPolicies
type ResponseHeadersPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResponseHeadersPolicy `json:"items"`
}

// Repository type metadata.
var (
	ResponseHeadersPolicyKind             = "ResponseHeadersPolicy"
	ResponseHeadersPolicyGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResponseHeadersPolicyKind}.String()
	ResponseHeadersPolicyKindAPIVersion   = ResponseHeadersPolicyKind + "." + GroupVersion.String()
	ResponseHeadersPolicyGroupVersionKind = GroupVersion.WithKind(ResponseHeadersPolicyKind)
)

func init() {
	SchemeBuilder.Register(&ResponseHeadersPolicy{}, &ResponseHeadersPolicyList{})
}
