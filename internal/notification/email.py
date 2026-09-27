from dataclasses import dataclass


@dataclass
class MailConfig:
    host: str
    port: int
    username: str
    password: str


class Mail():
    def __init__(self, config: MailConfig) -> None:
        self.config = config

    def send():
        ...
