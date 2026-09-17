import datetime as dt
from dataclasses import dataclass
from enum import StrEnum
from pyexpat import expat_CAPI
from time import timezone


class OTPChannel(StrEnum):
    EMAIL = "EMAIL"
    SMS = "SMS"


class OTPPurpose(StrEnum):
    VERIFY_EMAIL = "VERIFY_EMAIL"
    RESET_PASSWORD = "RESET_PASSWORD"



@dataclass
class OTP:
    recipient: str
    code_hash: str
    purpose: OTPPurpose
    channel: OTPChannel
    expires_at: dt.datetime
    created_at: dt.datetime
    attempts: int = 0
    max_attempts: int = 5

    @classmethod
    def create(
        cls,
        recipient: str,
        code_hash: str,
        purpose: OTPPurpose,
        channel: OTPChannel,
        expire_min=10,
    ):
        if not recipient or not code_hash:
            raise ValueError("all fields are required")

        return cls(
            recipient=recipient,
            code_hash=code_hash,
            purpose=purpose,
            channel=channel,
            created_at=dt.datetime.now(dt.UTC),
            expires_at=dt.datetime.now(dt.UTC) + dt.timedelta(minutes=expire_min),
        )

    def increment_attempt(self):
        self.attempts += 1

    def is_valid(self) -> bool:
        return (
            self.expires_at >= dt.datetime.now(dt.UTC)
            and self.attempts < self.max_attempts
        )
