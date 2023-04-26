# Transfer Flow

## Preconditions

We deployed Bridge contracts to each supported chain, added supported tokens with chain id to address pairings and provided liquidity for these tokens on each chain.

We deployed System contracts to our blockchain to coordinate validators actions and collect signatures for transfer approvals.

## Transfer

1. User deposits ERC-20 tokens to Bridge and specifies destination chain and receiver
    - tokens are deposited to LiqudityPool contract
    - event of deposit is being emitted, containing amount, destination chain and receiver
2. Validators see event of deposit and create distributed signature on System contract
3. Once distributed signature is created and it is valid, one of validators initiate call of transfer method from Bridge contract on destination chain
    - bridge contract on destination chain checks if signature is correct and corresponds to validators
    - checks if bridge liquidity is enough for transfer
    - deducts transfer fee and deposits it to liquidity pool
    - makes final transfer to reveiver
