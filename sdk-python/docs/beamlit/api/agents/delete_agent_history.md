Module beamlit.api.agents.delete_agent_history
==============================================

Functions
---------

`asyncio(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.agent_history.AgentHistory | None`
:   Args:
        agent_name (str):
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        AgentHistory

`asyncio_detailed(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.agent_history.AgentHistory]`
:   Args:
        agent_name (str):
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[AgentHistory]

`sync(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.agent_history.AgentHistory | None`
:   Args:
        agent_name (str):
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        AgentHistory

`sync_detailed(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.agent_history.AgentHistory]`
:   Args:
        agent_name (str):
        request_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[AgentHistory]