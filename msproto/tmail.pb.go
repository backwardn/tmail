// Code generated by protoc-gen-go.
// source: tmail.proto
// DO NOT EDIT!

/*
Package msproto is a generated protocol buffer package.

It is generated from these files:
	tmail.proto

It has these top-level messages:
	SmtpResponse
	SmtpdTelemetry
	SmtpdNewClientQuery
	SmtpdNewClientResponse
	SmtpdHeloQuery
	SmtpdHeloResponse
	SmtpdRcptToQuery
	SmtpdRcptToResponse
	SmtpdDataQuery
	SmtpdDataResponse
	SmtpdBeforeQueueingQuery
	SmtpdBeforeQueueingResponse
	DeliverdTelemetry
	DeliverdGetRoutesQuery
	DeliverdGetRoutesResponse
*/
package msproto

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// SMTP response for smtpd hooks
type SmtpResponse struct {
	Code             *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Msg              *string `protobuf:"bytes,2,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmtpResponse) Reset()         { *m = SmtpResponse{} }
func (m *SmtpResponse) String() string { return proto.CompactTextString(m) }
func (*SmtpResponse) ProtoMessage()    {}

func (m *SmtpResponse) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *SmtpResponse) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// smtpd telemetry
type SmtpdTelemetry struct {
	ServerId         *string  `protobuf:"bytes,1,req,name=server_id" json:"server_id,omitempty"`
	SessionId        *string  `protobuf:"bytes,2,req,name=session_id" json:"session_id,omitempty"`
	RemoteAddress    *string  `protobuf:"bytes,3,req,name=remote_address" json:"remote_address,omitempty"`
	Success          *bool    `protobuf:"varint,4,req,name=success" json:"success,omitempty"`
	SmtpResponseCode *uint32  `protobuf:"varint,5,req,name=smtp_response_code" json:"smtp_response_code,omitempty"`
	EnvMailfrom      *string  `protobuf:"bytes,6,req,name=env_mailfrom" json:"env_mailfrom,omitempty"`
	EnvRcptto        []string `protobuf:"bytes,7,rep,name=env_rcptto" json:"env_rcptto,omitempty"`
	MessageSize      *uint32  `protobuf:"varint,8,req,name=message_size" json:"message_size,omitempty"`
	IsTls            *bool    `protobuf:"varint,9,req,name=is_tls" json:"is_tls,omitempty"`
	Concurrency      *uint32  `protobuf:"varint,10,req,name=concurrency" json:"concurrency,omitempty"`
	ExecTime         *uint32  `protobuf:"varint,11,req,name=exec_time" json:"exec_time,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmtpdTelemetry) Reset()         { *m = SmtpdTelemetry{} }
func (m *SmtpdTelemetry) String() string { return proto.CompactTextString(m) }
func (*SmtpdTelemetry) ProtoMessage()    {}

func (m *SmtpdTelemetry) GetServerId() string {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return ""
}

func (m *SmtpdTelemetry) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdTelemetry) GetRemoteAddress() string {
	if m != nil && m.RemoteAddress != nil {
		return *m.RemoteAddress
	}
	return ""
}

func (m *SmtpdTelemetry) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *SmtpdTelemetry) GetSmtpResponseCode() uint32 {
	if m != nil && m.SmtpResponseCode != nil {
		return *m.SmtpResponseCode
	}
	return 0
}

func (m *SmtpdTelemetry) GetEnvMailfrom() string {
	if m != nil && m.EnvMailfrom != nil {
		return *m.EnvMailfrom
	}
	return ""
}

func (m *SmtpdTelemetry) GetEnvRcptto() []string {
	if m != nil {
		return m.EnvRcptto
	}
	return nil
}

func (m *SmtpdTelemetry) GetMessageSize() uint32 {
	if m != nil && m.MessageSize != nil {
		return *m.MessageSize
	}
	return 0
}

func (m *SmtpdTelemetry) GetIsTls() bool {
	if m != nil && m.IsTls != nil {
		return *m.IsTls
	}
	return false
}

func (m *SmtpdTelemetry) GetConcurrency() uint32 {
	if m != nil && m.Concurrency != nil {
		return *m.Concurrency
	}
	return 0
}

func (m *SmtpdTelemetry) GetExecTime() uint32 {
	if m != nil && m.ExecTime != nil {
		return *m.ExecTime
	}
	return 0
}

