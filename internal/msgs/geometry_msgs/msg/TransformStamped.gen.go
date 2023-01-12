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
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/transform_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/TransformStamped", TransformStampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/TransformStamped", TransformStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewTransformStamped
// function instead.
type TransformStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`// The frame id in the header is used as the reference frame of this transform.
	ChildFrameId string `yaml:"child_frame_id"`// The frame id of the child frame to which this transform points.
	Transform Transform `yaml:"transform"`// Translation and rotation in 3-dimensions of child_frame_id from header.frame_id.
}

// NewTransformStamped creates a new TransformStamped with default values.
func NewTransformStamped() *TransformStamped {
	self := TransformStamped{}
	self.SetDefaults()
	return &self
}

func (t *TransformStamped) Clone() *TransformStamped {
	c := &TransformStamped{}
	c.Header = *t.Header.Clone()
	c.ChildFrameId = t.ChildFrameId
	c.Transform = *t.Transform.Clone()
	return c
}

func (t *TransformStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TransformStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.ChildFrameId = ""
	t.Transform.SetDefaults()
}

func (t *TransformStamped) GetTypeSupport() types.MessageTypeSupport {
	return TransformStampedTypeSupport
}

// TransformStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TransformStampedPublisher struct {
	*rclgo.Publisher
}

// NewTransformStampedPublisher creates and returns a new publisher for the
// TransformStamped
func NewTransformStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TransformStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, TransformStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TransformStampedPublisher{pub}, nil
}

func (p *TransformStampedPublisher) Publish(msg *TransformStamped) error {
	return p.Publisher.Publish(msg)
}

// TransformStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TransformStampedSubscription struct {
	*rclgo.Subscription
}

// TransformStampedSubscriptionCallback type is used to provide a subscription
// handler function for a TransformStampedSubscription.
type TransformStampedSubscriptionCallback func(msg *TransformStamped, info *rclgo.RmwMessageInfo, err error)

// NewTransformStampedSubscription creates and returns a new subscription for the
// TransformStamped
func NewTransformStampedSubscription(node *rclgo.Node, topic_name string, subscriptionCallback TransformStampedSubscriptionCallback) (*TransformStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg TransformStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TransformStampedTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &TransformStampedSubscription{sub}, nil
}

func (s *TransformStampedSubscription) TakeMessage(out *TransformStamped) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTransformStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTransformStampedSlice(dst, src []TransformStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TransformStampedTypeSupport types.MessageTypeSupport = _TransformStampedTypeSupport{}

type _TransformStampedTypeSupport struct{}

func (t _TransformStampedTypeSupport) New() types.Message {
	return NewTransformStamped()
}

func (t _TransformStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__TransformStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__TransformStamped__create())
}

func (t _TransformStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__TransformStamped__destroy((*C.geometry_msgs__msg__TransformStamped)(pointer_to_free))
}

func (t _TransformStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TransformStamped)
	mem := (*C.geometry_msgs__msg__TransformStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.child_frame_id), m.ChildFrameId)
	TransformTypeSupport.AsCStruct(unsafe.Pointer(&mem.transform), &m.Transform)
}

func (t _TransformStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TransformStamped)
	mem := (*C.geometry_msgs__msg__TransformStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.StringAsGoStruct(&m.ChildFrameId, unsafe.Pointer(&mem.child_frame_id))
	TransformTypeSupport.AsGoStruct(&m.Transform, unsafe.Pointer(&mem.transform))
}

func (t _TransformStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__TransformStamped())
}

type CTransformStamped = C.geometry_msgs__msg__TransformStamped
type CTransformStamped__Sequence = C.geometry_msgs__msg__TransformStamped__Sequence

func TransformStamped__Sequence_to_Go(goSlice *[]TransformStamped, cSlice CTransformStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TransformStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		TransformStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func TransformStamped__Sequence_to_C(cSlice *CTransformStamped__Sequence, goSlice []TransformStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__TransformStamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__TransformStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		TransformStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func TransformStamped__Array_to_Go(goSlice []TransformStamped, cSlice []CTransformStamped) {
	for i := 0; i < len(cSlice); i++ {
		TransformStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TransformStamped__Array_to_C(cSlice []CTransformStamped, goSlice []TransformStamped) {
	for i := 0; i < len(goSlice); i++ {
		TransformStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
