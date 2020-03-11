package models

import (
	"time"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/types"
)

//HelloLog :
type HelloLog struct {
	ID        types.NullInt    `json:"id" field:"id" type:"int" increment:"auto" pkey:"true"`
	Message   types.NullString `json:"message" field:"message" type:"string"`
	Origin    types.NullString `json:"origin" field:"origin" type:"string"`
	Timestamp types.NullTime   `json:"timestamp" field:"timestamp" type:"datetime"`
	IsSent    types.NullBool   `json:"is_sent" field:"is_sent" type:"boolean"`
}

//GetTableName :
func (l *HelloLog) GetTableName() string {
	return "hello_log"
}

//SetID :
func (l *HelloLog) SetID(val int64) *HelloLog {
	l.ID = types.NewNullInt(val)
	return l
}

//SetMessage :
func (l *HelloLog) SetMessage(val string) *HelloLog {
	l.Message = types.NewNullString(val)
	return l
}

//SetOrigin :
func (l *HelloLog) SetOrigin(val string) *HelloLog {
	l.Origin = types.NewNullString(val)
	return l
}

//SetTimestamp :
func (l *HelloLog) SetTimestamp(val time.Time) *HelloLog {
	l.Timestamp = types.NewNullTime(val)
	return l
}

//SetIsSent :
func (l *HelloLog) SetIsSent(val bool) *HelloLog {
	l.IsSent = types.NewNullBool(val)
	return l
}
