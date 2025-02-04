Module beamlit.api.history.list_agents_history
==============================================

Functions
---------

`asyncio(*, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.agent_history.AgentHistory] | None`
:   Get all history for all agents
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['AgentHistory']

`asyncio_detailed(*, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.agent_history.AgentHistory]]`
:   Get all history for all agents
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['AgentHistory']]

`sync(*, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.agent_history.AgentHistory] | None`
:   Get all history for all agents
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['AgentHistory']

`sync_detailed(*, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.agent_history.AgentHistory]]`
:   Get all history for all agents
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['AgentHistory']]