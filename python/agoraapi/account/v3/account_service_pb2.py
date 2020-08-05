# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: account/v3/account_service.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from common.v3 import model_pb2 as common_dot_v3_dot_model__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='account/v3/account_service.proto',
  package='kin.agora.account.v3',
  syntax='proto3',
  serialized_options=b'\n\034org.kin.agora.gen.account.v3Z=github.com/kinecosystem/agora-api/genproto/account/v3;account\242\002\014APBAccountV3',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n account/v3/account_service.proto\x12\x14kin.agora.account.v3\x1a\x17validate/validate.proto\x1a\x15\x63ommon/v3/model.proto\"|\n\x0b\x41\x63\x63ountInfo\x12\x43\n\naccount_id\x18\x01 \x01(\x0b\x32%.kin.agora.common.v3.StellarAccountIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x17\n\x0fsequence_number\x18\x02 \x01(\x03\x12\x0f\n\x07\x62\x61lance\x18\x03 \x01(\x03\"[\n\x14\x43reateAccountRequest\x12\x43\n\naccount_id\x18\x01 \x01(\x0b\x32%.kin.agora.common.v3.StellarAccountIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\xb2\x01\n\x15\x43reateAccountResponse\x12\x42\n\x06result\x18\x01 \x01(\x0e\x32\x32.kin.agora.account.v3.CreateAccountResponse.Result\x12\x37\n\x0c\x61\x63\x63ount_info\x18\x02 \x01(\x0b\x32!.kin.agora.account.v3.AccountInfo\"\x1c\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\n\n\x06\x45XISTS\x10\x01\"\\\n\x15GetAccountInfoRequest\x12\x43\n\naccount_id\x18\x01 \x01(\x0b\x32%.kin.agora.common.v3.StellarAccountIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\xb7\x01\n\x16GetAccountInfoResponse\x12\x43\n\x06result\x18\x01 \x01(\x0e\x32\x33.kin.agora.account.v3.GetAccountInfoResponse.Result\x12\x37\n\x0c\x61\x63\x63ount_info\x18\x02 \x01(\x0b\x32!.kin.agora.account.v3.AccountInfo\"\x1f\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\r\n\tNOT_FOUND\x10\x01\"W\n\x10GetEventsRequest\x12\x43\n\naccount_id\x18\x01 \x01(\x0b\x32%.kin.agora.common.v3.StellarAccountIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\x96\x01\n\x06\x45vents\x12\x33\n\x06result\x18\x01 \x01(\x0e\x32#.kin.agora.account.v3.Events.Result\x12\x36\n\x06\x65vents\x18\x02 \x03(\x0b\x32\x1b.kin.agora.account.v3.EventB\t\xfa\x42\x06\x92\x01\x03\x10\x80\x01\"\x1f\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\r\n\tNOT_FOUND\x10\x01\"\x9e\x01\n\x05\x45vent\x12H\n\x14\x61\x63\x63ount_update_event\x18\x01 \x01(\x0b\x32(.kin.agora.account.v3.AccountUpdateEventH\x00\x12\x43\n\x11transaction_event\x18\x02 \x01(\x0b\x32&.kin.agora.account.v3.TransactionEventH\x00\x42\x06\n\x04type\"W\n\x12\x41\x63\x63ountUpdateEvent\x12\x41\n\x0c\x61\x63\x63ount_info\x18\x01 \x01(\x0b\x32!.kin.agora.account.v3.AccountInfoB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"T\n\x10TransactionEvent\x12 \n\x0c\x65nvelope_xdr\x18\x01 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80P\x12\x1e\n\nresult_xdr\x18\x02 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x00\x18\x80P2\xb5\x02\n\x07\x41\x63\x63ount\x12h\n\rCreateAccount\x12*.kin.agora.account.v3.CreateAccountRequest\x1a+.kin.agora.account.v3.CreateAccountResponse\x12k\n\x0eGetAccountInfo\x12+.kin.agora.account.v3.GetAccountInfoRequest\x1a,.kin.agora.account.v3.GetAccountInfoResponse\x12S\n\tGetEvents\x12&.kin.agora.account.v3.GetEventsRequest\x1a\x1c.kin.agora.account.v3.Events0\x01\x42l\n\x1corg.kin.agora.gen.account.v3Z=github.com/kinecosystem/agora-api/genproto/account/v3;account\xa2\x02\x0c\x41PBAccountV3b\x06proto3'
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,common_dot_v3_dot_model__pb2.DESCRIPTOR,])



