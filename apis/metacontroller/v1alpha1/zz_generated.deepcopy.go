// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConfig) DeepCopyInto(out *ClientConfig) {
	*out = *in
	out.Service = in.Service
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConfig.
func (in *ClientConfig) DeepCopy() *ClientConfig {
	if in == nil {
		return nil
	}
	out := new(ClientConfig)
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
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerHooks) DeepCopyInto(out *CompositeControllerHooks) {
	*out = *in
	out.Sync = in.Sync
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
	out.ListMeta = in.ListMeta
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
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeControllerSpec) DeepCopyInto(out *CompositeControllerSpec) {
	*out = *in
	out.ParentResource = in.ParentResource
	if in.ChildResources != nil {
		in, out := &in.ChildResources, &out.ChildResources
		*out = make([]ResourceRule, len(*in))
		copy(*out, *in)
	}
	out.ClientConfig = in.ClientConfig
	out.Hooks = in.Hooks
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
func (in *CompositeControllerSyncHook) DeepCopyInto(out *CompositeControllerSyncHook) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeControllerSyncHook.
func (in *CompositeControllerSyncHook) DeepCopy() *CompositeControllerSyncHook {
	if in == nil {
		return nil
	}
	out := new(CompositeControllerSyncHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerController) DeepCopyInto(out *InitializerController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerController.
func (in *InitializerController) DeepCopy() *InitializerController {
	if in == nil {
		return nil
	}
	out := new(InitializerController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InitializerController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerControllerHooks) DeepCopyInto(out *InitializerControllerHooks) {
	*out = *in
	out.Init = in.Init
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerControllerHooks.
func (in *InitializerControllerHooks) DeepCopy() *InitializerControllerHooks {
	if in == nil {
		return nil
	}
	out := new(InitializerControllerHooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerControllerInitHook) DeepCopyInto(out *InitializerControllerInitHook) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerControllerInitHook.
func (in *InitializerControllerInitHook) DeepCopy() *InitializerControllerInitHook {
	if in == nil {
		return nil
	}
	out := new(InitializerControllerInitHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerControllerList) DeepCopyInto(out *InitializerControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InitializerController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerControllerList.
func (in *InitializerControllerList) DeepCopy() *InitializerControllerList {
	if in == nil {
		return nil
	}
	out := new(InitializerControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InitializerControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerControllerSpec) DeepCopyInto(out *InitializerControllerSpec) {
	*out = *in
	if in.UninitializedResources != nil {
		in, out := &in.UninitializedResources, &out.UninitializedResources
		*out = make([]ResourceRule, len(*in))
		copy(*out, *in)
	}
	out.ClientConfig = in.ClientConfig
	out.Hooks = in.Hooks
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerControllerSpec.
func (in *InitializerControllerSpec) DeepCopy() *InitializerControllerSpec {
	if in == nil {
		return nil
	}
	out := new(InitializerControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerControllerStatus) DeepCopyInto(out *InitializerControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerControllerStatus.
func (in *InitializerControllerStatus) DeepCopy() *InitializerControllerStatus {
	if in == nil {
		return nil
	}
	out := new(InitializerControllerStatus)
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
