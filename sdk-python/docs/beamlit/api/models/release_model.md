Module beamlit.api.models.release_model
=======================================

Functions
---------

`asyncio(model_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.model_release.ModelRelease | None`
:   Release model from an environment
    
     Make a release for a model from an environment to another.
    
    Args:
        model_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ModelRelease

`asyncio_detailed(model_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.model_release.ModelRelease]`
:   Release model from an environment
    
     Make a release for a model from an environment to another.
    
    Args:
        model_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ModelRelease]

`sync(model_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.model_release.ModelRelease | None`
:   Release model from an environment
    
     Make a release for a model from an environment to another.
    
    Args:
        model_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        ModelRelease

`sync_detailed(model_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.model_release.ModelRelease]`
:   Release model from an environment
    
     Make a release for a model from an environment to another.
    
    Args:
        model_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[ModelRelease]