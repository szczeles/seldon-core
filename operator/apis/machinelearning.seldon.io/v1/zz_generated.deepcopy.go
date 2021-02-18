// +build !ignore_autogenerated

/*
Copyright 2019 The Seldon Authors.

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

package v1

import (
	"github.com/kedacore/keda/api/v1alpha1"
	"k8s.io/api/autoscaling/v2beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Explainer) DeepCopyInto(out *Explainer) {
	*out = *in
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Explainer.
func (in *Explainer) DeepCopy() *Explainer {
	if in == nil {
		return nil
	}
	out := new(Explainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logger) DeepCopyInto(out *Logger) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logger.
func (in *Logger) DeepCopy() *Logger {
	if in == nil {
		return nil
	}
	out := new(Logger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.DeletionTimestamp != nil {
		in, out := &in.DeletionTimestamp, &out.DeletionTimestamp
		*out = (*in).DeepCopy()
	}
	if in.DeletionGracePeriodSeconds != nil {
		in, out := &in.DeletionGracePeriodSeconds, &out.DeletionGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]metav1.OwnerReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictiveUnit) DeepCopyInto(out *PredictiveUnit) {
	*out = *in
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]PredictiveUnit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(PredictiveUnitType)
		**out = **in
	}
	if in.Implementation != nil {
		in, out := &in.Implementation, &out.Implementation
		*out = new(PredictiveUnitImplementation)
		**out = **in
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = new([]PredictiveUnitMethod)
		if **in != nil {
			in, out := *in, *out
			*out = make([]PredictiveUnitMethod, len(*in))
			copy(*out, *in)
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		copy(*out, *in)
	}
	if in.Logger != nil {
		in, out := &in.Logger, &out.Logger
		*out = new(Logger)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictiveUnit.
func (in *PredictiveUnit) DeepCopy() *PredictiveUnit {
	if in == nil {
		return nil
	}
	out := new(PredictiveUnit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorImageConfig) DeepCopyInto(out *PredictorImageConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorImageConfig.
func (in *PredictorImageConfig) DeepCopy() *PredictorImageConfig {
	if in == nil {
		return nil
	}
	out := new(PredictorImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorProtocolsConfig) DeepCopyInto(out *PredictorProtocolsConfig) {
	*out = *in
	if in.Seldon != nil {
		in, out := &in.Seldon, &out.Seldon
		*out = new(PredictorImageConfig)
		**out = **in
	}
	if in.KFServing != nil {
		in, out := &in.KFServing, &out.KFServing
		*out = new(PredictorImageConfig)
		**out = **in
	}
	if in.Tensorflow != nil {
		in, out := &in.Tensorflow, &out.Tensorflow
		*out = new(PredictorImageConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorProtocolsConfig.
func (in *PredictorProtocolsConfig) DeepCopy() *PredictorProtocolsConfig {
	if in == nil {
		return nil
	}
	out := new(PredictorProtocolsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorServerConfig) DeepCopyInto(out *PredictorServerConfig) {
	*out = *in
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make(map[Protocol]PredictorImageConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorServerConfig.
func (in *PredictorServerConfig) DeepCopy() *PredictorServerConfig {
	if in == nil {
		return nil
	}
	out := new(PredictorServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorSpec) DeepCopyInto(out *PredictorSpec) {
	*out = *in
	in.Graph.DeepCopyInto(&out.Graph)
	if in.ComponentSpecs != nil {
		in, out := &in.ComponentSpecs, &out.ComponentSpecs
		*out = make([]*SeldonPodSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SeldonPodSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.EngineResources.DeepCopyInto(&out.EngineResources)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.SvcOrchSpec.DeepCopyInto(&out.SvcOrchSpec)
	if in.Explainer != nil {
		in, out := &in.Explainer, &out.Explainer
		*out = new(Explainer)
		(*in).DeepCopyInto(*out)
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(SSL)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorSpec.
func (in *PredictorSpec) DeepCopy() *PredictorSpec {
	if in == nil {
		return nil
	}
	out := new(PredictorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSL) DeepCopyInto(out *SSL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSL.
func (in *SSL) DeepCopy() *SSL {
	if in == nil {
		return nil
	}
	out := new(SSL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonAddressable) DeepCopyInto(out *SeldonAddressable) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonAddressable.
func (in *SeldonAddressable) DeepCopy() *SeldonAddressable {
	if in == nil {
		return nil
	}
	out := new(SeldonAddressable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonDeployment) DeepCopyInto(out *SeldonDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonDeployment.
func (in *SeldonDeployment) DeepCopy() *SeldonDeployment {
	if in == nil {
		return nil
	}
	out := new(SeldonDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SeldonDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonDeploymentList) DeepCopyInto(out *SeldonDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SeldonDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonDeploymentList.
func (in *SeldonDeploymentList) DeepCopy() *SeldonDeploymentList {
	if in == nil {
		return nil
	}
	out := new(SeldonDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SeldonDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonDeploymentSpec) DeepCopyInto(out *SeldonDeploymentSpec) {
	*out = *in
	if in.Predictors != nil {
		in, out := &in.Predictors, &out.Predictors
		*out = make([]PredictorSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonDeploymentSpec.
func (in *SeldonDeploymentSpec) DeepCopy() *SeldonDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(SeldonDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonDeploymentStatus) DeepCopyInto(out *SeldonDeploymentStatus) {
	*out = *in
	if in.DeploymentStatus != nil {
		in, out := &in.DeploymentStatus, &out.DeploymentStatus
		*out = make(map[string]DeploymentStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceStatus != nil {
		in, out := &in.ServiceStatus, &out.ServiceStatus
		*out = make(map[string]ServiceStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(SeldonAddressable)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonDeploymentStatus.
func (in *SeldonDeploymentStatus) DeepCopy() *SeldonDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(SeldonDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonHpaSpec) DeepCopyInto(out *SeldonHpaSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta1.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonHpaSpec.
func (in *SeldonHpaSpec) DeepCopy() *SeldonHpaSpec {
	if in == nil {
		return nil
	}
	out := new(SeldonHpaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonPdbSpec) DeepCopyInto(out *SeldonPdbSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonPdbSpec.
func (in *SeldonPdbSpec) DeepCopy() *SeldonPdbSpec {
	if in == nil {
		return nil
	}
	out := new(SeldonPdbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonPodSpec) DeepCopyInto(out *SeldonPodSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.HpaSpec != nil {
		in, out := &in.HpaSpec, &out.HpaSpec
		*out = new(SeldonHpaSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.KedaSpec != nil {
		in, out := &in.KedaSpec, &out.KedaSpec
		*out = new(SeldonScaledObjectSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PdbSpec != nil {
		in, out := &in.PdbSpec, &out.PdbSpec
		*out = new(SeldonPdbSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonPodSpec.
func (in *SeldonPodSpec) DeepCopy() *SeldonPodSpec {
	if in == nil {
		return nil
	}
	out := new(SeldonPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeldonScaledObjectSpec) DeepCopyInto(out *SeldonScaledObjectSpec) {
	*out = *in
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(v1alpha1.AdvancedConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]v1alpha1.ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeldonScaledObjectSpec.
func (in *SeldonScaledObjectSpec) DeepCopy() *SeldonScaledObjectSpec {
	if in == nil {
		return nil
	}
	out := new(SeldonScaledObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SvcOrchSpec) DeepCopyInto(out *SvcOrchSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]*corev1.EnvVar, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.EnvVar)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SvcOrchSpec.
func (in *SvcOrchSpec) DeepCopy() *SvcOrchSpec {
	if in == nil {
		return nil
	}
	out := new(SvcOrchSpec)
	in.DeepCopyInto(out)
	return out
}
