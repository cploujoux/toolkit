Module beamlit.api.accounts.get_account
=======================================

Functions
---------

`asyncio(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.account.Account | None`
:   Get account by name
    
     Returns an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Account

`asyncio_detailed(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.account.Account]`
:   Get account by name
    
     Returns an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Account]

`sync(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.account.Account | None`
:   Get account by name
    
     Returns an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Account

`sync_detailed(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.account.Account]`
:   Get account by name
    
     Returns an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Account]