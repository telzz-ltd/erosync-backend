from typing import Any

from joserfc import jwk, jwt

key = jwk.import_key("my-insecure-jwt-secret", "oct")


def encode(payload: dict[str, Any]) -> str:
    return jwt.encode({"alg": "HS256"}, payload, key)


def decode(token_str: str) -> dict[str, Any]:
    token = jwt.decode(token_str, key)
    jwt.JWTClaimsRegistry().validate(token.claims)
    return token.claims
