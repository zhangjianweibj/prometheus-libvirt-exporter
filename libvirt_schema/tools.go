package libvirt_schema

type DomainState uint8

const (
	//no state
	DOMAIN_NOSTATE = DomainState(0)
	//the domain is running
	DOMAIN_RUNNING = DomainState(1)
	//the domain is blocked on resource
	DOMAIN_BLOCKED = DomainState(2)
	//the domain is paused by user
	DOMAIN_PAUSED = DomainState(3)
	//the domain is being shut down
	DOMAIN_SHUTDOWN = DomainState(4)
	//the domain is shut off
	DOMAIN_SHUTOFF = DomainState(5)
	//the domain is crashed
	DOMAIN_CRASHED = DomainState(6)
	//the domain is suspended by guest power management
	DOMAIN_PMSUSPENDED = DomainState(7)
	//this enum value will increase over time as new events are added to the libvirt API. It reflects the last state supported by this version of the libvirt API.
	DOMAIN_LAST = DomainState(8)
)
