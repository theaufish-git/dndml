# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: dndml/roll.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10\x64ndml/roll.proto\x12\x05\x64ndml\"j\n\x04Roll\x12\x0c\n\x04\x64ice\x18\x01 \x01(\r\x12\x0b\n\x03\x64ie\x18\x02 \x01(\r\x12\x1d\n\x06reroll\x18\x03 \x03(\x0b\x32\r.dndml.Reroll\x12\x19\n\x04keep\x18\x04 \x01(\x0b\x32\x0b.dndml.Keep\x12\r\n\x05\x62onus\x18\x05 \x01(\r\"%\n\x04Keep\x12\r\n\x05\x63ount\x18\x01 \x01(\r\x12\x0e\n\x06lowest\x18\x02 \x01(\x08\"#\n\x06Reroll\x12\x0b\n\x03\x64ie\x18\x01 \x01(\r\x12\x0c\n\x04once\x18\x02 \x01(\x08\x42*Z(github.com/theaufish-git/dndml/pkg/dndmlb\x06proto3')



_ROLL = DESCRIPTOR.message_types_by_name['Roll']
_KEEP = DESCRIPTOR.message_types_by_name['Keep']
_REROLL = DESCRIPTOR.message_types_by_name['Reroll']
Roll = _reflection.GeneratedProtocolMessageType('Roll', (_message.Message,), {
  'DESCRIPTOR' : _ROLL,
  '__module__' : 'dndml.roll_pb2'
  # @@protoc_insertion_point(class_scope:dndml.Roll)
  })
_sym_db.RegisterMessage(Roll)

Keep = _reflection.GeneratedProtocolMessageType('Keep', (_message.Message,), {
  'DESCRIPTOR' : _KEEP,
  '__module__' : 'dndml.roll_pb2'
  # @@protoc_insertion_point(class_scope:dndml.Keep)
  })
_sym_db.RegisterMessage(Keep)

Reroll = _reflection.GeneratedProtocolMessageType('Reroll', (_message.Message,), {
  'DESCRIPTOR' : _REROLL,
  '__module__' : 'dndml.roll_pb2'
  # @@protoc_insertion_point(class_scope:dndml.Reroll)
  })
_sym_db.RegisterMessage(Reroll)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z(github.com/theaufish-git/dndml/pkg/dndml'
  _ROLL._serialized_start=27
  _ROLL._serialized_end=133
  _KEEP._serialized_start=135
  _KEEP._serialized_end=172
  _REROLL._serialized_start=174
  _REROLL._serialized_end=209
# @@protoc_insertion_point(module_scope)