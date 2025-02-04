Module beamlit.api.environments.list_environments
=================================================

Functions
---------

`asyncio(*, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.environment.Environment] | None`
:   List environments
    
     Returns a list of all environments in the workspace.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['Environment']

`asyncio_detailed(*, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.environment.Environment]]`
:   List environments
    
     Returns a list of all environments in the workspace.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['Environment']]

`sync(*, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.environment.Environment] | None`
:   List environments
    
     Returns a list of all environments in the workspace.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['Environment']

`sync_detailed(*, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.environment.Environment]]`
:   List environments
    
     Returns a list of all environments in the workspace.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['Environment']]