_CREATEACCOUNTRESPONSE_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='kin.agora.account.v3.CreateAccountResponse.Result',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OK', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EXISTS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=476,
  serialized_end=504,
)
_sym_db.RegisterEnumDescriptor(_CREATEACCOUNTRESPONSE_RESULT)

_GETACCOUNTINFORESPONSE_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='kin.agora.account.v3.GetAccountInfoResponse.Result',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OK', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NOT_FOUND', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=753,
  serialized_end=784,
)
_sym_db.RegisterEnumDescriptor(_GETACCOUNTINFORESPONSE_RESULT)

_EVENTS_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='kin.agora.account.v3.Events.Result',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OK', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NOT_FOUND', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=753,
  serialized_end=784,
)
_sym_db.RegisterEnumDescriptor(_EVENTS_RESULT)


_ACCOUNTINFO = _descriptor.Descriptor(
  name='AccountInfo',
  full_name='kin.agora.account.v3.AccountInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='kin.agora.account.v3.AccountInfo.account_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sequence_number', full_name='kin.agora.account.v3.AccountInfo.sequence_number', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='balance', full_name='kin.agora.account.v3.AccountInfo.balance', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=106,
  serialized_end=230,
)


_CREATEACCOUNTREQUEST = _descriptor.Descriptor(
  name='CreateAccountRequest',
  full_name='kin.agora.account.v3.CreateAccountRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='kin.agora.account.v3.CreateAccountRequest.account_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=232,
  serialized_end=323,
)


_CREATEACCOUNTRESPONSE = _descriptor.Descriptor(
  name='CreateAccountResponse',
  full_name='kin.agora.account.v3.CreateAccountResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='kin.agora.account.v3.CreateAccountResponse.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='account_info', full_name='kin.agora.account.v3.CreateAccountResponse.account_info', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _CREATEACCOUNTRESPONSE_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=326,
  serialized_end=504,
)


_GETACCOUNTINFOREQUEST = _descriptor.Descriptor(
  name='GetAccountInfoRequest',
  full_name='kin.agora.account.v3.GetAccountInfoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='kin.agora.account.v3.GetAccountInfoRequest.account_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=506,
  serialized_end=598,
)


_GETACCOUNTINFORESPONSE = _descriptor.Descriptor(
  name='GetAccountInfoResponse',
  full_name='kin.agora.account.v3.GetAccountInfoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='kin.agora.account.v3.GetAccountInfoResponse.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='account_info', full_name='kin.agora.account.v3.GetAccountInfoResponse.account_info', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _GETACCOUNTINFORESPONSE_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=601,
  serialized_end=784,
)


_GETEVENTSREQUEST = _descriptor.Descriptor(
  name='GetEventsRequest',
  full_name='kin.agora.account.v3.GetEventsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='kin.agora.account.v3.GetEventsRequest.account_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=786,
  serialized_end=873,
)


_EVENTS = _descriptor.Descriptor(
  name='Events',
  full_name='kin.agora.account.v3.Events',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='kin.agora.account.v3.Events.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='events', full_name='kin.agora.account.v3.Events.events', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006\222\001\003\020\200\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _EVENTS_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=876,
  serialized_end=1026,
)


