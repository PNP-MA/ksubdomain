package runner

import (
	"github.com/phayes/freeport"
	"ksubdomain/core"
	"ksubdomain/core/gologger"
	"net"
	"time"
)

func TestSpeed(ether core.EthTable) {
	ether.DstMac = net.HardwareAddr{0x5c, 0xc9, 0x09, 0x33, 0x34, 0x80} // 指定一个错误的dstmac地址，包会经过本机网卡，但是发不出去
	var index int64 = 0
	start := time.Now().UnixNano() / 1e6
	timeSince := int64(15) // 15s
	var dnsid uint16 = 0x2021
	tmpFreeport, err := freeport.GetFreePort()
	if err != nil {
		gologger.Fatalf("freeport error:" + err.Error())
		return
	}
	handle, err := core.PcapInit(ether.Device)
	if err != nil {
		gologger.Fatalf("初始化pcap失败,error:" + err.Error())
		return
	}
	var now int64
	for {
		send("seebug.org", "8.8.8.8", ether, dnsid, uint16(tmpFreeport), handle)
		index++
		now = time.Now().UnixNano() / 1e6
		tickTime := (now - start) / 1000
		if tickTime >= timeSince {
			break
		}
		if (now-start)%1000 == 0 && now-start >= 900 {
			tickIndex := index / tickTime
			gologger.Printf("\r %ds 总发送:%d Packet 平均每秒速度:%dpps", tickTime, index, tickIndex)
		}
	}
	now = time.Now().UnixNano() / 1e6
	tickTime := (now - start) / 1000
	tickIndex := index / tickTime
	gologger.Printf("\r %ds 总发送:%d Packet 平均每秒速度:%dpps\n", tickTime, index, tickIndex)
}