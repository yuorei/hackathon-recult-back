from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Meta(_message.Message):
    __slots__ = ["name", "desc"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESC_FIELD_NUMBER: _ClassVar[int]
    name: str
    desc: str
    def __init__(self, name: _Optional[str] = ..., desc: _Optional[str] = ...) -> None: ...

class UploadRequest(_message.Message):
    __slots__ = ["meta", "data"]
    META_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    meta: Meta
    data: bytes
    def __init__(self, meta: _Optional[_Union[Meta, _Mapping]] = ..., data: _Optional[bytes] = ...) -> None: ...

class UploadResponse(_message.Message):
    __slots__ = ["image_url"]
    IMAGE_URL_FIELD_NUMBER: _ClassVar[int]
    image_url: str
    def __init__(self, image_url: _Optional[str] = ...) -> None: ...
