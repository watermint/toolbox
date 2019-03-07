package copy_ref

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_file"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type CopyRef struct {
	Raw           json.RawMessage
	CopyReference string `path:"copy_reference" json:"copy_reference"`
	Expires       string `path:"expires" json:"expires"`
}

type CopyRefGet struct {
	AsMemberId string                                        `json:"-"`
	PathRoot   interface{}                                   `json:"-"`
	OnError    func(annotation dbx_api.ErrorAnnotation) bool `json:"-"`
	OnEntry    func(ref CopyRef) bool                        `json:"-"`
}

func (z *CopyRefGet) Get(c *dbx_api.Context, path string) bool {
	p := struct {
		Path string `json:"path"`
	}{
		Path: path,
	}
	req := dbx_rpc.RpcRequest{
		AsMemberId: z.AsMemberId,
		Endpoint:   "files/copy_reference/get",
		PathRoot:   z.PathRoot,
		Param:      p,
	}
	res, ea, _ := req.Call(c)
	if ea.IsFailure() {
		return z.OnError(ea)
	}
	rj := gjson.Parse(res.Body)
	if !rj.Exists() {
		c.Log().Debug("unable to parse JSON", zap.String("body", res.Body))
		return z.OnError(dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     errors.New("unable to parse json data"),
		})
	}
	cr := CopyRef{}
	err := c.ParseModel(&cr, rj)
	if err != nil {
		c.Log().Debug("unable to parse model", zap.String("body", res.Body))
		return z.OnError(dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		})
	}
	return z.OnEntry(cr)
}

type CopyRefSave struct {
	AsMemberId string                                        `json:"-"`
	PathRoot   interface{}                                   `json:"-"`
	OnError    func(annotation dbx_api.ErrorAnnotation) bool `json:"-"`
	OnFolder   func(folder *dbx_file.Folder) bool            `json:"-"`
	OnFile     func(file *dbx_file.File) bool                `json:"-"`
}

func (z *CopyRefSave) Save(c *dbx_api.Context, ref CopyRef, path string) error {
	p := struct {
		CopyReference string `json:"copy_reference"`
		Path          string `json:"path"`
	}{
		CopyReference: ref.CopyReference,
		Path:          path,
	}
	ep := dbx_file.EntryParser{
		Logger:   c.Log(),
		OnError:  z.OnError,
		OnFolder: z.OnFolder,
		OnFile:   z.OnFile,
		OnDelete: func(deleted *dbx_file.Deleted) bool {
			c.Log().Debug("deleted file found", zap.Any("deleted", deleted))
			return true
		},
	}
	req := dbx_rpc.RpcRequest{
		AsMemberId: z.AsMemberId,
		PathRoot:   z.PathRoot,
		Endpoint:   "files/copy_reference/save",
		Param:      p,
	}
	res, ea, err := req.Call(c)
	if ea.IsFailure() {
		z.OnError(ea)
		return err
	}
	rj := gjson.Parse(res.Body)
	if !rj.Exists() {
		err = errors.New("unable to parse json data")
		c.Log().Debug("unable to parse JSON", zap.String("body", res.Body))
		z.OnError(dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		})
		return err
	}
	m := rj.Get("metadata")
	if !m.Exists() {
		c.Log().Debug("could not found `metadata`", zap.String("body", res.Body))
		err = errors.New("unable to parse metadata")
		z.OnError(dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		})
		return err
	}

	ep.Parse(m)
	return nil
}
