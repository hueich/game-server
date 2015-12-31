// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/duration_proto/duration.proto
// DO NOT EDIT!

/*
Package google_protobuf is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/duration_proto/duration.proto

It has these top-level messages:
	Duration
*/
package google_protobuf

import proto "github.com/hueich/game-server/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A Duration represents a signed, fixed-length span of time represented
// as a count of seconds and fractions of seconds at nanosecond
// resolution. It is independent of any calendar and concepts like "day"
// or "month". It is related to Timestamp in that the difference between
// two Timestamp values is a Duration and it can be added or subtracted
// from a Timestamp. Range is approximately +-10,000 years.
//
// Example 1: Compute Duration from two Timestamps in pseudo code.
//
//     Timestamp start = ...;
//     Timestamp end = ...;
//     Duration duration = ...;
//
//     duration.seconds = end.seconds - start.seconds;
//     duration.nanos = end.nanos - start.nanos;
//
//     if (duration.seconds < 0 && duration.nanos > 0) {
//       duration.seconds += 1;
//       duration.nanos -= 1000000000;
//     } else if (durations.seconds > 0 && duration.nanos < 0) {
//       duration.seconds -= 1;
//       duration.nanos += 1000000000;
//     }
//
// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
//
//     Timestamp start = ...;
//     Duration duration = ...;
//     Timestamp end = ...;
//
//     end.seconds = start.seconds + duration.seconds;
//     end.nanos = start.nanos + duration.nanos;
//
//     if (end.nanos < 0) {
//       end.seconds -= 1;
//       end.nanos += 1000000000;
//     } else if (end.nanos >= 1000000000) {
//       end.seconds += 1;
//       end.nanos -= 1000000000;
//     }
//
type Duration struct {
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Duration) Reset()                    { *m = Duration{} }
func (m *Duration) String() string            { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Duration)(nil), "google.protobuf.Duration")
}

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0xca, 0x4c, 0x2f, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0xcf, 0xcc, 0x2b,
	0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0x29, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0x87, 0x73, 0xf5, 0xc0, 0x5c, 0x21, 0x7e, 0xa8, 0x41, 0x60, 0x5e, 0x52,
	0x69, 0x9a, 0x92, 0x16, 0x17, 0x87, 0x0b, 0x54, 0x89, 0x10, 0x3f, 0x17, 0x7b, 0x71, 0x6a, 0x72,
	0x7e, 0x5e, 0x4a, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb3, 0x10, 0x2f, 0x17, 0x6b, 0x5e, 0x62,
	0x5e, 0x7e, 0xb1, 0x04, 0x13, 0x90, 0xcb, 0xea, 0xa4, 0xc9, 0x25, 0x9c, 0x9c, 0x9f, 0xab, 0x87,
	0x66, 0x84, 0x13, 0x2f, 0xcc, 0x80, 0x00, 0x90, 0x48, 0x00, 0xe3, 0x02, 0x46, 0xc6, 0x24, 0x36,
	0xb0, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x14, 0xb7, 0x46, 0xb9, 0x00, 0x00, 0x00,
}
