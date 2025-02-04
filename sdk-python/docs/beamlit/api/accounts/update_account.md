Module beamlit.api.accounts.update_account
==========================================

Functions
---------

`asyncio(account_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.models.account.Account | None`
:   Update account
    
     Updates an account by name.
    
    Args:
        account_id (str):
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Account

`asyncio_detailed(account_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.types.Response[beamlit.models.account.Account]`
:   Update account
    
     Updates an account by name.
    
    Args:
        account_id (str):
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Account]

`sync(account_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.models.account.Account | None`
:   Update account
    
     Updates an account by name.
    
    Args:
        account_id (str):
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Account

`sync_detailed(account_id: str, *, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.types.Response[beamlit.models.account.Account]`
:   Update account
    
     Updates an account by name.
    
    Args:
        account_id (str):
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Account]