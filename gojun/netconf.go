package main

import (
	"encoding/xml"
	"os"
	"sync"

	"github.com/Juniper/go-netconf/netconf"
)

func readRouterList() (*[]string, error) {
	list := make([]string, 0)
	return &list, nil
}

func getBGPNei(s *netconf.Session) (*BgpInformation, error) {
	var bgpInfo BgpInformation
	reply, err := s.Exec(netconf.RawMethod("<get-bgp-neighbor-information></get-bgp-neighbor-information>"))
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal([]byte(reply.Data), &bgpInfo)
	if err != nil {
		return nil, err
	}
	return &bgpInfo, nil
}

func pollJunipers(hosts []string) []BgpInformation {
	var bgpInfo []BgpInformation
	wg := &sync.WaitGroup{}
	wg.Add(len(hosts))
	for _, host := range hosts {
		go pollJuniper(wg, host, &bgpInfo)
	}
	wg.Wait()
	return bgpInfo
}

func pollJuniper(wg *sync.WaitGroup, host string, jInfo *[]BgpInformation) {
	defer wg.Done()
	// session, err := netconf.DialSSH(host, netconf.SSHConfigPassword(nc.user, nc.pass))
	session, err := netconf.DialSSH(host, netconf.SSHConfigPassword(os.Args[1], os.Args[2]))
	if err != nil {
		logTopicErr("dial", "dial failure", err)
		*jInfo = append(*jInfo, BgpInformation{Host: host, Err: err.Error()})
		return
	}

	defer session.Close()
	bgp, err := getBGPNei(session)
	bgp.Host = host
	if err != nil {
		logTopicErr("bgpNei", "bgpNei collection failure", err)
		bgp.Err = err.Error()
		*jInfo = append(*jInfo, *bgp)
		return
	}
	if bgp != nil {
		*jInfo = append(*jInfo, *bgp)
	}
}
