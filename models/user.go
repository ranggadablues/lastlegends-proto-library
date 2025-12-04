package models

import (
	"time"

	"github.com/ranggadablues/lastlegends-proto-library/user-proto/pb"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	Id        bson.ObjectID `bson:"_id" json:"id"`
	Firstname string        `bson:"firstname" json:"firstname"`
	Lastname  string        `bson:"lastname" json:"lastname"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"password"`
	AccessKey string        `bson:"accesskey" json:"accesskey"`
	Isactive  bool          `bson:"isactive" json:"isactive"`
	Createby  string        `bson:"createby" json:"createby"`
	CreatedAt time.Time     `bson:"createdat" json:"createdat"`
	Updateby  string        `bson:"updateby" json:"updateby"`
	UpdatedAt time.Time     `bson:"updatedat" json:"updatedat"`
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:        u.Id.Hex(),
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Password:  u.Password,
		Accesskey: u.AccessKey,
		Isactive:  u.Isactive,
		Createdat: timestamppb.New(u.CreatedAt),
		Updatedat: timestamppb.New(u.UpdatedAt),
	}
}

func (u *User) FromProto(user *pb.User) {
	u.Id, _ = bson.ObjectIDFromHex(user.Id)
	u.Firstname = user.Firstname
	u.Lastname = user.Lastname
	u.Email = user.Email
	u.Password = user.Password
	u.AccessKey = user.Accesskey
	u.Isactive = user.Isactive
	u.CreatedAt = user.Createdat.AsTime()
	u.UpdatedAt = user.Updatedat.AsTime()
}

func (u *User) Collection() string {
	return "users"
}
