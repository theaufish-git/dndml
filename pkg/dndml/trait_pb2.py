# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: dndml/trait.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from dndml import uses_pb2 as dndml_dot_uses__pb2
from dndml.enums import enums_pb2 as dndml_dot_enums_dot_enums__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11\x64ndml/trait.proto\x12\x05\x64ndml\x1a\x10\x64ndml/uses.proto\x1a\x17\x64ndml/enums/enums.proto\"\xbd\x01\n\x05Trait\x12\x13\n\x06source\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x15\n\x08template\x18\x02 \x01(\tH\x01\x88\x01\x01\x12!\n\x04kind\x18\x03 \x01(\x0e\x32\x13.dndml.enums.Object\x12\x0f\n\x07version\x18\x04 \x01(\t\x12\x0c\n\x04name\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x19\n\x04uses\x18\x07 \x01(\x0b\x32\x0b.dndml.UsesB\t\n\x07_sourceB\x0b\n\t_templateB*Z(github.com/theaufish-git/dndml/pkg/dndmlb\x06proto3')



_TRAIT = DESCRIPTOR.message_types_by_name['Trait']
Trait = _reflection.GeneratedProtocolMessageType('Trait', (_message.Message,), {
  'DESCRIPTOR' : _TRAIT,
  '__module__' : 'dndml.trait_pb2'
  # @@protoc_insertion_point(class_scope:dndml.Trait)
  })
_sym_db.RegisterMessage(Trait)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z(github.com/theaufish-git/dndml/pkg/dndml'
  _TRAIT._serialized_start=72
  _TRAIT._serialized_end=261
# @@protoc_insertion_point(module_scope)