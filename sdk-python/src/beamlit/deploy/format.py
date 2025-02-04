"""
This module provides utility functions to format deployment configurations into YAML-compatible strings.
It includes functions to convert arguments, parameters, dictionaries, and agent chains into properly formatted JSON or YAML strings.
"""

import ast

from beamlit.models import AgentChain, StoreFunctionParameter


def arg_to_list(arg: ast.List):
    value = []
    for v in arg.elts:
        value.append(format_value(v))
    return value

def format_value(v):
    """
    Formats an AST node value into its Python equivalent.

    Args:
        v (ast.AST): The AST node to format.

    Returns:
        Any: The formatted Python value.
    """
    if isinstance(v, ast.Constant):
        return v.value
    elif isinstance(v, ast.Dict):
        return arg_to_dict(v)
    elif isinstance(v, ast.List):
        return arg_to_list(v)

def arg_to_dict(arg: ast.keyword):
    """
    Converts an AST keyword argument to a dictionary.

    Args:
        arg (ast.keyword): The AST keyword argument.

    Returns:
        dict: The resulting dictionary.
    """
    value = {}
    for k, v in zip(arg.keys, arg.values):
        if isinstance(k, ast.Constant):
            value[k.value] = format_value(v)
    return value

def format_parameters(parameters: list[StoreFunctionParameter]) -> str:
    """
    Formats function parameters into a YAML-compatible string.

    Args:
        parameters (list[StoreFunctionParameter]): List of parameter objects.

    Returns:
        str: YAML-formatted string of parameters.
    """
    if not parameters:
        return "[]"

    formatted = []
    for param in parameters:
        formatted.append(f"""
    - name: {param.name}
      type: {param.type_}
      required: {str(param.required).lower()}
      description: {param.description}""")

    return "\n".join(formatted)

def format_dict(obj: dict) -> str:
    """
    Converts a dictionary to a YAML-compatible string.

    Args:
        obj (dict): The dictionary to format.

    Returns:
        str: YAML-formatted string representation of the dictionary.
    """
    if not obj:
        return "null"
    ret = ""
    for k, v in obj.items():
        if not v:
            ret += f"{k}: null\n"
        else:
            ret += f"{k}: {v}\n"
    return ret

def format_agent_chain(agentChain: list[AgentChain]) -> str:
    """
    Formats agent chain configuration into a YAML-compatible string.

    Args:
        agentChain (list[AgentChain]): List of agent chain configurations.

    Returns:
        str: YAML-formatted string of agent chain.
    """
    if not agentChain:
        return "[]"
    formatted = []

    for agent in agentChain:
        formatted.append(f"""
  - agent: {agent.name}
    enabled: {agent.enabled}""")
        if agent.description:
            formatted.append(f"        description: {agent.description}")
    return "\n".join(formatted)