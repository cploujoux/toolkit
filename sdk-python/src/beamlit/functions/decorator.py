"""
This module provides decorators for creating function tools with Beamlit and LangChain integration.
It includes decorators to generate tools, manage toolkits, and handle remote interactions.
"""

import ast
import asyncio
import functools
import importlib.util
import os
from collections.abc import Callable
from logging import getLogger
from typing import Union

from beamlit.authentication import new_client
from beamlit.client import AuthenticatedClient
from beamlit.common import slugify
from beamlit.common.settings import get_settings
from beamlit.functions.mcp.mcp import MCPClient, MCPToolkit
from beamlit.functions.remote.remote import RemoteToolkit
from beamlit.models import AgentChain, Function, FunctionKit
from fastapi import Request
from langchain_core.tools import StructuredTool
from langchain_core.tools.base import create_schema_from_function

logger = getLogger(__name__)

def get_functions(
    mcp_hub:Union[list[str], None]=None,
    remote_functions:Union[list[str], None]=None,
    client:Union[AuthenticatedClient, None]=None,
    dir:Union[str, None]=None,
    chain:Union[list[AgentChain], None]=None,
    remote_functions_empty:bool=True,
    from_decorator:str="function",
    warning:bool=True,
):
    """
    Retrieves a list of function tools for Beamlit, including local, MCP server, and remote functions.

    Args:
        mcp_hub (Union[list[str], None]): List of MCP server URLs to fetch tools from.
        remote_functions (Union[list[str], None]): List of remote function URLs to fetch tools from.
        client (Union[AuthenticatedClient, None]): Authenticated client for making API requests.
        dir (Union[str, None]): Directory path to scan for local function definitions.
        chain (Union[list[AgentChain], None]): List of agent chains to integrate with function tools.
        remote_functions_empty (bool): Whether to warn if remote functions are missing.
        from_decorator (str): Name of the decorator to look for in function definitions.
        warning (bool): Whether to show warnings during processing.

    Returns:
        list[BaseTool]: A list of configured tools for use in Beamlit.
    """
    from beamlit.agents.chain import ChainToolkit

    settings = get_settings()
    if client is None:
        client = new_client()
    if dir is None:
        dir = settings.agent.functions_directory

    functions = []
    logger = getLogger(__name__)
    settings = get_settings()

    # Walk through all Python files in functions directory and subdirectories
    if not os.path.exists(dir):
        if remote_functions_empty and warning:
            logger.warn(f"Functions directory {dir} not found")
    if os.path.exists(dir):
        for root, _, files in os.walk(dir):
            for file in files:
                if file.endswith(".py"):
                    file_path = os.path.join(root, file)
                    # Read and compile the file content
                    with open(file_path) as f:
                        try:
                            file_content = f.read()
                            # Parse the file content to find decorated functions
                            tree = ast.parse(file_content)

                            # Look for function definitions with decorators
                            for node in ast.walk(tree):
                                if (
                                    not isinstance(node, ast.FunctionDef)
                                    and not isinstance(node, ast.AsyncFunctionDef)
                                ) or len(node.decorator_list) == 0:
                                    continue
                                decorator = node.decorator_list[0]

                                decorator_name = ""
                                if isinstance(decorator, ast.Call):
                                    decorator_name = decorator.func.id
                                if isinstance(decorator, ast.Name):
                                    decorator_name = decorator.id
                                if decorator_name == from_decorator:
                                    # Get the function name and decorator name
                                    func_name = node.name

                                    # Import the module to get the actual function
                                    spec = importlib.util.spec_from_file_location(func_name, file_path)
                                    module = importlib.util.module_from_spec(spec)
                                    spec.loader.exec_module(module)
                                    # Check if kit=True in the decorator arguments
                                    is_kit = False
                                    if isinstance(decorator, ast.Call):
                                        for keyword in decorator.keywords:
                                            if keyword.arg == "kit" and isinstance(
                                                keyword.value, ast.Constant
                                            ):
                                                is_kit = keyword.value.value
                                    if is_kit and not settings.remote:
                                        kit_functions = get_functions(
                                            client=client,
                                            dir=os.path.join(root),
                                            remote_functions_empty=remote_functions_empty,
                                            from_decorator="kit",
                                        )
                                        functions.extend(kit_functions)

                                    # Get the decorated function
                                    if not is_kit and hasattr(module, func_name):
                                        func = getattr(module, func_name)
                                        if settings.remote:
                                            toolkit = RemoteToolkit(client, slugify(func.__name__))
                                            toolkit.initialize()
                                            functions.extend(toolkit.get_tools())
                                        else:
                                            if asyncio.iscoroutinefunction(func):
                                                functions.append(
                                                    StructuredTool(
                                                        name=func.__name__,
                                                        description=func.__doc__,
                                                        func=func,
                                                        coroutine=func,
                                                        args_schema=create_schema_from_function(func.__name__, func)
                                                    )
                                                )
                                            else:

                                                functions.append(
                                                    StructuredTool(
                                                        name=func.__name__,
                                                        description=func.__doc__,
                                                        func=func,
                                                        args_schema=create_schema_from_function(func.__name__, func)
                                                    )
                                                )
                        except Exception as e:
                            logger.warning(f"Error processing {file_path}: {e!s}")

    if mcp_hub:
        for server in mcp_hub:
            try:
                mcp_client = MCPClient(client, server)
                toolkit = MCPToolkit(client=mcp_client)
                toolkit.initialize()
                functions.extend(toolkit.get_tools())
            except Exception as e:
                logger.warn(f"Failed to initialize MCP server {server}: {e!s}")

    if remote_functions:
        for function in remote_functions:
            try:
                toolkit = RemoteToolkit(client, function)
                toolkit.initialize()
                functions.extend(toolkit.get_tools())
            except Exception as e:
                logger.warn(f"Failed to initialize remote function {function}: {e!s}")

    if chain:
        toolkit = ChainToolkit(client, chain)
        toolkit.initialize()
        functions.extend(toolkit.get_tools())

    return functions



