Module beamlit.api.default.create_account
=========================================

Functions
---------

`asyncio(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.models.account.Account | None`
:   Create account
    
     Creates an account.
    
    Args:
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Account

`asyncio_detailed(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.types.Response[beamlit.models.account.Account]`
:   Create account
    
     Creates an account.
    
    Args:
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Account]

`sync(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.models.account.Account | None`
:   Create account
    
     Creates an account.
    
    Args:
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Account

`sync_detailed(*, client: beamlit.client.AuthenticatedClient, body: beamlit.models.account.Account) ‑> beamlit.types.Response[beamlit.models.account.Account]`
:   Create account
    
     Creates an account.
    
    Args:
        body (Account): Account
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Account]