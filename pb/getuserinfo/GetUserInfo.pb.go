// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: GetUserInfo.proto

package getuserinfo

import (
	_ "github.com/PterX/go-futu-api/pb/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateType int32

const (
	UpdateType_UpdateType_None   UpdateType = 0 //无需升级
	UpdateType_UpdateType_Advice UpdateType = 1 //建议升级
	UpdateType_UpdateType_Force  UpdateType = 2 //强制升级
)

// Enum value maps for UpdateType.
var (
	UpdateType_name = map[int32]string{
		0: "UpdateType_None",
		1: "UpdateType_Advice",
		2: "UpdateType_Force",
	}
	UpdateType_value = map[string]int32{
		"UpdateType_None":   0,
		"UpdateType_Advice": 1,
		"UpdateType_Force":  2,
	}
)

func (x UpdateType) Enum() *UpdateType {
	p := new(UpdateType)
	*p = x
	return p
}

func (x UpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_GetUserInfo_proto_enumTypes[0].Descriptor()
}

func (UpdateType) Type() protoreflect.EnumType {
	return &file_GetUserInfo_proto_enumTypes[0]
}

func (x UpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UpdateType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UpdateType(num)
	return nil
}

// Deprecated: Use UpdateType.Descriptor instead.
func (UpdateType) EnumDescriptor() ([]byte, []int) {
	return file_GetUserInfo_proto_rawDescGZIP(), []int{0}
}

type UserInfoField int32

const (
	UserInfoField_UserInfoField_Basic      UserInfoField = 1  //昵称，用户头像，牛牛号
	UserInfoField_UserInfoField_API        UserInfoField = 2  //API权限信息
	UserInfoField_UserInfoField_QotRight   UserInfoField = 4  //市场的行情权限
	UserInfoField_UserInfoField_Disclaimer UserInfoField = 8  //免责
	UserInfoField_UserInfoField_Update     UserInfoField = 16 //升级类型
	UserInfoField_UserInfoField_WebKey     UserInfoField = 2048
)

// Enum value maps for UserInfoField.
var (
	UserInfoField_name = map[int32]string{
		1:    "UserInfoField_Basic",
		2:    "UserInfoField_API",
		4:    "UserInfoField_QotRight",
		8:    "UserInfoField_Disclaimer",
		16:   "UserInfoField_Update",
		2048: "UserInfoField_WebKey",
	}
	UserInfoField_value = map[string]int32{
		"UserInfoField_Basic":      1,
		"UserInfoField_API":        2,
		"UserInfoField_QotRight":   4,
		"UserInfoField_Disclaimer": 8,
		"UserInfoField_Update":     16,
		"UserInfoField_WebKey":     2048,
	}
)

func (x UserInfoField) Enum() *UserInfoField {
	p := new(UserInfoField)
	*p = x
	return p
}

func (x UserInfoField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInfoField) Descriptor() protoreflect.EnumDescriptor {
	return file_GetUserInfo_proto_enumTypes[1].Descriptor()
}

func (UserInfoField) Type() protoreflect.EnumType {
	return &file_GetUserInfo_proto_enumTypes[1]
}

func (x UserInfoField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UserInfoField) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UserInfoField(num)
	return nil
}

