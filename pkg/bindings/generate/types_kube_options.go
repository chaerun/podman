// Code generated by go generate; DO NOT EDIT.
package generate

import (
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *KubeOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *KubeOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithService set field Service to given value
func (o *KubeOptions) WithService(value bool) *KubeOptions {
	o.Service = &value
	return o
}

// GetService returns value of field Service
func (o *KubeOptions) GetService() bool {
	if o.Service == nil {
		var z bool
		return z
	}
	return *o.Service
}

// WithType set field Type to given value
func (o *KubeOptions) WithType(value string) *KubeOptions {
	o.Type = &value
	return o
}

// GetType returns value of field Type
func (o *KubeOptions) GetType() string {
	if o.Type == nil {
		var z string
		return z
	}
	return *o.Type
}

// WithReplicas set field Replicas to given value
func (o *KubeOptions) WithReplicas(value int32) *KubeOptions {
	o.Replicas = &value
	return o
}

// GetReplicas returns value of field Replicas
func (o *KubeOptions) GetReplicas() int32 {
	if o.Replicas == nil {
		var z int32
		return z
	}
	return *o.Replicas
}
