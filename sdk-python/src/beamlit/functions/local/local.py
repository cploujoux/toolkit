import asyncio
import warnings
from dataclasses import dataclass
from typing import Callable

import pydantic
import typing_extensions as t
from langchain_core.tools.base import BaseTool, ToolException

from beamlit.api.functions import get_function, list_functions
from beamlit.authentication.authentication import AuthenticatedClient
from beamlit.common.settings import get_settings
from beamlit.errors import UnexpectedStatus
from beamlit.functions.mcp.mcp import MCPClient, MCPToolkit
from beamlit.models import Function, StoreFunctionParameter
from beamlit.run import RunClient

@dataclass
class LocalToolKit:
    """
    Toolkit for managing local tools.

    Attributes:
        client (AuthenticatedClient): The authenticated client instance.
        function (str): The name of the local function to integrate.
        _function (Function | None): Cached Function object after initialization.
    """
    client: AuthenticatedClient
    local_function: pydantic.ConfigDict(arbitrary_types_allowed=True)
    _function: Function | None = None
    model_config = pydantic.ConfigDict(arbitrary_types_allowed=True)

    def initialize(self) -> None:
        """Initialize the session and retrieve the local function details."""
        if self._function is None:
            try:
                # For local functions, we directly create the Function object
                # based on the local function name
                self._function = Function(
                    metadata={"name": self.local_function['name']},
                    spec={
                        "configurations": {                            
                            "url": self.local_function['url'],
                            "sse": self.local_function['sse'],
                        },
                        "description": self.local_function['description'] or "",
                    }
                )
            except Exception as e:
                raise RuntimeError(f"Failed to initialize local function: {e}")

    def get_tools(self) -> list[BaseTool]:
        settings = get_settings()
        mcp_client = MCPClient(self.client, self._function.spec["configurations"]["url"])
        mcp_toolkit = MCPToolkit(client=mcp_client, sse=self._function.spec["configurations"]["sse"])
        mcp_toolkit.initialize()
        return mcp_toolkit.get_tools()