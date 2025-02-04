Module beamlit.api.agents.put_agent_history
===========================================

Functions
---------

`asyncio(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.agent_history.AgentHistory) ‑> beamlit.models.agent_history.AgentHistory | None`
:   Update agent's history by request ID
    
    Args:
        agent_name (str):
        request_id (str):
        body (AgentHistory): Agent deployment history
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        AgentHistory

`asyncio_detailed(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.agent_history.AgentHistory) ‑> beamlit.types.Response[beamlit.models.agent_history.AgentHistory]`
:   Update agent's history by request ID
    
    Args:
        agent_name (str):
        request_id (str):
        body (AgentHistory): Agent deployment history
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[AgentHistory]

`sync(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.agent_history.AgentHistory) ‑> beamlit.models.agent_history.AgentHistory | None`
:   Update agent's history by request ID
    
    Args:
        agent_name (str):
        request_id (str):
        body (AgentHistory): Agent deployment history
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        AgentHistory

`sync_detailed(agent_name: str, request_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.agent_history.AgentHistory) ‑> beamlit.types.Response[beamlit.models.agent_history.AgentHistory]`
:   Update agent's history by request ID
    
    Args:
        agent_name (str):
        request_id (str):
        body (AgentHistory): Agent deployment history
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[AgentHistory]