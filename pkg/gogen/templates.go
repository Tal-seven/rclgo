/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package gogen

import (
	"strings"
	"text/template"
)

var templateFuncMap template.FuncMap = template.FuncMap{
	"lc":                  strings.ToLower,
	"camelToSnake":        camelToSnake,
	"cSerializationCode":  cSerializationCode,
	"goSerializationCode": goSerializationCode,
	"defaultCode":         defaultCode,
	"ucFirst":             ucFirst,
}

var ros2MsgToGolangTypeTemplate = template.Must(template.New("ros2MsgToGolangTypeTemplate").Funcs(templateFuncMap).Parse(
	`/*{{ $Md := . }}
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package {{ $Md.RosPackage }}
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	{{range $path, $name := $Md.GoImports -}}
	{{$name}} "{{$path}}"
	{{""}}{{- end}}
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -l{{$Md.RosPackage}}__rosidl_typesupport_c -l{{$Md.RosPackage}}__rosidl_generator_c
{{range $k, $v := $Md.CImports -}}
#cgo LDFLAGS: -l{{$k}}__rosidl_typesupport_c -l{{$k}}__rosidl_generator_c
{{""}}
{{- end}}
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <{{$Md.RosPackage}}/msg/{{$Md.RosMsgName | camelToSnake}}.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("{{.RosPackage}}/{{.RosMsgName}}", &{{.RosMsgName}}{})
}

{{- if $Md.Constants }}
const (
{{- range $Md.Constants }}
	{{$Md.RosMsgName}}_{{.RosName}} {{.PkgReference}}{{.GoType}} = {{.Value}}{{if .Comment -}} // {{.Comment}}{{- end}}
{{- end }}
)
{{- end }}

type {{.RosMsgName}} struct {
	{{- range $k, $v := .Fields }}
	{{$v.GoName }} {{$v.TypeArray}}{{$v.PkgReference}}{{$v.GoType}}` +
		"{{\"\"}} `yaml:\"{{$v.RosName}}\"`" + `{{if .Comment -}} // {{.Comment}}{{- end}}
	{{- end }}
}

func New{{.RosMsgName}}() *{{.RosMsgName}} {
	self := {{.RosMsgName}}{}
	self.SetDefaults(nil)
	return &self
}

func (t *{{.RosMsgName}}) SetDefaults(d interface{}) ros2types.ROS2Msg {
	{{""}}
	{{- range $k, $v := .Fields -}}
	{{defaultCode $v}}
	{{- end -}}
	{{""}}
	return t
{{"}"}}

func (t *{{.RosMsgName}}) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__{{.RosPackage}}__msg__{{.RosMsgName}}())
}
func (t *{{.RosMsgName}}) PrepareMemory() unsafe.Pointer { //returns *C.{{.RosPackage}}__msg__{{.RosMsgName}}
	return (unsafe.Pointer)(C.{{.RosPackage}}__msg__{{.RosMsgName}}__create())
}
func (t *{{.RosMsgName}}) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.{{.RosPackage}}__msg__{{.RosMsgName}}__destroy((*C.{{.RosPackage}}__msg__{{.RosMsgName}})(pointer_to_free))
}
func (t *{{.RosMsgName}}) AsCStruct() unsafe.Pointer {
	mem := (*C.{{.RosPackage}}__msg__{{.RosMsgName}})(t.PrepareMemory())
	{{- range .Fields }}
	{{cSerializationCode . $Md}}
	{{- end }}
	return unsafe.Pointer(mem)
}
func (t *{{.RosMsgName}}) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	{{if .Fields -}}
	mem := (*C.{{.RosPackage}}__msg__{{.RosMsgName}})(ros2_message_buffer)
	{{- range .Fields }}
	{{goSerializationCode . $Md}}
	{{- end }}
	{{- end }}
}
func (t *{{.RosMsgName}}) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type C{{.RosMsgName}} = C.{{.RosPackage}}__msg__{{.RosMsgName}}
type C{{.RosMsgName}}__Sequence = C.{{.RosPackage}}__msg__{{.RosMsgName}}__Sequence

func {{.RosMsgName}}__Sequence_to_Go(goSlice *[]{{.RosMsgName}}, cSlice C{{.RosMsgName}}__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]{{.RosMsgName}}, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.{{.RosPackage}}__msg__{{.RosMsgName}}__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_{{.RosPackage}}__msg__{{.RosMsgName}} * uintptr(i)),
		))
		(*goSlice)[i] = {{.RosMsgName}}{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func {{.RosMsgName}}__Sequence_to_C(cSlice *C{{.RosMsgName}}__Sequence, goSlice []{{.RosMsgName}}) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.{{.RosPackage}}__msg__{{.RosMsgName}})(C.malloc((C.size_t)(C.sizeof_struct_{{.RosPackage}}__msg__{{.RosMsgName}} * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.{{.RosPackage}}__msg__{{.RosMsgName}})(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_{{.RosPackage}}__msg__{{.RosMsgName}} * uintptr(i)),
		))
		*cIdx = *(*C.{{.RosPackage}}__msg__{{.RosMsgName}})(v.AsCStruct())
	}
}
func {{.RosMsgName}}__Array_to_Go(goSlice []{{.RosMsgName}}, cSlice []C{{.RosMsgName}}) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func {{.RosMsgName}}__Array_to_C(cSlice []C{{.RosMsgName}}, goSlice []{{.RosMsgName}}) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.{{.RosPackage}}__msg__{{.RosMsgName}})(goSlice[i].AsCStruct())
	}
}


`))

