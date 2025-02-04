Module beamlit.api.workspaces.workspace_quotas_request
======================================================

Functions
---------

`asyncio_detailed(workspace_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[typing.Any]`
:   Request workspace quotas
    
     Send a request to update quotas for a workspace.
    
    Args:
        workspace_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Any]

`sync_detailed(workspace_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[typing.Any]`
:   Request workspace quotas
    
     Send a request to update quotas for a workspace.
    
    Args:
        workspace_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Any]