// Hook SMTPd New client
// SmtpdNewClientQuery
type SmtpdNewClientQuery struct {
	SessionId        *string `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	RemoteIp         *string `protobuf:"bytes,2,req,name=remote_ip" json:"remote_ip,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmtpdNewClientQuery) Reset()         { *m = SmtpdNewClientQuery{} }
func (m *SmtpdNewClientQuery) String() string { return proto.CompactTextString(m) }
func (*SmtpdNewClientQuery) ProtoMessage()    {}

func (m *SmtpdNewClientQuery) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdNewClientQuery) GetRemoteIp() string {
	if m != nil && m.RemoteIp != nil {
		return *m.RemoteIp
	}
	return ""
}

// SmtpdNewClientResponse
type SmtpdNewClientResponse struct {
	SessionId        *string       `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	SmtpResponse     *SmtpResponse `protobuf:"bytes,2,opt,name=smtp_response" json:"smtp_response,omitempty"`
	DropConnection   *bool         `protobuf:"varint,3,opt,name=drop_connection" json:"drop_connection,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SmtpdNewClientResponse) Reset()         { *m = SmtpdNewClientResponse{} }
func (m *SmtpdNewClientResponse) String() string { return proto.CompactTextString(m) }
func (*SmtpdNewClientResponse) ProtoMessage()    {}

func (m *SmtpdNewClientResponse) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdNewClientResponse) GetSmtpResponse() *SmtpResponse {
	if m != nil {
		return m.SmtpResponse
	}
	return nil
}

func (m *SmtpdNewClientResponse) GetDropConnection() bool {
	if m != nil && m.DropConnection != nil {
		return *m.DropConnection
	}
	return false
}

// hook SMTPd HELO/EHLO
// smtpdHeloQuery
type SmtpdHeloQuery struct {
	SessionId        *string `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	Helo             *string `protobuf:"bytes,2,req,name=helo" json:"helo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmtpdHeloQuery) Reset()         { *m = SmtpdHeloQuery{} }
func (m *SmtpdHeloQuery) String() string { return proto.CompactTextString(m) }
func (*SmtpdHeloQuery) ProtoMessage()    {}

func (m *SmtpdHeloQuery) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdHeloQuery) GetHelo() string {
	if m != nil && m.Helo != nil {
		return *m.Helo
	}
	return ""
}

// SmtpdHeloResponse
type SmtpdHeloResponse struct {
	SessionId        *string       `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	SmtpResponse     *SmtpResponse `protobuf:"bytes,2,opt,name=smtp_response" json:"smtp_response,omitempty"`
	DropConnection   *bool         `protobuf:"varint,3,opt,name=drop_connection" json:"drop_connection,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SmtpdHeloResponse) Reset()         { *m = SmtpdHeloResponse{} }
func (m *SmtpdHeloResponse) String() string { return proto.CompactTextString(m) }
func (*SmtpdHeloResponse) ProtoMessage()    {}

func (m *SmtpdHeloResponse) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdHeloResponse) GetSmtpResponse() *SmtpResponse {
	if m != nil {
		return m.SmtpResponse
	}
	return nil
}

func (m *SmtpdHeloResponse) GetDropConnection() bool {
	if m != nil && m.DropConnection != nil {
		return *m.DropConnection
	}
	return false
}

// Hook smtpd RCPT TO command
// smtpdRcptToQuery
type SmtpdRcptToQuery struct {
	SessionId        *string `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	Rcptto           *string `protobuf:"bytes,2,req,name=rcptto" json:"rcptto,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmtpdRcptToQuery) Reset()         { *m = SmtpdRcptToQuery{} }
func (m *SmtpdRcptToQuery) String() string { return proto.CompactTextString(m) }
func (*SmtpdRcptToQuery) ProtoMessage()    {}

func (m *SmtpdRcptToQuery) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdRcptToQuery) GetRcptto() string {
	if m != nil && m.Rcptto != nil {
		return *m.Rcptto
	}
	return ""
}

// SmtpdRcpttoAccessIsGrantedResponse
type SmtpdRcptToResponse struct {
	SessionId        *string       `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	SmtpResponse     *SmtpResponse `protobuf:"bytes,2,opt,name=smtp_response" json:"smtp_response,omitempty"`
	DropConnection   *bool         `protobuf:"varint,3,opt,name=drop_connection" json:"drop_connection,omitempty"`
	RelayGranted     *bool         `protobuf:"varint,4,opt,name=relay_granted" json:"relay_granted,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SmtpdRcptToResponse) Reset()         { *m = SmtpdRcptToResponse{} }
func (m *SmtpdRcptToResponse) String() string { return proto.CompactTextString(m) }
func (*SmtpdRcptToResponse) ProtoMessage()    {}

func (m *SmtpdRcptToResponse) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdRcptToResponse) GetSmtpResponse() *SmtpResponse {
	if m != nil {
		return m.SmtpResponse
	}
	return nil
}

