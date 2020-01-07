package mikrotik

type ipv6 struct {
	Firewall firewallv6CMD
	ND       ndv6CMD
	Settings cfg
}

// ====================================
//
// Firewall
//
// ====================================

type firewallv6CMD struct {
	AddressList cmd
	Connection  sshcmds
	Mangle      firewallOptionsCMD
	Filter      firewallOptionsCMD
	RAW         firewallOptionsCMD
}

// ====================================
//
// ND
//
// ====================================

type ndv6CMD struct {
	cmd
	Prefix ndPrefixCMD
}

type ndPrefixCMD struct {
	cmd
	Default cfg
}

// ====================================
//
// Settings
//
// ====================================
