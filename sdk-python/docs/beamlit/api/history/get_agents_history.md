Module beamlit.api.history.get_agents_history
=============================================

Functions
---------

`asyncio(request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.agent_history.AgentHistory] | None`
:   Get all history for a specific request ID from all agents
    
    Args:
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['AgentHistory']

`asyncio_detailed(request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.agent_history.AgentHistory]]`
:   Get all history for a specific request ID from all agents
    
    Args:
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['AgentHistory']]

`sync(request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.agent_history.AgentHistory] | None`
:   Get all history for a specific request ID from all agents
    
    Args:
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['AgentHistory']

`sync_detailed(request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.agent_history.AgentHistory]]`
:   Get all history for a specific request ID from all agents
    
    Args:
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['AgentHistory']]