func (m *SmtpdRcptToResponse) GetDropConnection() bool {
	if m != nil && m.DropConnection != nil {
		return *m.DropConnection
	}
	return false
}

func (m *SmtpdRcptToResponse) GetRelayGranted() bool {
	if m != nil && m.RelayGranted != nil {
		return *m.RelayGranted
	}
	return false
}

// Hook smtpd DATA
// smtpdDataMsg
type SmtpdDataQuery struct {
	SessionId        *string `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	DataLink         *string `protobuf:"bytes,2,req,name=data_link" json:"data_link,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmtpdDataQuery) Reset()         { *m = SmtpdDataQuery{} }
func (m *SmtpdDataQuery) String() string { return proto.CompactTextString(m) }
func (*SmtpdDataQuery) ProtoMessage()    {}

func (m *SmtpdDataQuery) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdDataQuery) GetDataLink() string {
	if m != nil && m.DataLink != nil {
		return *m.DataLink
	}
	return ""
}

// smtpdDataMsg
type SmtpdDataResponse struct {
	SessionId        *string       `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	SmtpResponse     *SmtpResponse `protobuf:"bytes,2,opt,name=smtp_response" json:"smtp_response,omitempty"`
	DataLink         *string       `protobuf:"bytes,3,opt,name=data_link" json:"data_link,omitempty"`
	DropConnection   *bool         `protobuf:"varint,4,opt,name=drop_connection" json:"drop_connection,omitempty"`
	ExtraHeaders     []string      `protobuf:"bytes,5,rep,name=extra_headers" json:"extra_headers,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SmtpdDataResponse) Reset()         { *m = SmtpdDataResponse{} }
func (m *SmtpdDataResponse) String() string { return proto.CompactTextString(m) }
func (*SmtpdDataResponse) ProtoMessage()    {}

func (m *SmtpdDataResponse) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdDataResponse) GetSmtpResponse() *SmtpResponse {
	if m != nil {
		return m.SmtpResponse
	}
	return nil
}

func (m *SmtpdDataResponse) GetDataLink() string {
	if m != nil && m.DataLink != nil {
		return *m.DataLink
	}
	return ""
}

func (m *SmtpdDataResponse) GetDropConnection() bool {
	if m != nil && m.DropConnection != nil {
		return *m.DropConnection
	}
	return false
}

func (m *SmtpdDataResponse) GetExtraHeaders() []string {
	if m != nil {
		return m.ExtraHeaders
	}
	return nil
}

// SmtpdBeforeQueing Query
type SmtpdBeforeQueueingQuery struct {
	SessionId        *string  `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	MailFrom         *string  `protobuf:"bytes,2,req,name=mail_from" json:"mail_from,omitempty"`
	RcptTo           []string `protobuf:"bytes,3,rep,name=rcpt_to" json:"rcpt_to,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmtpdBeforeQueueingQuery) Reset()         { *m = SmtpdBeforeQueueingQuery{} }
func (m *SmtpdBeforeQueueingQuery) String() string { return proto.CompactTextString(m) }
func (*SmtpdBeforeQueueingQuery) ProtoMessage()    {}

func (m *SmtpdBeforeQueueingQuery) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdBeforeQueueingQuery) GetMailFrom() string {
	if m != nil && m.MailFrom != nil {
		return *m.MailFrom
	}
	return ""
}

func (m *SmtpdBeforeQueueingQuery) GetRcptTo() []string {
	if m != nil {
		return m.RcptTo
	}
	return nil
}

// SmtpdBefore queuing Response
type SmtpdBeforeQueueingResponse struct {
	SessionId        *string       `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	MailFrom         *string       `protobuf:"bytes,2,opt,name=mail_from" json:"mail_from,omitempty"`
	RcptTo           []string      `protobuf:"bytes,3,rep,name=rcpt_to" json:"rcpt_to,omitempty"`
	SmtpResponse     *SmtpResponse `protobuf:"bytes,4,opt,name=smtp_response" json:"smtp_response,omitempty"`
	DropConnection   *bool         `protobuf:"varint,5,opt,name=drop_connection" json:"drop_connection,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SmtpdBeforeQueueingResponse) Reset()         { *m = SmtpdBeforeQueueingResponse{} }
func (m *SmtpdBeforeQueueingResponse) String() string { return proto.CompactTextString(m) }
func (*SmtpdBeforeQueueingResponse) ProtoMessage()    {}

func (m *SmtpdBeforeQueueingResponse) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdBeforeQueueingResponse) GetMailFrom() string {
	if m != nil && m.MailFrom != nil {
		return *m.MailFrom
	}
	return ""
}

