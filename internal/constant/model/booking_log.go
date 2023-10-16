package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CheckIn  = "CHECK_IN"
	CheckOut = "CHECK_OUT"
)

type BookingLog struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BookingLogID      string             `bson:"booking_log_id,omitempty" json:"booking_log_id,omitempty"`
	StaffID           string             `bson:"staff_id,omitempty" json:"staff_id,omitempty"`
	ExpiredInviteCode bool               `bson:"expired_invite_code,omitempty" json:"expired_invite_code,omitempty"`
	IsDefaulted       bool               `bson:"is_defaulted,omitempty" json:"is_defaulted,omitempty"`
	Type              string             `bson:"type,omitempty" json:"type,omitempty"`
	InviteID          string             `bson:"invite_id,omitempty" json:"invite_id,omitempty"`
	InviteCode        string             `bson:"-" json:"-,omitempty"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

func (es BookingLog) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.StaffID, validation.Required.Error("staff_id is required")),
		validation.Field(&es.Type, validation.Required.Error("type is required")),
		validation.Field(&es.Type, validation.In(CheckIn, CheckOut)),
	)
}
