/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/vector3.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Vector3", Vector3TypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/Vector3", Vector3TypeSupport)
}

// Do not create instances of this type directly. Always use NewVector3
// function instead.
type Vector3 struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
}

// NewVector3 creates a new Vector3 with default values.
func NewVector3() *Vector3 {
	self := Vector3{}
	self.SetDefaults()
	return &self
}

func (t *Vector3) Clone() *Vector3 {
	c := &Vector3{}
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	return c
}

func (t *Vector3) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Vector3) SetDefaults() {
	t.X = 0
	t.Y = 0
	t.Z = 0
}

func (t *Vector3) GetTypeSupport() types.MessageTypeSupport {
	return Vector3TypeSupport
}

// Vector3Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type Vector3Publisher struct {
	*rclgo.Publisher
}

// NewVector3Publisher creates and returns a new publisher for the
// Vector3
func NewVector3Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Vector3Publisher, error) {
	pub, err := node.NewPublisher(topic_name, Vector3TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Vector3Publisher{pub}, nil
}

func (p *Vector3Publisher) Publish(msg *Vector3) error {
	return p.Publisher.Publish(msg)
}

// Vector3Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type Vector3Subscription struct {
	*rclgo.Subscription
}

// Vector3SubscriptionCallback type is used to provide a subscription
// handler function for a Vector3Subscription.
type Vector3SubscriptionCallback func(msg *Vector3, info *rclgo.RmwMessageInfo, err error)

// NewVector3Subscription creates and returns a new subscription for the
// Vector3
func NewVector3Subscription(node *rclgo.Node, topic_name string, subscriptionCallback Vector3SubscriptionCallback) (*Vector3Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Vector3
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Vector3TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Vector3Subscription{sub}, nil
}

func (s *Vector3Subscription) TakeMessage(out *Vector3) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneVector3Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVector3Slice(dst, src []Vector3) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Vector3TypeSupport types.MessageTypeSupport = _Vector3TypeSupport{}

type _Vector3TypeSupport struct{}

func (t _Vector3TypeSupport) New() types.Message {
	return NewVector3()
}

func (t _Vector3TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Vector3
	return (unsafe.Pointer)(C.geometry_msgs__msg__Vector3__create())
}

func (t _Vector3TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Vector3__destroy((*C.geometry_msgs__msg__Vector3)(pointer_to_free))
}

func (t _Vector3TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Vector3)
	mem := (*C.geometry_msgs__msg__Vector3)(dst)
	mem.x = C.double(m.X)
	mem.y = C.double(m.Y)
	mem.z = C.double(m.Z)
}

func (t _Vector3TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Vector3)
	mem := (*C.geometry_msgs__msg__Vector3)(ros2_message_buffer)
	m.X = float64(mem.x)
	m.Y = float64(mem.y)
	m.Z = float64(mem.z)
}

func (t _Vector3TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Vector3())
}

type CVector3 = C.geometry_msgs__msg__Vector3
type CVector3__Sequence = C.geometry_msgs__msg__Vector3__Sequence

func Vector3__Sequence_to_Go(goSlice *[]Vector3, cSlice CVector3__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Vector3, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Vector3TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Vector3__Sequence_to_C(cSlice *CVector3__Sequence, goSlice []Vector3) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Vector3)(C.malloc(C.sizeof_struct_geometry_msgs__msg__Vector3 * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Vector3TypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Vector3__Array_to_Go(goSlice []Vector3, cSlice []CVector3) {
	for i := 0; i < len(cSlice); i++ {
		Vector3TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Vector3__Array_to_C(cSlice []CVector3, goSlice []Vector3) {
	for i := 0; i < len(goSlice); i++ {
		Vector3TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