// Deprecated: Use UserInfoField.Descriptor instead.
func (UserInfoField) EnumDescriptor() ([]byte, []int) {
	return file_GetUserInfo_proto_rawDescGZIP(), []int{1}
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag *int32 `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"` //UserInfoField集合，不设置默认返回全部信息
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetUserInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_GetUserInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_GetUserInfo_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetFlag() int32 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName              *string `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`                            //用户昵称
	AvatarUrl             *string `protobuf:"bytes,2,opt,name=avatarUrl" json:"avatarUrl,omitempty"`                          //用户头像url
	ApiLevel              *string `protobuf:"bytes,3,opt,name=apiLevel" json:"apiLevel,omitempty"`                            //api用户等级描述, 已在2.10版本之后废弃
	HkQotRight            *int32  `protobuf:"varint,4,opt,name=hkQotRight" json:"hkQotRight,omitempty"`                       //港股行情权限, QotRight
	UsQotRight            *int32  `protobuf:"varint,5,opt,name=usQotRight" json:"usQotRight,omitempty"`                       //美股行情权限, QotRight
	CnQotRight            *int32  `protobuf:"varint,6,opt,name=cnQotRight" json:"cnQotRight,omitempty"`                       //A股行情权限, QotRight // 废弃，使用shQotRight和szQotRight
	IsNeedAgreeDisclaimer *bool   `protobuf:"varint,7,opt,name=isNeedAgreeDisclaimer" json:"isNeedAgreeDisclaimer,omitempty"` //已开户用户需要同意免责声明，未开户或已同意的用户返回false
	UserID                *int64  `protobuf:"varint,8,opt,name=userID" json:"userID,omitempty"`                               //用户牛牛号
	UpdateType            *int32  `protobuf:"varint,9,opt,name=updateType" json:"updateType,omitempty"`                       //升级类型，UpdateType
	WebKey                *string `protobuf:"bytes,10,opt,name=webKey" json:"webKey,omitempty"`
	WebJumpUrlHead        *string `protobuf:"bytes,18,opt,name=webJumpUrlHead" json:"webJumpUrlHead,omitempty"`
	HkOptionQotRight      *int32  `protobuf:"varint,11,opt,name=hkOptionQotRight" json:"hkOptionQotRight,omitempty"`           //港股期权行情权限, Qot_Common.QotRight
	HasUSOptionQotRight   *bool   `protobuf:"varint,12,opt,name=hasUSOptionQotRight" json:"hasUSOptionQotRight,omitempty"`     //是否有美股期权行情权限
	HkFutureQotRight      *int32  `protobuf:"varint,13,opt,name=hkFutureQotRight" json:"hkFutureQotRight,omitempty"`           //港股期货行情权限, Qot_Common.QotRight
	SubQuota              *int32  `protobuf:"varint,14,opt,name=subQuota" json:"subQuota,omitempty"`                           //订阅额度
	HistoryKLQuota        *int32  `protobuf:"varint,15,opt,name=historyKLQuota" json:"historyKLQuota,omitempty"`               //历史K线额度
	UsFutureQotRight      *int32  `protobuf:"varint,16,opt,name=usFutureQotRight" json:"usFutureQotRight,omitempty"`           //美股期货行情权限, Qot_Common.QotRight(已废弃)
	UsOptionQotRight      *int32  `protobuf:"varint,17,opt,name=usOptionQotRight" json:"usOptionQotRight,omitempty"`           //美股期货行情权限, Qot_Common.QotRight
	UserAttribution       *int32  `protobuf:"varint,19,opt,name=userAttribution" json:"userAttribution,omitempty"`             //用户注册归属地, Common.UserAttribution
	UpdateWhatsNew        *string `protobuf:"bytes,20,opt,name=updateWhatsNew" json:"updateWhatsNew,omitempty"`                //升级提示
	UsIndexQotRight       *int32  `protobuf:"varint,21,opt,name=usIndexQotRight" json:"usIndexQotRight,omitempty"`             //美股指数行情权限, Qot_Common.QotRight
	UsOtcQotRight         *int32  `protobuf:"varint,22,opt,name=usOtcQotRight" json:"usOtcQotRight,omitempty"`                 //美股OTC市场行情权限, Qot_Common.QotRight
	UsCMEFutureQotRight   *int32  `protobuf:"varint,23,opt,name=usCMEFutureQotRight" json:"usCMEFutureQotRight,omitempty"`     //美股CME期货行情权限, Qot_Common.QotRight
	UsCBOTFutureQotRight  *int32  `protobuf:"varint,24,opt,name=usCBOTFutureQotRight" json:"usCBOTFutureQotRight,omitempty"`   //美股CBOT期货行情权限, Qot_Common.QotRight
	UsNYMEXFutureQotRight *int32  `protobuf:"varint,25,opt,name=usNYMEXFutureQotRight" json:"usNYMEXFutureQotRight,omitempty"` //美股NYMEX期货行情权限, Qot_Common.QotRight
	UsCOMEXFutureQotRight *int32  `protobuf:"varint,26,opt,name=usCOMEXFutureQotRight" json:"usCOMEXFutureQotRight,omitempty"` //美股COMEX期货行情权限, Qot_Common.QotRight
	UsCBOEFutureQotRight  *int32  `protobuf:"varint,27,opt,name=usCBOEFutureQotRight" json:"usCBOEFutureQotRight,omitempty"`   //美股CBOE期货行情权限, Qot_Common.QotRight
	SgFutureQotRight      *int32  `protobuf:"varint,28,opt,name=sgFutureQotRight" json:"sgFutureQotRight,omitempty"`           //新加坡市场期货行情权限, Qot_Common.QotRight
	JpFutureQotRight      *int32  `protobuf:"varint,29,opt,name=jpFutureQotRight" json:"jpFutureQotRight,omitempty"`           //日本市场期货行情权限, Qot_Common.QotRight
	IsAppNNOrMM           *bool   `protobuf:"varint,30,opt,name=isAppNNOrMM" json:"isAppNNOrMM,omitempty"`                     //true:NN false:MM
	ShQotRight            *int32  `protobuf:"varint,31,opt,name=shQotRight" json:"shQotRight,omitempty"`                       //上海市场行情权限, Qot_Common.QotRight
	SzQotRight            *int32  `protobuf:"varint,32,opt,name=szQotRight" json:"szQotRight,omitempty"`                       //深圳市场行情权限, Qot_Common.QotRight
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetUserInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_GetUserInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_GetUserInfo_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *S2C) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

