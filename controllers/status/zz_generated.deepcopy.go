//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package status

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentReleaseStatus) DeepCopyInto(out *ComponentReleaseStatus) {
	*out = *in
	in.Version.DeepCopyInto(&out.Version)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentReleaseStatus.
func (in *ComponentReleaseStatus) DeepCopy() *ComponentReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
	in.UpstreamReleases.DeepCopyInto(&out.UpstreamReleases)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamReleases) DeepCopyInto(out *UpstreamReleases) {
	*out = *in
	if in.UpstreamRelease != nil {
		in, out := &in.UpstreamRelease, &out.UpstreamRelease
		*out = make([]ComponentReleaseStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamReleases.
func (in *UpstreamReleases) DeepCopy() *UpstreamReleases {
	if in == nil {
		return nil
	}
	out := new(UpstreamReleases)
	in.DeepCopyInto(out)
	return out
}
