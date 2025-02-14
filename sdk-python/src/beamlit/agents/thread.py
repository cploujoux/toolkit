"""Module: thread

Defines threading capabilities for agents.
"""
import jwt
from fastapi import Request


def get_default_thread(request: Request) -> str:
    """
    Extracts the default thread identifier from an incoming HTTP request.
    Prioritizes the `X-Beamlit-Sub` header and falls back to decoding the JWT
    from the `Authorization` or `X-Beamlit-Authorization` headers.

    Parameters:
        request (Request): The incoming HTTP request object.

    Returns:
        str: The extracted thread identifier. Returns an empty string if no valid identifier is found.
    """
    if request.headers.get("X-Beamlit-Thread-Id"):
        return request.headers.get("X-Beamlit-Thread-Id")
    if request.headers.get("Thread-Id"):
        return request.headers.get("Thread-Id")
    return ""
