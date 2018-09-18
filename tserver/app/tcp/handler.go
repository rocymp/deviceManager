package tcp

import (
	"log"
	"time"

	"github.com/rocymp/zero"
)

type DMHandler struct {
	ss   *zero.SocketService
	addr string
}

func InitDMHandler(addr string, interval int) (dm *DMHandler) {
	ss, err := zero.NewSocketService(addr, interval)
	if err != nil {
		log.Printf("zero NewSocketService err %#v\n", err)
		return nil
	}
	dm = &DMHandler{
		ss:   ss,
		addr: addr,
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
	log.Printf("[MESSAGE]\tFrom:[%s] CMDID:[%d] DATA:[%s]\n", s.GetConn().GetName(), msg.GetCMD(), string(msg.GetData()))
}

func (dm *DMHandler) HandleDisconnect(s *zero.Session, err error) {
	log.Printf("[OFFLINE]\tFrom:[%s]\n", s.GetConn().GetName())
}

func (dm *DMHandler) HandleConnect(s *zero.Session) {
	log.Printf("[ONLINE]\tFrom:[%s]\n", s.GetConn().GetName())
}
