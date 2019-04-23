// Package openstack extracts OpenStack metadata from install
// configurations.
package openstack

import (
	"github.com/openshift-metal3/kni-installer/pkg/types"
	"github.com/openshift-metal3/kni-installer/pkg/types/openstack"
)

// Metadata converts an install configuration to OpenStack metadata.
func Metadata(infraID string, config *types.InstallConfig) *openstack.Metadata {
	return &openstack.Metadata{
		Region: config.Platform.OpenStack.Region,
		Cloud:  config.Platform.OpenStack.Cloud,
		Identifier: map[string]string{
			"openshiftClusterID": infraID,
		},
	}
}