func (x *S2C) GetApiLevel() string {
	if x != nil && x.ApiLevel != nil {
		return *x.ApiLevel
	}
	return ""
}

func (x *S2C) GetHkQotRight() int32 {
	if x != nil && x.HkQotRight != nil {
		return *x.HkQotRight
	}
	return 0
}

func (x *S2C) GetUsQotRight() int32 {
	if x != nil && x.UsQotRight != nil {
		return *x.UsQotRight
	}
	return 0
}

func (x *S2C) GetCnQotRight() int32 {
	if x != nil && x.CnQotRight != nil {
		return *x.CnQotRight
	}
	return 0
}

func (x *S2C) GetIsNeedAgreeDisclaimer() bool {
	if x != nil && x.IsNeedAgreeDisclaimer != nil {
		return *x.IsNeedAgreeDisclaimer
	}
	return false
}

func (x *S2C) GetUserID() int64 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *S2C) GetUpdateType() int32 {
	if x != nil && x.UpdateType != nil {
		return *x.UpdateType
	}
	return 0
}

func (x *S2C) GetWebKey() string {
	if x != nil && x.WebKey != nil {
		return *x.WebKey
	}
	return ""
}

func (x *S2C) GetWebJumpUrlHead() string {
	if x != nil && x.WebJumpUrlHead != nil {
		return *x.WebJumpUrlHead
	}
	return ""
}

func (x *S2C) GetHkOptionQotRight() int32 {
	if x != nil && x.HkOptionQotRight != nil {
		return *x.HkOptionQotRight
	}
	return 0
}

func (x *S2C) GetHasUSOptionQotRight() bool {
	if x != nil && x.HasUSOptionQotRight != nil {
		return *x.HasUSOptionQotRight
	}
	return false
}

func (x *S2C) GetHkFutureQotRight() int32 {
	if x != nil && x.HkFutureQotRight != nil {
		return *x.HkFutureQotRight
	}
	return 0
}

func (x *S2C) GetSubQuota() int32 {
	if x != nil && x.SubQuota != nil {
		return *x.SubQuota
	}
	return 0
}

func (x *S2C) GetHistoryKLQuota() int32 {
	if x != nil && x.HistoryKLQuota != nil {
		return *x.HistoryKLQuota
	}
	return 0
}

func (x *S2C) GetUsFutureQotRight() int32 {
	if x != nil && x.UsFutureQotRight != nil {
		return *x.UsFutureQotRight
	}
	return 0
}

func (x *S2C) GetUsOptionQotRight() int32 {
	if x != nil && x.UsOptionQotRight != nil {
		return *x.UsOptionQotRight
	}
	return 0
}

func (x *S2C) GetUserAttribution() int32 {
	if x != nil && x.UserAttribution != nil {
		return *x.UserAttribution
	}
	return 0
}

