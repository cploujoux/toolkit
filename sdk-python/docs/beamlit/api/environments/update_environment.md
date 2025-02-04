Module beamlit.api.environments.update_environment
==================================================

Functions
---------

`asyncio(environment_name: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.models.environment.Environment | None`
:   Update environment
    
     Updates an environment.
    
    Args:
        environment_name (str):
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Environment

`asyncio_detailed(environment_name: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.types.Response[beamlit.models.environment.Environment]`
:   Update environment
    
     Updates an environment.
    
    Args:
        environment_name (str):
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Environment]

`sync(environment_name: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.models.environment.Environment | None`
:   Update environment
    
     Updates an environment.
    
    Args:
        environment_name (str):
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Environment

`sync_detailed(environment_name: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.types.Response[beamlit.models.environment.Environment]`
:   Update environment
    
     Updates an environment.
    
    Args:
        environment_name (str):
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Environment]