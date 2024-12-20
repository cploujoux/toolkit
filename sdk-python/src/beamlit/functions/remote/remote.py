import asyncio
import urllib.parse
import warnings
from typing import Any, Callable

import pydantic
import pydantic_core
import typing_extensions as t
from beamlit.api.functions import get_function
from beamlit.authentication.authentication import AuthenticatedClient
from beamlit.models.function import Function
from langchain_core.tools.base import BaseTool, BaseToolkit, ToolException
from pydantic.json_schema import JsonSchemaValue
from pydantic_core import core_schema as cs


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


class RemoteTool(BaseTool):
    """
    Remote tool
    """

    client: AuthenticatedClient
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

class RemoteToolkit(BaseToolkit):
    """
    Remote toolkit
    """

    client: AuthenticatedClient
    function: str
    _function: Function | None = None

    model_config = pydantic.ConfigDict(arbitrary_types_allowed=True)

    def initialize(self) -> None:
        """Initialize the session and retrieve tools list"""
        if self._function is None:
            self._function = get_function(self.function, client=self.client)

    @t.override
    def get_tools(self) -> list[BaseTool]:
        if self._tools is None:
            raise RuntimeError("Must initialize the toolkit first")

        if self._function.spec.kit:
            return [
                RemoteTool(
                client=self.client,
                name=func.name,
                description=func.description or "",
                args_schema=create_schema_model(func.parameters),
                )
                for func in self._function.spec.kit
            ]

        return [
            RemoteTool(
                client=self.client,
                name=self._function.metadata.name,
                description=self._function.spec.description or "",
                args_schema=create_schema_model(self._function.spec.parameters),
            )
        ]