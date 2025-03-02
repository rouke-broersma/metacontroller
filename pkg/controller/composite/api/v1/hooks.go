package v1

import (
	"metacontroller/pkg/apis/metacontroller/v1alpha1"
	v1 "metacontroller/pkg/controller/common/api/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// SyncHookRequest is the object sent as JSON to the sync and finalize hooks.
type SyncHookRequest struct {
	Controller *v1alpha1.CompositeController `json:"controller"`
	Parent     *unstructured.Unstructured    `json:"parent"`
	Children   v1.RelativeObjectMap          `json:"children"`
	Related    v1.RelativeObjectMap          `json:"related"`
	Finalizing bool                          `json:"finalizing"`
}

// SyncHookResponse is the expected format of the JSON response from the sync and finalize hooks.
type SyncHookResponse struct {
	Status   map[string]interface{}       `json:"status"`
	Children []*unstructured.Unstructured `json:"children"`

	ResyncAfterSeconds float64 `json:"resyncAfterSeconds"`

	// Finalized is only used by the finalize hook.
	Finalized bool `json:"finalized"`
}
