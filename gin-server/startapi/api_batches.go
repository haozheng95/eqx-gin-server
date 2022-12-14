/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package startapi

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDeviceBatch - Create a devices batch
func CreateDeviceBatch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteBatch - Delete the Batch
func DeleteBatch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindBatchById - Retrieve a Batch
func FindBatchById(c *gin.Context) {
	id := c.Param("id")
	resp, r, err := apiClient.BatchesApi.FindBatchById(context.Background(), id).Include(include).Exclude(exclude).Execute()
	if !ErrorCheck(err, r, c) {
		return
	}
	c.JSON(http.StatusOK, resp)
}

// FindBatchesByProject - Retrieve all batches by project
func FindBatchesByProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
