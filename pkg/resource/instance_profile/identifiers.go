// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package instance_profile

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

// resourceIdentifiers implements the
// `aws-service-operator-k8s/pkg/types.AWSResourceIdentifiers` interface
type resourceIdentifiers struct {
	meta *ackv1alpha1.ResourceMetadata
}

// ARN returns the AWS Resource Name for the backend AWS resource. If nil,
// this means the resource has not yet been created in the backend AWS
// service.
func (ri *resourceIdentifiers) ARN() *ackv1alpha1.AWSResourceName {
	if ri.meta != nil {
		return ri.meta.ARN
	}
	return nil
}

// OwnerAccountID returns the AWS account identifier in which the
// backend AWS resource resides, or nil if this information is not known
// for the resource
func (ri *resourceIdentifiers) OwnerAccountID() *ackv1alpha1.AWSAccountID {
	if ri.meta != nil {
		return ri.meta.OwnerAccountID
	}
	return nil
}

// Region returns the AWS region in which the resource exists, or
// nil if this information is not known.
func (ri *resourceIdentifiers) Region() *ackv1alpha1.AWSRegion {
	if ri.meta != nil {
		return ri.meta.Region
	}
	return nil
}