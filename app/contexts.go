// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "adder": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/shogo82148/goa-adder/design
// --out=$(GOPATH)/src/github.com/shogo82148/goa-adder
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// DeleteOperandsContext provides the operands delete action context.
type DeleteOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewDeleteOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller delete action.
func NewDeleteOperandsContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteOperandsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// PatchOperandsContext provides the operands patch action context.
type PatchOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewPatchOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller patch action.
func NewPatchOperandsContext(ctx context.Context, r *http.Request, service *goa.Service) (*PatchOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PatchOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *PatchOperandsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// PutOperandsContext provides the operands put action context.
type PutOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewPutOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller put action.
func NewPutOperandsContext(ctx context.Context, r *http.Request, service *goa.Service) (*PutOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PutOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *PutOperandsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
