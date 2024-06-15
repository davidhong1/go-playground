package geecache

import "github.com/davidhong1/go-playground/day7-cache/geecache/geecachepb"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implmented by a peer
type PeerGetter interface {
	Get(in *geecachepb.Request, out *geecachepb.Response) error
}
