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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/px4_io_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/Px4IoStatus", &Px4IoStatus{})
}

// Do not create instances of this type directly. Always use NewPx4IoStatus
// function instead.
type Px4IoStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	FreeMemoryBytes uint16 `yaml:"free_memory_bytes"`
	VoltageV float32 `yaml:"voltage_v"`// Servo rail voltage in volts
	RssiV float32 `yaml:"rssi_v"`// RSSI pin voltage in volts
	StatusOutputsArmed bool `yaml:"status_outputs_armed"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusOverride bool `yaml:"status_override"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusRcOk bool `yaml:"status_rc_ok"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusRcPpm bool `yaml:"status_rc_ppm"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusRcDsm bool `yaml:"status_rc_dsm"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusRcSbus bool `yaml:"status_rc_sbus"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusFmuOk bool `yaml:"status_fmu_ok"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusRawPwm bool `yaml:"status_raw_pwm"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusMixerOk bool `yaml:"status_mixer_ok"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusArmSync bool `yaml:"status_arm_sync"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusInitOk bool `yaml:"status_init_ok"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusFailsafe bool `yaml:"status_failsafe"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusSafetyOff bool `yaml:"status_safety_off"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusFmuInitialized bool `yaml:"status_fmu_initialized"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusRcSt24 bool `yaml:"status_rc_st24"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	StatusRcSumd bool `yaml:"status_rc_sumd"`// PX4IO status flags (PX4IO_P_STATUS_FLAGS)
	AlarmVbattLow bool `yaml:"alarm_vbatt_low"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	AlarmTemperature bool `yaml:"alarm_temperature"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	AlarmServoCurrent bool `yaml:"alarm_servo_current"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	AlarmAccCurrent bool `yaml:"alarm_acc_current"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	AlarmFmuLost bool `yaml:"alarm_fmu_lost"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	AlarmRcLost bool `yaml:"alarm_rc_lost"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	AlarmPwmError bool `yaml:"alarm_pwm_error"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	AlarmVservoFault bool `yaml:"alarm_vservo_fault"`// PX4IO alarms (PX4IO_P_STATUS_ALARMS)
	ArmingIoArmOk bool `yaml:"arming_io_arm_ok"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingFmuArmed bool `yaml:"arming_fmu_armed"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingFmuPrearmed bool `yaml:"arming_fmu_prearmed"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingManualOverrideOk bool `yaml:"arming_manual_override_ok"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingFailsafeCustom bool `yaml:"arming_failsafe_custom"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingInairRestartOk bool `yaml:"arming_inair_restart_ok"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingAlwaysPwmEnable bool `yaml:"arming_always_pwm_enable"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingRcHandlingDisabled bool `yaml:"arming_rc_handling_disabled"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingLockdown bool `yaml:"arming_lockdown"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingForceFailsafe bool `yaml:"arming_force_failsafe"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingTerminationFailsafe bool `yaml:"arming_termination_failsafe"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	ArmingOverrideImmediate bool `yaml:"arming_override_immediate"`// PX4IO arming (PX4IO_P_SETUP_ARMING)
	Actuators [8]int16 `yaml:"actuators"`
	Servos [8]uint16 `yaml:"servos"`
	RawInputs [18]uint16 `yaml:"raw_inputs"`
}

// NewPx4IoStatus creates a new Px4IoStatus with default values.
func NewPx4IoStatus() *Px4IoStatus {
	self := Px4IoStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *Px4IoStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Px4IoStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Px4IoStatus())
}
func (t *Px4IoStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Px4IoStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__Px4IoStatus__create())
}
func (t *Px4IoStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Px4IoStatus__destroy((*C.px4_msgs__msg__Px4IoStatus)(pointer_to_free))
}
func (t *Px4IoStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__Px4IoStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.free_memory_bytes = C.uint16_t(t.FreeMemoryBytes)
	mem.voltage_v = C.float(t.VoltageV)
	mem.rssi_v = C.float(t.RssiV)
	mem.status_outputs_armed = C.bool(t.StatusOutputsArmed)
	mem.status_override = C.bool(t.StatusOverride)
	mem.status_rc_ok = C.bool(t.StatusRcOk)
	mem.status_rc_ppm = C.bool(t.StatusRcPpm)
	mem.status_rc_dsm = C.bool(t.StatusRcDsm)
	mem.status_rc_sbus = C.bool(t.StatusRcSbus)
	mem.status_fmu_ok = C.bool(t.StatusFmuOk)
	mem.status_raw_pwm = C.bool(t.StatusRawPwm)
	mem.status_mixer_ok = C.bool(t.StatusMixerOk)
	mem.status_arm_sync = C.bool(t.StatusArmSync)
	mem.status_init_ok = C.bool(t.StatusInitOk)
	mem.status_failsafe = C.bool(t.StatusFailsafe)
	mem.status_safety_off = C.bool(t.StatusSafetyOff)
	mem.status_fmu_initialized = C.bool(t.StatusFmuInitialized)
	mem.status_rc_st24 = C.bool(t.StatusRcSt24)
	mem.status_rc_sumd = C.bool(t.StatusRcSumd)
	mem.alarm_vbatt_low = C.bool(t.AlarmVbattLow)
	mem.alarm_temperature = C.bool(t.AlarmTemperature)
	mem.alarm_servo_current = C.bool(t.AlarmServoCurrent)
	mem.alarm_acc_current = C.bool(t.AlarmAccCurrent)
	mem.alarm_fmu_lost = C.bool(t.AlarmFmuLost)
	mem.alarm_rc_lost = C.bool(t.AlarmRcLost)
	mem.alarm_pwm_error = C.bool(t.AlarmPwmError)
	mem.alarm_vservo_fault = C.bool(t.AlarmVservoFault)
	mem.arming_io_arm_ok = C.bool(t.ArmingIoArmOk)
	mem.arming_fmu_armed = C.bool(t.ArmingFmuArmed)
	mem.arming_fmu_prearmed = C.bool(t.ArmingFmuPrearmed)
	mem.arming_manual_override_ok = C.bool(t.ArmingManualOverrideOk)
	mem.arming_failsafe_custom = C.bool(t.ArmingFailsafeCustom)
	mem.arming_inair_restart_ok = C.bool(t.ArmingInairRestartOk)
	mem.arming_always_pwm_enable = C.bool(t.ArmingAlwaysPwmEnable)
	mem.arming_rc_handling_disabled = C.bool(t.ArmingRcHandlingDisabled)
	mem.arming_lockdown = C.bool(t.ArmingLockdown)
	mem.arming_force_failsafe = C.bool(t.ArmingForceFailsafe)
	mem.arming_termination_failsafe = C.bool(t.ArmingTerminationFailsafe)
	mem.arming_override_immediate = C.bool(t.ArmingOverrideImmediate)
	cSlice_actuators := mem.actuators[:]
	rosidl_runtime_c.Int16__Array_to_C(*(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_actuators)), t.Actuators[:])
	cSlice_servos := mem.servos[:]
	rosidl_runtime_c.Uint16__Array_to_C(*(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_servos)), t.Servos[:])
	cSlice_raw_inputs := mem.raw_inputs[:]
	rosidl_runtime_c.Uint16__Array_to_C(*(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_raw_inputs)), t.RawInputs[:])
	return unsafe.Pointer(mem)
}
func (t *Px4IoStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__Px4IoStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.FreeMemoryBytes = uint16(mem.free_memory_bytes)
	t.VoltageV = float32(mem.voltage_v)
	t.RssiV = float32(mem.rssi_v)
	t.StatusOutputsArmed = bool(mem.status_outputs_armed)
	t.StatusOverride = bool(mem.status_override)
	t.StatusRcOk = bool(mem.status_rc_ok)
	t.StatusRcPpm = bool(mem.status_rc_ppm)
	t.StatusRcDsm = bool(mem.status_rc_dsm)
	t.StatusRcSbus = bool(mem.status_rc_sbus)
	t.StatusFmuOk = bool(mem.status_fmu_ok)
	t.StatusRawPwm = bool(mem.status_raw_pwm)
	t.StatusMixerOk = bool(mem.status_mixer_ok)
	t.StatusArmSync = bool(mem.status_arm_sync)
	t.StatusInitOk = bool(mem.status_init_ok)
	t.StatusFailsafe = bool(mem.status_failsafe)
	t.StatusSafetyOff = bool(mem.status_safety_off)
	t.StatusFmuInitialized = bool(mem.status_fmu_initialized)
	t.StatusRcSt24 = bool(mem.status_rc_st24)
	t.StatusRcSumd = bool(mem.status_rc_sumd)
	t.AlarmVbattLow = bool(mem.alarm_vbatt_low)
	t.AlarmTemperature = bool(mem.alarm_temperature)
	t.AlarmServoCurrent = bool(mem.alarm_servo_current)
	t.AlarmAccCurrent = bool(mem.alarm_acc_current)
	t.AlarmFmuLost = bool(mem.alarm_fmu_lost)
	t.AlarmRcLost = bool(mem.alarm_rc_lost)
	t.AlarmPwmError = bool(mem.alarm_pwm_error)
	t.AlarmVservoFault = bool(mem.alarm_vservo_fault)
	t.ArmingIoArmOk = bool(mem.arming_io_arm_ok)
	t.ArmingFmuArmed = bool(mem.arming_fmu_armed)
	t.ArmingFmuPrearmed = bool(mem.arming_fmu_prearmed)
	t.ArmingManualOverrideOk = bool(mem.arming_manual_override_ok)
	t.ArmingFailsafeCustom = bool(mem.arming_failsafe_custom)
	t.ArmingInairRestartOk = bool(mem.arming_inair_restart_ok)
	t.ArmingAlwaysPwmEnable = bool(mem.arming_always_pwm_enable)
	t.ArmingRcHandlingDisabled = bool(mem.arming_rc_handling_disabled)
	t.ArmingLockdown = bool(mem.arming_lockdown)
	t.ArmingForceFailsafe = bool(mem.arming_force_failsafe)
	t.ArmingTerminationFailsafe = bool(mem.arming_termination_failsafe)
	t.ArmingOverrideImmediate = bool(mem.arming_override_immediate)
	cSlice_actuators := mem.actuators[:]
	rosidl_runtime_c.Int16__Array_to_Go(t.Actuators[:], *(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_actuators)))
	cSlice_servos := mem.servos[:]
	rosidl_runtime_c.Uint16__Array_to_Go(t.Servos[:], *(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_servos)))
	cSlice_raw_inputs := mem.raw_inputs[:]
	rosidl_runtime_c.Uint16__Array_to_Go(t.RawInputs[:], *(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_raw_inputs)))
}
func (t *Px4IoStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CPx4IoStatus = C.px4_msgs__msg__Px4IoStatus
type CPx4IoStatus__Sequence = C.px4_msgs__msg__Px4IoStatus__Sequence

func Px4IoStatus__Sequence_to_Go(goSlice *[]Px4IoStatus, cSlice CPx4IoStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Px4IoStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Px4IoStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Px4IoStatus * uintptr(i)),
		))
		(*goSlice)[i] = Px4IoStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Px4IoStatus__Sequence_to_C(cSlice *CPx4IoStatus__Sequence, goSlice []Px4IoStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Px4IoStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Px4IoStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Px4IoStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Px4IoStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__Px4IoStatus)(v.AsCStruct())
	}
}
func Px4IoStatus__Array_to_Go(goSlice []Px4IoStatus, cSlice []CPx4IoStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Px4IoStatus__Array_to_C(cSlice []CPx4IoStatus, goSlice []Px4IoStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__Px4IoStatus)(goSlice[i].AsCStruct())
	}
}

