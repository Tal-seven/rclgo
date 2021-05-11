/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <geometry_msgs/msg/pose_array.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/PoseArray", &PoseArray{})
}

// Do not create instances of this type directly. Always use NewPoseArray
// function instead.
type PoseArray struct {
	Header std_msgs.Header `yaml:"header"`
	Poses []Pose `yaml:"poses"`
}

// NewPoseArray creates a new PoseArray with default values.
func NewPoseArray() *PoseArray {
	self := PoseArray{}
	self.SetDefaults(nil)
	return &self
}

func (t *PoseArray) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	
	return t
}

func (t *PoseArray) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PoseArray())
}
func (t *PoseArray) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PoseArray
	return (unsafe.Pointer)(C.geometry_msgs__msg__PoseArray__create())
}
func (t *PoseArray) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PoseArray__destroy((*C.geometry_msgs__msg__PoseArray)(pointer_to_free))
}
func (t *PoseArray) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__PoseArray)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	Pose__Sequence_to_C(&mem.poses, t.Poses)
	return unsafe.Pointer(mem)
}
func (t *PoseArray) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__PoseArray)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	Pose__Sequence_to_Go(&t.Poses, mem.poses)
}
func (t *PoseArray) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CPoseArray = C.geometry_msgs__msg__PoseArray
type CPoseArray__Sequence = C.geometry_msgs__msg__PoseArray__Sequence

func PoseArray__Sequence_to_Go(goSlice *[]PoseArray, cSlice CPoseArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PoseArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__PoseArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PoseArray * uintptr(i)),
		))
		(*goSlice)[i] = PoseArray{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func PoseArray__Sequence_to_C(cSlice *CPoseArray__Sequence, goSlice []PoseArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PoseArray)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__PoseArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__PoseArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PoseArray * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__PoseArray)(v.AsCStruct())
	}
}
func PoseArray__Array_to_Go(goSlice []PoseArray, cSlice []CPoseArray) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func PoseArray__Array_to_C(cSlice []CPoseArray, goSlice []PoseArray) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__PoseArray)(goSlice[i].AsCStruct())
	}
}

