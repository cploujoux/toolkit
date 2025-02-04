Module beamlit.api.accounts.list_accounts
=========================================

Functions
---------

`asyncio(*, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.account.Account] | None`
:   List accounts
    
     Returns a list of all accounts.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['Account']

`asyncio_detailed(*, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.account.Account]]`
:   List accounts
    
     Returns a list of all accounts.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['Account']]

`sync(*, client: beamlit.client.AuthenticatedClient) ‑> list[beamlit.models.account.Account] | None`
:   List accounts
    
     Returns a list of all accounts.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        list['Account']

`sync_detailed(*, client: beamlit.client.AuthenticatedClient) ‑> beamlit.types.Response[list[beamlit.models.account.Account]]`
:   List accounts
    
     Returns a list of all accounts.
    
    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.
    
    Returns:
        Response[list['Account']]