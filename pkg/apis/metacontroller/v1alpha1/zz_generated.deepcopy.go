//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnotationSelector) DeepCopyInto(out *AnnotationSelector) {
	*out = *in
	if in.MatchAnnotations != nil {
		in, out := &in.MatchAnnotations, &out.MatchAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]v1.LabelSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnotationSelector.
func (in *AnnotationSelector) DeepCopy() *AnnotationSelector {
	if in == nil {
		return nil
	}
	out := new(AnnotationSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChildUpdateStatusChecks) DeepCopyInto(out *ChildUpdateStatusChecks) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StatusConditionCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChildUpdateStatusChecks.
func (in *ChildUpdateStatusChecks) DeepCopy() *ChildUpdateStatusChecks {
	if in == nil {
		return nil
	}
	out := new(ChildUpdateStatusChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeController) DeepCopyInto(out *CompositeController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeController.
func (in *CompositeController) DeepCopy() *CompositeController {
	if in == nil {
		return nil
	}
	out := new(CompositeController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositeController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerChildResourceRule) DeepCopyInto(out *CompositeControllerChildResourceRule) {
	*out = *in
	out.ResourceRule = in.ResourceRule
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(CompositeControllerChildUpdateStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerChildResourceRule.
func (in *CompositeControllerChildResourceRule) DeepCopy() *CompositeControllerChildResourceRule {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerChildResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerChildUpdateStrategy) DeepCopyInto(out *CompositeControllerChildUpdateStrategy) {
	*out = *in
	in.StatusChecks.DeepCopyInto(&out.StatusChecks)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerChildUpdateStrategy.
func (in *CompositeControllerChildUpdateStrategy) DeepCopy() *CompositeControllerChildUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerChildUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerHooks) DeepCopyInto(out *CompositeControllerHooks) {
	*out = *in
	if in.Customize != nil {
		in, out := &in.Customize, &out.Customize
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	if in.Finalize != nil {
		in, out := &in.Finalize, &out.Finalize
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	if in.PreUpdateChild != nil {
		in, out := &in.PreUpdateChild, &out.PreUpdateChild
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	if in.PostUpdateChild != nil {
		in, out := &in.PostUpdateChild, &out.PostUpdateChild
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerHooks.
func (in *CompositeControllerHooks) DeepCopy() *CompositeControllerHooks {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerHooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerList) DeepCopyInto(out *CompositeControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CompositeController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerList.
func (in *CompositeControllerList) DeepCopy() *CompositeControllerList {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositeControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerParentResourceRule) DeepCopyInto(out *CompositeControllerParentResourceRule) {
	*out = *in
	out.ResourceRule = in.ResourceRule
	if in.RevisionHistory != nil {
		in, out := &in.RevisionHistory, &out.RevisionHistory
		*out = new(CompositeControllerRevisionHistory)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerParentResourceRule.
func (in *CompositeControllerParentResourceRule) DeepCopy() *CompositeControllerParentResourceRule {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerParentResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerRevisionHistory) DeepCopyInto(out *CompositeControllerRevisionHistory) {
	*out = *in
	if in.FieldPaths != nil {
		in, out := &in.FieldPaths, &out.FieldPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerRevisionHistory.
func (in *CompositeControllerRevisionHistory) DeepCopy() *CompositeControllerRevisionHistory {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerRevisionHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerSpec) DeepCopyInto(out *CompositeControllerSpec) {
	*out = *in
	in.ParentResource.DeepCopyInto(&out.ParentResource)
	if in.ChildResources != nil {
		in, out := &in.ChildResources, &out.ChildResources
		*out = make([]CompositeControllerChildResourceRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = new(CompositeControllerHooks)
		(*in).DeepCopyInto(*out)
	}
	if in.ResyncPeriodSeconds != nil {
		in, out := &in.ResyncPeriodSeconds, &out.ResyncPeriodSeconds
		*out = new(int32)
		**out = **in
	}
	if in.GenerateSelector != nil {
		in, out := &in.GenerateSelector, &out.GenerateSelector
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerSpec.
func (in *CompositeControllerSpec) DeepCopy() *CompositeControllerSpec {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerStatus) DeepCopyInto(out *CompositeControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerStatus.
func (in *CompositeControllerStatus) DeepCopy() *CompositeControllerStatus {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevision) DeepCopyInto(out *ControllerRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ParentPatch.DeepCopyInto(&out.ParentPatch)
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]ControllerRevisionChildren, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevision.
func (in *ControllerRevision) DeepCopy() *ControllerRevision {
	if in == nil {
		return nil
	}
	out := new(ControllerRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevisionChildren) DeepCopyInto(out *ControllerRevisionChildren) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevisionChildren.
func (in *ControllerRevisionChildren) DeepCopy() *ControllerRevisionChildren {
	if in == nil {
		return nil
	}
	out := new(ControllerRevisionChildren)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevisionList) DeepCopyInto(out *ControllerRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevisionList.
func (in *ControllerRevisionList) DeepCopy() *ControllerRevisionList {
	if in == nil {
		return nil
	}
	out := new(ControllerRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorController) DeepCopyInto(out *DecoratorController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorController.
func (in *DecoratorController) DeepCopy() *DecoratorController {
	if in == nil {
		return nil
	}
	out := new(DecoratorController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DecoratorController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorControllerAttachmentRule) DeepCopyInto(out *DecoratorControllerAttachmentRule) {
	*out = *in
	out.ResourceRule = in.ResourceRule
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(DecoratorControllerAttachmentUpdateStrategy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorControllerAttachmentRule.
func (in *DecoratorControllerAttachmentRule) DeepCopy() *DecoratorControllerAttachmentRule {
	if in == nil {
		return nil
	}
	out := new(DecoratorControllerAttachmentRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorControllerAttachmentUpdateStrategy) DeepCopyInto(out *DecoratorControllerAttachmentUpdateStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorControllerAttachmentUpdateStrategy.
func (in *DecoratorControllerAttachmentUpdateStrategy) DeepCopy() *DecoratorControllerAttachmentUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(DecoratorControllerAttachmentUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorControllerHooks) DeepCopyInto(out *DecoratorControllerHooks) {
	*out = *in
	if in.Customize != nil {
		in, out := &in.Customize, &out.Customize
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	if in.Finalize != nil {
		in, out := &in.Finalize, &out.Finalize
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorControllerHooks.
func (in *DecoratorControllerHooks) DeepCopy() *DecoratorControllerHooks {
	if in == nil {
		return nil
	}
	out := new(DecoratorControllerHooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorControllerList) DeepCopyInto(out *DecoratorControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DecoratorController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorControllerList.
func (in *DecoratorControllerList) DeepCopy() *DecoratorControllerList {
	if in == nil {
		return nil
	}
	out := new(DecoratorControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DecoratorControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorControllerResourceRule) DeepCopyInto(out *DecoratorControllerResourceRule) {
	*out = *in
	out.ResourceRule = in.ResourceRule
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AnnotationSelector != nil {
		in, out := &in.AnnotationSelector, &out.AnnotationSelector
		*out = new(AnnotationSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorControllerResourceRule.
func (in *DecoratorControllerResourceRule) DeepCopy() *DecoratorControllerResourceRule {
	if in == nil {
		return nil
	}
	out := new(DecoratorControllerResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorControllerSpec) DeepCopyInto(out *DecoratorControllerSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]DecoratorControllerResourceRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Attachments != nil {
		in, out := &in.Attachments, &out.Attachments
		*out = make([]DecoratorControllerAttachmentRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = new(DecoratorControllerHooks)
		(*in).DeepCopyInto(*out)
	}
	if in.ResyncPeriodSeconds != nil {
		in, out := &in.ResyncPeriodSeconds, &out.ResyncPeriodSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorControllerSpec.
func (in *DecoratorControllerSpec) DeepCopy() *DecoratorControllerSpec {
	if in == nil {
		return nil
	}
	out := new(DecoratorControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecoratorControllerStatus) DeepCopyInto(out *DecoratorControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecoratorControllerStatus.
func (in *DecoratorControllerStatus) DeepCopy() *DecoratorControllerStatus {
	if in == nil {
		return nil
	}
	out := new(DecoratorControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hook) DeepCopyInto(out *Hook) {
	*out = *in
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(Webhook)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hook.
func (in *Hook) DeepCopy() *Hook {
	if in == nil {
		return nil
	}
	out := new(Hook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelatedResourceRule) DeepCopyInto(out *RelatedResourceRule) {
	*out = *in
	out.ResourceRule = in.ResourceRule
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelatedResourceRule.
func (in *RelatedResourceRule) DeepCopy() *RelatedResourceRule {
	if in == nil {
		return nil
	}
	out := new(RelatedResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRule) DeepCopyInto(out *ResourceRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRule.
func (in *ResourceRule) DeepCopy() *ResourceRule {
	if in == nil {
		return nil
	}
	out := new(ResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceReference) DeepCopyInto(out *ServiceReference) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceReference.
func (in *ServiceReference) DeepCopy() *ServiceReference {
	if in == nil {
		return nil
	}
	out := new(ServiceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusConditionCheck) DeepCopyInto(out *StatusConditionCheck) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConditionCheck.
func (in *StatusConditionCheck) DeepCopy() *StatusConditionCheck {
	if in == nil {
		return nil
	}
	out := new(StatusConditionCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceReference)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhook.
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}
