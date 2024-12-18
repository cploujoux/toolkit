import asyncio
import urllib.parse
import warnings
from typing import Any, Callable

import pydantic
import pydantic_core
import requests
import typing_extensions as t
from beamlit.authentication import new_client
from beamlit.authentication.authentication import (AuthenticatedClient,
                                                   new_client)
from langchain_core.tools.base import BaseTool, BaseToolkit, ToolException
from mcp import ListToolsResult
from pydantic.json_schema import JsonSchemaValue
from pydantic_core import core_schema as cs

client = new_client()

RUN_URL = "http://localhost:8787"
SUB_ID = "test"


def create_schema_model(schema: dict[str, t.Any]) -> type[pydantic.BaseModel]:
    # Create a new model class that returns our JSON schema.
    # LangChain requires a BaseModel class.
    class Schema(pydantic.BaseModel):
        model_config = pydantic.ConfigDict(extra="allow")

        @t.override
        @classmethod
        def __get_pydantic_json_schema__(
            cls, core_schema: cs.CoreSchema, handler: pydantic.GetJsonSchemaHandler
        ) -> JsonSchemaValue:
            return schema

    return Schema


class MCPClient:
    def __init__(self, client: AuthenticatedClient, server_name: str):
        self.client = client
        self.server_name = server_name

    def list_tools(self) -> requests.Response:
        client = self.client.get_httpx_client()
        url = urllib.parse.urljoin(RUN_URL, f"{self.server_name}/{SUB_ID}/tools/list")

        response = client.request("GET", url)
        response.raise_for_status()
        return response

    def call_tool(self, tool_name: str, arguments: dict[str, Any] = None) -> requests.Response:
        client = self.client.get_httpx_client()
        url = urllib.parse.urljoin(RUN_URL, f"{self.server_name}/{SUB_ID}/tools/call")

        response = client.request("POST", url, json={"name": tool_name, "arguments": arguments})
        response.raise_for_status()
        return response

class MCPTool(BaseTool):
    """
    MCP server tool
    """

    client: MCPClient
    handle_tool_error: bool | str | Callable[[ToolException], str] | None = True

    @t.override
    def _run(self, *args: t.Any, **kwargs: t.Any) -> t.Any:
        warnings.warn(
            "Invoke this tool asynchronousely using `ainvoke`. This method exists only to satisfy standard tests.",
            stacklevel=1,
        )
        return asyncio.run(self._arun(*args, **kwargs))

    @t.override
    async def _arun(self, *args: t.Any, **kwargs: t.Any) -> t.Any:
        result = self.client.call_tool(self.name, arguments=kwargs)
        response = result.json()
        content = pydantic_core.to_json(response["content"]).decode()
        if response["isError"]:
            raise ToolException(content)
        return content

    @t.override
    @property
    def tool_call_schema(self) -> type[pydantic.BaseModel]:
        assert self.args_schema is not None  # noqa: S101
        return self.args_schema

class MCPToolkit(BaseToolkit):
    """
    MCP server toolkit
    """

    client: MCPClient
    """The MCP session used to obtain the tools"""

    _tools: ListToolsResult | None = None

    model_config = pydantic.ConfigDict(arbitrary_types_allowed=True)

    def initialize(self) -> None:
        """Initialize the session and retrieve tools list"""
        if self._tools is None:
            response = self.client.list_tools()
            self._tools = ListToolsResult(**response.json())

    @t.override
    def get_tools(self) -> list[BaseTool]:
        if self._tools is None:
            raise RuntimeError("Must initialize the toolkit first")

        return [
            MCPTool(
                client=self.client,
                name=tool.name,
                description=tool.description or "",
                args_schema=create_schema_model(tool.inputSchema),
            )
            # list_tools returns a PaginatedResult, but I don't see a way to pass the cursor to retrieve more tools
            for tool in self._tools.tools
        ]


def get_generic_tools(name):
    mcp_client = McpClient(client, name)
    return MCPToolkit(client=mcp_client)
