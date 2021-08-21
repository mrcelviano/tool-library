package dbrsession

import (
	"context"
	"github.com/gocraft/dbr"
)

type key int64

const (
	PostgreSessionKey key = iota
)

func newContext(ctx context.Context, key key, db *dbr.Connection) context.Context {
	return context.WithValue(ctx, key, db.NewSession(&eventReceiver{}))
}

func NewContext(ctx context.Context, db *dbr.Connection) context.Context {
	return newContext(ctx, PostgreSessionKey, db)
}

func fromContext(ctx context.Context, key key) dbr.SessionRunner {
	if ctx == nil {
		panic("context is nil")
	}
	if val, ok := ctx.Value(key).(dbr.SessionRunner); ok {
		return val
	}
	panic("session is not defined in context")
}

func FromContext(ctx context.Context) dbr.SessionRunner {
	return fromContext(ctx, PostgreSessionKey)
}
