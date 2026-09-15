from .domain import OTP, OTPChannel, OTPPurpose
from .protocols import OTPStore


class OTPService:
    store: OTPStore

    def create(
        self,
        recipient: str,
        channel: OTPChannel,
        purpose: OTPPurpose,
    ) -> str:
        otp = OTP.create(recipient, "", purpose, channel)
        self.store.save(otp)
        return ""

    def validate(
        self,
        code_hash: str,
        recipient: str,
        channel: OTPChannel,
        purpose: OTPPurpose,
    ) -> bool:
        return False
