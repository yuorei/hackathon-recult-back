# frozen_string_literal: true
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

require 'google/protobuf'


descriptor_data = "\n\nuser.proto\x12\x04user\"\x1c\n\x0eGetUserRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"\x80\x01\n\x0fGetUserResponse\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x1c\n\x06gender\x18\x04 \x01(\x0e\x32\x0c.user.Gender\x12&\n\x0b\x61\x66\x66iliation\x18\x05 \x01(\x0e\x32\x11.user.Affiliation\"\x88\x01\n\x11\x43reateUserRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\x1c\n\x06gender\x18\x04 \x01(\x0e\x32\x0c.user.Gender\x12&\n\x0b\x61\x66\x66iliation\x18\x05 \x01(\x0e\x32\x11.user.Affiliation\"\x83\x01\n\x12\x43reateUserResponse\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x1c\n\x06gender\x18\x04 \x01(\x0e\x32\x0c.user.Gender\x12&\n\x0b\x61\x66\x66iliation\x18\x05 \x01(\x0e\x32\x11.user.Affiliation\"6\n\x05Group\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\"E\n\x05Skill\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\r\n\x05level\x18\x04 \x01(\x05*.\n\x06Gender\x12\x07\n\x03MAN\x10\x00\x12\t\n\x05WOMAN\x10\x01\x12\x10\n\x0cGENDER_OTHER\x10\x02*1\n\x0b\x41\x66\x66iliation\x12\x0b\n\x07STUDENT\x10\x00\x12\x15\n\x11\x41\x46\x46ILIATION_OTHER\x10\x01\x32\x8a\x01\n\x0bUserService\x12\x38\n\x07GetUser\x12\x14.user.GetUserRequest\x1a\x15.user.GetUserResponse\"\x00\x12\x41\n\nCreateUser\x12\x17.user.CreateUserRequest\x1a\x18.user.CreateUserResponse\"\x00\x42\x07Z\x05/userb\x06proto3"

pool = Google::Protobuf::DescriptorPool.generated_pool

begin
  pool.add_serialized_file(descriptor_data)
rescue TypeError => e
  # Compatibility code: will be removed in the next major version.
  require 'google/protobuf/descriptor_pb'
  parsed = Google::Protobuf::FileDescriptorProto.decode(descriptor_data)
  parsed.clear_dependency
  serialized = parsed.class.encode(parsed)
  file = pool.add_serialized_file(serialized)
  warn "Warning: Protobuf detected an import path issue while loading generated file #{__FILE__}"
  imports = [
  ]
  imports.each do |type_name, expected_filename|
    import_file = pool.lookup(type_name).file_descriptor
    if import_file.name != expected_filename
      warn "- #{file.name} imports #{expected_filename}, but that import was loaded as #{import_file.name}"
    end
  end
  warn "Each proto file must use a consistent fully-qualified name."
  warn "This will become an error in the next major version."
end

module User
  GetUserRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.GetUserRequest").msgclass
  GetUserResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.GetUserResponse").msgclass
  CreateUserRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.CreateUserRequest").msgclass
  CreateUserResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.CreateUserResponse").msgclass
  Group = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.Group").msgclass
  Skill = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.Skill").msgclass
  Gender = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.Gender").enummodule
  Affiliation = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.Affiliation").enummodule
end
