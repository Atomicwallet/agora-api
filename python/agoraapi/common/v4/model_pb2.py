# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/v4/model.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='common/v4/model.proto',
  package='kin.agora.common.v4',
  syntax='proto3',
  serialized_options=b'\n\033org.kin.agora.gen.common.v4Z;github.com/kinecosystem/agora-api/genproto/common/v4;common\242\002\013APBCommonV4',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x15\x63ommon/v4/model.proto\x12\x13kin.agora.common.v4\x1a\x17validate/validate.proto\"+\n\x0fSolanaAccountId\x12\x18\n\x05value\x18\x01 \x01(\x0c\x42\t\xfa\x42\x06z\x04\x10 \x18 \")\n\rTransactionId\x12\x18\n\x05value\x18\x01 \x01(\x0c\x42\t\xfa\x42\x06z\x04\x10 \x18@\"%\n\tBlockhash\x12\x18\n\x05value\x18\x01 \x01(\x0c\x42\t\xfa\x42\x06z\x04\x10 \x18 \"0\n\x14TransactionSignature\x12\x18\n\x05value\x18\x01 \x01(\x0c\x42\t\xfa\x42\x06z\x04\x10@\x18@\"(\n\x0bTransaction\x12\x19\n\x05value\x18\x01 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\xd0\t\"\xf3\x01\n\x10TransactionError\x12<\n\x06reason\x18\x01 \x01(\x0e\x32,.kin.agora.common.v4.TransactionError.Reason\x12\x19\n\x11instruction_index\x18\x02 \x01(\x05\x12\x17\n\x03raw\x18\x03 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80P\"m\n\x06Reason\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07UNKNOWN\x10\x01\x12\x10\n\x0cUNAUTHORIZED\x10\x02\x12\r\n\tBAD_NONCE\x10\x03\x12\x16\n\x12INSUFFICIENT_FUNDS\x10\x04\x12\x13\n\x0fINVALID_ACCOUNT\x10\x05\"V\n\x12StellarTransaction\x12\x1e\n\nresult_xdr\x18\x01 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80P\x12 \n\x0c\x65nvelope_xdr\x18\x02 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80P*7\n\nCommitment\x12\n\n\x06RECENT\x10\x00\x12\n\n\x06SINGLE\x10\x01\x12\x08\n\x04ROOT\x10\x02\x12\x07\n\x03MAX\x10\x03\x42h\n\x1borg.kin.agora.gen.common.v4Z;github.com/kinecosystem/agora-api/genproto/common/v4;common\xa2\x02\x0b\x41PBCommonV4b\x06proto3'
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,])

_COMMITMENT = _descriptor.EnumDescriptor(
  name='Commitment',
  full_name='kin.agora.common.v4.Commitment',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='RECENT', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SINGLE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ROOT', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='MAX', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=624,
  serialized_end=679,
)
_sym_db.RegisterEnumDescriptor(_COMMITMENT)

Commitment = enum_type_wrapper.EnumTypeWrapper(_COMMITMENT)
RECENT = 0
SINGLE = 1
ROOT = 2
MAX = 3


_TRANSACTIONERROR_REASON = _descriptor.EnumDescriptor(
  name='Reason',
  full_name='kin.agora.common.v4.TransactionError.Reason',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UNAUTHORIZED', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='BAD_NONCE', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='INSUFFICIENT_FUNDS', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='INVALID_ACCOUNT', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=425,
  serialized_end=534,
)
_sym_db.RegisterEnumDescriptor(_TRANSACTIONERROR_REASON)


_SOLANAACCOUNTID = _descriptor.Descriptor(
  name='SolanaAccountId',
  full_name='kin.agora.common.v4.SolanaAccountId',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='kin.agora.common.v4.SolanaAccountId.value', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006z\004\020 \030 ', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=71,
  serialized_end=114,
)


_TRANSACTIONID = _descriptor.Descriptor(
  name='TransactionId',
  full_name='kin.agora.common.v4.TransactionId',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='kin.agora.common.v4.TransactionId.value', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006z\004\020 \030@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=116,
  serialized_end=157,
)


_BLOCKHASH = _descriptor.Descriptor(
  name='Blockhash',
  full_name='kin.agora.common.v4.Blockhash',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='kin.agora.common.v4.Blockhash.value', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006z\004\020 \030 ', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=159,
  serialized_end=196,
)


_TRANSACTIONSIGNATURE = _descriptor.Descriptor(
  name='TransactionSignature',
  full_name='kin.agora.common.v4.TransactionSignature',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='kin.agora.common.v4.TransactionSignature.value', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006z\004\020@\030@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=198,
  serialized_end=246,
)


