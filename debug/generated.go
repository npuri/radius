// Code generated by generate_main.go. DO NOT EDIT.

package debug

import "github.com/holgermetschulat/radius/dictionary"

var IncludedDictionary = &dictionary.Dictionary{Attributes:[]*dictionary.Attribute{&dictionary.Attribute{Name:"User-Name",OID:dictionary.OID{1},Type:1,},&dictionary.Attribute{Name:"User-Password",OID:dictionary.OID{2},Type:1,FlagEncrypt:dictionary.IntFlag{Int:1, Valid:true},},&dictionary.Attribute{Name:"CHAP-Password",OID:dictionary.OID{3},Type:2,},&dictionary.Attribute{Name:"NAS-IP-Address",OID:dictionary.OID{4},Type:3,},&dictionary.Attribute{Name:"NAS-Port",OID:dictionary.OID{5},Type:5,},&dictionary.Attribute{Name:"Service-Type",OID:dictionary.OID{6},Type:5,},&dictionary.Attribute{Name:"Framed-Protocol",OID:dictionary.OID{7},Type:5,},&dictionary.Attribute{Name:"Framed-IP-Address",OID:dictionary.OID{8},Type:3,},&dictionary.Attribute{Name:"Framed-IP-Netmask",OID:dictionary.OID{9},Type:3,},&dictionary.Attribute{Name:"Framed-Routing",OID:dictionary.OID{10},Type:5,},&dictionary.Attribute{Name:"Filter-Id",OID:dictionary.OID{11},Type:1,},&dictionary.Attribute{Name:"Framed-MTU",OID:dictionary.OID{12},Type:5,},&dictionary.Attribute{Name:"Framed-Compression",OID:dictionary.OID{13},Type:5,},&dictionary.Attribute{Name:"Login-IP-Host",OID:dictionary.OID{14},Type:3,},&dictionary.Attribute{Name:"Login-Service",OID:dictionary.OID{15},Type:5,},&dictionary.Attribute{Name:"Login-TCP-Port",OID:dictionary.OID{16},Type:5,},&dictionary.Attribute{Name:"Reply-Message",OID:dictionary.OID{18},Type:1,},&dictionary.Attribute{Name:"Callback-Number",OID:dictionary.OID{19},Type:1,},&dictionary.Attribute{Name:"Callback-Id",OID:dictionary.OID{20},Type:1,},&dictionary.Attribute{Name:"Framed-Route",OID:dictionary.OID{22},Type:1,},&dictionary.Attribute{Name:"Framed-IPX-Network",OID:dictionary.OID{23},Type:3,},&dictionary.Attribute{Name:"State",OID:dictionary.OID{24},Type:2,},&dictionary.Attribute{Name:"Class",OID:dictionary.OID{25},Type:2,},&dictionary.Attribute{Name:"Vendor-Specific",OID:dictionary.OID{26},Type:10,},&dictionary.Attribute{Name:"Session-Timeout",OID:dictionary.OID{27},Type:5,},&dictionary.Attribute{Name:"Idle-Timeout",OID:dictionary.OID{28},Type:5,},&dictionary.Attribute{Name:"Termination-Action",OID:dictionary.OID{29},Type:5,},&dictionary.Attribute{Name:"Called-Station-Id",OID:dictionary.OID{30},Type:1,},&dictionary.Attribute{Name:"Calling-Station-Id",OID:dictionary.OID{31},Type:1,},&dictionary.Attribute{Name:"NAS-Identifier",OID:dictionary.OID{32},Type:1,},&dictionary.Attribute{Name:"Proxy-State",OID:dictionary.OID{33},Type:2,},&dictionary.Attribute{Name:"Login-LAT-Service",OID:dictionary.OID{34},Type:1,},&dictionary.Attribute{Name:"Login-LAT-Node",OID:dictionary.OID{35},Type:1,},&dictionary.Attribute{Name:"Login-LAT-Group",OID:dictionary.OID{36},Type:2,},&dictionary.Attribute{Name:"Framed-AppleTalk-Link",OID:dictionary.OID{37},Type:5,},&dictionary.Attribute{Name:"Framed-AppleTalk-Network",OID:dictionary.OID{38},Type:5,},&dictionary.Attribute{Name:"Framed-AppleTalk-Zone",OID:dictionary.OID{39},Type:1,},&dictionary.Attribute{Name:"CHAP-Challenge",OID:dictionary.OID{60},Type:2,},&dictionary.Attribute{Name:"NAS-Port-Type",OID:dictionary.OID{61},Type:5,},&dictionary.Attribute{Name:"Port-Limit",OID:dictionary.OID{62},Type:5,},&dictionary.Attribute{Name:"Login-LAT-Port",OID:dictionary.OID{63},Type:1,},&dictionary.Attribute{Name:"Acct-Status-Type",OID:dictionary.OID{40},Type:5,},&dictionary.Attribute{Name:"Acct-Delay-Time",OID:dictionary.OID{41},Type:5,},&dictionary.Attribute{Name:"Acct-Input-Octets",OID:dictionary.OID{42},Type:5,},&dictionary.Attribute{Name:"Acct-Output-Octets",OID:dictionary.OID{43},Type:5,},&dictionary.Attribute{Name:"Acct-Session-Id",OID:dictionary.OID{44},Type:1,},&dictionary.Attribute{Name:"Acct-Authentic",OID:dictionary.OID{45},Type:5,},&dictionary.Attribute{Name:"Acct-Session-Time",OID:dictionary.OID{46},Type:5,},&dictionary.Attribute{Name:"Acct-Input-Packets",OID:dictionary.OID{47},Type:5,},&dictionary.Attribute{Name:"Acct-Output-Packets",OID:dictionary.OID{48},Type:5,},&dictionary.Attribute{Name:"Acct-Terminate-Cause",OID:dictionary.OID{49},Type:5,},&dictionary.Attribute{Name:"Acct-Multi-Session-Id",OID:dictionary.OID{50},Type:1,},&dictionary.Attribute{Name:"Acct-Link-Count",OID:dictionary.OID{51},Type:5,},&dictionary.Attribute{Name:"Acct-Tunnel-Connection",OID:dictionary.OID{68},Type:1,},&dictionary.Attribute{Name:"Acct-Tunnel-Packets-Lost",OID:dictionary.OID{86},Type:5,},&dictionary.Attribute{Name:"Acct-Input-Gigawords",OID:dictionary.OID{52},Type:5,},&dictionary.Attribute{Name:"Acct-Output-Gigawords",OID:dictionary.OID{53},Type:5,},&dictionary.Attribute{Name:"Event-Timestamp",OID:dictionary.OID{55},Type:4,},&dictionary.Attribute{Name:"ARAP-Password",OID:dictionary.OID{70},Type:2,Size:dictionary.IntFlag{Int:16, Valid:true},},&dictionary.Attribute{Name:"ARAP-Features",OID:dictionary.OID{71},Type:2,Size:dictionary.IntFlag{Int:14, Valid:true},},&dictionary.Attribute{Name:"ARAP-Zone-Access",OID:dictionary.OID{72},Type:5,},&dictionary.Attribute{Name:"ARAP-Security",OID:dictionary.OID{73},Type:5,},&dictionary.Attribute{Name:"ARAP-Security-Data",OID:dictionary.OID{74},Type:1,},&dictionary.Attribute{Name:"Password-Retry",OID:dictionary.OID{75},Type:5,},&dictionary.Attribute{Name:"Prompt",OID:dictionary.OID{76},Type:5,},&dictionary.Attribute{Name:"Connect-Info",OID:dictionary.OID{77},Type:1,},&dictionary.Attribute{Name:"Configuration-Token",OID:dictionary.OID{78},Type:1,},&dictionary.Attribute{Name:"EAP-Message",OID:dictionary.OID{79},Type:2,FlagConcat:dictionary.BoolFlag{Bool:true, Valid:true},},&dictionary.Attribute{Name:"Message-Authenticator",OID:dictionary.OID{80},Type:2,},&dictionary.Attribute{Name:"ARAP-Challenge-Response",OID:dictionary.OID{84},Type:2,Size:dictionary.IntFlag{Int:8, Valid:true},},&dictionary.Attribute{Name:"Acct-Interim-Interval",OID:dictionary.OID{85},Type:5,},&dictionary.Attribute{Name:"NAS-Port-Id",OID:dictionary.OID{87},Type:1,},&dictionary.Attribute{Name:"Framed-Pool",OID:dictionary.OID{88},Type:1,},&dictionary.Attribute{Name:"NAS-IPv6-Address",OID:dictionary.OID{95},Type:6,},&dictionary.Attribute{Name:"Framed-Interface-Id",OID:dictionary.OID{96},Type:8,},&dictionary.Attribute{Name:"Framed-IPv6-Prefix",OID:dictionary.OID{97},Type:7,},&dictionary.Attribute{Name:"Login-IPv6-Host",OID:dictionary.OID{98},Type:6,},&dictionary.Attribute{Name:"Framed-IPv6-Route",OID:dictionary.OID{99},Type:1,},&dictionary.Attribute{Name:"Framed-IPv6-Pool",OID:dictionary.OID{100},Type:1,},&dictionary.Attribute{Name:"Error-Cause",OID:dictionary.OID{101},Type:5,},},Values:[]*dictionary.Value{&dictionary.Value{Attribute:"Service-Type", Name:"Login-User", Number:0x1},&dictionary.Value{Attribute:"Service-Type", Name:"Framed-User", Number:0x2},&dictionary.Value{Attribute:"Service-Type", Name:"Callback-Login-User", Number:0x3},&dictionary.Value{Attribute:"Service-Type", Name:"Callback-Framed-User", Number:0x4},&dictionary.Value{Attribute:"Service-Type", Name:"Outbound-User", Number:0x5},&dictionary.Value{Attribute:"Service-Type", Name:"Administrative-User", Number:0x6},&dictionary.Value{Attribute:"Service-Type", Name:"NAS-Prompt-User", Number:0x7},&dictionary.Value{Attribute:"Service-Type", Name:"Authenticate-Only", Number:0x8},&dictionary.Value{Attribute:"Service-Type", Name:"Callback-NAS-Prompt", Number:0x9},&dictionary.Value{Attribute:"Service-Type", Name:"Call-Check", Number:0xa},&dictionary.Value{Attribute:"Service-Type", Name:"Callback-Administrative", Number:0xb},&dictionary.Value{Attribute:"Framed-Protocol", Name:"PPP", Number:0x1},&dictionary.Value{Attribute:"Framed-Protocol", Name:"SLIP", Number:0x2},&dictionary.Value{Attribute:"Framed-Protocol", Name:"ARAP", Number:0x3},&dictionary.Value{Attribute:"Framed-Protocol", Name:"Gandalf-SLML", Number:0x4},&dictionary.Value{Attribute:"Framed-Protocol", Name:"Xylogics-IPX-SLIP", Number:0x5},&dictionary.Value{Attribute:"Framed-Protocol", Name:"X.75-Synchronous", Number:0x6},&dictionary.Value{Attribute:"Framed-Routing", Name:"None", Number:0x0},&dictionary.Value{Attribute:"Framed-Routing", Name:"Broadcast", Number:0x1},&dictionary.Value{Attribute:"Framed-Routing", Name:"Listen", Number:0x2},&dictionary.Value{Attribute:"Framed-Routing", Name:"Broadcast-Listen", Number:0x3},&dictionary.Value{Attribute:"Framed-Compression", Name:"None", Number:0x0},&dictionary.Value{Attribute:"Framed-Compression", Name:"Van-Jacobson-TCP-IP", Number:0x1},&dictionary.Value{Attribute:"Framed-Compression", Name:"IPX-Header-Compression", Number:0x2},&dictionary.Value{Attribute:"Framed-Compression", Name:"Stac-LZS", Number:0x3},&dictionary.Value{Attribute:"Login-Service", Name:"Telnet", Number:0x0},&dictionary.Value{Attribute:"Login-Service", Name:"Rlogin", Number:0x1},&dictionary.Value{Attribute:"Login-Service", Name:"TCP-Clear", Number:0x2},&dictionary.Value{Attribute:"Login-Service", Name:"PortMaster", Number:0x3},&dictionary.Value{Attribute:"Login-Service", Name:"LAT", Number:0x4},&dictionary.Value{Attribute:"Login-Service", Name:"X25-PAD", Number:0x5},&dictionary.Value{Attribute:"Login-Service", Name:"X25-T3POS", Number:0x6},&dictionary.Value{Attribute:"Login-Service", Name:"TCP-Clear-Quiet", Number:0x8},&dictionary.Value{Attribute:"Login-TCP-Port", Name:"Telnet", Number:0x17},&dictionary.Value{Attribute:"Login-TCP-Port", Name:"Rlogin", Number:0x201},&dictionary.Value{Attribute:"Login-TCP-Port", Name:"Rsh", Number:0x202},&dictionary.Value{Attribute:"Termination-Action", Name:"Default", Number:0x0},&dictionary.Value{Attribute:"Termination-Action", Name:"RADIUS-Request", Number:0x1},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"Async", Number:0x0},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"Sync", Number:0x1},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"ISDN", Number:0x2},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"ISDN-V120", Number:0x3},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"ISDN-V110", Number:0x4},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"Virtual", Number:0x5},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"PIAFS", Number:0x6},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"HDLC-Clear-Channel", Number:0x7},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"X.25", Number:0x8},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"X.75", Number:0x9},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"G.3-Fax", Number:0xa},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"SDSL", Number:0xb},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"ADSL-CAP", Number:0xc},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"ADSL-DMT", Number:0xd},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"IDSL", Number:0xe},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"Ethernet", Number:0xf},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"xDSL", Number:0x10},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"Cable", Number:0x11},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"Wireless-Other", Number:0x12},&dictionary.Value{Attribute:"NAS-Port-Type", Name:"Wireless-802.11", Number:0x13},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Start", Number:0x1},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Stop", Number:0x2},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Alive", Number:0x3},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Interim-Update", Number:0x3},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Accounting-On", Number:0x7},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Accounting-Off", Number:0x8},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Failed", Number:0xf},&dictionary.Value{Attribute:"Acct-Authentic", Name:"RADIUS", Number:0x1},&dictionary.Value{Attribute:"Acct-Authentic", Name:"Local", Number:0x2},&dictionary.Value{Attribute:"Acct-Authentic", Name:"Remote", Number:0x3},&dictionary.Value{Attribute:"Acct-Authentic", Name:"Diameter", Number:0x4},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"User-Request", Number:0x1},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Lost-Carrier", Number:0x2},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Lost-Service", Number:0x3},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Idle-Timeout", Number:0x4},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Session-Timeout", Number:0x5},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Admin-Reset", Number:0x6},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Admin-Reboot", Number:0x7},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Port-Error", Number:0x8},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"NAS-Error", Number:0x9},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"NAS-Request", Number:0xa},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"NAS-Reboot", Number:0xb},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Port-Unneeded", Number:0xc},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Port-Preempted", Number:0xd},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Port-Suspended", Number:0xe},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Service-Unavailable", Number:0xf},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Callback", Number:0x10},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"User-Error", Number:0x11},&dictionary.Value{Attribute:"Acct-Terminate-Cause", Name:"Host-Request", Number:0x12},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Tunnel-Start", Number:0x9},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Tunnel-Stop", Number:0xa},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Tunnel-Reject", Number:0xb},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Tunnel-Link-Start", Number:0xc},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Tunnel-Link-Stop", Number:0xd},&dictionary.Value{Attribute:"Acct-Status-Type", Name:"Tunnel-Link-Reject", Number:0xe},&dictionary.Value{Attribute:"ARAP-Zone-Access", Name:"Default-Zone", Number:0x1},&dictionary.Value{Attribute:"ARAP-Zone-Access", Name:"Zone-Filter-Inclusive", Number:0x2},&dictionary.Value{Attribute:"ARAP-Zone-Access", Name:"Zone-Filter-Exclusive", Number:0x4},&dictionary.Value{Attribute:"Prompt", Name:"No-Echo", Number:0x0},&dictionary.Value{Attribute:"Prompt", Name:"Echo", Number:0x1},&dictionary.Value{Attribute:"Service-Type", Name:"Authorize-Only", Number:0x11},&dictionary.Value{Attribute:"Error-Cause", Name:"Residual-Context-Removed", Number:0xc9},&dictionary.Value{Attribute:"Error-Cause", Name:"Invalid-EAP-Packet", Number:0xca},&dictionary.Value{Attribute:"Error-Cause", Name:"Unsupported-Attribute", Number:0x191},&dictionary.Value{Attribute:"Error-Cause", Name:"Missing-Attribute", Number:0x192},&dictionary.Value{Attribute:"Error-Cause", Name:"NAS-Identification-Mismatch", Number:0x193},&dictionary.Value{Attribute:"Error-Cause", Name:"Invalid-Request", Number:0x194},&dictionary.Value{Attribute:"Error-Cause", Name:"Unsupported-Service", Number:0x195},&dictionary.Value{Attribute:"Error-Cause", Name:"Unsupported-Extension", Number:0x196},&dictionary.Value{Attribute:"Error-Cause", Name:"Administratively-Prohibited", Number:0x1f5},&dictionary.Value{Attribute:"Error-Cause", Name:"Proxy-Request-Not-Routable", Number:0x1f6},&dictionary.Value{Attribute:"Error-Cause", Name:"Session-Context-Not-Found", Number:0x1f7},&dictionary.Value{Attribute:"Error-Cause", Name:"Session-Context-Not-Removable", Number:0x1f8},&dictionary.Value{Attribute:"Error-Cause", Name:"Proxy-Processing-Error", Number:0x1f9},&dictionary.Value{Attribute:"Error-Cause", Name:"Resources-Unavailable", Number:0x1fa},&dictionary.Value{Attribute:"Error-Cause", Name:"Request-Initiated", Number:0x1fb},&dictionary.Value{Attribute:"Error-Cause", Name:"Invalid-Attribute-Value", Number:0x197},&dictionary.Value{Attribute:"Error-Cause", Name:"Multiple-Session-Selection-Unsupported", Number:0x1fc},},}
