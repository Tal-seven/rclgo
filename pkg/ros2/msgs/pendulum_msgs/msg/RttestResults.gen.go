/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package pendulum_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	builtin_interfaces "github.com/tiiuae/rclgo/pkg/ros2/msgs/builtin_interfaces/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpendulum_msgs__rosidl_typesupport_c -lpendulum_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <pendulum_msgs/msg/rttest_results.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("pendulum_msgs/RttestResults", &RttestResults{})
}

// Do not create instances of this type directly. Always use NewRttestResults
// function instead.
type RttestResults struct {
	Stamp builtin_interfaces.Time `yaml:"stamp"`
	Command JointCommand `yaml:"command"`
	State JointState `yaml:"state"`
	CurLatency uint64 `yaml:"cur_latency"`
	MeanLatency float64 `yaml:"mean_latency"`
	MinLatency uint64 `yaml:"min_latency"`
	MaxLatency uint64 `yaml:"max_latency"`
	MinorPagefaults uint64 `yaml:"minor_pagefaults"`
	MajorPagefaults uint64 `yaml:"major_pagefaults"`
}

// NewRttestResults creates a new RttestResults with default values.
func NewRttestResults() *RttestResults {
	self := RttestResults{}
	self.SetDefaults(nil)
	return &self
}

func (t *RttestResults) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Stamp.SetDefaults(nil)
	t.Command.SetDefaults(nil)
	t.State.SetDefaults(nil)
	
	return t
}

func (t *RttestResults) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pendulum_msgs__msg__RttestResults())
}
func (t *RttestResults) PrepareMemory() unsafe.Pointer { //returns *C.pendulum_msgs__msg__RttestResults
	return (unsafe.Pointer)(C.pendulum_msgs__msg__RttestResults__create())
}
func (t *RttestResults) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pendulum_msgs__msg__RttestResults__destroy((*C.pendulum_msgs__msg__RttestResults)(pointer_to_free))
}
func (t *RttestResults) AsCStruct() unsafe.Pointer {
	mem := (*C.pendulum_msgs__msg__RttestResults)(t.PrepareMemory())
	mem.stamp = *(*C.builtin_interfaces__msg__Time)(t.Stamp.AsCStruct())
	mem.command = *(*C.pendulum_msgs__msg__JointCommand)(t.Command.AsCStruct())
	mem.state = *(*C.pendulum_msgs__msg__JointState)(t.State.AsCStruct())
	mem.cur_latency = C.uint64_t(t.CurLatency)
	mem.mean_latency = C.double(t.MeanLatency)
	mem.min_latency = C.uint64_t(t.MinLatency)
	mem.max_latency = C.uint64_t(t.MaxLatency)
	mem.minor_pagefaults = C.uint64_t(t.MinorPagefaults)
	mem.major_pagefaults = C.uint64_t(t.MajorPagefaults)
	return unsafe.Pointer(mem)
}
func (t *RttestResults) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.pendulum_msgs__msg__RttestResults)(ros2_message_buffer)
	t.Stamp.AsGoStruct(unsafe.Pointer(&mem.stamp))
	t.Command.AsGoStruct(unsafe.Pointer(&mem.command))
	t.State.AsGoStruct(unsafe.Pointer(&mem.state))
	t.CurLatency = uint64(mem.cur_latency)
	t.MeanLatency = float64(mem.mean_latency)
	t.MinLatency = uint64(mem.min_latency)
	t.MaxLatency = uint64(mem.max_latency)
	t.MinorPagefaults = uint64(mem.minor_pagefaults)
	t.MajorPagefaults = uint64(mem.major_pagefaults)
}
func (t *RttestResults) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CRttestResults = C.pendulum_msgs__msg__RttestResults
type CRttestResults__Sequence = C.pendulum_msgs__msg__RttestResults__Sequence

func RttestResults__Sequence_to_Go(goSlice *[]RttestResults, cSlice CRttestResults__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RttestResults, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pendulum_msgs__msg__RttestResults__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__RttestResults * uintptr(i)),
		))
		(*goSlice)[i] = RttestResults{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func RttestResults__Sequence_to_C(cSlice *CRttestResults__Sequence, goSlice []RttestResults) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pendulum_msgs__msg__RttestResults)(C.malloc((C.size_t)(C.sizeof_struct_pendulum_msgs__msg__RttestResults * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pendulum_msgs__msg__RttestResults)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__RttestResults * uintptr(i)),
		))
		*cIdx = *(*C.pendulum_msgs__msg__RttestResults)(v.AsCStruct())
	}
}
func RttestResults__Array_to_Go(goSlice []RttestResults, cSlice []CRttestResults) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func RttestResults__Array_to_C(cSlice []CRttestResults, goSlice []RttestResults) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.pendulum_msgs__msg__RttestResults)(goSlice[i].AsCStruct())
	}
}