_TRANSACTION = _descriptor.Descriptor(
  name='Transaction',
  full_name='kin.agora.common.v4.Transaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='kin.agora.common.v4.Transaction.value', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\320\t', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=248,
  serialized_end=288,
)


_TRANSACTIONERROR = _descriptor.Descriptor(
  name='TransactionError',
  full_name='kin.agora.common.v4.TransactionError',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='reason', full_name='kin.agora.common.v4.TransactionError.reason', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='instruction_index', full_name='kin.agora.common.v4.TransactionError.instruction_index', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='raw', full_name='kin.agora.common.v4.TransactionError.raw', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TRANSACTIONERROR_REASON,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=291,
  serialized_end=534,
)


_STELLARTRANSACTION = _descriptor.Descriptor(
  name='StellarTransaction',
  full_name='kin.agora.common.v4.StellarTransaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result_xdr', full_name='kin.agora.common.v4.StellarTransaction.result_xdr', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='envelope_xdr', full_name='kin.agora.common.v4.StellarTransaction.envelope_xdr', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=536,
  serialized_end=622,
)

_TRANSACTIONERROR.fields_by_name['reason'].enum_type = _TRANSACTIONERROR_REASON
_TRANSACTIONERROR_REASON.containing_type = _TRANSACTIONERROR
DESCRIPTOR.message_types_by_name['SolanaAccountId'] = _SOLANAACCOUNTID
DESCRIPTOR.message_types_by_name['TransactionId'] = _TRANSACTIONID
DESCRIPTOR.message_types_by_name['Blockhash'] = _BLOCKHASH
DESCRIPTOR.message_types_by_name['TransactionSignature'] = _TRANSACTIONSIGNATURE
DESCRIPTOR.message_types_by_name['Transaction'] = _TRANSACTION
DESCRIPTOR.message_types_by_name['TransactionError'] = _TRANSACTIONERROR
DESCRIPTOR.message_types_by_name['StellarTransaction'] = _STELLARTRANSACTION
DESCRIPTOR.enum_types_by_name['Commitment'] = _COMMITMENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SolanaAccountId = _reflection.GeneratedProtocolMessageType('SolanaAccountId', (_message.Message,), {
  'DESCRIPTOR' : _SOLANAACCOUNTID,
  '__module__' : 'common.v4.model_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.common.v4.SolanaAccountId)
  })
_sym_db.RegisterMessage(SolanaAccountId)

TransactionId = _reflection.GeneratedProtocolMessageType('TransactionId', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTIONID,
  '__module__' : 'common.v4.model_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.common.v4.TransactionId)
  })
_sym_db.RegisterMessage(TransactionId)

Blockhash = _reflection.GeneratedProtocolMessageType('Blockhash', (_message.Message,), {
  'DESCRIPTOR' : _BLOCKHASH,
  '__module__' : 'common.v4.model_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.common.v4.Blockhash)
  })
_sym_db.RegisterMessage(Blockhash)

TransactionSignature = _reflection.GeneratedProtocolMessageType('TransactionSignature', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTIONSIGNATURE,
  '__module__' : 'common.v4.model_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.common.v4.TransactionSignature)
  })
_sym_db.RegisterMessage(TransactionSignature)

Transaction = _reflection.GeneratedProtocolMessageType('Transaction', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTION,
  '__module__' : 'common.v4.model_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.common.v4.Transaction)
  })
_sym_db.RegisterMessage(Transaction)

TransactionError = _reflection.GeneratedProtocolMessageType('TransactionError', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTIONERROR,
  '__module__' : 'common.v4.model_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.common.v4.TransactionError)
  })
_sym_db.RegisterMessage(TransactionError)

StellarTransaction = _reflection.GeneratedProtocolMessageType('StellarTransaction', (_message.Message,), {
  'DESCRIPTOR' : _STELLARTRANSACTION,
  '__module__' : 'common.v4.model_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.common.v4.StellarTransaction)
  })
_sym_db.RegisterMessage(StellarTransaction)


DESCRIPTOR._options = None
_SOLANAACCOUNTID.fields_by_name['value']._options = None
_TRANSACTIONID.fields_by_name['value']._options = None
_BLOCKHASH.fields_by_name['value']._options = None
_TRANSACTIONSIGNATURE.fields_by_name['value']._options = None
_TRANSACTION.fields_by_name['value']._options = None
_TRANSACTIONERROR.fields_by_name['raw']._options = None
_STELLARTRANSACTION.fields_by_name['result_xdr']._options = None
_STELLARTRANSACTION.fields_by_name['envelope_xdr']._options = None
# @@protoc_insertion_point(module_scope)