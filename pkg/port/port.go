package port

import "github.com/meshplus/bitxhub-model/pb"


// port 类型：主要是sider peer、plugin、blockchain peer。
type Type int

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

type port struct {}

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

