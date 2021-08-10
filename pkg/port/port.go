package port

import "github.com/meshplus/bitxhub-model/pb"

// port 类型：主要是sider peer、plugin、blockchain peer。
type Type int

const (
	Hub          Type = iota //hub: 同步数据，同步元数据等。
	SiderCarPeer             //SiderCar节点
	Client                   //其它区块链客户端
	Unknown                  //未知
)


// 设计一套port管理机制：包括各种的管理模块。以组合的行驶。
// 设计一套，管理机制。
// 与中继交互的是单独完整的机制。并且注册到路由表中。或者更加类型，这样就限制一个pier最多只能连接一个hub。避免网络风暴。或者只是一个转发功能。转发到指定节点。
// 先是从转发开始完成。
// 协议实现
// 路由策略
// 客户度
// 验证器（上链）


type MangerPort struct {
	// peer manger
	// sidercar manger
}

type Manger interface {
	Start()
	Close()
	Remove()
	Add()
	Query()
}

// 代表每一个路由端点
// 对router来说，只需要体现两个作用：1、did 唯一标识，2：接受一个ibtp数据函数。Send、recive（不对外开放）
// client 可是代表是sdk rpc 这些东西。
// 是否要做一个管理层，管理整个port.以及plugin。
// Port
type Port interface {
	ID() string
	Type() Type
	Name() string

	Send(ibtp *pb.IBTP) error
}

type Plugin interface {
}

type plugin struct {
}

func NewPlugin() plugin {
	return plugin{}
}

type peer struct {
}

func NewPeer() port {
	return port{}
}

type port struct{}

func (p port) ID() string {
	panic("implement me")
}

func (p port) Send(ibtp *pb.IBTP) error {
	panic("implement me")
}

func (p port) Type() Type {
	panic("implement me")
}

func (p port) Name() string {
	panic("implement me")
}