def kit(bl_kit: FunctionKit = None, **kwargs: dict) -> Callable:
    """
    Decorator to create function tools with Beamlit and LangChain integration.

    Args:
        bl_kit (FunctionKit | None): Optional FunctionKit to associate with the function.
        **kwargs (dict): Additional keyword arguments for function configuration.

    Returns:
        Callable: The decorated function.
    """

    def wrapper(func: Callable) -> Callable:
        if bl_kit and not func.__doc__ and bl_kit.description:
            func.__doc__ = bl_kit.description
        return func

    return wrapper


def function(*args, function: Function | dict = None, kit=False, **kwargs: dict) -> Callable:
    """
    Decorator to create function tools with Beamlit and LangChain integration.

    Args:
        function (Function | dict): Function metadata or a dictionary representing it.
        kit (bool): Whether to associate a function kit.
        **kwargs (dict): Additional keyword arguments for function configuration.

    Returns:
        Callable: The decorated function.
    """
    if function is not None and not isinstance(function, dict):
        raise Exception(
            'function must be a dictionary, example: @function(function={"metadata": {"name": "my_function"}})'
        )
    if isinstance(function, dict):
        function = Function(**function)

    def wrapper(func: Callable) -> Callable:
        if function and not func.__doc__ and function.spec and function.spec.description:
            func.__doc__ = function.spec.description

        @functools.wraps(func)
        async def wrapped(*args, **kwargs):
            if len(args) > 0 and isinstance(args[0], Request):
                body = await args[0].json()
                args = [body.get(param) for param in func.__code__.co_varnames[:func.__code__.co_argcount]]
            return await func(*args, **kwargs) if asyncio.iscoroutinefunction(func) else func(*args, **kwargs)
        return wrapped

    return wrapper
