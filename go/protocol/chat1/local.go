// Auto-generated by avdl-compiler v1.3.7 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/local.avdl

package chat1

import (
	"errors"
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type MessageText struct {
	Body string `codec:"body" json:"body"`
}

type MessageConversationMetadata struct {
	ConversationTitle string `codec:"conversationTitle" json:"conversationTitle"`
}

type MessageEdit struct {
	MessageID MessageID `codec:"messageID" json:"messageID"`
	Body      string    `codec:"body" json:"body"`
}

type MessageDelete struct {
	MessageID MessageID `codec:"messageID" json:"messageID"`
}

type MessageAttachment struct {
	Path string `codec:"path" json:"path"`
}

type MessageBody struct {
	MessageType__ MessageType                  `codec:"messageType" json:"messageType"`
	Text__        *MessageText                 `codec:"text,omitempty" json:"text,omitempty"`
	Attachment__  *MessageAttachment           `codec:"attachment,omitempty" json:"attachment,omitempty"`
	Edit__        *MessageEdit                 `codec:"edit,omitempty" json:"edit,omitempty"`
	Delete__      *MessageDelete               `codec:"delete,omitempty" json:"delete,omitempty"`
	Metadata__    *MessageConversationMetadata `codec:"metadata,omitempty" json:"metadata,omitempty"`
}

func (o *MessageBody) MessageType() (ret MessageType, err error) {
	switch o.MessageType__ {
	case MessageType_TEXT:
		if o.Text__ == nil {
			err = errors.New("unexpected nil value for Text__")
			return ret, err
		}
	case MessageType_ATTACHMENT:
		if o.Attachment__ == nil {
			err = errors.New("unexpected nil value for Attachment__")
			return ret, err
		}
	case MessageType_EDIT:
		if o.Edit__ == nil {
			err = errors.New("unexpected nil value for Edit__")
			return ret, err
		}
	case MessageType_DELETE:
		if o.Delete__ == nil {
			err = errors.New("unexpected nil value for Delete__")
			return ret, err
		}
	case MessageType_METADATA:
		if o.Metadata__ == nil {
			err = errors.New("unexpected nil value for Metadata__")
			return ret, err
		}
	}
	return o.MessageType__, nil
}

func (o MessageBody) Text() MessageText {
	if o.MessageType__ != MessageType_TEXT {
		panic("wrong case accessed")
	}
	if o.Text__ == nil {
		return MessageText{}
	}
	return *o.Text__
}

func (o MessageBody) Attachment() MessageAttachment {
	if o.MessageType__ != MessageType_ATTACHMENT {
		panic("wrong case accessed")
	}
	if o.Attachment__ == nil {
		return MessageAttachment{}
	}
	return *o.Attachment__
}

func (o MessageBody) Edit() MessageEdit {
	if o.MessageType__ != MessageType_EDIT {
		panic("wrong case accessed")
	}
	if o.Edit__ == nil {
		return MessageEdit{}
	}
	return *o.Edit__
}

func (o MessageBody) Delete() MessageDelete {
	if o.MessageType__ != MessageType_DELETE {
		panic("wrong case accessed")
	}
	if o.Delete__ == nil {
		return MessageDelete{}
	}
	return *o.Delete__
}

func (o MessageBody) Metadata() MessageConversationMetadata {
	if o.MessageType__ != MessageType_METADATA {
		panic("wrong case accessed")
	}
	if o.Metadata__ == nil {
		return MessageConversationMetadata{}
	}
	return *o.Metadata__
}

func NewMessageBodyWithText(v MessageText) MessageBody {
	return MessageBody{
		MessageType__: MessageType_TEXT,
		Text__:        &v,
	}
}

func NewMessageBodyWithAttachment(v MessageAttachment) MessageBody {
	return MessageBody{
		MessageType__: MessageType_ATTACHMENT,
		Attachment__:  &v,
	}
}

func NewMessageBodyWithEdit(v MessageEdit) MessageBody {
	return MessageBody{
		MessageType__: MessageType_EDIT,
		Edit__:        &v,
	}
}

func NewMessageBodyWithDelete(v MessageDelete) MessageBody {
	return MessageBody{
		MessageType__: MessageType_DELETE,
		Delete__:      &v,
	}
}

func NewMessageBodyWithMetadata(v MessageConversationMetadata) MessageBody {
	return MessageBody{
		MessageType__: MessageType_METADATA,
		Metadata__:    &v,
	}
}

type MessagePlaintextVersion int

const (
	MessagePlaintextVersion_V1 MessagePlaintextVersion = 1
)

var MessagePlaintextVersionMap = map[string]MessagePlaintextVersion{
	"V1": 1,
}

var MessagePlaintextVersionRevMap = map[MessagePlaintextVersion]string{
	1: "V1",
}

type MessagePlaintextV1 struct {
	ClientHeader MessageClientHeader `codec:"clientHeader" json:"clientHeader"`
	MessageBody  MessageBody         `codec:"messageBody" json:"messageBody"`
}

type MessagePlaintext struct {
	Version__ MessagePlaintextVersion `codec:"version" json:"version"`
	V1__      *MessagePlaintextV1     `codec:"v1,omitempty" json:"v1,omitempty"`
}

func (o *MessagePlaintext) Version() (ret MessagePlaintextVersion, err error) {
	switch o.Version__ {
	case MessagePlaintextVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	}
	return o.Version__, nil
}

func (o MessagePlaintext) V1() MessagePlaintextV1 {
	if o.Version__ != MessagePlaintextVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return MessagePlaintextV1{}
	}
	return *o.V1__
}

