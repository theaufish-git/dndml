# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: dndml/v5/creature.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from dndml import roll_pb2 as dndml_dot_roll__pb2
from dndml.v5.enums import enums_pb2 as dndml_dot_v5_dot_enums_dot_enums__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17\x64ndml/v5/creature.proto\x12\x08\x64ndml.v5\x1a\x10\x64ndml/roll.proto\x1a\x1a\x64ndml/v5/enums/enums.proto\"D\n\x08\x43reature\x12\x0c\n\x04name\x18\x01 \x01(\t\x12*\n\x04type\x18\x02 \x01(\x0e\x32\x1c.dndml.v5.enums.CreatureTypeB-Z+github.com/theaufish-git/dndml/pkg/dndml/v5b\x06proto3')



_CREATURE = DESCRIPTOR.message_types_by_name['Creature']
Creature = _reflection.GeneratedProtocolMessageType('Creature', (_message.Message,), {
  'DESCRIPTOR' : _CREATURE,
  '__module__' : 'dndml.v5.creature_pb2'
  # @@protoc_insertion_point(class_scope:dndml.v5.Creature)
  })
_sym_db.RegisterMessage(Creature)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z+github.com/theaufish-git/dndml/pkg/dndml/v5'
  _CREATURE._serialized_start=83
  _CREATURE._serialized_end=151
# @@protoc_insertion_point(module_scope)
