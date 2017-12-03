package main

import (
	"encoding/xml"
)

type BgpBfd struct {
	BfdConfigurationState string   `xml:"bfd-configuration-state" json:"bfd-configuration-state"`
	BfdOperationalState   string   `xml:"bfd-operational-state" json:"bfd-operational-state"`
	XMLName               xml.Name `xml:"bgp-bfd" json:"bgp-bfd"`
}

type BgpError struct {
	Name         string   `xml:"name" json:"name"`
	ReceiveCount int32    `xml:"receive-count" json:"receive-count"`
	SendCount    int32    `xml:"send-count" json:"send-count"`
	XMLName      xml.Name `xml:"bgp-error" json:"bgp-error"`
}

type BgpInformation struct {
	Host      string     `xml:"host" json:"host"`
	AttrXmlns string     `xml:" xmlns,attr"  json:""`
	BgpPeer   []*BgpPeer `xml:"bgp-peer" json:"bgp-peer"`
	XMLName   xml.Name   `xml:"bgp-information" json:"bgp-information"`
	Err       string     `xml:"err" json:"err"`
}

type BgpOptionInformation struct {
	AttrXmlns                string   `xml:" xmlns,attr"  json:""`
	AddressFamilies          string   `xml:"address-families" json:"address-families"`
	AuthenticationAlgorithm  string   `xml:"authentication-algorithm" json:"authentication-algorithm"`
	AuthenticationConfigured string   `xml:"authentication-configured" json:"authentication-configured"`
	AuthenticationKeyChain   string   `xml:"authentication-key-chain" json:"authentication-key-chain"`
	BgpOptions               string   `xml:"bgp-options" json:"bgp-options"`
	BgpOptions2              string   `xml:"bgp-options2" json:"bgp-options2"`
	BgpOptionsExtended       string   `xml:"bgp-options-extended" json:"bgp-options-extended"`
	ExportPolicy             string   `xml:"export-policy" json:"export-policy"`
	Holdtime                 int32    `xml:"holdtime" json:"holdtime"`
	ImportPolicy             string   `xml:"import-policy" json:"import-policy"`
	LocalAddress             string   `xml:"local-address" json:"local-address"`
	LocalAs                  string   `xml:"local-as" json:"local-as"`
	LocalSystemAs            string   `xml:"local-system-as" json:"local-system-as"`
	Preference               int32    `xml:"preference" json:"preference"`
	XMLName                  xml.Name `xml:"bgp-option-information" json:"bgp-option-information"`
}

type BgpOutputQueue struct {
	Count      int32    `xml:"count" json:"count"`
	Number     int32    `xml:"number" json:"number"`
	RibAdvNlri string   `xml:"rib-adv-nlri" json:"rib-adv-nlri"`
	TableName  string   `xml:"table-name" json:"table-name"`
	XMLName    xml.Name `xml:"bgp-output-queue" json:"bgp-output-queue"`
}

