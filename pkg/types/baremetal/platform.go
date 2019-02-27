package baremetal

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	// URI is the identifier for the libvirtd connection.  It must be
	// reachable from the host where the installer is run.
	// +optional
	// Default is qemu:///system
	URI string `json:"URI,omitempty"`

	// MasterVIP is the VIP for the master API.
	// +optional
	// FIXME: bare metal - add a default, use net.IP type
	MasterVIP string `json:"masterVIP,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on bare metal for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`
}