func NewMessagePlaintextWithV1(v MessagePlaintextV1) MessagePlaintext {
	return MessagePlaintext{
		Version__: MessagePlaintextVersion_V1,
		V1__:      &v,
	}
}

type HeaderPlaintextVersion int

const (
	HeaderPlaintextVersion_V1 HeaderPlaintextVersion = 1
)

var HeaderPlaintextVersionMap = map[string]HeaderPlaintextVersion{
	"V1": 1,
}

var HeaderPlaintextVersionRevMap = map[HeaderPlaintextVersion]string{
	1: "V1",
}

type HeaderPlaintextV1 struct {
	Conv            ConversationIDTriple     `codec:"conv" json:"conv"`
	TlfName         string                   `codec:"tlfName" json:"tlfName"`
	TlfPublic       bool                     `codec:"tlfPublic" json:"tlfPublic"`
	MessageType     MessageType              `codec:"messageType" json:"messageType"`
	Prev            []MessagePreviousPointer `codec:"prev" json:"prev"`
	Sender          gregor1.UID              `codec:"sender" json:"sender"`
	SenderDevice    gregor1.DeviceID         `codec:"senderDevice" json:"senderDevice"`
	BodyHash        Hash                     `codec:"bodyHash" json:"bodyHash"`
	HeaderSignature *SignatureInfo           `codec:"headerSignature,omitempty" json:"headerSignature,omitempty"`
}

type HeaderPlaintext struct {
	Version__ HeaderPlaintextVersion `codec:"version" json:"version"`
	V1__      *HeaderPlaintextV1     `codec:"v1,omitempty" json:"v1,omitempty"`
}

func (o *HeaderPlaintext) Version() (ret HeaderPlaintextVersion, err error) {
	switch o.Version__ {
	case HeaderPlaintextVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	}
	return o.Version__, nil
}

func (o HeaderPlaintext) V1() HeaderPlaintextV1 {
	if o.Version__ != HeaderPlaintextVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return HeaderPlaintextV1{}
	}
	return *o.V1__
}

func NewHeaderPlaintextWithV1(v HeaderPlaintextV1) HeaderPlaintext {
	return HeaderPlaintext{
		Version__: HeaderPlaintextVersion_V1,
		V1__:      &v,
	}
}

type BodyPlaintextVersion int

const (
	BodyPlaintextVersion_V1 BodyPlaintextVersion = 1
)

var BodyPlaintextVersionMap = map[string]BodyPlaintextVersion{
	"V1": 1,
}

var BodyPlaintextVersionRevMap = map[BodyPlaintextVersion]string{
	1: "V1",
}

type BodyPlaintextV1 struct {
	MessageBody MessageBody `codec:"messageBody" json:"messageBody"`
}

type BodyPlaintext struct {
	Version__ BodyPlaintextVersion `codec:"version" json:"version"`
	V1__      *BodyPlaintextV1     `codec:"v1,omitempty" json:"v1,omitempty"`
}

func (o *BodyPlaintext) Version() (ret BodyPlaintextVersion, err error) {
	switch o.Version__ {
	case BodyPlaintextVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	}
	return o.Version__, nil
}

func (o BodyPlaintext) V1() BodyPlaintextV1 {
	if o.Version__ != BodyPlaintextVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return BodyPlaintextV1{}
	}
	return *o.V1__
}

func NewBodyPlaintextWithV1(v BodyPlaintextV1) BodyPlaintext {
	return BodyPlaintext{
		Version__: BodyPlaintextVersion_V1,
		V1__:      &v,
	}
}

