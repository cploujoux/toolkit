Module beamlit.api.environments.get_environment_metrics
=======================================================

Functions
---------

`asyncio(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.environment_metrics.EnvironmentMetrics | None`
:   Get environment metrics
    
     Returns metrics for an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        EnvironmentMetrics

`asyncio_detailed(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.environment_metrics.EnvironmentMetrics]`
:   Get environment metrics
    
     Returns metrics for an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[EnvironmentMetrics]

`sync(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.environment_metrics.EnvironmentMetrics | None`
:   Get environment metrics
    
     Returns metrics for an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        EnvironmentMetrics

`sync_detailed(environment_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.environment_metrics.EnvironmentMetrics]`
:   Get environment metrics
    
     Returns metrics for an environment by name.
    
    Args:
        environment_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[EnvironmentMetrics]