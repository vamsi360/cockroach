// Code generated by protoc-gen-gogo.
// source: cockroach/sql/driver/wire.proto
// DO NOT EDIT!

/*
	Package driver is a generated protocol buffer package.

	It is generated from these files:
		cockroach/sql/driver/wire.proto

	It has these top-level messages:
		Datum
		Result
		Request
		Response
*/
package driver

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Datum struct {
	// Using explicit proto types provides convenient access when using json. If
	// we used a Kind+Bytes approach the json interface would involve base64
	// encoded data.
	//
	// Types that are valid to be assigned to Payload:
	//	*Datum_BoolVal
	//	*Datum_IntVal
	//	*Datum_FloatVal
	//	*Datum_BytesVal
	//	*Datum_StringVal
	//	*Datum_TimeVal
	Payload isDatum_Payload `protobuf_oneof:"payload"`
}

func (m *Datum) Reset()      { *m = Datum{} }
func (*Datum) ProtoMessage() {}

type isDatum_Payload interface {
	isDatum_Payload()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Datum_BoolVal struct {
	BoolVal bool `protobuf:"varint,1,opt,name=bool_val,oneof"`
}
type Datum_IntVal struct {
	IntVal int64 `protobuf:"varint,2,opt,name=int_val,oneof"`
}
type Datum_FloatVal struct {
	FloatVal float64 `protobuf:"fixed64,3,opt,name=float_val,oneof"`
}
type Datum_BytesVal struct {
	BytesVal []byte `protobuf:"bytes,4,opt,name=bytes_val,oneof"`
}
type Datum_StringVal struct {
	StringVal string `protobuf:"bytes,5,opt,name=string_val,oneof"`
}
type Datum_TimeVal struct {
	TimeVal *Datum_Timestamp `protobuf:"bytes,6,opt,name=time_val,oneof"`
}

func (*Datum_BoolVal) isDatum_Payload()   {}
func (*Datum_IntVal) isDatum_Payload()    {}
func (*Datum_FloatVal) isDatum_Payload()  {}
func (*Datum_BytesVal) isDatum_Payload()  {}
func (*Datum_StringVal) isDatum_Payload() {}
func (*Datum_TimeVal) isDatum_Payload()   {}

func (m *Datum) GetPayload() isDatum_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Datum) GetBoolVal() bool {
	if x, ok := m.GetPayload().(*Datum_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (m *Datum) GetIntVal() int64 {
	if x, ok := m.GetPayload().(*Datum_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (m *Datum) GetFloatVal() float64 {
	if x, ok := m.GetPayload().(*Datum_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

func (m *Datum) GetBytesVal() []byte {
	if x, ok := m.GetPayload().(*Datum_BytesVal); ok {
		return x.BytesVal
	}
	return nil
}

func (m *Datum) GetStringVal() string {
	if x, ok := m.GetPayload().(*Datum_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (m *Datum) GetTimeVal() *Datum_Timestamp {
	if x, ok := m.GetPayload().(*Datum_TimeVal); ok {
		return x.TimeVal
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Datum) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Datum_OneofMarshaler, _Datum_OneofUnmarshaler, []interface{}{
		(*Datum_BoolVal)(nil),
		(*Datum_IntVal)(nil),
		(*Datum_FloatVal)(nil),
		(*Datum_BytesVal)(nil),
		(*Datum_StringVal)(nil),
		(*Datum_TimeVal)(nil),
	}
}

func _Datum_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Datum)
	// payload
	switch x := m.Payload.(type) {
	case *Datum_BoolVal:
		t := uint64(0)
		if x.BoolVal {
			t = 1
		}
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *Datum_IntVal:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.IntVal))
	case *Datum_FloatVal:
		_ = b.EncodeVarint(3<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.FloatVal))
	case *Datum_BytesVal:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.BytesVal)
	case *Datum_StringVal:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.StringVal)
	case *Datum_TimeVal:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeVal); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Datum.Payload has unexpected type %T", x)
	}
	return nil
}

func _Datum_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Datum)
	switch tag {
	case 1: // payload.bool_val
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Payload = &Datum_BoolVal{x != 0}
		return true, err
	case 2: // payload.int_val
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Payload = &Datum_IntVal{int64(x)}
		return true, err
	case 3: // payload.float_val
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Payload = &Datum_FloatVal{math.Float64frombits(x)}
		return true, err
	case 4: // payload.bytes_val
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &Datum_BytesVal{x}
		return true, err
	case 5: // payload.string_val
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &Datum_StringVal{x}
		return true, err
	case 6: // payload.time_val
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Datum_Timestamp)
		err := b.DecodeMessage(msg)
		m.Payload = &Datum_TimeVal{msg}
		return true, err
	default:
		return false, nil
	}
}

