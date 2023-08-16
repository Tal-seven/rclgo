/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/multi_echo_laser_scan.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/MultiEchoLaserScan", MultiEchoLaserScanTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/MultiEchoLaserScan", MultiEchoLaserScanTypeSupport)
}

// Do not create instances of this type directly. Always use NewMultiEchoLaserScan
// function instead.
type MultiEchoLaserScan struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp in the header is the acquisition time of
	AngleMin float32 `yaml:"angle_min"`// start angle of the scan [rad]
	AngleMax float32 `yaml:"angle_max"`// end angle of the scan [rad]
	AngleIncrement float32 `yaml:"angle_increment"`// angular distance between measurements [rad]
	TimeIncrement float32 `yaml:"time_increment"`// time between measurements [seconds] - if your scanner
	ScanTime float32 `yaml:"scan_time"`// time between scans [seconds]. is moving, this will be used in interpolating positionof 3d points
	RangeMin float32 `yaml:"range_min"`// minimum range value [m]
	RangeMax float32 `yaml:"range_max"`// maximum range value [m]
	Ranges []LaserEcho `yaml:"ranges"`// range data [m]
	Intensities []LaserEcho `yaml:"intensities"`// intensity data [device-specific units].  If your. (Note: NaNs, values < range_min or > range_max should be discarded)+Inf measurements are out of range-Inf measurements are too close to determine exact distance.
}

// NewMultiEchoLaserScan creates a new MultiEchoLaserScan with default values.
func NewMultiEchoLaserScan() *MultiEchoLaserScan {
	self := MultiEchoLaserScan{}
	self.SetDefaults()
	return &self
}

func (t *MultiEchoLaserScan) Clone() *MultiEchoLaserScan {
	c := &MultiEchoLaserScan{}
	c.Header = *t.Header.Clone()
	c.AngleMin = t.AngleMin
	c.AngleMax = t.AngleMax
	c.AngleIncrement = t.AngleIncrement
	c.TimeIncrement = t.TimeIncrement
	c.ScanTime = t.ScanTime
	c.RangeMin = t.RangeMin
	c.RangeMax = t.RangeMax
	if t.Ranges != nil {
		c.Ranges = make([]LaserEcho, len(t.Ranges))
		CloneLaserEchoSlice(c.Ranges, t.Ranges)
	}
	if t.Intensities != nil {
		c.Intensities = make([]LaserEcho, len(t.Intensities))
		CloneLaserEchoSlice(c.Intensities, t.Intensities)
	}
	return c
}

func (t *MultiEchoLaserScan) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MultiEchoLaserScan) SetDefaults() {
	t.Header.SetDefaults()
	t.AngleMin = 0
	t.AngleMax = 0
	t.AngleIncrement = 0
	t.TimeIncrement = 0
	t.ScanTime = 0
	t.RangeMin = 0
	t.RangeMax = 0
	t.Ranges = nil
	t.Intensities = nil
}

func (t *MultiEchoLaserScan) GetTypeSupport() types.MessageTypeSupport {
	return MultiEchoLaserScanTypeSupport
}

// MultiEchoLaserScanPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MultiEchoLaserScanPublisher struct {
	*rclgo.Publisher
}

// NewMultiEchoLaserScanPublisher creates and returns a new publisher for the
// MultiEchoLaserScan
func NewMultiEchoLaserScanPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MultiEchoLaserScanPublisher, error) {
	pub, err := node.NewPublisher(topic_name, MultiEchoLaserScanTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MultiEchoLaserScanPublisher{pub}, nil
}

func (p *MultiEchoLaserScanPublisher) Publish(msg *MultiEchoLaserScan) error {
	return p.Publisher.Publish(msg)
}

// MultiEchoLaserScanSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MultiEchoLaserScanSubscription struct {
	*rclgo.Subscription
}

// MultiEchoLaserScanSubscriptionCallback type is used to provide a subscription
// handler function for a MultiEchoLaserScanSubscription.
type MultiEchoLaserScanSubscriptionCallback func(msg *MultiEchoLaserScan, info *rclgo.MessageInfo, err error)

