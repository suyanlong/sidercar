package plugins

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/meshplus/bitxhub-model/pb"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var (
	Handshake = plugin.HandshakeConfig{
		ProtocolVersion:  4,
		MagicCookieKey:   "PIER_APPCHAIN_PLUGIN",
		MagicCookieValue: "PIER",
	}
	PluginName = "appchain-plugin"
)

// 插件进程在启动时设置Plugins，即ServeConfig中设置Plugins时，会指明其实现者；
// 宿主机进程在启动时也设置Plugins，即ClientConfig中设置Plugins时，不需要指明其实现者。
// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	PluginName: &AppchainGRPCPlugin{}, //宿主机进程的插件集
}

// 插件实现
// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type AppchainGRPCPlugin struct {
	plugin.Plugin //要嵌入（net/rpc）插件接口，是反射的作用。整个接口体需要实现：GRPCPlugin、plugin.Plugin两个接口，而plugin.Plugin并没有使用，嵌入进去即可。
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Client
}

// GRPCPlugin实现
func (p *AppchainGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterAppchainPluginServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *AppchainGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client:      pb.NewAppchainPluginClient(c), //创建gRPC客户端的方法是自动生成的
		doneContext: ctx,
	}, nil
}

//type GRPCPlugin interface {
//	// 此方法被插件进程调用
//	// 你需要向其提供的grpc.ServergRPC参数，注册服务的实现（服务器端存根）
//	// 由于gRPC下服务器端是单例模式，因此该方法仅调用一次
//	GRPCServer(*GRPCBroker, *grpc.Server) error
//
//	// 此方法被宿主进程调用
//	// 你需要返回一个业务接口的实现（客户端存根），此实现直接将请求转给gRPC客户端即可
//	// 传入的context对象会在插件进程销毁时取消
//	GRPCClient(context.Context, *GRPCBroker, *grpc.ClientConn) (interface{}, error)
//}
