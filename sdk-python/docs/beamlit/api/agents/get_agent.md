Module beamlit.api.agents.get_agent
===================================

Functions
---------

`asyncio(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.agent.Agent | None`
:   Get agent by name
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Agent

`asyncio_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.agent.Agent]`
:   Get agent by name
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Agent]

`sync(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.models.agent.Agent | None`
:   Get agent by name
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Agent

`sync_detailed(agent_name: str, *, client: beamlit.client.AuthenticatedClient, environment: beamlit.types.Unset | str = <beamlit.types.Unset object>) ‑> beamlit.types.Response[beamlit.models.agent.Agent]`
:   Get agent by name
    
    Args:
        agent_name (str):
        environment (Union[Unset, str]):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Agent]