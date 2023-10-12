from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class StressLevelRequest(_message.Message):
    __slots__ = ["heart_rate", "systolic_blood_pressure", "diastolic_blood_pressure", "oxygen_saturation", "body_temp_c"]
    HEART_RATE_FIELD_NUMBER: _ClassVar[int]
    SYSTOLIC_BLOOD_PRESSURE_FIELD_NUMBER: _ClassVar[int]
    DIASTOLIC_BLOOD_PRESSURE_FIELD_NUMBER: _ClassVar[int]
    OXYGEN_SATURATION_FIELD_NUMBER: _ClassVar[int]
    BODY_TEMP_C_FIELD_NUMBER: _ClassVar[int]
    heart_rate: int
    systolic_blood_pressure: int
    diastolic_blood_pressure: int
    oxygen_saturation: int
    body_temp_c: float
    def __init__(self, heart_rate: _Optional[int] = ..., systolic_blood_pressure: _Optional[int] = ..., diastolic_blood_pressure: _Optional[int] = ..., oxygen_saturation: _Optional[int] = ..., body_temp_c: _Optional[float] = ...) -> None: ...

class StressLevelResponse(_message.Message):
    __slots__ = ["stress_level_index", "stress_level_label"]
    STRESS_LEVEL_INDEX_FIELD_NUMBER: _ClassVar[int]
    STRESS_LEVEL_LABEL_FIELD_NUMBER: _ClassVar[int]
    stress_level_index: int
    stress_level_label: str
    def __init__(self, stress_level_index: _Optional[int] = ..., stress_level_label: _Optional[str] = ...) -> None: ...
