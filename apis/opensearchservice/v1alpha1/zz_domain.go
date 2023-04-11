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

// DomainParameters defines the desired state of Domain
type DomainParameters struct {
	// Region is which region the Domain will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Identity and Access Management (IAM) policy document specifying the access
	// policies for the new domain.
	AccessPolicies *string `json:"accessPolicies,omitempty"`
	// Key-value pairs to specify advanced configuration options. The following
	// key-value pairs are supported:
	//
	//    * "rest.action.multi.allow_explicit_index": "true" | "false" - Note the
	//    use of a string rather than a boolean. Specifies whether explicit references
	//    to indexes are allowed inside the body of HTTP requests. If you want to
	//    configure access policies for domain sub-resources, such as specific indexes
	//    and domain APIs, you must disable this property. Default is true.
	//
	//    * "indices.fielddata.cache.size": "80" - Note the use of a string rather
	//    than a boolean. Specifies the percentage of heap space allocated to field
	//    data. Default is unbounded.
	//
	//    * "indices.query.bool.max_clause_count": "1024" - Note the use of a string
	//    rather than a boolean. Specifies the maximum number of clauses allowed
	//    in a Lucene boolean query. Default is 1,024. Queries with more than the
	//    permitted number of clauses result in a TooManyClauses error.
	//
	//    * "override_main_response_version": "true" | "false" - Note the use of
	//    a string rather than a boolean. Specifies whether the domain reports its
	//    version as 7.10 to allow Elasticsearch OSS clients and plugins to continue
	//    working with it. Default is false when creating a domain and true when
	//    upgrading a domain.
	//
	// For more information, see Advanced cluster parameters (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options).
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty"`
	// Options for fine-grained access control.
	AdvancedSecurityOptions *AdvancedSecurityOptionsInput `json:"advancedSecurityOptions,omitempty"`
	// Options for Auto-Tune.
	AutoTuneOptions *AutoTuneOptionsInput `json:"autoTuneOptions,omitempty"`
	// Container for the cluster configuration of a domain.
	ClusterConfig *ClusterConfig `json:"clusterConfig,omitempty"`
	// Key-value pairs to configure Amazon Cognito authentication. For more information,
	// see Configuring Amazon Cognito authentication for OpenSearch Dashboards (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).
	CognitoOptions *CognitoOptions `json:"cognitoOptions,omitempty"`
	// Additional options for the domain endpoint, such as whether to require HTTPS
	// for all traffic.
	DomainEndpointOptions *DomainEndpointOptions `json:"domainEndpointOptions,omitempty"`
	// Container for the parameters required to enable EBS-based storage for an
	// OpenSearch Service domain.
	EBSOptions *EBSOptions `json:"ebsOptions,omitempty"`
	// String of format Elasticsearch_X.Y or OpenSearch_X.Y to specify the engine
	// version for the OpenSearch Service domain. For example, OpenSearch_1.0 or
	// Elasticsearch_7.9. For more information, see Creating and managing Amazon
	// OpenSearch Service domains (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).
	EngineVersion *string `json:"engineVersion,omitempty"`
	// Key-value pairs to configure slow log publishing.
	LogPublishingOptions map[string]*LogPublishingOption `json:"logPublishingOptions,omitempty"`
	// Enables node-to-node encryption.
	NodeToNodeEncryptionOptions *NodeToNodeEncryptionOptions `json:"nodeToNodeEncryptionOptions,omitempty"`
	// List of tags to add to the domain upon creation.
	Tags                   []*Tag `json:"tags,omitempty"`
	CustomDomainParameters `json:",inline"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DomainParameters `json:"forProvider"`
}

// DomainObservation defines the observed state of Domain
type DomainObservation struct {
	// The Amazon Resource Name (ARN) of the domain. For more information, see IAM
	// identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html)
	// in the AWS Identity and Access Management User Guide.
	ARN *string `json:"arn,omitempty"`
	// Identity and Access Management (IAM) policy document specifying the access
	// policies for the domain.
	AccessPolicies *string `json:"accessPolicies,omitempty"`
	// Key-value pairs that specify advanced configuration options.
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty"`
	// Settings for fine-grained access control.
	AdvancedSecurityOptions *AdvancedSecurityOptions `json:"advancedSecurityOptions,omitempty"`
	// Auto-Tune settings for the domain.
	AutoTuneOptions *AutoTuneOptionsOutput `json:"autoTuneOptions,omitempty"`
	// Information about a configuration change happening on the domain.
	ChangeProgressDetails *ChangeProgressDetails `json:"changeProgressDetails,omitempty"`
	// Container for the cluster configuration of the domain.
	ClusterConfig *ClusterConfig `json:"clusterConfig,omitempty"`
	// Key-value pairs to configure Amazon Cognito authentication for OpenSearch
	// Dashboards.
	CognitoOptions *CognitoOptions `json:"cognitoOptions,omitempty"`
	// Creation status of an OpenSearch Service domain. True if domain creation
	// is complete. False if domain creation is still in progress.
	Created *bool `json:"created,omitempty"`
	// Deletion status of an OpenSearch Service domain. True if domain deletion
	// is complete. False if domain deletion is still in progress. Once deletion
	// is complete, the status of the domain is no longer returned.
	Deleted *bool `json:"deleted,omitempty"`
	// Additional options for the domain endpoint, such as whether to require HTTPS
	// for all traffic.
	DomainEndpointOptions *DomainEndpointOptions `json:"domainEndpointOptions,omitempty"`
	// Unique identifier for the domain.
	DomainID *string `json:"domainID,omitempty"`
	// Name of the domain. Domain names are unique across all domains owned by the
	// same account within an Amazon Web Services Region.
	DomainName *string `json:"domainName,omitempty"`
	// Encryption at rest settings for the domain.
	EncryptionAtRestOptions *EncryptionAtRestOptions `json:"encryptionAtRestOptions,omitempty"`
	// Domain-specific endpoint used to submit index, search, and data upload requests
	// to the domain.
	Endpoint *string `json:"endpoint,omitempty"`
	// The key-value pair that exists if the OpenSearch Service domain uses VPC
	// endpoints.. Example key, value: 'vpc','vpc-endpoint-h2dsd34efgyghrtguk5gt6j2foh4.us-east-1.es.amazonaws.com'.
	Endpoints map[string]*string `json:"endpoints,omitempty"`
	// Version of OpenSearch or Elasticsearch that the domain is running, in the
	// format Elasticsearch_X.Y or OpenSearch_X.Y.
	EngineVersion *string `json:"engineVersion,omitempty"`
	// Whether node-to-node encryption is enabled or disabled.
	NodeToNodeEncryptionOptions *NodeToNodeEncryptionOptions `json:"nodeToNodeEncryptionOptions,omitempty"`
	// The status of the domain configuration. True if OpenSearch Service is processing
	// configuration changes. False if the configuration is active.
	Processing *bool `json:"processing,omitempty"`
	// The current status of the domain's service software.
	ServiceSoftwareOptions *ServiceSoftwareOptions `json:"serviceSoftwareOptions,omitempty"`
	// DEPRECATED. Container for parameters required to configure automated snapshots
	// of domain indexes.
	SnapshotOptions *SnapshotOptions `json:"snapshotOptions,omitempty"`
	// The status of a domain version upgrade to a new version of OpenSearch or
	// Elasticsearch. True if OpenSearch Service is in the process of a version
	// upgrade. False if the configuration is active.
	UpgradeProcessing *bool `json:"upgradeProcessing,omitempty"`
	// The VPC configuration for the domain.
	VPCOptions *VPCDerivedInfo `json:"vpcOptions,omitempty"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	DomainKind             = "Domain"
	DomainGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainKind}.String()
	DomainKindAPIVersion   = DomainKind + "." + GroupVersion.String()
	DomainGroupVersionKind = GroupVersion.WithKind(DomainKind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
