Module beamlit.api.agents.list_agent_history
============================================

Functions
---------

`asyncio(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.agent_history.AgentHistory] | None`
:   Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['AgentHistory']

`asyncio_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.agent_history.AgentHistory]]`
:   Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['AgentHistory']]

`sync(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.agent_history.AgentHistory] | None`
:   Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['AgentHistory']

`sync_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.agent_history.AgentHistory]]`
:   Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['AgentHistory']]