// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	libvirttfvars "github.com/openshift-metal3/kni-installer/pkg/tfvars/libvirt"
	"github.com/pkg/errors"
	"github.com/rodaine/hclencoder"
)

type config struct {
	LibvirtURI      string `hcl:"libvirt_uri,omitempty"`
	IronicURI       string `hcl:"ironic_uri,omitempty"`
	Image           string `hcl:"os_image,omitempty"`
	BareMetalBridge string `hcl:"baremetal_bridge,omitempty"`
	OverCloudBridge string `hcl:"overcloud_bridge,omitempty"`

	// Data required for masters deployment - several maps per master, because of terraform's
	// limitation that maps cannot be strings
	MasterNodes interface{} `hcl:"master_nodes"`
	Properties  interface{} `hcl:"properties"`
	RootDevices interface{} `hcl:"root_devices"`
	DriverInfos interface{} `hcl:"driver_infos"`

	MasterConfiguration map[string]interface{} `hcl:"master_configuration"`
}

// TFVars generates bare metal specific Terraform variables.
func TFVars(libvirtURI, ironicURI, osImage, baremetalBridge, overcloudBridge string, nodes map[string]interface{}, configuration map[string]interface{}) ([]byte, error) {
	osImage, err := libvirttfvars.CachedImage(osImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached libvirt image")
	}

	cfg := &config{
		LibvirtURI:          libvirtURI,
		IronicURI:           ironicURI,
		Image:               osImage,
		BareMetalBridge:     baremetalBridge,
		OverCloudBridge:     overcloudBridge,
		MasterNodes:         nodes["master_nodes"],
		Properties:          nodes["properties"],
		RootDevices:         nodes["root_devices"],
		DriverInfos:         nodes["driver_infos"],
		MasterConfiguration: configuration,
	}

	return hclencoder.Encode(cfg)
}
