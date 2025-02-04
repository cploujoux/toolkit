Module beamlit.api.environments.get_environment
===============================================

Functions
---------

`asyncio(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.environment.Environment | None`
:   Get environment
    
     Returns an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Environment

`asyncio_detailed(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.environment.Environment]`
:   Get environment
    
     Returns an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Environment]

`sync(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.environment.Environment | None`
:   Get environment
    
     Returns an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Environment

`sync_detailed(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.environment.Environment]`
:   Get environment
    
     Returns an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Environment]