_EVENT = _descriptor.Descriptor(
  name='Event',
  full_name='kin.agora.account.v3.Event',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_update_event', full_name='kin.agora.account.v3.Event.account_update_event', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='transaction_event', full_name='kin.agora.account.v3.Event.transaction_event', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
    _descriptor.OneofDescriptor(
      name='type', full_name='kin.agora.account.v3.Event.type',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=1029,
  serialized_end=1187,
)


_ACCOUNTUPDATEEVENT = _descriptor.Descriptor(
  name='AccountUpdateEvent',
  full_name='kin.agora.account.v3.AccountUpdateEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_info', full_name='kin.agora.account.v3.AccountUpdateEvent.account_info', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=1189,
  serialized_end=1276,
)


_TRANSACTIONEVENT = _descriptor.Descriptor(
  name='TransactionEvent',
  full_name='kin.agora.account.v3.TransactionEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='envelope_xdr', full_name='kin.agora.account.v3.TransactionEvent.envelope_xdr', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='result_xdr', full_name='kin.agora.account.v3.TransactionEvent.result_xdr', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\000\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=1278,
  serialized_end=1362,
)

_ACCOUNTINFO.fields_by_name['account_id'].message_type = common_dot_v3_dot_model__pb2._STELLARACCOUNTID
_CREATEACCOUNTREQUEST.fields_by_name['account_id'].message_type = common_dot_v3_dot_model__pb2._STELLARACCOUNTID
_CREATEACCOUNTRESPONSE.fields_by_name['result'].enum_type = _CREATEACCOUNTRESPONSE_RESULT
_CREATEACCOUNTRESPONSE.fields_by_name['account_info'].message_type = _ACCOUNTINFO
_CREATEACCOUNTRESPONSE_RESULT.containing_type = _CREATEACCOUNTRESPONSE
_GETACCOUNTINFOREQUEST.fields_by_name['account_id'].message_type = common_dot_v3_dot_model__pb2._STELLARACCOUNTID
_GETACCOUNTINFORESPONSE.fields_by_name['result'].enum_type = _GETACCOUNTINFORESPONSE_RESULT
_GETACCOUNTINFORESPONSE.fields_by_name['account_info'].message_type = _ACCOUNTINFO
_GETACCOUNTINFORESPONSE_RESULT.containing_type = _GETACCOUNTINFORESPONSE
_GETEVENTSREQUEST.fields_by_name['account_id'].message_type = common_dot_v3_dot_model__pb2._STELLARACCOUNTID
_EVENTS.fields_by_name['result'].enum_type = _EVENTS_RESULT
_EVENTS.fields_by_name['events'].message_type = _EVENT
_EVENTS_RESULT.containing_type = _EVENTS
_EVENT.fields_by_name['account_update_event'].message_type = _ACCOUNTUPDATEEVENT
_EVENT.fields_by_name['transaction_event'].message_type = _TRANSACTIONEVENT
_EVENT.oneofs_by_name['type'].fields.append(
  _EVENT.fields_by_name['account_update_event'])
_EVENT.fields_by_name['account_update_event'].containing_oneof = _EVENT.oneofs_by_name['type']
_EVENT.oneofs_by_name['type'].fields.append(
  _EVENT.fields_by_name['transaction_event'])
_EVENT.fields_by_name['transaction_event'].containing_oneof = _EVENT.oneofs_by_name['type']
_ACCOUNTUPDATEEVENT.fields_by_name['account_info'].message_type = _ACCOUNTINFO
DESCRIPTOR.message_types_by_name['AccountInfo'] = _ACCOUNTINFO
DESCRIPTOR.message_types_by_name['CreateAccountRequest'] = _CREATEACCOUNTREQUEST
DESCRIPTOR.message_types_by_name['CreateAccountResponse'] = _CREATEACCOUNTRESPONSE
DESCRIPTOR.message_types_by_name['GetAccountInfoRequest'] = _GETACCOUNTINFOREQUEST
DESCRIPTOR.message_types_by_name['GetAccountInfoResponse'] = _GETACCOUNTINFORESPONSE
DESCRIPTOR.message_types_by_name['GetEventsRequest'] = _GETEVENTSREQUEST
DESCRIPTOR.message_types_by_name['Events'] = _EVENTS
DESCRIPTOR.message_types_by_name['Event'] = _EVENT
DESCRIPTOR.message_types_by_name['AccountUpdateEvent'] = _ACCOUNTUPDATEEVENT
DESCRIPTOR.message_types_by_name['TransactionEvent'] = _TRANSACTIONEVENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AccountInfo = _reflection.GeneratedProtocolMessageType('AccountInfo', (_message.Message,), {
  'DESCRIPTOR' : _ACCOUNTINFO,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.AccountInfo)
  })
