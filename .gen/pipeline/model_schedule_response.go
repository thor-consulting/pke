/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// ScheduleResponse struct for ScheduleResponse
type ScheduleResponse struct {
	Uid string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Schedule string `json:"schedule,omitempty"`
	Ttl string `json:"ttl,omitempty"`
	Labels Labels `json:"labels,omitempty"`
	Options BackupOptions `json:"options,omitempty"`
	Status string `json:"status,omitempty"`
	LastBackup string `json:"lastBackup,omitempty"`
}
