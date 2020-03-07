package models

import (
	"time"
	"gopkg.in/guregu/null.v3"
)

//NullInt : Integer Type with Null
type NullInt struct {
	null.Int
}

//NewNullInt : 
func NewNullInt(val int64) NullInt {
	return NullInt{null.NewInt(val,true)}
}

//NullFloat : Float Type with Null
type NullFloat struct {
	null.Float
}

//NewNullFloat : 
func NewNullFloat(val float64) NullFloat {
	return NullFloat{null.NewFloat(val,true)}
}

//NullBool : Boolean Type with Null
type NullBool struct {
	null.Bool
}

//NewNullBool : 
func NewNullBool(val bool) NullBool {
	return NullBool{null.NewBool(val,true)}
}

//NullString : String Type with Null
type NullString struct {
	null.String
}

//NewNullString : 
func NewNullString(val string) NullString {
	return NullString{null.NewString(val,true)}
}

//NullTime : DateTime Type with Null
type NullTime struct {
	null.Time
}

//NewNullTime : 
func NewNullTime(val time.Time) NullTime {
	return NullTime{null.NewTime(val,true)}
}
