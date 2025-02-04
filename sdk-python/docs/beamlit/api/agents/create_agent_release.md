Module beamlit.api.agents.create_agent_release
==============================================

Functions
---------

`asyncio(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.agent_release.AgentRelease | None`
:   Create release for a agent from an environment
    
    Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        AgentRelease

`asyncio_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.agent_release.AgentRelease]`
:   Create release for a agent from an environment
    
    Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[AgentRelease]

`sync(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.agent_release.AgentRelease | None`
:   Create release for a agent from an environment
    
    Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        AgentRelease

`sync_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.agent_release.AgentRelease]`
:   Create release for a agent from an environment
    
    Args:
        agent_name (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[AgentRelease]