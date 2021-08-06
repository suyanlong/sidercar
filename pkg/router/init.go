package router

import (
	"context"
	"github.com/link33/sidercar/pkg/port"
	appchainmgr "github.com/meshplus/bitxhub-core/appchain-mgr"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"

	"github.com/meshplus/bitxhub-kit/storage"
)

// 程序动态注入与随机的删除。
var routerMap map[string]port.Port

func Register(id string, port port.Port) {
	routerMap[id] = port
}

type router struct {
	logger logrus.FieldLogger
	ctx    context.Context
	cancel context.CancelFunc
	store  storage.Storage
}

func (r router) Start() error {
	panic("implement me")
}

func (r router) Stop() error {
	panic("implement me")
}

func (r router) Broadcast(ids []string) error {
	panic("implement me")
}

// TODO 路由规则、路由优先级
// 一个连接一个goroutine
func (r router) Route(ibtp *pb.IBTP) error {
	_,to := ibtp.From,ibtp.To
	if p, is := routerMap[to];is {
		return p.Send(ibtp)
	}else {
		r.firstRoute(ibtp)
		return nil
	}
	panic("implement me")
}

func (r router) firstRoute(ibtp *pb.IBTP)  {
	panic("implement me")
}

func (r router) ExistAppchain(id string) bool {
	panic("implement me")
}

func (r router) AddAppchains(appchains []*appchainmgr.Appchain) error {
	panic("implement me")
}

func (r router) InPut(ibtp *pb.IBTP) chan *pb.IBTP {
	panic("implement me")
}

func (r router) OutPut(ibtp *pb.IBTP) chan *pb.IBTP {
	panic("implement me")
}
