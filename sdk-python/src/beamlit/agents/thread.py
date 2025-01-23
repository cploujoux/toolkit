
import jwt
from fastapi import Request


def get_default_thread(request: Request) -> str:
    if request.headers.get("X-Beamlit-Sub"):
        return request.headers.get("X-Beamlit-Sub")
    authorization = request.headers.get("Authorization", request.headers.get("X-Beamlit-Authorization"))
    if authorization and len(authorization.split("Bearer ")) > 1:
        token = authorization.split(" ")[1]
        decoded = jwt.decode(token, options={"verify_signature": False})
        return decoded["sub"]
    return ""