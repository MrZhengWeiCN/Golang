package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connector //嵌入接口
}

type Connector interface {
	Connector()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connector() {
	fmt.Println("Connect", pc.name)
}

func main() {
	a := PhoneConnector{"PhoneConnector"}
	a.Connector()
	//不需要显示的申明实现了某一接口，如果实现了接口的所有方法，就默认实现了此接口
	DisConnect(a)
}

func DisConnect(usb interface{}) {
	//由系统去猜测是什么类型
	//typeswitch
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("DisConnect:", v.name)
	default:
		fmt.Println("Unknown Divice.")
	}
	/*if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("DisConnect:", pc.name)
		return
	}
	fmt.Println("Unknown Divice.")*/
}
