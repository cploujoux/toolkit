Module beamlit.api.functions.get_function_metrics
=================================================

Functions
---------

`asyncio(function_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics | None`
:   Get function metrics
    
    Args:
        function_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ResourceEnvironmentMetrics

`asyncio_detailed(function_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics]`
:   Get function metrics
    
    Args:
        function_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ResourceEnvironmentMetrics]

`sync(function_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics | None`
:   Get function metrics
    
    Args:
        function_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ResourceEnvironmentMetrics

`sync_detailed(function_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.resource_environment_metrics.ResourceEnvironmentMetrics]`
:   Get function metrics
    
    Args:
        function_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ResourceEnvironmentMetrics]