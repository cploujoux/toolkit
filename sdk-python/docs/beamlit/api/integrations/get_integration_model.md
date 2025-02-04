Module beamlit.api.integrations.get_integration_model
=====================================================

Functions
---------

`asyncio_detailed(integration_name: str, model_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[typing.Any]`
:   Get integration model
    
     Returns a model for an integration by ID.
    
    Args:
        integration_name (str):
        model_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Any]

`sync_detailed(integration_name: str, model_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[typing.Any]`
:   Get integration model
    
     Returns a model for an integration by ID.
    
    Args:
        integration_name (str):
        model_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Any]