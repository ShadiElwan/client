package chat1

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/keybase/client/go/protocol/gregor1"
)

// Eq compares two TLFIDs
func (id TLFID) Eq(other TLFID) bool {
	return bytes.Equal([]byte(id), []byte(other))
}

// EqString is like EqualsTo, except that it accepts a fmt.Stringer. This
// can be useful for comparing keybase1.TLFID and chat1.TLFID.
func (id TLFID) EqString(other fmt.Stringer) bool {
	return hex.EncodeToString(id) == other.String()
}

func (id TLFID) String() string {
	return hex.EncodeToString(id)
}

func MakeConvID(val string) (ConversationID, error) {
	return hex.DecodeString(val)
}

func (cid ConversationID) String() string {
	return hex.EncodeToString(cid)
}

func (cid ConversationID) IsNil() bool {
	return len(cid) == 0
}

func (cid ConversationID) Eq(c ConversationID) bool {
	return bytes.Equal(cid, c)
}

func (cid ConversationID) Less(c ConversationID) bool {
	return bytes.Compare(cid, c) < 0
}

// DbShortForm should only be used when interacting with the database, and should
// never leave Gregor
func (cid ConversationID) DbShortForm() []byte {
	return cid[:10]
}

func MakeTLFID(val string) (TLFID, error) {
	return hex.DecodeString(val)
}

func MakeTopicID(val string) (TopicID, error) {
	return hex.DecodeString(val)
}

func MakeTopicType(val int64) TopicType {
	return TopicType(val)
}

func (mid MessageID) String() string {
	return strconv.FormatUint(uint64(mid), 10)
}

func (t MessageType) String() string {
	switch t {
	case MessageType_NONE:
		return "NONE"
	case MessageType_TEXT:
		return "TEXT"
	case MessageType_ATTACHMENT:
		return "ATTACHMENT"
	case MessageType_EDIT:
		return "EDIT"
	case MessageType_DELETE:
		return "DELETE"
	case MessageType_METADATA:
		return "METADATA"
	case MessageType_ATTACHMENTUPLOADED:
		return "ATTACHMENTUPLOADED"
	default:
		return "UNKNOWN"
	}
}

func (t TopicType) String() string {
	switch t {
	case TopicType_NONE:
		return "NONE"
	case TopicType_CHAT:
		return "CHAT"
	case TopicType_DEV:
		return "DEV"
	default:
		return "UNKNOWN"
	}
}

func (t TopicID) String() string {
	return hex.EncodeToString(t)
}

func (me ConversationIDTriple) Eq(other ConversationIDTriple) bool {
	return me.Tlfid.Eq(other.Tlfid) &&
		bytes.Equal([]byte(me.TopicID), []byte(other.TopicID)) &&
		me.TopicType == other.TopicType
}

func (hash Hash) String() string {
	return hex.EncodeToString(hash)
}

func (hash Hash) Eq(other Hash) bool {
	return bytes.Equal(hash, other)
}

func (m MessageUnboxed) GetMessageID() MessageID {
	if state, err := m.State(); err == nil {
		if state == MessageUnboxedState_VALID {
			return m.Valid().ServerHeader.MessageID
		}
		if state == MessageUnboxedState_ERROR {
			return m.Error().MessageID
		}
	}
	return 0
}

func (m MessageUnboxed) GetMessageType() MessageType {
	if state, err := m.State(); err == nil {
		if state == MessageUnboxedState_VALID {
			return m.Valid().ClientHeader.MessageType
		}
		if state == MessageUnboxedState_ERROR {
			return m.Error().MessageType
		}
	}
	return MessageType_NONE
}

func (m MessageUnboxed) IsValid() bool {
	if state, err := m.State(); err == nil {
		return state == MessageUnboxedState_VALID
	}
	return false
}

func (m MessageBoxed) GetMessageID() MessageID {
	return m.ServerHeader.MessageID
}

func (m MessageBoxed) GetMessageType() MessageType {
	return m.ClientHeader.MessageType
}

var ConversationStatusGregorMap = map[ConversationStatus]string{
	ConversationStatus_UNFILED:  "unfiled",
	ConversationStatus_FAVORITE: "favorite",
	ConversationStatus_IGNORED:  "ignored",
	ConversationStatus_BLOCKED:  "blocked",
	ConversationStatus_MUTED:    "muted",
}

var ConversationStatusGregorRevMap = map[string]ConversationStatus{
	"unfiled":  ConversationStatus_UNFILED,
	"favorite": ConversationStatus_FAVORITE,
	"ignored":  ConversationStatus_IGNORED,
	"blocked":  ConversationStatus_BLOCKED,
	"muted":    ConversationStatus_MUTED,
}

func (t ConversationIDTriple) Hash() []byte {
	h := sha256.New()
	h.Write(t.Tlfid)
	h.Write(t.TopicID)
	h.Write([]byte(strconv.Itoa(int(t.TopicType))))
	hash := h.Sum(nil)

	return hash
}

func (t ConversationIDTriple) ToConversationID(shardID [2]byte) ConversationID {
	h := t.Hash()
	h[0], h[1] = shardID[0], shardID[1]
	return ConversationID(h)
}

func (t ConversationIDTriple) Derivable(cid ConversationID) bool {
	h := t.Hash()
	return bytes.Equal(h[2:], []byte(cid[2:]))
}

func (o OutboxID) Eq(r OutboxID) bool {
	return bytes.Equal(o, r)
}

func (o OutboxID) String() string {
	return hex.EncodeToString(o)
}

