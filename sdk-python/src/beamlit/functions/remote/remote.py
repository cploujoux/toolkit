import asyncio
import warnings
from typing import Any, Callable

import pydantic
import typing_extensions as t
from dataclasses import dataclass
from beamlit.api.functions import get_function
from beamlit.authentication.authentication import AuthenticatedClient
from beamlit.run import RunClient
from beamlit.common.settings import get_settings
from beamlit.models import Function, StoreFunctionParameter
from langchain_core.tools.base import BaseTool, ToolException
from pydantic.json_schema import JsonSchemaValue
from pydantic_core import core_schema as cs


def create_schema_model(parameters: list[StoreFunctionParameter]) -> type[pydantic.BaseModel]:
    # Create a new model class that returns our JSON schema.
    # LangChain requires a BaseModel class.
    class Schema(pydantic.BaseModel):
        model_config = pydantic.ConfigDict(extra="allow")

        @t.override
        @classmethod
        def __get_pydantic_json_schema__(
            cls, core_schema: cs.CoreSchema, handler: pydantic.GetJsonSchemaHandler
        ) -> JsonSchemaValue:
            schema = {
                "type": "object",
                "properties": {},
                "required": [],
            }
            for parameter in parameters:
                schema["properties"][parameter.name] = {
                    "type": parameter.type_,
                    "description": parameter.description,
                }
                if parameter.required:
                    schema["required"].append(parameter.name)
            return schema

    return Schema


class RemoteTool(BaseTool):
    """
    Remote tool
    """

    client: RunClient
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
        settings = get_settings()
        result = self.client.run(
            "function",
            self.name,
            settings.environment,
            "POST",
            kwargs,
        )
        return result.text

    @t.override
    @property
    def tool_call_schema(self) -> type[pydantic.BaseModel]:
        assert self.args_schema is not None  # noqa: S101
        return self.args_schema

@dataclass
class RemoteToolkit:
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
            self._function = get_function.sync_detailed(self.function, client=self.client).parsed

    @t.override
    def get_tools(self) -> list[BaseTool]:
        if self._function is None:
            raise RuntimeError("Must initialize the toolkit first")

        if self._function.spec.kit:
            return [
                RemoteTool(
                    client=RunClient(self.client),
                    name=func.name,
                    description=func.description or "",
                    args_schema=create_schema_model(func.parameters),
                )
                for func in self._function.spec.kit
            ]

        return [
            RemoteTool(
                client=RunClient(self.client),
                name=self._function.metadata.name,
                description=self._function.spec.description or "",
                args_schema=create_schema_model(self._function.spec.parameters),
            )
        ]