_sym_db.RegisterMessage(AccountInfo)

CreateAccountRequest = _reflection.GeneratedProtocolMessageType('CreateAccountRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEACCOUNTREQUEST,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.CreateAccountRequest)
  })
_sym_db.RegisterMessage(CreateAccountRequest)

CreateAccountResponse = _reflection.GeneratedProtocolMessageType('CreateAccountResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEACCOUNTRESPONSE,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.CreateAccountResponse)
  })
_sym_db.RegisterMessage(CreateAccountResponse)

GetAccountInfoRequest = _reflection.GeneratedProtocolMessageType('GetAccountInfoRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETACCOUNTINFOREQUEST,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.GetAccountInfoRequest)
  })
_sym_db.RegisterMessage(GetAccountInfoRequest)

GetAccountInfoResponse = _reflection.GeneratedProtocolMessageType('GetAccountInfoResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETACCOUNTINFORESPONSE,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.GetAccountInfoResponse)
  })
_sym_db.RegisterMessage(GetAccountInfoResponse)

GetEventsRequest = _reflection.GeneratedProtocolMessageType('GetEventsRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETEVENTSREQUEST,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.GetEventsRequest)
  })
_sym_db.RegisterMessage(GetEventsRequest)

Events = _reflection.GeneratedProtocolMessageType('Events', (_message.Message,), {
  'DESCRIPTOR' : _EVENTS,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.Events)
  })
_sym_db.RegisterMessage(Events)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), {
  'DESCRIPTOR' : _EVENT,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.Event)
  })
_sym_db.RegisterMessage(Event)

AccountUpdateEvent = _reflection.GeneratedProtocolMessageType('AccountUpdateEvent', (_message.Message,), {
  'DESCRIPTOR' : _ACCOUNTUPDATEEVENT,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.AccountUpdateEvent)
  })
_sym_db.RegisterMessage(AccountUpdateEvent)

TransactionEvent = _reflection.GeneratedProtocolMessageType('TransactionEvent', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTIONEVENT,
  '__module__' : 'account.v3.account_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.account.v3.TransactionEvent)
  })
_sym_db.RegisterMessage(TransactionEvent)


DESCRIPTOR._options = None
_ACCOUNTINFO.fields_by_name['account_id']._options = None
_CREATEACCOUNTREQUEST.fields_by_name['account_id']._options = None
_GETACCOUNTINFOREQUEST.fields_by_name['account_id']._options = None
_GETEVENTSREQUEST.fields_by_name['account_id']._options = None
_EVENTS.fields_by_name['events']._options = None
_ACCOUNTUPDATEEVENT.fields_by_name['account_info']._options = None
_TRANSACTIONEVENT.fields_by_name['envelope_xdr']._options = None
_TRANSACTIONEVENT.fields_by_name['result_xdr']._options = None

_ACCOUNT = _descriptor.ServiceDescriptor(
  name='Account',
  full_name='kin.agora.account.v3.Account',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1365,
  serialized_end=1674,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateAccount',
    full_name='kin.agora.account.v3.Account.CreateAccount',
    index=0,
    containing_service=None,
    input_type=_CREATEACCOUNTREQUEST,
    output_type=_CREATEACCOUNTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetAccountInfo',
    full_name='kin.agora.account.v3.Account.GetAccountInfo',
    index=1,
    containing_service=None,
    input_type=_GETACCOUNTINFOREQUEST,
    output_type=_GETACCOUNTINFORESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetEvents',
    full_name='kin.agora.account.v3.Account.GetEvents',
    index=2,
    containing_service=None,
    input_type=_GETEVENTSREQUEST,
    output_type=_EVENTS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ACCOUNT)

DESCRIPTOR.services_by_name['Account'] = _ACCOUNT

# @@protoc_insertion_point(module_scope)
