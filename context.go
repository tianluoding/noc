package noc

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Context struct {
	R *http.Request
	W http.ResponseWriter
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{R: r, W: w}
}

func (ctx *Context) ReadJSON(data interface{}) error {
	err := json.NewDecoder(ctx.R.Body).Decode(&data)
	if err != nil {
		return errors.New("read json error")
	}
	defer ctx.R.Body.Close()
	return nil
}

func (ctx *Context) WriteJSON(code int, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx.W.Header().Set("Content-Type", "application/json")
	ctx.W.WriteHeader(code)
	_, err = ctx.W.Write(bs)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) BadRequestJSON() error {
	return ctx.WriteJSON(http.StatusBadRequest, nil)
}
