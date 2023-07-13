package geecache

import "geecache/geecachepb"

// PeerPicker is the interface that must be implemented to
// locate the peer that owns a specific key
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer
type PeerGetter interface {
	Get(in *geecachepb.Request, out *geecachepb.Response) error // 使用 pb 文件创建的结构体
}
