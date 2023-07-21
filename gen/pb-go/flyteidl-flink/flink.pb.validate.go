// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl-flink/flink.proto

package flyteidl_flink

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _flink_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Resource) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCpu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "Cpu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMemory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "Memory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPersistentVolume()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "PersistentVolume",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on JobManager with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JobManager) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JobManagerValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// JobManagerValidationError is the validation error returned by
// JobManager.Validate if the designated constraints aren't met.
type JobManagerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobManagerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobManagerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobManagerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobManagerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobManagerValidationError) ErrorName() string { return "JobManagerValidationError" }

// Error satisfies the builtin error interface
func (e JobManagerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobManager.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobManagerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobManagerValidationError{}

// Validate checks the field values on TaskManager with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TaskManager) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskManagerValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetReplicas() < 0 {
		return TaskManagerValidationError{
			field:  "Replicas",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// TaskManagerValidationError is the validation error returned by
// TaskManager.Validate if the designated constraints aren't met.
type TaskManagerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskManagerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskManagerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskManagerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskManagerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskManagerValidationError) ErrorName() string { return "TaskManagerValidationError" }

// Error satisfies the builtin error interface
func (e TaskManagerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskManager.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskManagerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskManagerValidationError{}

// Validate checks the field values on JFlyte with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JFlyte) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IndexFileLocation

	for idx, item := range m.GetArtifacts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JFlyteValidationError{
					field:  fmt.Sprintf("Artifacts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// JFlyteValidationError is the validation error returned by JFlyte.Validate if
// the designated constraints aren't met.
type JFlyteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JFlyteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JFlyteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JFlyteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JFlyteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JFlyteValidationError) ErrorName() string { return "JFlyteValidationError" }

// Error satisfies the builtin error interface
func (e JFlyteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJFlyte.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JFlyteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JFlyteValidationError{}

// Validate checks the field values on FlinkJob with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FlinkJob) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetJarFiles() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return FlinkJobValidationError{
				field:  fmt.Sprintf("JarFiles[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

		if _, err := url.Parse(item); err != nil {
			return FlinkJobValidationError{
				field:  fmt.Sprintf("JarFiles[%v]", idx),
				reason: "value must be a valid URI",
				cause:  err,
			}
		}

	}

	if utf8.RuneCountInString(m.GetMainClass()) < 1 {
		return FlinkJobValidationError{
			field:  "MainClass",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for FlinkProperties

	if v, ok := interface{}(m.GetJobManager()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FlinkJobValidationError{
				field:  "JobManager",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTaskManager()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FlinkJobValidationError{
				field:  "TaskManager",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ServiceAccount

	// no validation rules for Image

	// no validation rules for FlinkVersion

	if m.GetParallelism() < 0 {
		return FlinkJobValidationError{
			field:  "Parallelism",
			reason: "value must be greater than or equal to 0",
		}
	}

	// no validation rules for KubernetesClusterName

	// no validation rules for Region

	if v, ok := interface{}(m.GetJflyte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FlinkJobValidationError{
				field:  "Jflyte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FlinkJobValidationError is the validation error returned by
// FlinkJob.Validate if the designated constraints aren't met.
type FlinkJobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FlinkJobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FlinkJobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FlinkJobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FlinkJobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FlinkJobValidationError) ErrorName() string { return "FlinkJobValidationError" }

// Error satisfies the builtin error interface
func (e FlinkJobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFlinkJob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FlinkJobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FlinkJobValidationError{}

// Validate checks the field values on JobExecutionInfo with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *JobExecutionInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// JobExecutionInfoValidationError is the validation error returned by
// JobExecutionInfo.Validate if the designated constraints aren't met.
type JobExecutionInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobExecutionInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobExecutionInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobExecutionInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobExecutionInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobExecutionInfoValidationError) ErrorName() string { return "JobExecutionInfoValidationError" }

// Error satisfies the builtin error interface
func (e JobExecutionInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobExecutionInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobExecutionInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobExecutionInfoValidationError{}

// Validate checks the field values on JobManagerExecutionInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *JobManagerExecutionInfo) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// JobManagerExecutionInfoValidationError is the validation error returned by
// JobManagerExecutionInfo.Validate if the designated constraints aren't met.
type JobManagerExecutionInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobManagerExecutionInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobManagerExecutionInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobManagerExecutionInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobManagerExecutionInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobManagerExecutionInfoValidationError) ErrorName() string {
	return "JobManagerExecutionInfoValidationError"
}

// Error satisfies the builtin error interface
func (e JobManagerExecutionInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobManagerExecutionInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobManagerExecutionInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobManagerExecutionInfoValidationError{}

// Validate checks the field values on FlinkExecutionInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FlinkExecutionInfo) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetJob()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FlinkExecutionInfoValidationError{
				field:  "Job",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetJobManager()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FlinkExecutionInfoValidationError{
				field:  "JobManager",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FlinkExecutionInfoValidationError is the validation error returned by
// FlinkExecutionInfo.Validate if the designated constraints aren't met.
type FlinkExecutionInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FlinkExecutionInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FlinkExecutionInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FlinkExecutionInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FlinkExecutionInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FlinkExecutionInfoValidationError) ErrorName() string {
	return "FlinkExecutionInfoValidationError"
}

// Error satisfies the builtin error interface
func (e FlinkExecutionInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFlinkExecutionInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FlinkExecutionInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FlinkExecutionInfoValidationError{}

// Validate checks the field values on Resource_Quantity with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Resource_Quantity) Validate() error {
	if m == nil {
		return nil
	}

	if !_Resource_Quantity_String__Pattern.MatchString(m.GetString_()) {
		return Resource_QuantityValidationError{
			field:  "String_",
			reason: "value does not match regex pattern \"^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$\"",
		}
	}

	return nil
}

// Resource_QuantityValidationError is the validation error returned by
// Resource_Quantity.Validate if the designated constraints aren't met.
type Resource_QuantityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Resource_QuantityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Resource_QuantityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Resource_QuantityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Resource_QuantityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Resource_QuantityValidationError) ErrorName() string {
	return "Resource_QuantityValidationError"
}

// Error satisfies the builtin error interface
func (e Resource_QuantityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource_Quantity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Resource_QuantityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Resource_QuantityValidationError{}

var _Resource_Quantity_String__Pattern = regexp.MustCompile("^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$")

// Validate checks the field values on Resource_PersistentVolume with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Resource_PersistentVolume) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := Resource_PersistentVolume_Type_name[int32(m.GetType())]; !ok {
		return Resource_PersistentVolumeValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Resource_PersistentVolumeValidationError{
				field:  "Size",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Resource_PersistentVolumeValidationError is the validation error returned by
// Resource_PersistentVolume.Validate if the designated constraints aren't met.
type Resource_PersistentVolumeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Resource_PersistentVolumeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Resource_PersistentVolumeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Resource_PersistentVolumeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Resource_PersistentVolumeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Resource_PersistentVolumeValidationError) ErrorName() string {
	return "Resource_PersistentVolumeValidationError"
}

// Error satisfies the builtin error interface
func (e Resource_PersistentVolumeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource_PersistentVolume.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Resource_PersistentVolumeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Resource_PersistentVolumeValidationError{}

// Validate checks the field values on JFlyte_Artifact with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *JFlyte_Artifact) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if utf8.RuneCountInString(m.GetLocation()) < 1 {
		return JFlyte_ArtifactValidationError{
			field:  "Location",
			reason: "value length must be at least 1 runes",
		}
	}

	if _, err := url.Parse(m.GetLocation()); err != nil {
		return JFlyte_ArtifactValidationError{
			field:  "Location",
			reason: "value must be a valid URI",
			cause:  err,
		}
	}

	return nil
}

// JFlyte_ArtifactValidationError is the validation error returned by
// JFlyte_Artifact.Validate if the designated constraints aren't met.
type JFlyte_ArtifactValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JFlyte_ArtifactValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JFlyte_ArtifactValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JFlyte_ArtifactValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JFlyte_ArtifactValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JFlyte_ArtifactValidationError) ErrorName() string { return "JFlyte_ArtifactValidationError" }

// Error satisfies the builtin error interface
func (e JFlyte_ArtifactValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJFlyte_Artifact.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JFlyte_ArtifactValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JFlyte_ArtifactValidationError{}
