Module beamlit.api.accounts.delete_account
==========================================

Functions
---------

`asyncio(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.account.Account | Any | None`
:   Delete account
    
     Deletes an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Union[Account, Any]

`asyncio_detailed(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.account.Account | Any]`
:   Delete account
    
     Deletes an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Union[Account, Any]]

`sync(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.models.account.Account | Any | None`
:   Delete account
    
     Deletes an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Union[Account, Any]

`sync_detailed(account_id: str, *, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[beamlit.models.account.Account | Any]`
:   Delete account
    
     Deletes an account by name.
    
    Args:
        account_id (str):
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[Union[Account, Any]]