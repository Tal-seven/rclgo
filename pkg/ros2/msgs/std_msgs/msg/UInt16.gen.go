/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_msgs/msg/u_int16.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/UInt16", &UInt16{})
}

// Do not create instances of this type directly. Always use NewUInt16
// function instead.
type UInt16 struct {
	Data uint16 `yaml:"data"`
}

// NewUInt16 creates a new UInt16 with default values.
func NewUInt16() *UInt16 {
	self := UInt16{}
	self.SetDefaults(nil)
	return &self
}

func (t *UInt16) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *UInt16) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt16())
}
func (t *UInt16) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt16
	return (unsafe.Pointer)(C.std_msgs__msg__UInt16__create())
}
func (t *UInt16) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt16__destroy((*C.std_msgs__msg__UInt16)(pointer_to_free))
}
func (t *UInt16) AsCStruct() unsafe.Pointer {
	mem := (*C.std_msgs__msg__UInt16)(t.PrepareMemory())
	mem.data = C.uint16_t(t.Data)
	return unsafe.Pointer(mem)
}
func (t *UInt16) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.std_msgs__msg__UInt16)(ros2_message_buffer)
	t.Data = uint16(mem.data)
}
func (t *UInt16) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUInt16 = C.std_msgs__msg__UInt16
type CUInt16__Sequence = C.std_msgs__msg__UInt16__Sequence

func UInt16__Sequence_to_Go(goSlice *[]UInt16, cSlice CUInt16__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt16, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt16__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt16 * uintptr(i)),
		))
		(*goSlice)[i] = UInt16{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UInt16__Sequence_to_C(cSlice *CUInt16__Sequence, goSlice []UInt16) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt16)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt16 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt16)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt16 * uintptr(i)),
		))
		*cIdx = *(*C.std_msgs__msg__UInt16)(v.AsCStruct())
	}
}
func UInt16__Array_to_Go(goSlice []UInt16, cSlice []CUInt16) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UInt16__Array_to_C(cSlice []CUInt16, goSlice []UInt16) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.std_msgs__msg__UInt16)(goSlice[i].AsCStruct())
	}
}

