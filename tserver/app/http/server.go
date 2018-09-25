package http

import (
	"encoding/json"
	"log"

	"github.com/rocymp/zero"

	"github.com/labstack/echo"
	"github.com/rocymp/atx-server/proto"
	"github.com/rocymp/deviceManager/tserver/app/common"
	"github.com/rocymp/deviceManager/tserver/app/model"
	"github.com/rocymp/deviceManager/tserver/app/tcp"
)

var (
	DM *tcp.DMHandler
)

func Init(dm *tcp.DMHandler) {
	DM = dm
	e := echo.New()
	e.GET("/api/v1/listhost", ListHost)
	e.GET("/api/v1/listdevice", ListDevice)
	e.POST("/api/v1/startbyhost", StartByHost)
	e.POST("/api/v1/stopbyhost", StopByHost)
	log.Println(e.Start("0.0.0.0:19999"))
}

func ListHost(c echo.Context) error {
	sessions := DM.GetServer().GetSession()

	hlist := make([]model.HostInfo, 0)
	for _, se := range sessions {
		sdmap := se.GetSetting(common.DEVICEKEY)

		haddr := se.GetConn().GetName()
		dcount := len(*sdmap.(*map[string]proto.DeviceInfo))
		uuid := se.GetSessionID()
		hinfo := model.HostInfo{
			Address: haddr,
			UUID:    uuid,
			DNum:    dcount,
		}

		hlist = append(hlist, hinfo)
	}

	return common.JSONE(c, "", common.StatusOK, hlist)
}

func ListDeviceByHost(c echo.Context) error {
	sessions := DM.GetServer().GetSession()

	hlist := make([]model.HostInfo, 0)
	for _, se := range sessions {
		sdmap := se.GetSetting(common.DEVICEKEY)

		haddr := se.GetConn().GetName()
		dcount := len(*sdmap.(*map[string]proto.DeviceInfo))
		uuid := se.GetSessionID()
		hinfo := model.HostInfo{
			Address: haddr,
			UUID:    uuid,
			DNum:    dcount,
		}

		hlist = append(hlist, hinfo)
	}

	return common.JSONE(c, "", common.StatusOK, hlist)
}

func ListDevice(c echo.Context) error {
	sessions := DM.GetServer().GetSession()

	hlist := make([]proto.DeviceInfo, 0)
	for _, se := range sessions {
		sdmap := se.GetSetting(common.DEVICEKEY)
		for _, d := range *sdmap.(*map[string]proto.DeviceInfo) {
			hlist = append(hlist, d)
		}

	}

	return common.JSONE(c, "", common.StatusOK, hlist)
}

func StartByHost(c echo.Context) error {
	u := new(model.StartByHostReq)
	if err := c.Bind(u); err != nil {
		return common.JSONE(c, "参数错误", common.ParamsErr, nil)
	}

	rmsg := proto.RoomMessage{
		Rid:     u.Rid,
		Command: proto.StartRoom,
	}

	sessions := DM.GetServer().GetSession()

	b, err := json.Marshal(rmsg)
	if err != nil {
		return common.JSONE(c, "参数错误", common.ParamsErr, nil)
	}
	msg := zero.NewMessage(int32(proto.StartRoomMessage), b)
	for _, se := range sessions {
		se.GetConn().SendMessage(msg)
	}

	return common.JSONE(c, "启动成功", common.StatusOK, nil)
}

func StopByHost(c echo.Context) error {
	u := new(model.StartByHostReq)
	if err := c.Bind(u); err != nil {
		return common.JSONE(c, "参数错误", common.ParamsErr, nil)
	}

	rmsg := proto.RoomMessage{
		Rid:     u.Rid,
		Command: proto.StopRoom,
	}

	sessions := DM.GetServer().GetSession()

	b, err := json.Marshal(rmsg)
	if err != nil {
		return common.JSONE(c, "参数错误", common.ParamsErr, nil)
	}
	msg := zero.NewMessage(int32(proto.StartRoomMessage), b)
	for _, se := range sessions {
		se.GetConn().SendMessage(msg)
	}

	return common.JSONE(c, "启动成功", common.StatusOK, nil)
}
