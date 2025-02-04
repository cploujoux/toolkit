Module beamlit.functions
========================
Functions package providing function decorators and utilities for Beamlit integration.
It includes decorators for creating function tools and utilities for managing and retrieving functions.

Sub-modules
-----------
* beamlit.functions.common
* beamlit.functions.decorator
* beamlit.functions.mcp
* beamlit.functions.remote

Functions
---------

`function(*args, function: beamlit.models.function.Function | dict = None, kit=False, **kwargs: dict) ‑> <class 'collections.abc.Callable'>`
:   Decorator to create function tools with Beamlit and LangChain integration.
    
    Args:
        function (Function | dict): Function metadata or a dictionary representing it.
        kit (bool): Whether to associate a function kit.
        **kwargs (dict): Additional keyword arguments for function configuration.
    
    Returns:
        Callable: The decorated function.

`get_functions(remote_functions: list[str] | None = None, client: beamlit.client.AuthenticatedClient | None = None, dir: str | None = None, chain: list[beamlit.models.agent_chain.AgentChain] | None = None, remote_functions_empty: bool = True, from_decorator: str = 'function', warning: bool = True)`
:   

`kit(bl_kit: beamlit.models.function_kit.FunctionKit = None, **kwargs: dict) ‑> <class 'collections.abc.Callable'>`
:   Decorator to create function tools with Beamlit and LangChain integration.
    
    Args:
        bl_kit (FunctionKit | None): Optional FunctionKit to associate with the function.
        **kwargs (dict): Additional keyword arguments for function configuration.
    
    Returns:
        Callable: The decorated function.