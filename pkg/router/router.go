package router

import (
	appchainmgr "github.com/meshplus/bitxhub-core/appchain-mgr"
	"github.com/meshplus/bitxhub-model/pb"
)

//go:generate mockgen -destination mock_router/mock_router.go -package mock_router -source router.go
type Router interface {
	// Start starts the router module
	Start() error

	// Stop
	Stop() error

	//Broadcast broadcasts the registered appchain ids to the union network
	Broadcast(ids []string) error

	//Route sends ibtp to the union pier in target relay chain
	Route(ibtp *pb.IBTP) error

	//ExistAppchain returns if appchain id exist in route map
	ExistAppchain(id string) bool

	//AddAppchains adds appchains to route map and broadcast them to union network
	AddAppchains(appchains []*appchainmgr.Appchain) error

	//异步返回
	InPut(ibtp *pb.IBTP) chan *pb.IBTP
	OutPut(ibtp *pb.IBTP) chan *pb.IBTP
	//代替上面两个方法
	//Router(ibtp *pb.IBTP) chan *pb.IBTP
}

//路由规则，根据路由表，rules接口， 这个放到交易内部。可以让用户决定。并且签名。
// 1、广播
// 2、单一路由，指定路由标识。
// 3、p2p模式的路由，
// 4、中继路由
// 5、定制化：

//验证规则
//1、收集多个节点
//2、收集指定节点
//3、收集共识节点
// 即满足多个条件

//路由规则优先级：
//1、用户交易内部的路由规则最高。
//2、用户在程序设定。
//3、系统默认。
