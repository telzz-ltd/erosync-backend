from enum import StrEnum
from smtplib import SMTP

from jinja2 import Environment, FileSystemLoader, select_autoescape

from email.message import EmailMessage


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


def send_mail(template: str, template_args: dict, to: str, subject: str):
    msg = EmailMessage()
    msg["From"] = "support@erosyncng.com"
    msg["To"] = to
    msg["Subject"] = subject
    msg.set_content("Please view this email in an HTML-compatible client.")
    msg.add_alternative(
        load_template(template, **template_args),
        subtype="html",
    )

    with SMTP("127.0.0.1", 1025) as smtp:
        smtp.send_message(msg=msg)
