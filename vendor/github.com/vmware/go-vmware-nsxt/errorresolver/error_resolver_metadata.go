/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package errorresolver

// Error along with its metadata
type ErrorResolverMetadata struct {

	// The entity/node UUID where the error has occurred.
	EntityId string `json:"entity_id"`

	// The error id as reported by the entity where the error occurred.
	ErrorId int64 `json:"error_id"`

	// This can come from some external system like syslog collector
	SystemMetadata *ErrorResolverSystemMetadata `json:"system_metadata,omitempty"`

	// User supplied metadata that might be required by the resolver
	UserMetadata *ErrorResolverUserMetadata `json:"user_metadata,omitempty"`
}
