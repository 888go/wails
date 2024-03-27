package application

type EventType int

const (
	StartUp  EventType = iota //hs:常量_事件类型_启动前
	ShutDown                  //hs:常量_事件类型_应用退出
	DomReady                  //hs:常量_事件类型_DOM就绪
)
