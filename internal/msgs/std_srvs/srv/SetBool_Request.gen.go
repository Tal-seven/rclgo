/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_srvs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_srvs/srv/set_bool.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_srvs/SetBool_Request", SetBool_RequestTypeSupport)
	typemap.RegisterMessage("std_srvs/srv/SetBool_Request", SetBool_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetBool_Request
// function instead.
type SetBool_Request struct {
	Data bool `yaml:"data"`// e.g. for hardware enabling / disabling
}

// NewSetBool_Request creates a new SetBool_Request with default values.
func NewSetBool_Request() *SetBool_Request {
	self := SetBool_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetBool_Request) Clone() *SetBool_Request {
	c := &SetBool_Request{}
	c.Data = t.Data
	return c
}

func (t *SetBool_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetBool_Request) SetDefaults() {
	t.Data = false
}

func (t *SetBool_Request) GetTypeSupport() types.MessageTypeSupport {
	return SetBool_RequestTypeSupport
}

// SetBool_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetBool_RequestPublisher struct {
	*rclgo.Publisher
}

// NewSetBool_RequestPublisher creates and returns a new publisher for the
// SetBool_Request
func NewSetBool_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetBool_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetBool_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetBool_RequestPublisher{pub}, nil
}

func (p *SetBool_RequestPublisher) Publish(msg *SetBool_Request) error {
	return p.Publisher.Publish(msg)
}

// SetBool_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetBool_RequestSubscription struct {
	*rclgo.Subscription
}

// SetBool_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a SetBool_RequestSubscription.
type SetBool_RequestSubscriptionCallback func(msg *SetBool_Request, info *rclgo.RmwMessageInfo, err error)

// NewSetBool_RequestSubscription creates and returns a new subscription for the
// SetBool_Request
func NewSetBool_RequestSubscription(node *rclgo.Node, topic_name string, subscriptionCallback SetBool_RequestSubscriptionCallback) (*SetBool_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetBool_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetBool_RequestTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &SetBool_RequestSubscription{sub}, nil
}

func (s *SetBool_RequestSubscription) TakeMessage(out *SetBool_Request) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetBool_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetBool_RequestSlice(dst, src []SetBool_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetBool_RequestTypeSupport types.MessageTypeSupport = _SetBool_RequestTypeSupport{}

type _SetBool_RequestTypeSupport struct{}

func (t _SetBool_RequestTypeSupport) New() types.Message {
	return NewSetBool_Request()
}

func (t _SetBool_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_srvs__srv__SetBool_Request
	return (unsafe.Pointer)(C.std_srvs__srv__SetBool_Request__create())
}

func (t _SetBool_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_srvs__srv__SetBool_Request__destroy((*C.std_srvs__srv__SetBool_Request)(pointer_to_free))
}

func (t _SetBool_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetBool_Request)
	mem := (*C.std_srvs__srv__SetBool_Request)(dst)
	mem.data = C.bool(m.Data)
}

func (t _SetBool_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetBool_Request)
	mem := (*C.std_srvs__srv__SetBool_Request)(ros2_message_buffer)
	m.Data = bool(mem.data)
}

func (t _SetBool_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_srvs__srv__SetBool_Request())
}

type CSetBool_Request = C.std_srvs__srv__SetBool_Request
type CSetBool_Request__Sequence = C.std_srvs__srv__SetBool_Request__Sequence

func SetBool_Request__Sequence_to_Go(goSlice *[]SetBool_Request, cSlice CSetBool_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetBool_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetBool_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetBool_Request__Sequence_to_C(cSlice *CSetBool_Request__Sequence, goSlice []SetBool_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.std_srvs__srv__SetBool_Request)(C.malloc(C.sizeof_struct_std_srvs__srv__SetBool_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetBool_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetBool_Request__Array_to_Go(goSlice []SetBool_Request, cSlice []CSetBool_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetBool_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetBool_Request__Array_to_C(cSlice []CSetBool_Request, goSlice []SetBool_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetBool_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
