/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/add_two_ints.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/AddTwoInts_Response", AddTwoInts_ResponseTypeSupport)
	typemap.RegisterMessage("example_interfaces/srv/AddTwoInts_Response", AddTwoInts_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewAddTwoInts_Response
// function instead.
type AddTwoInts_Response struct {
	Sum int64 `yaml:"sum"`
}

// NewAddTwoInts_Response creates a new AddTwoInts_Response with default values.
func NewAddTwoInts_Response() *AddTwoInts_Response {
	self := AddTwoInts_Response{}
	self.SetDefaults()
	return &self
}

func (t *AddTwoInts_Response) Clone() *AddTwoInts_Response {
	c := &AddTwoInts_Response{}
	c.Sum = t.Sum
	return c
}

func (t *AddTwoInts_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *AddTwoInts_Response) SetDefaults() {
	t.Sum = 0
}

func (t *AddTwoInts_Response) GetTypeSupport() types.MessageTypeSupport {
	return AddTwoInts_ResponseTypeSupport
}

// AddTwoInts_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type AddTwoInts_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewAddTwoInts_ResponsePublisher creates and returns a new publisher for the
// AddTwoInts_Response
func NewAddTwoInts_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*AddTwoInts_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, AddTwoInts_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &AddTwoInts_ResponsePublisher{pub}, nil
}

func (p *AddTwoInts_ResponsePublisher) Publish(msg *AddTwoInts_Response) error {
	return p.Publisher.Publish(msg)
}

// AddTwoInts_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type AddTwoInts_ResponseSubscription struct {
	*rclgo.Subscription
}

// AddTwoInts_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a AddTwoInts_ResponseSubscription.
type AddTwoInts_ResponseSubscriptionCallback func(msg *AddTwoInts_Response, info *rclgo.RmwMessageInfo, err error)

// NewAddTwoInts_ResponseSubscription creates and returns a new subscription for the
// AddTwoInts_Response
func NewAddTwoInts_ResponseSubscription(node *rclgo.Node, topic_name string, subscriptionCallback AddTwoInts_ResponseSubscriptionCallback) (*AddTwoInts_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg AddTwoInts_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, AddTwoInts_ResponseTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &AddTwoInts_ResponseSubscription{sub}, nil
}

func (s *AddTwoInts_ResponseSubscription) TakeMessage(out *AddTwoInts_Response) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneAddTwoInts_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAddTwoInts_ResponseSlice(dst, src []AddTwoInts_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AddTwoInts_ResponseTypeSupport types.MessageTypeSupport = _AddTwoInts_ResponseTypeSupport{}

type _AddTwoInts_ResponseTypeSupport struct{}

func (t _AddTwoInts_ResponseTypeSupport) New() types.Message {
	return NewAddTwoInts_Response()
}

func (t _AddTwoInts_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__AddTwoInts_Response
	return (unsafe.Pointer)(C.example_interfaces__srv__AddTwoInts_Response__create())
}

func (t _AddTwoInts_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__AddTwoInts_Response__destroy((*C.example_interfaces__srv__AddTwoInts_Response)(pointer_to_free))
}

func (t _AddTwoInts_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AddTwoInts_Response)
	mem := (*C.example_interfaces__srv__AddTwoInts_Response)(dst)
	mem.sum = C.int64_t(m.Sum)
}

func (t _AddTwoInts_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AddTwoInts_Response)
	mem := (*C.example_interfaces__srv__AddTwoInts_Response)(ros2_message_buffer)
	m.Sum = int64(mem.sum)
}

func (t _AddTwoInts_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__AddTwoInts_Response())
}

type CAddTwoInts_Response = C.example_interfaces__srv__AddTwoInts_Response
type CAddTwoInts_Response__Sequence = C.example_interfaces__srv__AddTwoInts_Response__Sequence

func AddTwoInts_Response__Sequence_to_Go(goSlice *[]AddTwoInts_Response, cSlice CAddTwoInts_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AddTwoInts_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		AddTwoInts_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func AddTwoInts_Response__Sequence_to_C(cSlice *CAddTwoInts_Response__Sequence, goSlice []AddTwoInts_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.example_interfaces__srv__AddTwoInts_Response)(C.malloc(C.sizeof_struct_example_interfaces__srv__AddTwoInts_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		AddTwoInts_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func AddTwoInts_Response__Array_to_Go(goSlice []AddTwoInts_Response, cSlice []CAddTwoInts_Response) {
	for i := 0; i < len(cSlice); i++ {
		AddTwoInts_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AddTwoInts_Response__Array_to_C(cSlice []CAddTwoInts_Response, goSlice []AddTwoInts_Response) {
	for i := 0; i < len(goSlice); i++ {
		AddTwoInts_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
