# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x12\x04user\"\x1c\n\x0eGetUserRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"\x80\x01\n\x0fGetUserResponse\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x1c\n\x06gender\x18\x04 \x01(\x0e\x32\x0c.user.Gender\x12&\n\x0b\x61\x66\x66iliation\x18\x05 \x01(\x0e\x32\x11.user.Affiliation\"\xa3\x01\n\x11\x43reateUserRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\x1c\n\x06gender\x18\x04 \x01(\x0e\x32\x0c.user.Gender\x12&\n\x0b\x61\x66\x66iliation\x18\x05 \x01(\x0e\x32\x11.user.Affiliation\x12\x19\n\x11profile_image_url\x18\x08 \x01(\t\"\x83\x01\n\x12\x43reateUserResponse\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x1c\n\x06gender\x18\x04 \x01(\x0e\x32\x0c.user.Gender\x12&\n\x0b\x61\x66\x66iliation\x18\x05 \x01(\x0e\x32\x11.user.Affiliation\"6\n\x05Group\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\"E\n\x05Skill\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\r\n\x05level\x18\x04 \x01(\x05*.\n\x06Gender\x12\x07\n\x03MAN\x10\x00\x12\t\n\x05WOMAN\x10\x01\x12\x10\n\x0cGENDER_OTHER\x10\x02*1\n\x0b\x41\x66\x66iliation\x12\x0b\n\x07STUDENT\x10\x00\x12\x15\n\x11\x41\x46\x46ILIATION_OTHER\x10\x01\x32\x8a\x01\n\x0bUserService\x12\x38\n\x07GetUser\x12\x14.user.GetUserRequest\x1a\x15.user.GetUserResponse\"\x00\x12\x41\n\nCreateUser\x12\x17.user.CreateUserRequest\x1a\x18.user.CreateUserResponse\"\x00\x42\x07Z\x05/userb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\005/user'
  _globals['_GENDER']._serialized_start=608
  _globals['_GENDER']._serialized_end=654
  _globals['_AFFILIATION']._serialized_start=656
  _globals['_AFFILIATION']._serialized_end=705
  _globals['_GETUSERREQUEST']._serialized_start=20
  _globals['_GETUSERREQUEST']._serialized_end=48
  _globals['_GETUSERRESPONSE']._serialized_start=51
  _globals['_GETUSERRESPONSE']._serialized_end=179
  _globals['_CREATEUSERREQUEST']._serialized_start=182
  _globals['_CREATEUSERREQUEST']._serialized_end=345
  _globals['_CREATEUSERRESPONSE']._serialized_start=348
  _globals['_CREATEUSERRESPONSE']._serialized_end=479
  _globals['_GROUP']._serialized_start=481
  _globals['_GROUP']._serialized_end=535
  _globals['_SKILL']._serialized_start=537
  _globals['_SKILL']._serialized_end=606
  _globals['_USERSERVICE']._serialized_start=708
  _globals['_USERSERVICE']._serialized_end=846
# @@protoc_insertion_point(module_scope)