func (x *S2C) GetUpdateWhatsNew() string {
	if x != nil && x.UpdateWhatsNew != nil {
		return *x.UpdateWhatsNew
	}
	return ""
}

func (x *S2C) GetUsIndexQotRight() int32 {
	if x != nil && x.UsIndexQotRight != nil {
		return *x.UsIndexQotRight
	}
	return 0
}

func (x *S2C) GetUsOtcQotRight() int32 {
	if x != nil && x.UsOtcQotRight != nil {
		return *x.UsOtcQotRight
	}
	return 0
}

func (x *S2C) GetUsCMEFutureQotRight() int32 {
	if x != nil && x.UsCMEFutureQotRight != nil {
		return *x.UsCMEFutureQotRight
	}
	return 0
}

func (x *S2C) GetUsCBOTFutureQotRight() int32 {
	if x != nil && x.UsCBOTFutureQotRight != nil {
		return *x.UsCBOTFutureQotRight
	}
	return 0
}

func (x *S2C) GetUsNYMEXFutureQotRight() int32 {
	if x != nil && x.UsNYMEXFutureQotRight != nil {
		return *x.UsNYMEXFutureQotRight
	}
	return 0
}

func (x *S2C) GetUsCOMEXFutureQotRight() int32 {
	if x != nil && x.UsCOMEXFutureQotRight != nil {
		return *x.UsCOMEXFutureQotRight
	}
	return 0
}

func (x *S2C) GetUsCBOEFutureQotRight() int32 {
	if x != nil && x.UsCBOEFutureQotRight != nil {
		return *x.UsCBOEFutureQotRight
	}
	return 0
}

func (x *S2C) GetSgFutureQotRight() int32 {
	if x != nil && x.SgFutureQotRight != nil {
		return *x.SgFutureQotRight
	}
	return 0
}

func (x *S2C) GetJpFutureQotRight() int32 {
	if x != nil && x.JpFutureQotRight != nil {
		return *x.JpFutureQotRight
	}
	return 0
}

func (x *S2C) GetIsAppNNOrMM() bool {
	if x != nil && x.IsAppNNOrMM != nil {
		return *x.IsAppNNOrMM
	}
	return false
}

func (x *S2C) GetShQotRight() int32 {
	if x != nil && x.ShQotRight != nil {
		return *x.ShQotRight
	}
	return 0
}

