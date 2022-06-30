package srpc

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
)

// NewClientWithMuxedConn constructs a new client with a MuxedConn.
func NewClientWithMuxedConn(conn network.MuxedConn) Client {
	openStreamFn := NewOpenStreamWithMuxedConn(conn)
	return NewClient(openStreamFn)
}

// NewOpenStreamWithMuxedConn constructs a OpenStream func with a MuxedConn.
func NewOpenStreamWithMuxedConn(conn network.MuxedConn) OpenStreamFunc {
	return func(ctx context.Context, msgHandler PacketHandler) (Writer, error) {
		mstrm, err := conn.OpenStream(ctx)
		if err != nil {
			return nil, err
		}
		rw := NewPacketReadWriter(mstrm)
		go func() {
			err := rw.ReadPump(msgHandler)
			if err != nil {
				_ = rw.Close()
			}
		}()
		return rw, nil
	}
}