type MessageInfoLocal struct {
	IsNew            bool   `codec:"isNew" json:"isNew"`
	SenderUsername   string `codec:"senderUsername" json:"senderUsername"`
	SenderDeviceName string `codec:"senderDeviceName" json:"senderDeviceName"`
}

type Message struct {
	ServerHeader     MessageServerHeader `codec:"serverHeader" json:"serverHeader"`
	MessagePlaintext MessagePlaintext    `codec:"messagePlaintext" json:"messagePlaintext"`
	Info             *MessageInfoLocal   `codec:"info,omitempty" json:"info,omitempty"`
}

type ThreadView struct {
	Messages   []Message   `codec:"messages" json:"messages"`
	Pagination *Pagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type MessageSelector struct {
	MessageTypes  []MessageType    `codec:"MessageTypes" json:"MessageTypes"`
	Since         *string          `codec:"Since,omitempty" json:"Since,omitempty"`
	OnlyNew       bool             `codec:"onlyNew" json:"onlyNew"`
	Limit         int              `codec:"limit" json:"limit"`
	Conversations []ConversationID `codec:"conversations" json:"conversations"`
	MarkAsRead    bool             `codec:"markAsRead" json:"markAsRead"`
}

type ConversationInfoLocal struct {
	Id         ConversationID       `codec:"id" json:"id"`
	Triple     ConversationIDTriple `codec:"triple" json:"triple"`
	TlfName    string               `codec:"tlfName" json:"tlfName"`
	TopicName  string               `codec:"topicName" json:"topicName"`
	TopicType  TopicType            `codec:"topicType" json:"topicType"`
	Visibility TLFVisibility        `codec:"visibility" json:"visibility"`
}

type ConversationLocal struct {
	Info     *ConversationInfoLocal `codec:"info,omitempty" json:"info,omitempty"`
	Messages []Message              `codec:"messages" json:"messages"`
}

type GetInboxLocalRes struct {
	Inbox      InboxView   `codec:"inbox" json:"inbox"`
	RateLimits []RateLimit `codec:"rateLimits" json:"rateLimits"`
}

type GetThreadLocalRes struct {
	Thread     ThreadView  `codec:"thread" json:"thread"`
	RateLimits []RateLimit `codec:"rateLimits" json:"rateLimits"`
}

type PostLocalRes struct {
	RateLimits []RateLimit `codec:"rateLimits" json:"rateLimits"`
}

type ResolveConversationLocalRes struct {
	Convs      []ConversationInfoLocal `codec:"convs" json:"convs"`
	RateLimits []RateLimit             `codec:"rateLimits" json:"rateLimits"`
}

type NewConversationLocalRes struct {
	Conv       ConversationInfoLocal `codec:"conv" json:"conv"`
	RateLimits []RateLimit           `codec:"rateLimits" json:"rateLimits"`
}

type UpdateTopicNameLocalRes struct {
	RateLimits []RateLimit `codec:"rateLimits" json:"rateLimits"`
}

type GetMessagesLocalRes struct {
	Msgs       []ConversationLocal `codec:"msgs" json:"msgs"`
	RateLimits []RateLimit         `codec:"rateLimits" json:"rateLimits"`
}

type GetInboxSummaryLocalRes struct {
	Conversations []ConversationLocal `codec:"conversations" json:"conversations"`
	More          []ConversationLocal `codec:"more" json:"more"`
	MoreTotal     int                 `codec:"moreTotal" json:"moreTotal"`
	RateLimits    []RateLimit         `codec:"rateLimits" json:"rateLimits"`
}

type GetInboxLocalArg struct {
	Query      *GetInboxQuery `codec:"query,omitempty" json:"query,omitempty"`
	Pagination *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetThreadLocalArg struct {
	ConversationID ConversationID  `codec:"conversationID" json:"conversationID"`
	Query          *GetThreadQuery `codec:"query,omitempty" json:"query,omitempty"`
	Pagination     *Pagination     `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type PostLocalArg struct {
	ConversationID   ConversationID   `codec:"conversationID" json:"conversationID"`
	MessagePlaintext MessagePlaintext `codec:"messagePlaintext" json:"messagePlaintext"`
}

type ResolveConversationLocalArg struct {
	Conversation ConversationInfoLocal `codec:"conversation" json:"conversation"`
}

type NewConversationLocalArg struct {
	Conversation ConversationInfoLocal `codec:"conversation" json:"conversation"`
}

type UpdateTopicNameLocalArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	NewTopicName   string         `codec:"newTopicName" json:"newTopicName"`
}

type GetMessagesLocalArg struct {
	Selector MessageSelector `codec:"selector" json:"selector"`
}

type GetInboxSummaryLocalArg struct {
	TopicType  TopicType     `codec:"topicType" json:"topicType"`
	After      string        `codec:"after" json:"after"`
	Before     string        `codec:"before" json:"before"`
	Limit      int           `codec:"limit" json:"limit"`
	Visibility TLFVisibility `codec:"visibility" json:"visibility"`
}

type LocalInterface interface {
	GetInboxLocal(context.Context, GetInboxLocalArg) (GetInboxLocalRes, error)
	GetThreadLocal(context.Context, GetThreadLocalArg) (GetThreadLocalRes, error)
	PostLocal(context.Context, PostLocalArg) (PostLocalRes, error)
	ResolveConversationLocal(context.Context, ConversationInfoLocal) (ResolveConversationLocalRes, error)
	NewConversationLocal(context.Context, ConversationInfoLocal) (NewConversationLocalRes, error)
	UpdateTopicNameLocal(context.Context, UpdateTopicNameLocalArg) (UpdateTopicNameLocalRes, error)
	GetMessagesLocal(context.Context, MessageSelector) (GetMessagesLocalRes, error)
	GetInboxSummaryLocal(context.Context, GetInboxSummaryLocalArg) (GetInboxSummaryLocalRes, error)
}

func LocalProtocol(i LocalInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.local",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getInboxLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxLocalArg)(nil), args)
						return
					}
					ret, err = i.GetInboxLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getThreadLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetThreadLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetThreadLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetThreadLocalArg)(nil), args)
						return
					}
					ret, err = i.GetThreadLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postLocal": {
				MakeArg: func() interface{} {
					ret := make([]PostLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostLocalArg)(nil), args)
						return
					}
					ret, err = i.PostLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"resolveConversationLocal": {
				MakeArg: func() interface{} {
					ret := make([]ResolveConversationLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ResolveConversationLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ResolveConversationLocalArg)(nil), args)
						return
					}
					ret, err = i.ResolveConversationLocal(ctx, (*typedArgs)[0].Conversation)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationLocal": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationLocalArg)(nil), args)
						return
					}
					ret, err = i.NewConversationLocal(ctx, (*typedArgs)[0].Conversation)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"updateTopicNameLocal": {
				MakeArg: func() interface{} {
					ret := make([]UpdateTopicNameLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]UpdateTopicNameLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]UpdateTopicNameLocalArg)(nil), args)
						return
					}
					ret, err = i.UpdateTopicNameLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getMessagesLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetMessagesLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetMessagesLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetMessagesLocalArg)(nil), args)
						return
					}
					ret, err = i.GetMessagesLocal(ctx, (*typedArgs)[0].Selector)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getInboxSummaryLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxSummaryLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxSummaryLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxSummaryLocalArg)(nil), args)
						return
					}
					ret, err = i.GetInboxSummaryLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type LocalClient struct {
	Cli rpc.GenericClient
}

func (c LocalClient) GetInboxLocal(ctx context.Context, __arg GetInboxLocalArg) (res GetInboxLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.getInboxLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetThreadLocal(ctx context.Context, __arg GetThreadLocalArg) (res GetThreadLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.getThreadLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) PostLocal(ctx context.Context, __arg PostLocalArg) (res PostLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.postLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) ResolveConversationLocal(ctx context.Context, conversation ConversationInfoLocal) (res ResolveConversationLocalRes, err error) {
	__arg := ResolveConversationLocalArg{Conversation: conversation}
	err = c.Cli.Call(ctx, "chat.1.local.resolveConversationLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) NewConversationLocal(ctx context.Context, conversation ConversationInfoLocal) (res NewConversationLocalRes, err error) {
	__arg := NewConversationLocalArg{Conversation: conversation}
	err = c.Cli.Call(ctx, "chat.1.local.newConversationLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) UpdateTopicNameLocal(ctx context.Context, __arg UpdateTopicNameLocalArg) (res UpdateTopicNameLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.updateTopicNameLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetMessagesLocal(ctx context.Context, selector MessageSelector) (res GetMessagesLocalRes, err error) {
	__arg := GetMessagesLocalArg{Selector: selector}
	err = c.Cli.Call(ctx, "chat.1.local.getMessagesLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetInboxSummaryLocal(ctx context.Context, __arg GetInboxSummaryLocalArg) (res GetInboxSummaryLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.getInboxSummaryLocal", []interface{}{__arg}, &res)
	return
}