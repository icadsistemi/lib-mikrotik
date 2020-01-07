package mikrotik

// ====================================
//
// Neighbor Discovery
//
// ====================================

type IPV6NeighborDiscovery struct {
	ID                          string `mikrotik:".id"`
	AdvertiseDNS                bool   `json:"advertise-dns"`
	AdvertiseMacAddress         bool   `json:"advertise-mac-address"`
	Comment                     string `mikrotik:"comment"`
	Default                     bool   `json:"default"`
	Disabled                    bool   `json:"disabled"`
	HopLimit                    string `json:"hop-limit"`
	Interface                   string `json:"interface"`
	Invalid                     bool   `json:"invalid"`
	ManagedAddressConfiguration bool   `json:"managed-address-configuration"`
	Mtu                         string `json:"mtu"`
	OtherConfiguration          bool   `json:"other-configuration"`
	RaDelay                     string `json:"ra-delay"`
	RaInterval                  string `json:"ra-interval"`
	RaLifetime                  string `json:"ra-lifetime"`
	ReachableTime               string `json:"reachable-time"`
	RetransmitInterval          string `json:"retransmit-interval"`
}

type IPV6NeighborDiscoveryPrefix struct {
	ID                 string `json:".id"`
	SixToFourInterface string `json:"6to4-interface"`
	Autonomous         bool   `json:"autonomous"`
	Comment            string `mikrotik:"comment"`
	Disabled           bool   `json:"disabled"`
	Interface          string `json:"interface"`
	Invalid            bool   `json:"invalid"`
	OnLink             bool   `json:"on-link"`
	PreferredLifetime  string `json:"preferred-lifetime"`
	Prefix             string `json:"prefix"`
	ValidLifetime      string `json:"valid-lifetime"`
}

type IPV6NeighborDiscoveryPrefixDefault struct {
	Autonomous        bool   `json:"autonomous"`
	PreferredLifetime string `json:"preferred-lifetime"`
	ValidLifetime     string `json:"valid-lifetime"`
}

// ====================================
//
// Settings
//
// ====================================

type IPV6Settings struct {
	AcceptRedirects            string `json:"accept-redirects"`
	AcceptRouterAdvertisements string `json:"accept-router-advertisements"`
	Forward                    bool   `json:"forward"`
	MaxNeighborEntries         string `json:"max-neighbor-entries"`
}

// ====================================
//
// Firewall
//
// ====================================

type IPV6FirewallRaw struct {
	ID      string `json:".id"`
	Action  string `json:"action"`
	Bytes   int    `json:"bytes"`
	Chain   string `json:"chain"`
	Comment string `json:"comment"`
	Dynamic string `json:"dynamic"`
	Invalid string `json:"invalid"`
	Packets int    `json:"packets"`
}
