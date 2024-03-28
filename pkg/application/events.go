package application

type EventType int

const (
	X常量_事件类型_启动前 EventType = iota
	X常量_事件类型_应用退出 //hs:常量_事件类型_应用退出     
	X常量_事件类型_DOM就绪 //hs:常量_事件类型_DOM就绪     
)
