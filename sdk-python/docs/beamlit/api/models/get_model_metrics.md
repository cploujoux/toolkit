Module beamlit.api.models.get_model_metrics
===========================================

Functions
---------

`asyncio(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics | None`
:   Get model metrics
    
     Returns metrics for a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ResourceEnvironmentMetrics

`asyncio_detailed(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics]`
:   Get model metrics
    
     Returns metrics for a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ResourceEnvironmentMetrics]

`sync(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics | None`
:   Get model metrics
    
     Returns metrics for a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ResourceEnvironmentMetrics

`sync_detailed(model_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics]`
:   Get model metrics
    
     Returns metrics for a model by name.
    
    Args:
        model_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ResourceEnvironmentMetrics]