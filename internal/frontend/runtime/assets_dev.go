//go:build dev

package runtime

var RuntimeAssetsBundle = &RuntimeAssets{
	desktopIPC:       DesktopIPC,
	websocketIPC:     WebsocketIPC,
	runtimeDesktopJS: RuntimeDesktopJS,
}

type RuntimeAssets struct {
	desktopIPC       []byte
	websocketIPC     []byte
	runtimeDesktopJS []byte
}


// ff:
func (r *RuntimeAssets) DesktopIPC() []byte {
	return r.desktopIPC
}


// ff:
func (r *RuntimeAssets) WebsocketIPC() []byte {
	return r.websocketIPC
}


// ff:
func (r *RuntimeAssets) RuntimeDesktopJS() []byte {
	return r.runtimeDesktopJS
}
