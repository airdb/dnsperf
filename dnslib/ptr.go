package dnslib

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/miekg/dns"
)

const (
	MAX_CONCURRENCY = 10000000
	CHANNEL_CACHE   = 1000
)

var tmpChan = make(chan struct{}, CHANNEL_CACHE)
var SumChan = make(chan struct{}, MAX_CONCURRENCY)

func (ptr *PtrClient) StressPtr() {
	c := new(dns.Client)
	c.Timeout = 3 * time.Second


	for i := 0; i < MAX_CONCURRENCY; i++ {
		a := FakerData{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}

		tmpChan <- struct{}{}
		go ptr.request(c, a.IPV4)
	}

	for i := MAX_CONCURRENCY; i > 0; i-- {
		<-SumChan
	}
}

type PtrClient struct {
	DNSServer string
}

func (ptr *PtrClient)  request(c *dns.Client,  ip string) {
	msg := new(dns.Msg)

	domain := ip + ".in-addr.arpa."
	msg.SetQuestion(domain, dns.TypePTR)
	msg.SetEdns0(dns.DefaultMsgSize, false)

	upaddr := ptr.DNSServer
	if !strings.Contains( ptr.DNSServer, ":") {
		upaddr += ":53"
	}
	conn, _ := c.Dial(upaddr)

	resp, _, err := c.ExchangeWithConn(msg, conn)
	if err != nil || resp == nil {
		log.Println(err)
	} else {
		log.Println(upaddr, domain, resp.Id)
	}

	<-tmpChan
	SumChan <- struct{}{}
}

type FakerData struct {
	IPV4 string `faker:"ipv4"`
}