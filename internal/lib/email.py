from jinja2 import Environment, PackageLoader, select_autoescape

env = Environment(loader=PackageLoader("templates"), autoescape=select_autoescape())

template = env.get_template("send_email_verification_code.html")
