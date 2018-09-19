package tcp

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rocymp/atx-server/proto"
	"github.com/rocymp/zero"
)

type DMHandler struct {
	ss   *zero.SocketService
	addr string
}

func InitDMHandler(taddr string, interval int) (dm *DMHandler) {
	ss, err := zero.NewSocketService(taddr, interval)
	if err != nil {
		log.Printf("zero NewSocketService err %#v\n", err)
		return nil
	}
	dm = &DMHandler{
		ss:   ss,
		addr: taddr,
	}
	//注册服务
	ss.RegConnectHandler(dm.HandleConnect)
	ss.RegDisconnectHandler(dm.HandleDisconnect)
	ss.RegMessageHandler(dm.HandleMessage)

	return dm
}

func (dm *DMHandler) Run() {
	log.Println("Device Manger is Running on " + dm.addr)
	dm.ss.Serv()
}

func (dm *DMHandler) Tick() {
	timer := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-timer.C:
			log.Printf("[Tick] ConnsCount[%d]\n", dm.GetServer().GetConnsCount())
			msg := zero.NewMessage(23, []byte("from server"))
			dm.GetServer().Broadcast(msg)
		}
	}
}

func (dm *DMHandler) Stop() {
	log.Println("Device Manger is stoping. ")
	dm.ss.Stop("Stoping")
}

func (dm *DMHandler) GetServer() *zero.SocketService {
	return dm.ss
}

func (dm *DMHandler) HandleMessage(s *zero.Session, msg *zero.Message) {
	// 处理设备上报事件
	if msg.GetCMD() == int32(proto.UpReportMessage) {
		dis := make([]proto.DeviceInfo, 0)
		data := msg.GetData()

		err := json.Unmarshal([]byte(data), &dis)

		if err != nil {
			log.Printf("Unmarshal error %#v\n", err)
		}

		s.SetSetting("devices", dis)

		fmt.Printf("[Device Report] devices:[%#v]\n", dis)
	}
	// log.Printf("[MESSAGE]\tFrom:[%s] CMDID:[%d] DATA:[%s]\n", s.GetConn().GetName(), msg.GetCMD(), string(msg.GetData()))
}

func (dm *DMHandler) HandleDisconnect(s *zero.Session, err error) {
	log.Printf("[OFFLINE]\tFrom:[%s]\n", s.GetConn().GetName())
}

func (dm *DMHandler) HandleConnect(s *zero.Session) {
	log.Printf("[ONLINE]\tFrom:[%s]\n", s.GetConn().GetName())
}
