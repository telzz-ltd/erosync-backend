from typing import Protocol
from dataclasses import dataclass

from .domain import OTP


@dataclass
class FindOTPParam:
    recipient: str
    purpose: str
    channel: str


class OTPStore(Protocol):
    def save(self, otp: OTP) -> None: ...
    def find_one(self, param: FindOTPParam) -> OTP: ...
    def delete(self, otp: OTP) -> None: ...
