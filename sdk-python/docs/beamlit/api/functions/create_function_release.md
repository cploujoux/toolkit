Module beamlit.api.functions.create_function_release
====================================================

Functions
---------

`asyncio(function_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.function_release.FunctionRelease | None`
:   Create release for a function from an environment
    
    Args:
        function_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        FunctionRelease

`asyncio_detailed(function_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.function_release.FunctionRelease]`
:   Create release for a function from an environment
    
    Args:
        function_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[FunctionRelease]

`sync(function_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.function_release.FunctionRelease | None`
:   Create release for a function from an environment
    
    Args:
        function_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        FunctionRelease

`sync_detailed(function_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.function_release.FunctionRelease]`
:   Create release for a function from an environment
    
    Args:
        function_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[FunctionRelease]