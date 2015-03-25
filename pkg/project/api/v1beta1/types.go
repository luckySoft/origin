package v1beta1

import (
	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta3"
)

// ProjectList is a list of Project objects.
type ProjectList struct {
	kapi.TypeMeta `json:",inline"`
	kapi.ListMeta `json:"metadata,omitempty"`
	Items         []Project `json:"items"`
}

// ProjectSpec describes the attributes on a Project
type ProjectSpec struct {
}

// ProjectStatus is information about the current status of a Project
type ProjectStatus struct {
	Phase kapi.NamespacePhase `json:"phase,omitempty" description:"phase is the current lifecycle phase of the project"`
}

// Project is a logical top-level container for a set of origin resources
type Project struct {
	kapi.TypeMeta   `json:",inline"`
	kapi.ObjectMeta `json:"metadata,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`

	// Spec defines the behavior of the Namespace.
	Spec ProjectSpec `json:"spec,omitempty" description:"spec defines the behavior of the Project"`

	// Status describes the current status of a Namespace
	Status ProjectStatus `json:"status,omitempty" description:"status describes the current status of a Project; read-only"`
}
