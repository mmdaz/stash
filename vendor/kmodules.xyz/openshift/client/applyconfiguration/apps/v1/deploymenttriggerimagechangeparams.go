/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// DeploymentTriggerImageChangeParamsApplyConfiguration represents an declarative configuration of the DeploymentTriggerImageChangeParams type for use
// with apply.
type DeploymentTriggerImageChangeParamsApplyConfiguration struct {
	Automatic          *bool               `json:"automatic,omitempty"`
	ContainerNames     []string            `json:"containerNames,omitempty"`
	From               *v1.ObjectReference `json:"from,omitempty"`
	LastTriggeredImage *string             `json:"lastTriggeredImage,omitempty"`
}

// DeploymentTriggerImageChangeParamsApplyConfiguration constructs an declarative configuration of the DeploymentTriggerImageChangeParams type for use with
// apply.
func DeploymentTriggerImageChangeParams() *DeploymentTriggerImageChangeParamsApplyConfiguration {
	return &DeploymentTriggerImageChangeParamsApplyConfiguration{}
}

// WithAutomatic sets the Automatic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Automatic field is set to the value of the last call.
func (b *DeploymentTriggerImageChangeParamsApplyConfiguration) WithAutomatic(value bool) *DeploymentTriggerImageChangeParamsApplyConfiguration {
	b.Automatic = &value
	return b
}

// WithContainerNames adds the given value to the ContainerNames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ContainerNames field.
func (b *DeploymentTriggerImageChangeParamsApplyConfiguration) WithContainerNames(values ...string) *DeploymentTriggerImageChangeParamsApplyConfiguration {
	for i := range values {
		b.ContainerNames = append(b.ContainerNames, values[i])
	}
	return b
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *DeploymentTriggerImageChangeParamsApplyConfiguration) WithFrom(value v1.ObjectReference) *DeploymentTriggerImageChangeParamsApplyConfiguration {
	b.From = &value
	return b
}

// WithLastTriggeredImage sets the LastTriggeredImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTriggeredImage field is set to the value of the last call.
func (b *DeploymentTriggerImageChangeParamsApplyConfiguration) WithLastTriggeredImage(value string) *DeploymentTriggerImageChangeParamsApplyConfiguration {
	b.LastTriggeredImage = &value
	return b
}