func (m *SmtpdBeforeQueueingResponse) GetRcptTo() []string {
	if m != nil {
		return m.RcptTo
	}
	return nil
}

func (m *SmtpdBeforeQueueingResponse) GetSmtpResponse() *SmtpResponse {
	if m != nil {
		return m.SmtpResponse
	}
	return nil
}

func (m *SmtpdBeforeQueueingResponse) GetDropConnection() bool {
	if m != nil && m.DropConnection != nil {
		return *m.DropConnection
	}
	return false
}

// telemetry
type DeliverdTelemetry struct {
	ServerId               *string `protobuf:"bytes,1,req,name=server_id" json:"server_id,omitempty"`
	DeliverdId             *string `protobuf:"bytes,2,req,name=deliverd_id" json:"deliverd_id,omitempty"`
	Success                *bool   `protobuf:"varint,3,req,name=success" json:"success,omitempty"`
	ExecTime               *uint32 `protobuf:"varint,4,req,name=exec_time" json:"exec_time,omitempty"`
	MessagesInQueue        *uint32 `protobuf:"varint,5,req,name=messages_in_queue" json:"messages_in_queue,omitempty"`
	ConcurrencyRemote      *uint32 `protobuf:"varint,6,req,name=concurrency_remote" json:"concurrency_remote,omitempty"`
	ConcurrencyLocal       *uint32 `protobuf:"varint,7,req,name=concurrency_local" json:"concurrency_local,omitempty"`
	From                   *string `protobuf:"bytes,8,req,name=from" json:"from,omitempty"`
	To                     *string `protobuf:"bytes,9,req,name=to" json:"to,omitempty"`
	IsLocal                *bool   `protobuf:"varint,10,req,name=is_local" json:"is_local,omitempty"`
	LocalAddress           *string `protobuf:"bytes,11,opt,name=local_address" json:"local_address,omitempty"`
	RemoteAddress          *string `protobuf:"bytes,12,opt,name=remote_address" json:"remote_address,omitempty"`
	RemoteSmtpResponseCode *uint32 `protobuf:"varint,13,opt,name=remote_smtp_response_code" json:"remote_smtp_response_code,omitempty"`
	XXX_unrecognized       []byte  `json:"-"`
}

func (m *DeliverdTelemetry) Reset()         { *m = DeliverdTelemetry{} }
func (m *DeliverdTelemetry) String() string { return proto.CompactTextString(m) }
func (*DeliverdTelemetry) ProtoMessage()    {}

func (m *DeliverdTelemetry) GetServerId() string {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return ""
}

func (m *DeliverdTelemetry) GetDeliverdId() string {
	if m != nil && m.DeliverdId != nil {
		return *m.DeliverdId
	}
	return ""
}

func (m *DeliverdTelemetry) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *DeliverdTelemetry) GetExecTime() uint32 {
	if m != nil && m.ExecTime != nil {
		return *m.ExecTime
	}
	return 0
}

func (m *DeliverdTelemetry) GetMessagesInQueue() uint32 {
	if m != nil && m.MessagesInQueue != nil {
		return *m.MessagesInQueue
	}
	return 0
}

func (m *DeliverdTelemetry) GetConcurrencyRemote() uint32 {
	if m != nil && m.ConcurrencyRemote != nil {
		return *m.ConcurrencyRemote
	}
	return 0
}

func (m *DeliverdTelemetry) GetConcurrencyLocal() uint32 {
	if m != nil && m.ConcurrencyLocal != nil {
		return *m.ConcurrencyLocal
	}
	return 0
}

func (m *DeliverdTelemetry) GetFrom() string {
	if m != nil && m.From != nil {
		return *m.From
	}
	return ""
}

func (m *DeliverdTelemetry) GetTo() string {
	if m != nil && m.To != nil {
		return *m.To
	}
	return ""
}

func (m *DeliverdTelemetry) GetIsLocal() bool {
	if m != nil && m.IsLocal != nil {
		return *m.IsLocal
	}
	return false
}

func (m *DeliverdTelemetry) GetLocalAddress() string {
	if m != nil && m.LocalAddress != nil {
		return *m.LocalAddress
	}
	return ""
}

func (m *DeliverdTelemetry) GetRemoteAddress() string {
	if m != nil && m.RemoteAddress != nil {
		return *m.RemoteAddress
	}
	return ""
}

func (m *DeliverdTelemetry) GetRemoteSmtpResponseCode() uint32 {
	if m != nil && m.RemoteSmtpResponseCode != nil {
		return *m.RemoteSmtpResponseCode
	}
	return 0
}

