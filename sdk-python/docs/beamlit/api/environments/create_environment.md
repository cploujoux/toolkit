Module beamlit.api.environments.create_environment
==================================================

Functions
---------

`asyncio(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.models.environment.Environment | None`
:   Create environment
    
     Creates an environment.
    
    Args:
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Environment

`asyncio_detailed(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.types.Response[beamlit.models.environment.Environment]`
:   Create environment
    
     Creates an environment.
    
    Args:
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Environment]

`sync(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.models.environment.Environment | None`
:   Create environment
    
     Creates an environment.
    
    Args:
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Environment

`sync_detailed(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.environment.Environment) ‑> beamlit.types.Response[beamlit.models.environment.Environment]`
:   Create environment
    
     Creates an environment.
    
    Args:
        body (Environment): Environment on which deployments will be made (e.g. development,
            production), enforcing multiple policies at once.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Environment]