package application

type EventType int

const (
	StartUp EventType = iota
	ShutDown //hs:常量_事件类型_应用退出     
	DomReady //hs:常量_事件类型_DOM就绪     
)
