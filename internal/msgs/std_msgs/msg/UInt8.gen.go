/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int8.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/UInt8", UInt8TypeSupport)
	typemap.RegisterMessage("std_msgs/msg/UInt8", UInt8TypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt8
// function instead.
type UInt8 struct {
	Data uint8 `yaml:"data"`
}

// NewUInt8 creates a new UInt8 with default values.
func NewUInt8() *UInt8 {
	self := UInt8{}
	self.SetDefaults()
	return &self
}

func (t *UInt8) Clone() *UInt8 {
	c := &UInt8{}
	c.Data = t.Data
	return c
}

func (t *UInt8) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt8) SetDefaults() {
	t.Data = 0
}

func (t *UInt8) GetTypeSupport() types.MessageTypeSupport {
	return UInt8TypeSupport
}

// UInt8Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type UInt8Publisher struct {
	*rclgo.Publisher
}

// NewUInt8Publisher creates and returns a new publisher for the
// UInt8
func NewUInt8Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*UInt8Publisher, error) {
	pub, err := node.NewPublisher(topic_name, UInt8TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &UInt8Publisher{pub}, nil
}

func (p *UInt8Publisher) Publish(msg *UInt8) error {
	return p.Publisher.Publish(msg)
}

// UInt8Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type UInt8Subscription struct {
	*rclgo.Subscription
}

// UInt8SubscriptionCallback type is used to provide a subscription
// handler function for a UInt8Subscription.
type UInt8SubscriptionCallback func(msg *UInt8, info *rclgo.MessageInfo, err error)

// NewUInt8Subscription creates and returns a new subscription for the
// UInt8
func NewUInt8Subscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback UInt8SubscriptionCallback) (*UInt8Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg UInt8
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, UInt8TypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &UInt8Subscription{sub}, nil
}

func (s *UInt8Subscription) TakeMessage(out *UInt8) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneUInt8Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt8Slice(dst, src []UInt8) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt8TypeSupport types.MessageTypeSupport = _UInt8TypeSupport{}

type _UInt8TypeSupport struct{}

func (t _UInt8TypeSupport) New() types.Message {
	return NewUInt8()
}

func (t _UInt8TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt8
	return (unsafe.Pointer)(C.std_msgs__msg__UInt8__create())
}

func (t _UInt8TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt8__destroy((*C.std_msgs__msg__UInt8)(pointer_to_free))
}

func (t _UInt8TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt8)
	mem := (*C.std_msgs__msg__UInt8)(dst)
	mem.data = C.uint8_t(m.Data)
}

func (t _UInt8TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt8)
	mem := (*C.std_msgs__msg__UInt8)(ros2_message_buffer)
	m.Data = uint8(mem.data)
}

func (t _UInt8TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt8())
}

type CUInt8 = C.std_msgs__msg__UInt8
type CUInt8__Sequence = C.std_msgs__msg__UInt8__Sequence

func UInt8__Sequence_to_Go(goSlice *[]UInt8, cSlice CUInt8__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt8, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		UInt8TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func UInt8__Sequence_to_C(cSlice *CUInt8__Sequence, goSlice []UInt8) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt8)(C.malloc(C.sizeof_struct_std_msgs__msg__UInt8 * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		UInt8TypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func UInt8__Array_to_Go(goSlice []UInt8, cSlice []CUInt8) {
	for i := 0; i < len(cSlice); i++ {
		UInt8TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt8__Array_to_C(cSlice []CUInt8, goSlice []UInt8) {
	for i := 0; i < len(goSlice); i++ {
		UInt8TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