// Timestamp represents an absolute timestamp devoid of time-zone.
type Datum_Timestamp struct {
	// The time in seconds since, January 1, 1970 UTC (Unix time).
	Sec int64 `protobuf:"varint,1,opt,name=sec" json:"sec"`
	// nsec specifies a non-negative nanosecond offset within sec.
	// It must be in the range [0, 999999999].
	Nsec uint32 `protobuf:"varint,2,opt,name=nsec" json:"nsec"`
}

func (m *Datum_Timestamp) Reset()         { *m = Datum_Timestamp{} }
func (m *Datum_Timestamp) String() string { return proto.CompactTextString(m) }
func (*Datum_Timestamp) ProtoMessage()    {}

func (m *Datum_Timestamp) GetSec() int64 {
	if m != nil {
		return m.Sec
	}
	return 0
}

func (m *Datum_Timestamp) GetNsec() uint32 {
	if m != nil {
		return m.Nsec
	}
	return 0
}

// A Result is a collection of rows.
type Result struct {
	// Error is non-nil if an error occurred while executing the statement.
	Error *string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	// The names of the columns returned in the result set in the order specified
	// in the SQL statement. The number of columns will equal the number of
	// values in each Row.
	Columns []string `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	// The rows in the result set.
	Rows []Result_Row `protobuf:"bytes,3,rep,name=rows" json:"rows"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}

func (m *Result) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *Result) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Result) GetRows() []Result_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

// A Row is a collection of values representing a row in a result.
type Result_Row struct {
	Values []Datum `protobuf:"bytes,1,rep,name=values" json:"values"`
}

func (m *Result_Row) Reset()         { *m = Result_Row{} }
func (m *Result_Row) String() string { return proto.CompactTextString(m) }
func (*Result_Row) ProtoMessage()    {}

func (m *Result_Row) GetValues() []Datum {
	if m != nil {
		return m.Values
	}
	return nil
}

// An SQL request to cockroach. A transaction can consist of multiple
// requests.
type Request struct {
	// User is the originating user.
	User string `protobuf:"bytes,1,opt,name=user" json:"user"`
	// Session settings that were returned in the last response that
	// contained them, being reflected back to the server.
	Session []byte `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	// SQL statement(s) to be serially executed by the server. Multiple
	// statements are passed as a single string separated by semicolons.
	Sql string `protobuf:"bytes,3,opt,name=sql" json:"sql"`
	// Parameters referred to in the above SQL statement(s) using "?".
	Params []Datum `protobuf:"bytes,4,rep,name=params" json:"params"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Request) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Request) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *Request) GetParams() []Datum {
	if m != nil {
		return m.Params
	}
	return nil
}

type Response struct {
	// Setting that should be reflected back in all subsequent requests.
	// When not set, future requests should continue to use existing settings.
	Session []byte `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// The list of results. There is one result object per SQL statement in the
	// request.
	Results []Result `protobuf:"bytes,2,rep,name=results" json:"results"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Response) GetResults() []Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Datum) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Datum) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		nn1, err := m.Payload.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *Datum_BoolVal) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x8
	i++
	if m.BoolVal {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}
func (m *Datum_IntVal) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x10
	i++
	i = encodeVarintWire(data, i, uint64(m.IntVal))
	return i, nil
}
func (m *Datum_FloatVal) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x19
	i++
	i = encodeFixed64Wire(data, i, uint64(math.Float64bits(m.FloatVal)))
	return i, nil
}
func (m *Datum_BytesVal) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.BytesVal != nil {
		data[i] = 0x22
		i++
		i = encodeVarintWire(data, i, uint64(len(m.BytesVal)))
		i += copy(data[i:], m.BytesVal)
	}
	return i, nil
}
func (m *Datum_StringVal) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x2a
	i++
	i = encodeVarintWire(data, i, uint64(len(m.StringVal)))
	i += copy(data[i:], m.StringVal)
	return i, nil
}
func (m *Datum_TimeVal) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.TimeVal != nil {
		data[i] = 0x32
		i++
		i = encodeVarintWire(data, i, uint64(m.TimeVal.Size()))
		n2, err := m.TimeVal.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *Datum_Timestamp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Datum_Timestamp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintWire(data, i, uint64(m.Sec))
	data[i] = 0x10
	i++
	i = encodeVarintWire(data, i, uint64(m.Nsec))
	return i, nil
}

