package dnslib

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/miekg/dns"
)

type FakerDataRecordA struct {
	DomainName string `faker:"word"`
}


func (ptr *PtrClient) FakerDataRecordA() (*FakerDataRecordA, error){
	f := FakerDataRecordA{}
	err := faker.FakeData(&f)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &f, nil
}

func (ptr *PtrClient) StressRecordA() {
	c := new(dns.Client)
	c.Timeout = 3 * time.Second

	f, err := ptr.FakerDataRecordA()
	if err != nil {
		return
	}

	ptr.RequestA(c, f.DomainName)
}

func (ptr *PtrClient)  RequestA(c *dns.Client,  domain string) {
	msg := new(dns.Msg)

	msg.SetQuestion(domain, dns.TypeA)
	msg.SetEdns0(dns.DefaultMsgSize, false)

	upaddr := ptr.DNSServer
	if !strings.Contains( ptr.DNSServer, ":") {
		upaddr += ":53"
	}

	conn, err := c.Dial(upaddr)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	resp, _, err := c.ExchangeWithConn(msg, conn)
	if err != nil || resp == nil {
		log.Println(err)
	} else {
		log.Println(upaddr, domain, resp.Id)
	}

	<-tmpChan
	SumChan <- struct{}{}
}