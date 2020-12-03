/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 3.0.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package containerinventory

type ContainerInventoryObject struct {
	// Container object
	ContainerObject map[string]interface{} `json:"container_object"`
	// Object update type, one of CREATE, UPDATE, DELETE
	ObjectUpdateType string `json:"object_update_type"`
}