// NewMultiEchoLaserScanSubscription creates and returns a new subscription for the
// MultiEchoLaserScan
func NewMultiEchoLaserScanSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MultiEchoLaserScanSubscriptionCallback) (*MultiEchoLaserScanSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MultiEchoLaserScan
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MultiEchoLaserScanTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MultiEchoLaserScanSubscription{sub}, nil
}

func (s *MultiEchoLaserScanSubscription) TakeMessage(out *MultiEchoLaserScan) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMultiEchoLaserScanSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMultiEchoLaserScanSlice(dst, src []MultiEchoLaserScan) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MultiEchoLaserScanTypeSupport types.MessageTypeSupport = _MultiEchoLaserScanTypeSupport{}

type _MultiEchoLaserScanTypeSupport struct{}

func (t _MultiEchoLaserScanTypeSupport) New() types.Message {
	return NewMultiEchoLaserScan()
}

func (t _MultiEchoLaserScanTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__MultiEchoLaserScan
	return (unsafe.Pointer)(C.sensor_msgs__msg__MultiEchoLaserScan__create())
}

func (t _MultiEchoLaserScanTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__MultiEchoLaserScan__destroy((*C.sensor_msgs__msg__MultiEchoLaserScan)(pointer_to_free))
}

func (t _MultiEchoLaserScanTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultiEchoLaserScan)
	mem := (*C.sensor_msgs__msg__MultiEchoLaserScan)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.angle_min = C.float(m.AngleMin)
	mem.angle_max = C.float(m.AngleMax)
	mem.angle_increment = C.float(m.AngleIncrement)
	mem.time_increment = C.float(m.TimeIncrement)
	mem.scan_time = C.float(m.ScanTime)
	mem.range_min = C.float(m.RangeMin)
	mem.range_max = C.float(m.RangeMax)
	LaserEcho__Sequence_to_C(&mem.ranges, m.Ranges)
	LaserEcho__Sequence_to_C(&mem.intensities, m.Intensities)
}

func (t _MultiEchoLaserScanTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultiEchoLaserScan)
	mem := (*C.sensor_msgs__msg__MultiEchoLaserScan)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.AngleMin = float32(mem.angle_min)
	m.AngleMax = float32(mem.angle_max)
	m.AngleIncrement = float32(mem.angle_increment)
	m.TimeIncrement = float32(mem.time_increment)
	m.ScanTime = float32(mem.scan_time)
	m.RangeMin = float32(mem.range_min)
	m.RangeMax = float32(mem.range_max)
	LaserEcho__Sequence_to_Go(&m.Ranges, mem.ranges)
	LaserEcho__Sequence_to_Go(&m.Intensities, mem.intensities)
}

func (t _MultiEchoLaserScanTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__MultiEchoLaserScan())
}

type CMultiEchoLaserScan = C.sensor_msgs__msg__MultiEchoLaserScan
type CMultiEchoLaserScan__Sequence = C.sensor_msgs__msg__MultiEchoLaserScan__Sequence

func MultiEchoLaserScan__Sequence_to_Go(goSlice *[]MultiEchoLaserScan, cSlice CMultiEchoLaserScan__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiEchoLaserScan, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MultiEchoLaserScanTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MultiEchoLaserScan__Sequence_to_C(cSlice *CMultiEchoLaserScan__Sequence, goSlice []MultiEchoLaserScan) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__MultiEchoLaserScan)(C.malloc(C.sizeof_struct_sensor_msgs__msg__MultiEchoLaserScan * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MultiEchoLaserScanTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MultiEchoLaserScan__Array_to_Go(goSlice []MultiEchoLaserScan, cSlice []CMultiEchoLaserScan) {
	for i := 0; i < len(cSlice); i++ {
		MultiEchoLaserScanTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultiEchoLaserScan__Array_to_C(cSlice []CMultiEchoLaserScan, goSlice []MultiEchoLaserScan) {
	for i := 0; i < len(goSlice); i++ {
		MultiEchoLaserScanTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
