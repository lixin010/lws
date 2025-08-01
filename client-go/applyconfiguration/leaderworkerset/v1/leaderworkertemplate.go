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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	leaderworkersetv1 "sigs.k8s.io/lws/api/leaderworkerset/v1"
)

// LeaderWorkerTemplateApplyConfiguration represents a declarative configuration of the LeaderWorkerTemplate type for use
// with apply.
type LeaderWorkerTemplateApplyConfiguration struct {
	LeaderTemplate *corev1.PodTemplateSpecApplyConfiguration `json:"leaderTemplate,omitempty"`
	WorkerTemplate *corev1.PodTemplateSpecApplyConfiguration `json:"workerTemplate,omitempty"`
	Size           *int32                                    `json:"size,omitempty"`
	RestartPolicy  *leaderworkersetv1.RestartPolicyType      `json:"restartPolicy,omitempty"`
	SubGroupPolicy *SubGroupPolicyApplyConfiguration         `json:"subGroupPolicy,omitempty"`
	ResizePolicy   *leaderworkersetv1.ResizePolicyType       `json:"resizePolicy,omitempty"`
}

// LeaderWorkerTemplateApplyConfiguration constructs a declarative configuration of the LeaderWorkerTemplate type for use with
// apply.
func LeaderWorkerTemplate() *LeaderWorkerTemplateApplyConfiguration {
	return &LeaderWorkerTemplateApplyConfiguration{}
}

// WithLeaderTemplate sets the LeaderTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LeaderTemplate field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithLeaderTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *LeaderWorkerTemplateApplyConfiguration {
	b.LeaderTemplate = value
	return b
}

// WithWorkerTemplate sets the WorkerTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WorkerTemplate field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithWorkerTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *LeaderWorkerTemplateApplyConfiguration {
	b.WorkerTemplate = value
	return b
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithSize(value int32) *LeaderWorkerTemplateApplyConfiguration {
	b.Size = &value
	return b
}

// WithRestartPolicy sets the RestartPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartPolicy field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithRestartPolicy(value leaderworkersetv1.RestartPolicyType) *LeaderWorkerTemplateApplyConfiguration {
	b.RestartPolicy = &value
	return b
}

// WithSubGroupPolicy sets the SubGroupPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubGroupPolicy field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithSubGroupPolicy(value *SubGroupPolicyApplyConfiguration) *LeaderWorkerTemplateApplyConfiguration {
	b.SubGroupPolicy = value
	return b
}

// WithResizePolicy sets the ResizePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResizePolicy field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithResizePolicy(value leaderworkersetv1.ResizePolicyType) *LeaderWorkerTemplateApplyConfiguration {
	b.ResizePolicy = &value
	return b
}
