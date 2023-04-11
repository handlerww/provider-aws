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

// RouteParameters defines the desired state of Route
type RouteParameters struct {
	// Region is which region the Route will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The ID of the carrier gateway.
	//
	// You can only use this option when the VPC contains a subnet which is associated
	// with a Wavelength Zone.
	CarrierGatewayID *string `json:"carrierGatewayID,omitempty"`
	// The Amazon Resource Name (ARN) of the core network.
	CoreNetworkARN *string `json:"coreNetworkARN,omitempty"`
	// The IPv4 CIDR address block used for the destination match. Routing decisions
	// are based on the most specific match. We modify the specified CIDR block
	// to its canonical form; for example, if you specify 100.68.0.18/18, we modify
	// it to 100.68.0.0/18.
	DestinationCIDRBlock *string `json:"destinationCIDRBlock,omitempty"`
	// The IPv6 CIDR block used for the destination match. Routing decisions are
	// based on the most specific match.
	DestinationIPv6CIDRBlock *string `json:"destinationIPv6CIDRBlock,omitempty"`
	// The ID of a prefix list used for the destination match.
	DestinationPrefixListID *string `json:"destinationPrefixListID,omitempty"`
	// [IPv6 traffic only] The ID of an egress-only internet gateway.
	EgressOnlyInternetGatewayID *string `json:"egressOnlyInternetGatewayID,omitempty"`
	// The ID of the local gateway.
	LocalGatewayID *string `json:"localGatewayID,omitempty"`
	// The ID of a network interface.
	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty"`
	// The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.
	VPCEndpointID         *string `json:"vpcEndpointID,omitempty"`
	CustomRouteParameters `json:",inline"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RouteParameters `json:"forProvider"`
}

// RouteObservation defines the observed state of Route
type RouteObservation struct {
	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `json:"return_,omitempty"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	RouteKind             = "Route"
	RouteGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteKind}.String()
	RouteKindAPIVersion   = RouteKind + "." + GroupVersion.String()
	RouteGroupVersionKind = GroupVersion.WithKind(RouteKind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
