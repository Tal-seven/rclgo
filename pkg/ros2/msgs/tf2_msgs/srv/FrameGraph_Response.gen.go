/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package tf2_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltf2_msgs__rosidl_typesupport_c -ltf2_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <tf2_msgs/srv/frame_graph.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("tf2_msgs/FrameGraph_Response", &FrameGraph_Response{})
}

// Do not create instances of this type directly. Always use NewFrameGraph_Response
// function instead.
type FrameGraph_Response struct {
	FrameYaml rosidl_runtime_c.String `yaml:"frame_yaml"`
}

// NewFrameGraph_Response creates a new FrameGraph_Response with default values.
func NewFrameGraph_Response() *FrameGraph_Response {
	self := FrameGraph_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *FrameGraph_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.FrameYaml.SetDefaults("")
	
	return t
}

func (t *FrameGraph_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__tf2_msgs__srv__FrameGraph_Response())
}
func (t *FrameGraph_Response) PrepareMemory() unsafe.Pointer { //returns *C.tf2_msgs__srv__FrameGraph_Response
	return (unsafe.Pointer)(C.tf2_msgs__srv__FrameGraph_Response__create())
}
func (t *FrameGraph_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.tf2_msgs__srv__FrameGraph_Response__destroy((*C.tf2_msgs__srv__FrameGraph_Response)(pointer_to_free))
}
func (t *FrameGraph_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.tf2_msgs__srv__FrameGraph_Response)(t.PrepareMemory())
	mem.frame_yaml = *(*C.rosidl_runtime_c__String)(t.FrameYaml.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *FrameGraph_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.tf2_msgs__srv__FrameGraph_Response)(ros2_message_buffer)
	t.FrameYaml.AsGoStruct(unsafe.Pointer(&mem.frame_yaml))
}
func (t *FrameGraph_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CFrameGraph_Response = C.tf2_msgs__srv__FrameGraph_Response
type CFrameGraph_Response__Sequence = C.tf2_msgs__srv__FrameGraph_Response__Sequence

func FrameGraph_Response__Sequence_to_Go(goSlice *[]FrameGraph_Response, cSlice CFrameGraph_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FrameGraph_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.tf2_msgs__srv__FrameGraph_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_tf2_msgs__srv__FrameGraph_Response * uintptr(i)),
		))
		(*goSlice)[i] = FrameGraph_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func FrameGraph_Response__Sequence_to_C(cSlice *CFrameGraph_Response__Sequence, goSlice []FrameGraph_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.tf2_msgs__srv__FrameGraph_Response)(C.malloc((C.size_t)(C.sizeof_struct_tf2_msgs__srv__FrameGraph_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.tf2_msgs__srv__FrameGraph_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_tf2_msgs__srv__FrameGraph_Response * uintptr(i)),
		))
		*cIdx = *(*C.tf2_msgs__srv__FrameGraph_Response)(v.AsCStruct())
	}
}
func FrameGraph_Response__Array_to_Go(goSlice []FrameGraph_Response, cSlice []CFrameGraph_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func FrameGraph_Response__Array_to_C(cSlice []CFrameGraph_Response, goSlice []FrameGraph_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.tf2_msgs__srv__FrameGraph_Response)(goSlice[i].AsCStruct())
	}
}

