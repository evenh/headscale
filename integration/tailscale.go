package integration

import (
	"net/netip"

	"tailscale.com/ipn/ipnstate"
)

//nolint
type TailscaleClient interface {
	Hostname() string
	Shutdown() error
	Version() string
	Execute(command []string) (string, error)
	Up(loginServer, authKey string, enableSSH bool) error
	IPs() ([]netip.Addr, error)
	FQDN() (string, error)
	Status() (*ipnstate.Status, error)
	WaitForPeers(expected int) error
	Ping(hostnameOrIP string) error
	ID() string
}
