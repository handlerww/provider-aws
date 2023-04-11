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

// ConnectionParameters defines the desired state of Connection
type ConnectionParameters struct {
	// Region is which region the Connection will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The ID of the Data Catalog in which to create the connection. If none is
	// provided, the Amazon Web Services account ID is used by default.
	CatalogID *string `json:"catalogID,omitempty"`
	// The tags you assign to the connection.
	Tags                       map[string]*string `json:"tags,omitempty"`
	CustomConnectionParameters `json:",inline"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConnectionParameters `json:"forProvider"`
}

// ConnectionObservation defines the observed state of Connection
type ConnectionObservation struct {
	// The time that this connection definition was created.
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// The user, group, or role that last updated this connection definition.
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// The last time that this connection definition was updated.
	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	ConnectionKind             = "Connection"
	ConnectionGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionKind}.String()
	ConnectionKindAPIVersion   = ConnectionKind + "." + GroupVersion.String()
	ConnectionGroupVersionKind = GroupVersion.WithKind(ConnectionKind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