func (t TLFVisibility) Eq(r TLFVisibility) bool {
	return int(t) == int(r)
}

// Visibility is a helper to get around a nil pointer for visibility,
// and to get around TLFVisibility_ANY.  The default is PRIVATE.
// Note:  not sure why visibility is a pointer, or what TLFVisibility_ANY
// is for, but don't want to change the API.
func (q *GetInboxLocalQuery) Visibility() TLFVisibility {
	visibility := TLFVisibility_PRIVATE
	if q.TlfVisibility != nil && *q.TlfVisibility == TLFVisibility_PUBLIC {
		visibility = TLFVisibility_PUBLIC
	}
	return visibility
}

// Visibility is a helper to get around a nil pointer for visibility,
// and to get around TLFVisibility_ANY.  The default is PRIVATE.
// Note:  not sure why visibility is a pointer, or what TLFVisibility_ANY
// is for, but don't want to change the API.
func (q *GetInboxQuery) Visibility() TLFVisibility {
	visibility := TLFVisibility_PRIVATE
	if q.TlfVisibility != nil && *q.TlfVisibility == TLFVisibility_PUBLIC {
		visibility = TLFVisibility_PUBLIC
	}
	return visibility
}

// TLFNameExpanded returns a TLF name with a reset suffix if it exists.
// This version can be used in requests to lookup the TLF.
func (c ConversationInfoLocal) TLFNameExpanded() string {
	return ExpandTLFName(c.TlfName, c.FinalizeInfo)
}

// TLFNameExpandedSummary returns a TLF name with a summary of the
// account reset if there was one.
// This version is for display purposes only and connot be used to lookup the TLF.
func (c ConversationInfoLocal) TLFNameExpandedSummary() string {
	if c.FinalizeInfo == nil {
		return c.TlfName
	}
	return c.TlfName + " " + c.FinalizeInfo.BeforeSummary()
}

// TLFNameExpanded returns a TLF name with a reset suffix if it exists.
// This version can be used in requests to lookup the TLF.
func (h MessageClientHeader) TLFNameExpanded(finalizeInfo *ConversationFinalizeInfo) string {
	return ExpandTLFName(h.TlfName, finalizeInfo)
}

// ExpandTLFName returns a TLF name with a reset suffix if it exists.
// This version can be used in requests to lookup the TLF.
func ExpandTLFName(name string, finalizeInfo *ConversationFinalizeInfo) string {
	if finalizeInfo == nil {
		return name
	}
	if len(finalizeInfo.ResetFull) == 0 {
		return name
	}
	if strings.Contains(name, " account reset ") {
		return name
	}
	return name + " " + finalizeInfo.ResetFull
}

// BeforeSummary returns a summary of the finalize without "files" in it.
// The canonical name for a TLF after reset has a "(files before ... account reset...)" suffix
// which doesn't make much sense in other uses (like chat).
func (f *ConversationFinalizeInfo) BeforeSummary() string {
	return fmt.Sprintf("(before %s account reset %s)", f.ResetUser, f.ResetDate)
}

func (p Pagination) Eq(other Pagination) bool {
	return p.Last == other.Last && bytes.Equal(p.Next, other.Next) &&
		bytes.Equal(p.Previous, other.Previous) && p.Num == other.Num
}

func (c ConversationLocal) GetMtime() gregor1.Time {
	return c.ReaderInfo.Mtime
}

func (c ConversationLocal) GetConvID() ConversationID {
	return c.Info.Id
}

func (c Conversation) GetMtime() gregor1.Time {
	return c.ReaderInfo.Mtime
}

func (c Conversation) GetConvID() ConversationID {
	return c.Metadata.ConversationID
}

func (c Conversation) GetMaxMessage(typ MessageType) (MessageBoxed, error) {
	for _, msg := range c.MaxMsgs {
		if msg.GetMessageType() == typ {
			return msg, nil
		}
	}
	return MessageBoxed{}, fmt.Errorf("max message not found: %v", typ)
}

func (c Conversation) Includes(uid gregor1.UID) bool {
	for _, auid := range c.Metadata.ActiveList {
		if uid.Eq(auid) {
			return true
		}
	}
	return false
}

/*
func ConvertMessageBodyV1ToV2(v1 MessageBodyV1) (MessageBody, error) {
	t, err := v1.MessageType()
	if err != nil {
		return MessageBody{}, err
	}
	switch t {
	case MessageType_TEXT:
		return NewMessageBodyWithText(v1.Text()), nil
	case MessageType_ATTACHMENT:
		previous := v1.Attachment()
		upgraded := MessageAttachment{
			Object:   previous.Object,
			Metadata: previous.Metadata,
			Uploaded: true,
		}
		if previous.Preview != nil {
			upgraded.Previews = []Asset{*previous.Preview}
		}
		return NewMessageBodyWithAttachment(upgraded), nil
	case MessageType_EDIT:
		return NewMessageBodyWithEdit(v1.Edit()), nil
	case MessageType_DELETE:
		return NewMessageBodyWithDelete(v1.Delete()), nil
	case MessageType_METADATA:
		return NewMessageBodyWithMetadata(v1.Metadata()), nil
	case MessageType_HEADLINE:
		return NewMessageBodyWithHeadline(v1.Headline()), nil
	case MessageType_NONE:
		return MessageBody{MessageType__: MessageType_NONE}, nil
	}

	return MessageBody{}, fmt.Errorf("ConvertMessageBodyV1ToV2: unhandled message type %v", t)
}
*/