func (m *Result) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Result) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintWire(data, i, uint64(len(*m.Error)))
		i += copy(data[i:], *m.Error)
	}
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Rows) > 0 {
		for _, msg := range m.Rows {
			data[i] = 0x1a
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Result_Row) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Result_Row) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			data[i] = 0xa
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintWire(data, i, uint64(len(m.User)))
	i += copy(data[i:], m.User)
	if m.Session != nil {
		data[i] = 0x12
		i++
		i = encodeVarintWire(data, i, uint64(len(m.Session)))
		i += copy(data[i:], m.Session)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintWire(data, i, uint64(len(m.Sql)))
	i += copy(data[i:], m.Sql)
	if len(m.Params) > 0 {
		for _, msg := range m.Params {
			data[i] = 0x22
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Session != nil {
		data[i] = 0xa
		i++
		i = encodeVarintWire(data, i, uint64(len(m.Session)))
		i += copy(data[i:], m.Session)
	}
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			data[i] = 0x12
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Wire(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Wire(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWire(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Datum) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		n += m.Payload.Size()
	}
	return n
}

func (m *Datum_BoolVal) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *Datum_IntVal) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovWire(uint64(m.IntVal))
	return n
}
func (m *Datum_FloatVal) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}
func (m *Datum_BytesVal) Size() (n int) {
	var l int
	_ = l
	if m.BytesVal != nil {
		l = len(m.BytesVal)
		n += 1 + l + sovWire(uint64(l))
	}
	return n
}
func (m *Datum_StringVal) Size() (n int) {
	var l int
	_ = l
	l = len(m.StringVal)
	n += 1 + l + sovWire(uint64(l))
	return n
}
func (m *Datum_TimeVal) Size() (n int) {
	var l int
	_ = l
	if m.TimeVal != nil {
		l = m.TimeVal.Size()
		n += 1 + l + sovWire(uint64(l))
	}
	return n
}
func (m *Datum_Timestamp) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovWire(uint64(m.Sec))
	n += 1 + sovWire(uint64(m.Nsec))
	return n
}

func (m *Result) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = len(*m.Error)
		n += 1 + l + sovWire(uint64(l))
	}
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			l = len(s)
			n += 1 + l + sovWire(uint64(l))
		}
	}
	if len(m.Rows) > 0 {
		for _, e := range m.Rows {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func (m *Result_Row) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func (m *Request) Size() (n int) {
	var l int
	_ = l
	l = len(m.User)
	n += 1 + l + sovWire(uint64(l))
	if m.Session != nil {
		l = len(m.Session)
		n += 1 + l + sovWire(uint64(l))
	}
	l = len(m.Sql)
	n += 1 + l + sovWire(uint64(l))
	if len(m.Params) > 0 {
		for _, e := range m.Params {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Session != nil {
		l = len(m.Session)
		n += 1 + l + sovWire(uint64(l))
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func sovWire(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWire(x uint64) (n int) {
	return sovWire(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Datum) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Datum: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Datum: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolVal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Payload = &Datum_BoolVal{b}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntVal", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Payload = &Datum_IntVal{v}
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatVal", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Payload = &Datum_FloatVal{float64(math.Float64frombits(v))}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesVal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, data[iNdEx:postIndex])
			m.Payload = &Datum_BytesVal{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = &Datum_StringVal{string(data[iNdEx:postIndex])}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeVal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Datum_Timestamp{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &Datum_TimeVal{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Datum_Timestamp) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sec", wireType)
			}
			m.Sec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Sec |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nsec", wireType)
			}
			m.Nsec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Nsec |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Result) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Error = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rows = append(m.Rows, Result_Row{})
			if err := m.Rows[len(m.Rows)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Result_Row) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Row: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Row: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, Datum{})
			if err := m.Values[len(m.Values)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sql", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sql = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Params = append(m.Params, Datum{})
			if err := m.Params[len(m.Params)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWire(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWire
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWire
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWire
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWire
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWire
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWire(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWire = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWire   = fmt.Errorf("proto: integer overflow")
)