func (x *S2C) GetSzQotRight() int32 {
	if x != nil && x.SzQotRight != nil {
		return *x.SzQotRight
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetUserInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_GetUserInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_GetUserInfo_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //返回结果，参见Common.RetType的枚举定义
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`             //返回结果描述
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`          //错误码，客户端一般通过retType和retMsg来判断结果和详情，errCode只做日志记录，仅在个别协议失败时对账用
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetUserInfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_GetUserInfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_GetUserInfo_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_GetUserInfo_proto protoreflect.FileDescriptor

var file_GetUserInfo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19,
	0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22, 0xf1, 0x09, 0x0a, 0x03, 0x53, 0x32,
	0x43, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6b, 0x51, 0x6f, 0x74,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x6b, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x51, 0x6f, 0x74,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6e, 0x51, 0x6f, 0x74,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6e, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x73, 0x4e, 0x65, 0x65,
	0x64, 0x41, 0x67, 0x72, 0x65, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x41, 0x67,
	0x72, 0x65, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a,
	0x0e, 0x77, 0x65, 0x62, 0x4a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x65, 0x62, 0x4a, 0x75, 0x6d, 0x70, 0x55, 0x72,
	0x6c, 0x48, 0x65, 0x61, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x68, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x68, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x30, 0x0a, 0x13, 0x68, 0x61, 0x73, 0x55, 0x53, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x68, 0x61, 0x73, 0x55, 0x53, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x68, 0x6b, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x68,
	0x6b, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x73, 0x75, 0x62, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x75,
	0x73, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x2a, 0x0a, 0x10, 0x75, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x75, 0x73, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x68, 0x61, 0x74, 0x73, 0x4e, 0x65, 0x77, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x68, 0x61, 0x74, 0x73, 0x4e, 0x65, 0x77, 0x12, 0x28, 0x0a,
	0x0f, 0x75, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x75, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x4f, 0x74, 0x63,
	0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x75, 0x73, 0x4f, 0x74, 0x63, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a,
	0x13, 0x75, 0x73, 0x43, 0x4d, 0x45, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x75, 0x73, 0x43, 0x4d,
	0x45, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x32, 0x0a, 0x14, 0x75, 0x73, 0x43, 0x42, 0x4f, 0x54, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x75,
	0x73, 0x43, 0x42, 0x4f, 0x54, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x75, 0x73, 0x4e, 0x59, 0x4d, 0x45, 0x58, 0x46, 0x75,
	0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x75, 0x73, 0x4e, 0x59, 0x4d, 0x45, 0x58, 0x46, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x75, 0x73, 0x43,
	0x4f, 0x4d, 0x45, 0x58, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x75, 0x73, 0x43, 0x4f, 0x4d, 0x45,
	0x58, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x32, 0x0a, 0x14, 0x75, 0x73, 0x43, 0x42, 0x4f, 0x45, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x75,
	0x73, 0x43, 0x42, 0x4f, 0x45, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x67, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51,
	0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x73,
	0x67, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x2a, 0x0a, 0x10, 0x6a, 0x70, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6a, 0x70, 0x46, 0x75, 0x74,
	0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x73, 0x41, 0x70, 0x70, 0x4e, 0x4e, 0x4f, 0x72, 0x4d, 0x4d, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x41, 0x70, 0x70, 0x4e, 0x4e, 0x4f, 0x72, 0x4d, 0x4d, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x68, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x68, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x7a, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x7a, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2d, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x80, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x73,
	0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a,
	0x4e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x10, 0x02, 0x2a,
	0xae, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x41, 0x50, 0x49, 0x10,
	0x02, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x10, 0x04, 0x12, 0x1c, 0x0a,
	0x18, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x44,
	0x69, 0x73, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x10, 0x10, 0x12, 0x19, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x10, 0x80, 0x10,
	0x42, 0x42, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x74, 0x65, 0x72, 0x58, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74,
	0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x74, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x6e, 0x66, 0x6f,
}

var (
	file_GetUserInfo_proto_rawDescOnce sync.Once
	file_GetUserInfo_proto_rawDescData = file_GetUserInfo_proto_rawDesc
)

func file_GetUserInfo_proto_rawDescGZIP() []byte {
	file_GetUserInfo_proto_rawDescOnce.Do(func() {
		file_GetUserInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetUserInfo_proto_rawDescData)
	})
	return file_GetUserInfo_proto_rawDescData
}

var file_GetUserInfo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_GetUserInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_GetUserInfo_proto_goTypes = []any{
	(UpdateType)(0),    // 0: GetUserInfo.UpdateType
	(UserInfoField)(0), // 1: GetUserInfo.UserInfoField
	(*C2S)(nil),        // 2: GetUserInfo.C2S
	(*S2C)(nil),        // 3: GetUserInfo.S2C
	(*Request)(nil),    // 4: GetUserInfo.Request
	(*Response)(nil),   // 5: GetUserInfo.Response
}
var file_GetUserInfo_proto_depIdxs = []int32{
	2, // 0: GetUserInfo.Request.c2s:type_name -> GetUserInfo.C2S
	3, // 1: GetUserInfo.Response.s2c:type_name -> GetUserInfo.S2C
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetUserInfo_proto_init() }
func file_GetUserInfo_proto_init() {
	if File_GetUserInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetUserInfo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*C2S); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetUserInfo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*S2C); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetUserInfo_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetUserInfo_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GetUserInfo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetUserInfo_proto_goTypes,
		DependencyIndexes: file_GetUserInfo_proto_depIdxs,
		EnumInfos:         file_GetUserInfo_proto_enumTypes,
		MessageInfos:      file_GetUserInfo_proto_msgTypes,
	}.Build()
	File_GetUserInfo_proto = out.File
	file_GetUserInfo_proto_rawDesc = nil
	file_GetUserInfo_proto_goTypes = nil
	file_GetUserInfo_proto_depIdxs = nil
}
