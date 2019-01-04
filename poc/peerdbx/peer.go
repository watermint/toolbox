package peerdbx

import (
	"go.uber.org/zap"
)

type Session interface {
	AsMemberId(teamMemberId string) Session
	AsAdminId(teamMemberId string) Session
	Log() *zap.Logger
	Rpc(endpoint string) RpcRequest
}

type RpcRequest interface {
	WithParam(p interface{}) RpcRequest
	NoAuthHeader() RpcRequest
}

type sessionImpl struct {
	asMemberId string
	asAdminId  string
	log        *zap.Logger
	token      string
}

func (z *sessionImpl) AsMemberId(teamMemberId string) Session {
	return &sessionImpl{
		asMemberId: teamMemberId,
		asAdminId:  z.asAdminId,
		log:        z.log,
		token:      z.token,
	}
}

func (z *sessionImpl) AsAdminId(teamMemberId string) Session {
	return &sessionImpl{
		asMemberId: z.asMemberId,
		asAdminId:  teamMemberId,
		log:        z.log,
		token:      z.token,
	}
}

func (z *sessionImpl) Log() *zap.Logger {
	return z.log
}

func (z *sessionImpl) Rpc(endpoint string) RpcRequest {
	return &rpcRequestImpl{
		token: z.token,
	}
}

type rpcRequestImpl struct {
	token        string
	asMemberId   string
	asAdminId    string
	endpoint     string
	noAuthHeader bool
	param        interface{}
}

func (z *rpcRequestImpl) WithParam(p interface{}) RpcRequest {
	panic("implement me")
}

func (z *rpcRequestImpl) NoAuthHeader() RpcRequest {
	panic("implement me")
}
