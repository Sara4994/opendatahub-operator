/*
Copyright 2023.

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

// Package status contains different conditions, phases and progresses,
// being used by DataScienceCluster and DSCInitialization's controller
package status

import (
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	"github.com/operator-framework/api/pkg/lib/version"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// These constants represent the overall Phase as used by .Status.Phase.
const (
	// PhaseIgnored is used when a resource is ignored
	// is an example of a constant that is not used anywhere in the code.
	PhaseIgnored = "Ignored"
	// PhaseNotReady is used when waiting for system to be ready after reconcile is successful
	// is an example of a constant that is not used anywhere in the code.
	PhaseNotReady = "Not Ready"
	// PhaseClusterExpanding is used when cluster is expanding capacity
	// is an example of a constant that is not used anywhere in the code.
	PhaseClusterExpanding = "Expanding Capacity"
	// PhaseDeleting is used when cluster is deleting
	// is an example of a constant that is not used anywhere in the code.
	PhaseDeleting = "Deleting"
	// PhaseConnecting is used when cluster is connecting to external cluster
	// is an example of a constant that is not used anywhere in the code.
	PhaseConnecting = "Connecting"
	// PhaseOnboarding is used when consumer is Onboarding
	// is an example of a constant that is not used anywhere in the code.
	PhaseOnboarding = "Onboarding"

	// PhaseProgressing is used when SetProgressingCondition() is called.
	PhaseProgressing = "Progressing"
	// PhaseError is used when SetErrorCondition() is called.
	PhaseError = "Error"
	// PhaseReady is used when SetCompleteCondition is called.
	PhaseReady = "Ready"
)

// List of constants to show different reconciliation messages and statuses.
const (
	// ReconcileFailed is used when multiple DSCI instance exists or DSC reconcile failed/removal failed.
	ReconcileFailed                       = "ReconcileFailed"
	ReconcileInit                         = "ReconcileInit"
	ReconcileCompleted                    = "ReconcileCompleted"
	ReconcileCompletedWithComponentErrors = "ReconcileCompletedWithComponentErrors"
	ReconcileCompletedMessage             = "Reconcile completed successfully"

	// ConditionReconcileComplete represents extra Condition Type, used by .Condition.Type.
	ConditionReconcileComplete conditionsv1.ConditionType = "ReconcileComplete"
)

const (
	CapabilityServiceMesh              conditionsv1.ConditionType = "CapabilityServiceMesh"
	CapabilityServiceMeshAuthorization conditionsv1.ConditionType = "CapabilityServiceMeshAuthorization"
	CapabilityDSPv2Argo                conditionsv1.ConditionType = "CapabilityDSPv2Argo"
)

const (
	MissingOperatorReason string = "MissingOperator"
	ConfiguredReason      string = "Configured"
	RemovedReason         string = "Removed"
	CapabilityFailed      string = "CapabilityFailed"
	ArgoWorkflowExist     string = "ArgoWorkflowExist"
)

const (
	ReadySuffix = "Ready"
)

// SetProgressingCondition sets the ProgressingCondition to True and other conditions to false or
// Unknown. Used when we are just starting to reconcile, and there are no existing conditions.
func SetProgressingCondition(conditions *[]conditionsv1.Condition, reason string, message string) {
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    ConditionReconcileComplete,
		Status:  corev1.ConditionUnknown,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionAvailable,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionProgressing,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionDegraded,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionUpgradeable,
		Status:  corev1.ConditionUnknown,
		Reason:  reason,
		Message: message,
	})
}

// SetErrorCondition sets the ConditionReconcileComplete to False in case of any errors
// during the reconciliation process.
func SetErrorCondition(conditions *[]conditionsv1.Condition, reason string, message string) {
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    ConditionReconcileComplete,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionAvailable,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionProgressing,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionDegraded,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionUpgradeable,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
}

// SetCompleteCondition sets the ConditionReconcileComplete to True and other Conditions
// to indicate that the reconciliation process has completed successfully.
func SetCompleteCondition(conditions *[]conditionsv1.Condition, reason string, message string) {
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    ConditionReconcileComplete,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionAvailable,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionProgressing,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionDegraded,
		Status:  corev1.ConditionFalse,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionUpgradeable,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: message,
	})
	conditionsv1.RemoveStatusCondition(conditions, CapabilityDSPv2Argo)
}

// SetCondition is a general purpose function to update any type of condition.
func SetCondition(conditions *[]conditionsv1.Condition, conditionType string, reason string, message string, status corev1.ConditionStatus) {
	conditionsv1.SetStatusCondition(conditions, conditionsv1.Condition{
		Type:    conditionsv1.ConditionType(conditionType),
		Status:  status,
		Reason:  reason,
		Message: message,
	})
}

// SetComponentCondition appends Condition Type with const ReadySuffix for given component
// when component finished reconcile.
func SetComponentCondition(conditions *[]conditionsv1.Condition, component string, reason string, message string, status corev1.ConditionStatus) {
	SetCondition(conditions, component+ReadySuffix, reason, message, status)
}

// RemoveComponentCondition remove Condition of giving component.
func RemoveComponentCondition(conditions *[]conditionsv1.Condition, component string) {
	conditionsv1.RemoveStatusCondition(conditions, conditionsv1.ConditionType(component+ReadySuffix))
}

type Platform string

type ComponentObject interface {
	client.Object
	GetComponentsStatus() *ComponentsStatus
}

// Condition represents the state of the operator's
// reconciliation functionality.
// +k8s:deepcopy-gen=true
type ComponentReleaseStatus struct {
	Name        Platform                `json:"name,omitempty"`
	DisplayName string                  `json:"displayname,omitempty"`
	Version     version.OperatorVersion `json:"version,omitempty"`
	RepoURL     string                  `json:"repourl,omitempty"`
}

// +k8s:deepcopy-gen=true
type ComponentStatus struct {
	UpstreamRelease []ComponentReleaseStatus `json:"upstreamrelease,omitempty"`
}

// +k8s:deepcopy-gen=true
// DashboardStatus struct holds the status for the Dashboard component.
type DashboardStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// WorkbenchesStatus struct holds the status for the Workbenches component.
type WorkbenchesStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// ModelMeshServingStatus struct holds the status for the ModelMeshServing component.
type ModelMeshServingStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// DataSciencePipelinesStatus struct holds the status for the DataSciencePipelines component.
type DataSciencePipelinesStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// KserveStatus struct holds the status for the Kserve component.
type KserveStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// KueueStatus struct holds the status for the Kueue component.
type KueueStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// CodeFlareStatus struct holds the status for the CodeFlare component.
type CodeFlareStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// RayStatus struct holds the status for the Ray component.
type RayStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// TrustyAIStatus struct holds the status for the TrustyAIStatus component.
type TrustyAIStatus struct {
	ComponentStatus `json:",inline"`
}

// +k8s:deepcopy-gen=true
// ModelRegistryStatus struct holds the status for the ModelRegistry component.
type ModelRegistryStatus struct {
	ComponentStatus     `json:",inline"`
	RegistriesNamespace string `json:"registriesNamespace,omitempty"`
}

// +k8s:deepcopy-gen=true
// TrainingOperatorStatus struct holds the status for the TrainingOperator component.
type TrainingOperatorStatus struct {
	ComponentStatus `json:",inline"`
}

// ComponentsStatus defines the custom status of DataScienceCluster components.
// +k8s:deepcopy-gen=true
type ComponentsStatus struct {
	Dashboard            *DashboardStatus            `json:"dashboard,omitempty"`
	Workbenches          *WorkbenchesStatus          `json:"workbenches,omitempty"`
	ModelMeshServing     *ModelMeshServingStatus     `json:"modelmeshserving,omitempty"`
	DataSciencePipelines *DataSciencePipelinesStatus `json:"datasciencepipelines,omitempty"`
	Kserve               *KserveStatus               `json:"kserve,omitempty"`
	Kueue                *KueueStatus                `json:"kueue,omitempty"`
	CodeFlare            *CodeFlareStatus            `json:"codeflare,omitempty"`
	Ray                  *RayStatus                  `json:"ray,omitempty"`
	TrustyAI             *TrustyAIStatus             `json:"trustyai,omitempty"`
	ModelRegistry        *ModelRegistryStatus        `json:"modelregistry,omitempty"`
	TrainingOperator     *TrainingOperatorStatus     `json:"trainingoperator,omitempty"`
}
