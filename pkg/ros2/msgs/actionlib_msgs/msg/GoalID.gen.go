/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package actionlib_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	builtin_interfaces "github.com/tiiuae/rclgo/pkg/ros2/msgs/builtin_interfaces/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lactionlib_msgs__rosidl_typesupport_c -lactionlib_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <actionlib_msgs/msg/goal_id.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("actionlib_msgs/GoalID", &GoalID{})
}

// Do not create instances of this type directly. Always use NewGoalID
// function instead.
type GoalID struct {
	Stamp builtin_interfaces.Time `yaml:"stamp"`// The stamp should store the time at which this goal was requested.It is used by an action server when it tries to preempt allgoals that were requested before a certain time
	Id rosidl_runtime_c.String `yaml:"id"`// The id provides a way to associate feedback andresult message with specific goal requests. The idspecified must be unique.
}

// NewGoalID creates a new GoalID with default values.
func NewGoalID() *GoalID {
	self := GoalID{}
	self.SetDefaults(nil)
	return &self
}

func (t *GoalID) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Stamp.SetDefaults(nil)
	t.Id.SetDefaults("")
	
	return t
}

func (t *GoalID) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__actionlib_msgs__msg__GoalID())
}
func (t *GoalID) PrepareMemory() unsafe.Pointer { //returns *C.actionlib_msgs__msg__GoalID
	return (unsafe.Pointer)(C.actionlib_msgs__msg__GoalID__create())
}
func (t *GoalID) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.actionlib_msgs__msg__GoalID__destroy((*C.actionlib_msgs__msg__GoalID)(pointer_to_free))
}
func (t *GoalID) AsCStruct() unsafe.Pointer {
	mem := (*C.actionlib_msgs__msg__GoalID)(t.PrepareMemory())
	mem.stamp = *(*C.builtin_interfaces__msg__Time)(t.Stamp.AsCStruct())
	mem.id = *(*C.rosidl_runtime_c__String)(t.Id.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *GoalID) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.actionlib_msgs__msg__GoalID)(ros2_message_buffer)
	t.Stamp.AsGoStruct(unsafe.Pointer(&mem.stamp))
	t.Id.AsGoStruct(unsafe.Pointer(&mem.id))
}
func (t *GoalID) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CGoalID = C.actionlib_msgs__msg__GoalID
type CGoalID__Sequence = C.actionlib_msgs__msg__GoalID__Sequence

func GoalID__Sequence_to_Go(goSlice *[]GoalID, cSlice CGoalID__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GoalID, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.actionlib_msgs__msg__GoalID__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_actionlib_msgs__msg__GoalID * uintptr(i)),
		))
		(*goSlice)[i] = GoalID{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func GoalID__Sequence_to_C(cSlice *CGoalID__Sequence, goSlice []GoalID) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.actionlib_msgs__msg__GoalID)(C.malloc((C.size_t)(C.sizeof_struct_actionlib_msgs__msg__GoalID * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.actionlib_msgs__msg__GoalID)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_actionlib_msgs__msg__GoalID * uintptr(i)),
		))
		*cIdx = *(*C.actionlib_msgs__msg__GoalID)(v.AsCStruct())
	}
}
func GoalID__Array_to_Go(goSlice []GoalID, cSlice []CGoalID) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func GoalID__Array_to_C(cSlice []CGoalID, goSlice []GoalID) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.actionlib_msgs__msg__GoalID)(goSlice[i].AsCStruct())
	}
}

