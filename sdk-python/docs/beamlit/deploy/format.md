Module beamlit.deploy.format
============================
This module provides utility functions to format deployment configurations into YAML-compatible strings.
It includes functions to convert arguments, parameters, dictionaries, and agent chains into properly formatted JSON or YAML strings.

Functions
---------

`arg_to_dict(arg: ast.keyword)`
:   Converts an AST keyword argument to a dictionary.
    
    Args:
        arg (ast.keyword): The AST keyword argument.
    
    Returns:
        dict: The resulting dictionary.

`arg_to_list(arg: ast.List)`
:   

`format_agent_chain(agentChain: list[beamlit.models.agent_chain.AgentChain]) ‑> str`
:   Formats agent chain configuration into a YAML-compatible string.
    
    Args:
        agentChain (list[AgentChain]): List of agent chain configurations.
    
    Returns:
        str: YAML-formatted string of agent chain.

`format_dict(obj: dict) ‑> str`
:   Converts a dictionary to a YAML-compatible string.
    
    Args:
        obj (dict): The dictionary to format.
    
    Returns:
        str: YAML-formatted string representation of the dictionary.

`format_parameters(parameters: list[beamlit.models.store_function_parameter.StoreFunctionParameter]) ‑> str`
:   Formats function parameters into a YAML-compatible string.
    
    Args:
        parameters (list[StoreFunctionParameter]): List of parameter objects.
    
    Returns:
        str: YAML-formatted string of parameters.

`format_value(v)`
:   Formats an AST node value into its Python equivalent.
    
    Args:
        v (ast.AST): The AST node to format.
    
    Returns:
        Any: The formatted Python value.