var ros2rosidl_runtime_c_handlers = template.Must(template.New("ros2-rosidl-runtime-c-handlers-template").Funcs(templateFuncMap).Parse(
	`/*{{ $Md := . }}
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rosidl_runtime_c

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo CFLAGS: -I/opt/ros/foxy/include

#include "rosidl_runtime_c/string.h"
#include "rosidl_runtime_c/primitives_sequence.h"

*/
import "C"
import (
	"unsafe"
){{range $k, $v := .PMap -}}{{if .SkipAutogen}}{{- else -}}
{{""}}
{{""}}
// {{.RosType | ucFirst}}
type C{{.RosType | ucFirst}} = C.{{.CType}}
type C{{.RosType | ucFirst}}__Sequence = C.rosidl_runtime_c__{{.CStructName}}__Sequence

func {{.RosType | ucFirst}}__Sequence_to_Go(goSlice *[]{{.GoType}}, cSlice C{{.RosType | ucFirst}}__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]{{.GoType}}, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.{{.CType}})(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_{{.CType}} * uintptr(i)),
		))
		(*goSlice)[i] = {{.GoType}}(*cIdx)
	}
}
func {{.RosType | ucFirst}}__Sequence_to_C(cSlice *C{{.RosType | ucFirst}}__Sequence, goSlice []{{.GoType}}) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.{{.CType}})(C.malloc((C.size_t)(C.sizeof_{{.CType}} * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.{{.CType}})(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_{{.CType}} * uintptr(i)),
		))
		*cIdx = (C.{{.CType}})(v)
	}
}
func {{.RosType | ucFirst}}__Array_to_Go(goSlice []{{.GoType}}, cSlice []C{{.RosType | ucFirst}}) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = {{.GoType}}(cSlice[i])
	}
}
func {{.RosType | ucFirst}}__Array_to_C(cSlice []C{{.RosType | ucFirst}}, goSlice []{{.GoType}}) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.{{.CType}}(goSlice[i])
	}
}
{{- end}}{{- end}}
`))

var ros2MsgImportAllPackage = template.Must(template.New("ros2MsgToGolangTypeTemplate").Funcs(templateFuncMap).Parse(
	`/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package ros2msgs

import (
	{{range $md := . -}}
	_ "github.com/tiiuae/rclgo/pkg/ros2/msgs/{{$md.RosPackage}}/msg" //
	{{""}}{{- end}}
)
`))

var ros2ErrorCodes = template.Must(template.New("ros2ErrorCodes").Funcs(templateFuncMap).Parse(
	`/*{{ $P := . }}
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package ros2

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c -lrcutils -lrmw_implementation -lpx4_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: {{.rclcPath}}/lib/librclc.so
#cgo CFLAGS: -I/opt/ros/foxy/include -I{{.rclcPath}}/include/

#include <rcl/types.h>
#include <rmw/ret_types.h>

*/
import "C"
import (
	"runtime"
)

func ErrorsCastC(rcl_ret_t C.rcl_ret_t, context string) RCLError {
	stackTraceBuffer := make([]byte, 2048)
	runtime.Stack(stackTraceBuffer, false) // Get stack trace of the current running thread only

	// https://stackoverflow.com/questions/9928221/table-of-functions-vs-switch-in-golang
	// switch-case is faster thanks to compiler optimization than a dispatcher?
	switch rcl_ret_t {
	{{range $e := .ERRORS -}}{{if $e.Rcl_ret_t -}}{{if not (index $P.DEDUP_FILTER $e.Name) -}}
	case C.{{$e.Name}}:
		return &{{$e.Name}}{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: {{$e.Rcl_ret_t}}, trace: string(stackTraceBuffer), ctx: errorsBuildContext(&{{$e.Name}}{}, context, string(stackTraceBuffer))}}
	{{""}}
	{{- end}}{{- end}}{{- end}}
	default:
		return &RCL_RET_GOLANG_UNKNOWN_RET_TYPE{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: (int)(rcl_ret_t), ctx: context}}
	}
}

type RCL_RET_GOLANG_UNKNOWN_RET_TYPE struct {
	RCL_RET_struct
}

{{range $e := .ERRORS -}}{{if $e.Rcl_ret_t}}
// {{$e.Name}} {{$e.Comment}}
type {{$e.Name}} struct {
	RCL_RET_struct
}
{{""}}
{{- end}}{{- end}}

{{range $e := .ERRORS -}}{{if $e.Reference}}
// {{$e.Name}} {{$e.Comment}}
type {{$e.Name}} {{$e.Reference}}
{{""}}
{{- end}}{{- end}}

`))