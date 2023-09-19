from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Gender(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    MAN: _ClassVar[Gender]
    WOMAN: _ClassVar[Gender]
    GENDER_OTHER: _ClassVar[Gender]

class Affiliation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    STUDENT: _ClassVar[Affiliation]
    AFFILIATION_OTHER: _ClassVar[Affiliation]
MAN: Gender
WOMAN: Gender
GENDER_OTHER: Gender
STUDENT: Affiliation
AFFILIATION_OTHER: Affiliation

class GetUserRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class GetUserResponse(_message.Message):
    __slots__ = ["id", "name", "email", "gender", "affiliation"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    AFFILIATION_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    email: str
    gender: Gender
    affiliation: Affiliation
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., email: _Optional[str] = ..., gender: _Optional[_Union[Gender, str]] = ..., affiliation: _Optional[_Union[Affiliation, str]] = ...) -> None: ...

class CreateUserRequest(_message.Message):
    __slots__ = ["name", "email", "password", "gender", "affiliation", "profile_image_url"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    AFFILIATION_FIELD_NUMBER: _ClassVar[int]
    PROFILE_IMAGE_URL_FIELD_NUMBER: _ClassVar[int]
    name: str
    email: str
    password: str
    gender: Gender
    affiliation: Affiliation
    profile_image_url: str
    def __init__(self, name: _Optional[str] = ..., email: _Optional[str] = ..., password: _Optional[str] = ..., gender: _Optional[_Union[Gender, str]] = ..., affiliation: _Optional[_Union[Affiliation, str]] = ..., profile_image_url: _Optional[str] = ...) -> None: ...

class CreateUserResponse(_message.Message):
    __slots__ = ["id", "name", "email", "gender", "affiliation"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    AFFILIATION_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    email: str
    gender: Gender
    affiliation: Affiliation
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., email: _Optional[str] = ..., gender: _Optional[_Union[Gender, str]] = ..., affiliation: _Optional[_Union[Affiliation, str]] = ...) -> None: ...

class Group(_message.Message):
    __slots__ = ["id", "name", "description"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    description: str
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class Skill(_message.Message):
    __slots__ = ["id", "name", "description", "level"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    description: str
    level: int
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., level: _Optional[int] = ...) -> None: ...