// Get routes query
type DeliverdGetRoutesQuery struct {
	DeliverdId       *string `protobuf:"bytes,1,req,name=deliverd_id" json:"deliverd_id,omitempty"`
	Mailfrom         *string `protobuf:"bytes,2,req,name=mailfrom" json:"mailfrom,omitempty"`
	Rcptto           *string `protobuf:"bytes,3,req,name=rcptto" json:"rcptto,omitempty"`
	AuthentifiedUser *string `protobuf:"bytes,4,req,name=authentified_user" json:"authentified_user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeliverdGetRoutesQuery) Reset()         { *m = DeliverdGetRoutesQuery{} }
func (m *DeliverdGetRoutesQuery) String() string { return proto.CompactTextString(m) }
func (*DeliverdGetRoutesQuery) ProtoMessage()    {}

func (m *DeliverdGetRoutesQuery) GetDeliverdId() string {
	if m != nil && m.DeliverdId != nil {
		return *m.DeliverdId
	}
	return ""
}

func (m *DeliverdGetRoutesQuery) GetMailfrom() string {
	if m != nil && m.Mailfrom != nil {
		return *m.Mailfrom
	}
	return ""
}

func (m *DeliverdGetRoutesQuery) GetRcptto() string {
	if m != nil && m.Rcptto != nil {
		return *m.Rcptto
	}
	return ""
}

func (m *DeliverdGetRoutesQuery) GetAuthentifiedUser() string {
	if m != nil && m.AuthentifiedUser != nil {
		return *m.AuthentifiedUser
	}
	return ""
}

// get routes response
type DeliverdGetRoutesResponse struct {
	DeliverdId       *string                            `protobuf:"bytes,1,req,name=deliverd_id" json:"deliverd_id,omitempty"`
	Routes           []*DeliverdGetRoutesResponse_Route `protobuf:"bytes,2,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *DeliverdGetRoutesResponse) Reset()         { *m = DeliverdGetRoutesResponse{} }
func (m *DeliverdGetRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*DeliverdGetRoutesResponse) ProtoMessage()    {}

func (m *DeliverdGetRoutesResponse) GetDeliverdId() string {
	if m != nil && m.DeliverdId != nil {
		return *m.DeliverdId
	}
	return ""
}

func (m *DeliverdGetRoutesResponse) GetRoutes() []*DeliverdGetRoutesResponse_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type DeliverdGetRoutesResponse_Route struct {
	RemoteHost       *string `protobuf:"bytes,1,req,name=remote_host" json:"remote_host,omitempty"`
	RemotePort       *int64  `protobuf:"varint,2,req,name=remote_port" json:"remote_port,omitempty"`
	LocalIp          *string `protobuf:"bytes,3,opt,name=local_ip" json:"local_ip,omitempty"`
	Priority         *int32  `protobuf:"varint,4,opt,name=priority" json:"priority,omitempty"`
	SmtpauthLogin    *string `protobuf:"bytes,5,opt,name=smtpauth_login" json:"smtpauth_login,omitempty"`
	SmtpauthPassword *string `protobuf:"bytes,6,opt,name=smtpauth_password" json:"smtpauth_password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeliverdGetRoutesResponse_Route) Reset()         { *m = DeliverdGetRoutesResponse_Route{} }
func (m *DeliverdGetRoutesResponse_Route) String() string { return proto.CompactTextString(m) }
func (*DeliverdGetRoutesResponse_Route) ProtoMessage()    {}

func (m *DeliverdGetRoutesResponse_Route) GetRemoteHost() string {
	if m != nil && m.RemoteHost != nil {
		return *m.RemoteHost
	}
	return ""
}

func (m *DeliverdGetRoutesResponse_Route) GetRemotePort() int64 {
	if m != nil && m.RemotePort != nil {
		return *m.RemotePort
	}
	return 0
}

func (m *DeliverdGetRoutesResponse_Route) GetLocalIp() string {
	if m != nil && m.LocalIp != nil {
		return *m.LocalIp
	}
	return ""
}

func (m *DeliverdGetRoutesResponse_Route) GetPriority() int32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

func (m *DeliverdGetRoutesResponse_Route) GetSmtpauthLogin() string {
	if m != nil && m.SmtpauthLogin != nil {
		return *m.SmtpauthLogin
	}
	return ""
}

func (m *DeliverdGetRoutesResponse_Route) GetSmtpauthPassword() string {
	if m != nil && m.SmtpauthPassword != nil {
		return *m.SmtpauthPassword
	}
	return ""
}

func init() {
}