type BgpPeer struct {
	AttrJunosStyle                    string                `xml:"junos style,attr"  json:""`
	ActiveHoldtime                    int32                 `xml:"active-holdtime" json:"active-holdtime"`
	BgpBfd                            *BgpBfd               `xml:"bgp-bfd" json:"bgp-bfd"`
	BgpError                          []*BgpError           `xml:"bgp-error" json:"bgp-error"`
	BgpOptionInformation              *BgpOptionInformation `xml:"bgp-option-information" json:"bgp-option-information"`
	BgpOutputQueue                    []*BgpOutputQueue     `xml:"bgp-output-queue" json:"bgp-output-queue"`
	BgpRib                            []*BgpRib             `xml:"bgp-rib" json:"bgp-rib"`
	Description                       string                `xml:"description" json:"description"`
	FlapCount                         int32                 `xml:"flap-count" json:"flap-count"`
	GatherBufferedOctetsTx            int32                 `xml:"gather-buffered-octets-tx" json:"gather-buffered-octets-tx"`
	GroupIndex                        int32                 `xml:"group-index" json:"group-index"`
	InputMessages                     int32                 `xml:"input-messages" json:"input-messages"`
	InputOctets                       int64                 `xml:"input-octets" json:"input-octets"`
	InputRefreshes                    int32                 `xml:"input-refreshes" json:"input-refreshes"`
	InputUpdates                      int32                 `xml:"input-updates" json:"input-updates"`
	KeepaliveInterval                 int32                 `xml:"keepalive-interval" json:"keepalive-interval"`
	LastChecked                       int32                 `xml:"last-checked" json:"last-checked"`
	LastError                         string                `xml:"last-error" json:"last-error"`
	LastEvent                         string                `xml:"last-event" json:"last-event"`
	LastFlapEvent                     string                `xml:"last-flap-event" json:"last-flap-event"`
	LastReceived                      int32                 `xml:"last-received" json:"last-received"`
	LastSent                          int32                 `xml:"last-sent" json:"last-sent"`
	LastState                         string                `xml:"last-state" json:"last-state"`
	LocalAddress                      string                `xml:"local-address" json:"local-address"`
	LocalAs                           string                `xml:"local-as" json:"local-as"`
	LocalId                           string                `xml:"local-id" json:"local-id"`
	LocalInterfaceIndex               int32                 `xml:"local-interface-index" json:"local-interface-index"`
	LocalInterfaceName                string                `xml:"local-interface-name" json:"local-interface-name"`
	NlriTypePeer                      string                `xml:"nlri-type-peer" json:"nlri-type-peer"`
	NlriTypeSession                   string                `xml:"nlri-type-session" json:"nlri-type-session"`
	OutputMessages                    int32                 `xml:"output-messages" json:"output-messages"`
	OutputOctets                      int64                 `xml:"output-octets" json:"output-octets"`
	OutputRefreshes                   int32                 `xml:"output-refreshes" json:"output-refreshes"`
	OutputUpdates                     int32                 `xml:"output-updates" json:"output-updates"`
	Peer4byteAsCapabilityAdvertised   int32                 `xml:"peer-4byte-as-capability-advertised" json:"peer-4byte-as-capability-advertised"`
	Peer4byteAsCapabilityNotSupported string                `xml:"peer-4byte-as-capability-not-supported" json:"peer-4byte-as-capability-not-supported"`
	PeerAddpathNotSupported           string                `xml:"peer-addpath-not-supported" json:"peer-addpath-not-supported"`
	PeerAddress                       string                `xml:"peer-address" json:"peer-address"`
	PeerAs                            string                `xml:"peer-as" json:"peer-as"`
	PeerCfgRti                        string                `xml:"peer-cfg-rti" json:"peer-cfg-rti"`
	PeerEndOfRibReceived              string                `xml:"peer-end-of-rib-received" json:"peer-end-of-rib-received"`
	PeerEndOfRibScheduled             string                `xml:"peer-end-of-rib-scheduled" json:"peer-end-of-rib-scheduled"`
	PeerEndOfRibSent                  string                `xml:"peer-end-of-rib-sent" json:"peer-end-of-rib-sent"`
	PeerFlags                         string                `xml:"peer-flags" json:"peer-flags"`
	PeerGroup                         string                `xml:"peer-group" json:"peer-group"`
	PeerId                            string                `xml:"peer-id" json:"peer-id"`
	PeerIndex                         int32                 `xml:"peer-index" json:"peer-index"`
	PeerNoHelper                      string                `xml:"peer-no-helper" json:"peer-no-helper"`
	PeerNoRefresh                     string                `xml:"peer-no-refresh" json:"peer-no-refresh"`
	PeerNoRestart                     string                `xml:"peer-no-restart" json:"peer-no-restart"`
	PeerRefreshCapability             int32                 `xml:"peer-refresh-capability" json:"peer-refresh-capability"`
	PeerRestartFlagsReceived          string                `xml:"peer-restart-flags-received" json:"peer-restart-flags-received"`
	PeerRestartNlriCanSaveState       string                `xml:"peer-restart-nlri-can-save-state" json:"peer-restart-nlri-can-save-state"`
	PeerRestartNlriConfigured         string                `xml:"peer-restart-nlri-configured" json:"peer-restart-nlri-configured"`
	PeerRestartNlriNegotiated         string                `xml:"peer-restart-nlri-negotiated" json:"peer-restart-nlri-negotiated"`
	PeerRestartNlriReceived           string                `xml:"peer-restart-nlri-received" json:"peer-restart-nlri-received"`
	PeerRestartNlriStateSaved         string                `xml:"peer-restart-nlri-state-saved" json:"peer-restart-nlri-state-saved"`
	PeerRestartTimeReceived           int32                 `xml:"peer-restart-time-received" json:"peer-restart-time-received"`
	PeerStaleRouteTimeConfigured      int32                 `xml:"peer-stale-route-time-configured" json:"peer-stale-route-time-configured"`
	PeerState                         string                `xml:"peer-state" json:"peer-state"`
	PeerType                          string                `xml:"peer-type" json:"peer-type"`
	RouteReflectorClient              string                `xml:"route-reflector-client" json:"route-reflector-client"`
	XMLName                           xml.Name              `xml:"bgp-peer" json:"bgp-peer"`
}

type BgpRib struct {
	AttrJunosStyle        string   `xml:"junos style,attr"  json:""`
	AcceptedPrefixCount   int32    `xml:"accepted-prefix-count" json:"accepted-prefix-count"`
	ActivePrefixCount     int32    `xml:"active-prefix-count" json:"active-prefix-count"`
	AdvertisedPrefixCount int32    `xml:"advertised-prefix-count" json:"advertised-prefix-count"`
	BgpRibState           string   `xml:"bgp-rib-state" json:"bgp-rib-state"`
	Name                  string   `xml:"name" json:"name"`
	ReceivedPrefixCount   int32    `xml:"received-prefix-count" json:"received-prefix-count"`
	RibBit                string   `xml:"rib-bit" json:"rib-bit"`
	SendState             string   `xml:"send-state" json:"send-state"`
	SuppressedPrefixCount int32    `xml:"suppressed-prefix-count" json:"suppressed-prefix-count"`
	VpnRibState           string   `xml:"vpn-rib-state" json:"vpn-rib-state"`
	XMLName               xml.Name `xml:"bgp-rib" json:"bgp-rib"`
}
