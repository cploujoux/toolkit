"""
This module provides functionality for executing HTTP requests against Beamlit resources.
"""
import urllib.parse
from typing import Any

import requests

from beamlit.client import AuthenticatedClient
from beamlit.common import HTTPError, get_settings


class RunClient:
    """Provides functionality for executing HTTP requests against Beamlit resources.

    This module contains the RunClient class which handles authenticated HTTP requests to Beamlit
    resources. It allows users to interact with different resource types (like functions or services)
    in specific environments, supporting various HTTP methods and request parameters.

    Example:
        ```python
        client = new_client()
        run_client = RunClient(client)
        response = run_client.run(
            resource_type="function",
            resource_name="my-function",
            environment="prod",
            method="POST",
            json={"key": "value"}
        )
        ```

    Args:
        client (AuthenticatedClient): An authenticated client instance for making HTTP requests.
    """

    def __init__(self, client: AuthenticatedClient):
        self.client = client

    def run(
        self,
        resource_type: str,
        resource_name: str,
        environment: str,
        method: str,
        path: str = "",
        headers: dict[str, str] | None = None,
        json: dict[str, Any] | None = None,
        data: str | None = None,
        params: dict[str, str] | None = None,
    ) -> requests.Response:
        """Execute an HTTP request against a Beamlit resource.

        Args:
            resource_type (str): The type of resource to interact with (e.g., 'function', 'service').
            resource_name (str): The name of the specific resource.
            environment (str): The environment to execute the request in.
            method (str): The HTTP method to use (e.g., 'GET', 'POST', 'PUT', 'DELETE').
            path (str, optional): Additional path segments to append to the resource URL. Defaults to "".
            headers (dict[str, str] | None, optional): HTTP headers to include in the request. Defaults to None.
            json (dict[str, Any] | None, optional): JSON payload to send with the request. Defaults to None.
            data (str | None, optional): Raw data to send with the request. Defaults to None.
            params (dict[str, str] | None, optional): Query parameters to include in the URL. Defaults to None.

        Returns:
            requests.Response: The HTTP response from the server.

        Raises:
            HTTPError: If the server responds with a status code >= 400.
        """
        settings = get_settings()
        headers = headers or {}
        params = params or {}

        # Build the path
        if path:
            path = f"{settings.workspace}/{resource_type}s/{resource_name}/{path}"
        else:
            path = f"{settings.workspace}/{resource_type}s/{resource_name}"

        client = self.client.get_httpx_client()
        url = urllib.parse.urljoin(settings.run_url, path)

        kwargs = {
            "headers": headers,
            "params": {"environment": environment, **params},
        }
        if data:
            kwargs["data"] = data
        if json:
            kwargs["json"] = json

        response = client.request(method, url, **kwargs)
        if response.status_code >= 400:
            raise HTTPError(response.status_code, response.text)
        return response
