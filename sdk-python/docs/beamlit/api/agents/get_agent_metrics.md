Module beamlit.api.agents.get_agent_metrics
===========================================

Functions
---------

`asyncio(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics | None`
:   Get agent metrics
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ResourceEnvironmentMetrics

`asyncio_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics]`
:   Get agent metrics
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ResourceEnvironmentMetrics]

`sync(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics | None`
:   Get agent metrics
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ResourceEnvironmentMetrics

`sync_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics]`
:   Get agent metrics
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ResourceEnvironmentMetrics]