package defaults

import (
	"github.com/openshift-metal3/kni-installer/pkg/types"
	"github.com/openshift-metal3/kni-installer/pkg/types/libvirt"
)

// SetMachinePoolDefaults sets the defaults for the machine pool.
func SetMachinePoolDefaults(p *types.MachinePool, platform string) {
	defaultReplicaCount := int64(3)
	if platform == libvirt.Name {
		defaultReplicaCount = 1
	}
	if p.Replicas == nil {
		p.Replicas = &defaultReplicaCount
	}
	if p.Hyperthreading == "" {
		p.Hyperthreading = types.HyperthreadingEnabled
	}
}
