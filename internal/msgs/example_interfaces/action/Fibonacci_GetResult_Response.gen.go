/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/action/fibonacci.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Fibonacci_GetResult_Response", Fibonacci_GetResult_ResponseTypeSupport)
	typemap.RegisterMessage("example_interfaces/action/Fibonacci_GetResult_Response", Fibonacci_GetResult_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewFibonacci_GetResult_Response
// function instead.
type Fibonacci_GetResult_Response struct {
	Status int8 `yaml:"status"`
	Result Fibonacci_Result `yaml:"result"`
}

// NewFibonacci_GetResult_Response creates a new Fibonacci_GetResult_Response with default values.
func NewFibonacci_GetResult_Response() *Fibonacci_GetResult_Response {
	self := Fibonacci_GetResult_Response{}
	self.SetDefaults()
	return &self
}

func (t *Fibonacci_GetResult_Response) Clone() *Fibonacci_GetResult_Response {
	c := &Fibonacci_GetResult_Response{}
	c.Status = t.Status
	c.Result = *t.Result.Clone()
	return c
}

func (t *Fibonacci_GetResult_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Fibonacci_GetResult_Response) SetDefaults() {
	t.Status = 0
	t.Result.SetDefaults()
}

func (t *Fibonacci_GetResult_Response) GetTypeSupport() types.MessageTypeSupport {
	return Fibonacci_GetResult_ResponseTypeSupport
}

// Fibonacci_GetResult_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Fibonacci_GetResult_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewFibonacci_GetResult_ResponsePublisher creates and returns a new publisher for the
// Fibonacci_GetResult_Response
func NewFibonacci_GetResult_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Fibonacci_GetResult_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, Fibonacci_GetResult_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_GetResult_ResponsePublisher{pub}, nil
}

func (p *Fibonacci_GetResult_ResponsePublisher) Publish(msg *Fibonacci_GetResult_Response) error {
	return p.Publisher.Publish(msg)
}

// Fibonacci_GetResult_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Fibonacci_GetResult_ResponseSubscription struct {
	*rclgo.Subscription
}

// Fibonacci_GetResult_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a Fibonacci_GetResult_ResponseSubscription.
type Fibonacci_GetResult_ResponseSubscriptionCallback func(msg *Fibonacci_GetResult_Response, info *rclgo.MessageInfo, err error)

// NewFibonacci_GetResult_ResponseSubscription creates and returns a new subscription for the
// Fibonacci_GetResult_Response
func NewFibonacci_GetResult_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback Fibonacci_GetResult_ResponseSubscriptionCallback) (*Fibonacci_GetResult_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Fibonacci_GetResult_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Fibonacci_GetResult_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_GetResult_ResponseSubscription{sub}, nil
}

func (s *Fibonacci_GetResult_ResponseSubscription) TakeMessage(out *Fibonacci_GetResult_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFibonacci_GetResult_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFibonacci_GetResult_ResponseSlice(dst, src []Fibonacci_GetResult_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Fibonacci_GetResult_ResponseTypeSupport types.MessageTypeSupport = _Fibonacci_GetResult_ResponseTypeSupport{}

type _Fibonacci_GetResult_ResponseTypeSupport struct{}

func (t _Fibonacci_GetResult_ResponseTypeSupport) New() types.Message {
	return NewFibonacci_GetResult_Response()
}

func (t _Fibonacci_GetResult_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__action__Fibonacci_GetResult_Response
	return (unsafe.Pointer)(C.example_interfaces__action__Fibonacci_GetResult_Response__create())
}

func (t _Fibonacci_GetResult_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__action__Fibonacci_GetResult_Response__destroy((*C.example_interfaces__action__Fibonacci_GetResult_Response)(pointer_to_free))
}

func (t _Fibonacci_GetResult_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Fibonacci_GetResult_Response)
	mem := (*C.example_interfaces__action__Fibonacci_GetResult_Response)(dst)
	mem.status = C.int8_t(m.Status)
	Fibonacci_ResultTypeSupport.AsCStruct(unsafe.Pointer(&mem.result), &m.Result)
}

func (t _Fibonacci_GetResult_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Fibonacci_GetResult_Response)
	mem := (*C.example_interfaces__action__Fibonacci_GetResult_Response)(ros2_message_buffer)
	m.Status = int8(mem.status)
	Fibonacci_ResultTypeSupport.AsGoStruct(&m.Result, unsafe.Pointer(&mem.result))
}

func (t _Fibonacci_GetResult_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__action__Fibonacci_GetResult_Response())
}

type CFibonacci_GetResult_Response = C.example_interfaces__action__Fibonacci_GetResult_Response
type CFibonacci_GetResult_Response__Sequence = C.example_interfaces__action__Fibonacci_GetResult_Response__Sequence

func Fibonacci_GetResult_Response__Sequence_to_Go(goSlice *[]Fibonacci_GetResult_Response, cSlice CFibonacci_GetResult_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Fibonacci_GetResult_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Fibonacci_GetResult_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Fibonacci_GetResult_Response__Sequence_to_C(cSlice *CFibonacci_GetResult_Response__Sequence, goSlice []Fibonacci_GetResult_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.example_interfaces__action__Fibonacci_GetResult_Response)(C.malloc(C.sizeof_struct_example_interfaces__action__Fibonacci_GetResult_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Fibonacci_GetResult_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Fibonacci_GetResult_Response__Array_to_Go(goSlice []Fibonacci_GetResult_Response, cSlice []CFibonacci_GetResult_Response) {
	for i := 0; i < len(cSlice); i++ {
		Fibonacci_GetResult_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Fibonacci_GetResult_Response__Array_to_C(cSlice []CFibonacci_GetResult_Response, goSlice []Fibonacci_GetResult_Response) {
	for i := 0; i < len(goSlice); i++ {
		Fibonacci_GetResult_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
