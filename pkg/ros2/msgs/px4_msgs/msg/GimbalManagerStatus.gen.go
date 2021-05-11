/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/gimbal_manager_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/GimbalManagerStatus", &GimbalManagerStatus{})
}

// Do not create instances of this type directly. Always use NewGimbalManagerStatus
// function instead.
type GimbalManagerStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Flags uint32 `yaml:"flags"`
	GimbalDeviceId uint8 `yaml:"gimbal_device_id"`
	PrimaryControlSysid uint8 `yaml:"primary_control_sysid"`
	PrimaryControlCompid uint8 `yaml:"primary_control_compid"`
	SecondaryControlSysid uint8 `yaml:"secondary_control_sysid"`
	SecondaryControlCompid uint8 `yaml:"secondary_control_compid"`
}

// NewGimbalManagerStatus creates a new GimbalManagerStatus with default values.
func NewGimbalManagerStatus() *GimbalManagerStatus {
	self := GimbalManagerStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *GimbalManagerStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *GimbalManagerStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__GimbalManagerStatus())
}
func (t *GimbalManagerStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__GimbalManagerStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__GimbalManagerStatus__create())
}
func (t *GimbalManagerStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__GimbalManagerStatus__destroy((*C.px4_msgs__msg__GimbalManagerStatus)(pointer_to_free))
}
func (t *GimbalManagerStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__GimbalManagerStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.flags = C.uint32_t(t.Flags)
	mem.gimbal_device_id = C.uint8_t(t.GimbalDeviceId)
	mem.primary_control_sysid = C.uint8_t(t.PrimaryControlSysid)
	mem.primary_control_compid = C.uint8_t(t.PrimaryControlCompid)
	mem.secondary_control_sysid = C.uint8_t(t.SecondaryControlSysid)
	mem.secondary_control_compid = C.uint8_t(t.SecondaryControlCompid)
	return unsafe.Pointer(mem)
}
func (t *GimbalManagerStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__GimbalManagerStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Flags = uint32(mem.flags)
	t.GimbalDeviceId = uint8(mem.gimbal_device_id)
	t.PrimaryControlSysid = uint8(mem.primary_control_sysid)
	t.PrimaryControlCompid = uint8(mem.primary_control_compid)
	t.SecondaryControlSysid = uint8(mem.secondary_control_sysid)
	t.SecondaryControlCompid = uint8(mem.secondary_control_compid)
}
func (t *GimbalManagerStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CGimbalManagerStatus = C.px4_msgs__msg__GimbalManagerStatus
type CGimbalManagerStatus__Sequence = C.px4_msgs__msg__GimbalManagerStatus__Sequence

func GimbalManagerStatus__Sequence_to_Go(goSlice *[]GimbalManagerStatus, cSlice CGimbalManagerStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GimbalManagerStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__GimbalManagerStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerStatus * uintptr(i)),
		))
		(*goSlice)[i] = GimbalManagerStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func GimbalManagerStatus__Sequence_to_C(cSlice *CGimbalManagerStatus__Sequence, goSlice []GimbalManagerStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__GimbalManagerStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__GimbalManagerStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__GimbalManagerStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__GimbalManagerStatus)(v.AsCStruct())
	}
}
func GimbalManagerStatus__Array_to_Go(goSlice []GimbalManagerStatus, cSlice []CGimbalManagerStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func GimbalManagerStatus__Array_to_C(cSlice []CGimbalManagerStatus, goSlice []GimbalManagerStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__GimbalManagerStatus)(goSlice[i].AsCStruct())
	}
}

