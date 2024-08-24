package rc_value

import (
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
	"strings"
)

func NewValueConn(acceptType interface{}, factory func(peerName string) api_conn.Connection) rc_recipe.Value {
	return &ValueConnBase{
		peerName:   api_conn.DefaultPeerName,
		acceptType: reflect.TypeOf(acceptType).Elem(),
		factory:    factory,
		conn:       factory(api_conn.DefaultPeerName),
	}
}

type ValueConnBase struct {
	peerName   string
	conn       api_conn.Connection
	acceptType reflect.Type
	factory    func(peerName string) api_conn.Connection
}

func (z *ValueConnBase) ValueText() string {
	return z.peerName
}

func (z *ValueConnBase) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(z.acceptType) {
		return &ValueConnBase{
			peerName:   z.peerName,
			conn:       z.factory(z.peerName),
			acceptType: z.acceptType,
			factory:    z.factory,
		}
	}
	return nil
}

func (z *ValueConnBase) Bind() interface{} {
	return &z.peerName
}

func (z *ValueConnBase) Init() (v interface{}) {
	return z.conn
}

func (z *ValueConnBase) ApplyPreset(v0 interface{}) {
	z.conn = v0.(api_conn.Connection)
	z.peerName = z.conn.PeerName()
}

func (z *ValueConnBase) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueConnBase) Debug() interface{} {
	if s, ok := z.conn.(api_conn.ScopedConnection); ok {
		return map[string]string{
			"peerName":   z.peerName,
			"scopeLabel": s.ScopeLabel(),
			"scopes":     strings.Join(s.Scopes(), ","),
		}
	}
	return map[string]string{
		"peerName":   z.peerName,
		"scopeLabel": z.conn.ScopeLabel(),
	}
}

func (z *ValueConnBase) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueConnBase) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = z.factory(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueConnBase) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueConnBase) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueConnBase) Spec() (typeName string, typeAttr interface{}) {
	typeName = es_reflect.Key(z.conn)
	if s, ok := z.conn.(api_conn.ScopedConnection); ok {
		typeAttr = s.Scopes()
	} else {
		typeAttr = z.conn.ScopeLabel()
	}
	return
}

func (z *ValueConnBase) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}
