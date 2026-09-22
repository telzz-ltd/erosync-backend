from enum import StrEnum
from smtplib import SMTP

from jinja2 import Environment, FileSystemLoader, select_autoescape


class Template(StrEnum):
    VERIFICATION_CODE = "verification_code"
    PASSWORD_RESET = "password_reset"
    WELCOME = "welcome"


def load_template(template_name: str, **kwargs) -> str:
    env = Environment(
        loader=FileSystemLoader("templates/email"), autoescape=select_autoescape()
    )
    template = env.get_template(f"{template_name}.html")
    return template.render(kwargs)


def send_mail(template: str, args: dict, to: str):
    with SMTP("127.0.0.1") as smtp:
        smtp.noop()
        smtp.send_message(
            msg=load_template(template, **args),
            from_addr="support@erosyncng.com",
            to_addrs=to,
        )
