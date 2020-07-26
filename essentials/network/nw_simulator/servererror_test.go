package nw_simulator

import (
	"net/http"
	"testing"
)

func TestServerErrorClient_Call(t *testing.T) {
	{
		nc := NewServerError(&PanicClient{}, 100, http.StatusInternalServerError, NoDecorator)
		res := nc.Call(&MockApiContext{}, &MockReqBuilder{})
		if res.IsSuccess() {
			t.Error(res.IsSuccess())
		}
		if res.Code() != http.StatusInternalServerError {
			t.Error(res.Code())
		}
	}
	{
		nc := NewServerError(&PanicClient{}, 100, http.StatusBadGateway, NoDecorator)
		res := nc.Call(&MockApiContext{}, &MockReqBuilder{})
		if res.IsSuccess() {
			t.Error(res.IsSuccess())
		}
		if res.Code() != http.StatusBadGateway {
			t.Error(res.Code())
		}
	}

	{
		// status code should be fallback to 500
		nc := NewServerError(&PanicClient{}, 100, http.StatusOK, NoDecorator)
		res := nc.Call(&MockApiContext{}, &MockReqBuilder{})
		if res.IsSuccess() {
			t.Error(res.IsSuccess())
		}
		if res.Code() != http.StatusInternalServerError {
			t.Error(res.Code())
		}
	}
}
