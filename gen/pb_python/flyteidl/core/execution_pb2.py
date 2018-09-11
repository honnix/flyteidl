# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/execution.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/execution.proto',
  package='flyteidl.core',
  syntax='proto3',
  serialized_pb=_b('\n\x1d\x66lyteidl/core/execution.proto\x12\rflyteidl.core\"/\n\x0e\x45xecutionError\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t*o\n\x0e\x45xecutionPhase\x12\r\n\tUNDEFINED\x10\x00\x12\x0b\n\x07RUNNING\x10\x01\x12\r\n\tSUCCEEDED\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\x12\r\n\tTIMED_OUT\x10\x04\x12\x0b\n\x07\x41\x42ORTED\x10\x05\x12\n\n\x06QUEUED\x10\x06\x42\x32Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
)

_EXECUTIONPHASE = _descriptor.EnumDescriptor(
  name='ExecutionPhase',
  full_name='flyteidl.core.ExecutionPhase',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNDEFINED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RUNNING', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUCCEEDED', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FAILED', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TIMED_OUT', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ABORTED', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='QUEUED', index=6, number=6,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=97,
  serialized_end=208,
)
_sym_db.RegisterEnumDescriptor(_EXECUTIONPHASE)

ExecutionPhase = enum_type_wrapper.EnumTypeWrapper(_EXECUTIONPHASE)
UNDEFINED = 0
RUNNING = 1
SUCCEEDED = 2
FAILED = 3
TIMED_OUT = 4
ABORTED = 5
QUEUED = 6



_EXECUTIONERROR = _descriptor.Descriptor(
  name='ExecutionError',
  full_name='flyteidl.core.ExecutionError',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='flyteidl.core.ExecutionError.code', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='flyteidl.core.ExecutionError.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=48,
  serialized_end=95,
)

DESCRIPTOR.message_types_by_name['ExecutionError'] = _EXECUTIONERROR
DESCRIPTOR.enum_types_by_name['ExecutionPhase'] = _EXECUTIONPHASE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ExecutionError = _reflection.GeneratedProtocolMessageType('ExecutionError', (_message.Message,), dict(
  DESCRIPTOR = _EXECUTIONERROR,
  __module__ = 'flyteidl.core.execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.ExecutionError)
  ))
_sym_db.RegisterMessage(ExecutionError)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/core'))
# @@protoc_insertion_point(module_scope)