package peermgr

import (
	"github.com/libp2p/go-libp2p-core/peer"
	peermgr "github.com/link33/sidercar/internal/peermgr/proto"
	network "github.com/meshplus/go-lightp2p"
)

type MessageHandler func(network.Stream, *peermgr.Message)
type ConnectHandler func(string)

//go:generate mockgen -destination mock_peermgr/mock_peermgr.go -package mock_peermgr -source peermgr.go
type PeerManager interface {
	DHTManager

	// Start
	Start() error

	// Stop
	Stop() error

	// AsyncSend sends message to peer with peer info.
	AsyncSend(string, *peermgr.Message) error

	Connect(info *peer.AddrInfo) (string, error)

	// SendWithStream sends message using existed stream
	SendWithStream(network.Stream, *peermgr.Message) (*peermgr.Message, error)

	// AsyncSendWithStream sends message using existed stream
	AsyncSendWithStream(network.Stream, *peermgr.Message) error

	// Send sends message waiting response
	Send(string, *peermgr.Message) (*peermgr.Message, error)

	// Peers
	Peers() map[string]*peer.AddrInfo

	// RegisterMsgHandler
	RegisterMsgHandler(peermgr.Message_Type, MessageHandler) error

	// RegisterMultiMsgHandler
	RegisterMultiMsgHandler([]peermgr.Message_Type, MessageHandler) error

	// RegisterConnectHandler
	RegisterConnectHandler(ConnectHandler) error
}

type DHTManager interface {
	// Search for peers who are able to provide a given key
	FindProviders(id string) (string, error)

	// Provide adds the given cid to the content routing system. If 'true' is
	// passed, it also announces it, otherwise it is just kept in the local
	// accounting of which objects are being provided.
	Provider(string, bool) error
}
