from typing import Protocol

from .domain import OTP


class OTPStore(Protocol):
    def save(self, otp: OTP) -> None: ...
