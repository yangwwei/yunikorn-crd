/*

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PartitionSpec defines the desired state of Partition
type PartitionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Partition. Edit Partition_types.go to remove/update
	Queues []Queue
	NodeSortPolicy NodeSortingPolicy
	PlacementRules []PlacementRule           `yaml:",omitempty" json:",omitempty"`
	Limits         []Limit                   `yaml:",omitempty" json:",omitempty"`
	Preemption     PartitionPreemptionConfig `yaml:",omitempty" json:",omitempty"`
}

// PartitionStatus defines the observed state of Partition
type PartitionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Partition is the Schema for the partitions API
type Partition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PartitionSpec   `json:"spec,omitempty"`
	Status PartitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PartitionList contains a list of Partition
type PartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Partition `json:"items"`
}

// The queue object for each queue:
// - the name of the queue
// - a resources object to specify resource limits on the queue
// - the maximum number of applications that can run in the queue
// - a set of properties, exact definition of what can be set is not part of the yaml
// - ACL for submit and or admin access
// - a list of sub or child queues
// - a list of users specifying limits on a queue
type Queue struct {
	Name            string
	Parent          bool              `yaml:",omitempty" json:",omitempty"`
	Resources       Resources         `yaml:",omitempty" json:",omitempty"`
	MaxApplications uint64            `yaml:",omitempty" json:",omitempty"`
	Properties      map[string]string `yaml:",omitempty" json:",omitempty"`
	AdminACL        string            `yaml:",omitempty" json:",omitempty"`
	SubmitACL       string            `yaml:",omitempty" json:",omitempty"`
	Queues          []Queue     `yaml:",omitempty" json:",omitempty"`
	Limits          []Limit           `yaml:",omitempty" json:",omitempty"`
}

// The resource limits to set on the queue. The definition allows for an unlimited number of types to be used.
// The mapping to "known" resources is not handled here.
// - guaranteed resources
// - max resources
type Resources struct {
	Guaranteed map[string]string `yaml:",omitempty" json:",omitempty"`
	Max        map[string]string `yaml:",omitempty" json:",omitempty"`
}

// The queue placement rule definition
// - the name of the rule
// - create flag: can the rule create a queue
// - user and group filter to be applied on the callers
// - rule link to allow setting a rule to generate the parent
// - value a generic value interpreted depending on the rule type (i.e queue name for the "fixed" rule
// or the application label name for the "tag" rule)
type PlacementRule struct {
	Name   string
	Create bool           `yaml:",omitempty" json:",omitempty"`
	Filter Filter         `yaml:",omitempty" json:",omitempty"`
	Parent *PlacementRule `yaml:",omitempty" json:",omitempty"`
	Value  string         `yaml:",omitempty" json:",omitempty"`
}

// The user and group filter for a rule.
// - type of filter (allow or deny filter, empty means allow)
// - list of users to filter (maybe empty)
// - list of groups to filter (maybe empty)
// if the list of users or groups is exactly 1 long it is interpreted as a regular expression
type Filter struct {
	Type   string
	Users  []string `yaml:",omitempty" json:",omitempty"`
	Groups []string `yaml:",omitempty" json:",omitempty"`
}

// A list of limit objects to define limits for a partition or queue
type Limits struct {
	Limit []Limit
}

// The limit object to specify user and or group limits at different levels in the partition or queues
// Different limits for the same user or group may be defined at different levels in the hierarchy
// - limit description (optional)
// - list of users (maybe empty)
// - list of groups (maybe empty)
// - maximum resources as a resource object to allow for the user or group
// - maximum number of applications the user or group can have running
type Limit struct {
	Limit           string
	Users           []string          `yaml:",omitempty" json:",omitempty"`
	Groups          []string          `yaml:",omitempty" json:",omitempty"`
	MaxResources    map[string]string `yaml:",omitempty" json:",omitempty"`
	MaxApplications uint64            `yaml:",omitempty" json:",omitempty"`
}

// Global Node Sorting Policy section
// - type: different type of policies supported (binpacking, fair etc)
type NodeSortingPolicy struct {
	Type string
}

type PartitionPreemptionConfig struct {
	Enabled bool
}

func init() {
	SchemeBuilder.Register(&Partition{}, &PartitionList{})
}
