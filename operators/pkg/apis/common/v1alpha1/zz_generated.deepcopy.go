// +build !ignore_autogenerated

// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by main. DO NOT EDIT.

package v1alpha1

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagState) DeepCopyInto(out *FeatureFlagState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagState.
func (in *FeatureFlagState) DeepCopy() *FeatureFlagState {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FeatureFlags) DeepCopyInto(out *FeatureFlags) {
	{
		in := &in
		*out = make(FeatureFlags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlags.
func (in FeatureFlags) DeepCopy() FeatureFlags {
	if in == nil {
		return nil
	}
	out := new(FeatureFlags)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPConfig) DeepCopyInto(out *HTTPConfig) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
	in.TLS.DeepCopyInto(&out.TLS)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPConfig.
func (in *HTTPConfig) DeepCopy() *HTTPConfig {
	if in == nil {
		return nil
	}
	out := new(HTTPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPService) DeepCopyInto(out *HTTPService) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPService.
func (in *HTTPService) DeepCopy() *HTTPService {
	if in == nil {
		return nil
	}
	out := new(HTTPService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPServiceObjectMeta) DeepCopyInto(out *HTTPServiceObjectMeta) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPServiceObjectMeta.
func (in *HTTPServiceObjectMeta) DeepCopy() *HTTPServiceObjectMeta {
	if in == nil {
		return nil
	}
	out := new(HTTPServiceObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPServiceSpec) DeepCopyInto(out *HTTPServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPServiceSpec.
func (in *HTTPServiceSpec) DeepCopy() *HTTPServiceSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSelector) DeepCopyInto(out *ObjectSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSelector.
func (in *ObjectSelector) DeepCopy() *ObjectSelector {
	if in == nil {
		return nil
	}
	out := new(ObjectSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReconcilerStatus) DeepCopyInto(out *ReconcilerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReconcilerStatus.
func (in *ReconcilerStatus) DeepCopy() *ReconcilerStatus {
	if in == nil {
		return nil
	}
	out := new(ReconcilerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedCertificate) DeepCopyInto(out *SelfSignedCertificate) {
	*out = *in
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]SubjectAlternativeName, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedCertificate.
func (in *SelfSignedCertificate) DeepCopy() *SelfSignedCertificate {
	if in == nil {
		return nil
	}
	out := new(SelfSignedCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAlternativeName) DeepCopyInto(out *SubjectAlternativeName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAlternativeName.
func (in *SubjectAlternativeName) DeepCopy() *SubjectAlternativeName {
	if in == nil {
		return nil
	}
	out := new(SubjectAlternativeName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSOptions) DeepCopyInto(out *TLSOptions) {
	*out = *in
	if in.SelfSignedCertificate != nil {
		in, out := &in.SelfSignedCertificate, &out.SelfSignedCertificate
		*out = new(SelfSignedCertificate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSOptions.
func (in *TLSOptions) DeepCopy() *TLSOptions {
	if in == nil {
		return nil
	}
	out := new(TLSOptions)
	in.DeepCopyInto(out)
	return out
}
