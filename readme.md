# Pi-Bridge

Pi-Bridge is new generation layer2 cross-chain bridge, based on [Cosmos-SDK](https://github.com/cosmos/cosmos-sdk) development.

## Required

[Go 1.15+](https://golang.org/dl/)

## Install

Please make sure have already installed `Go` correctly, and set environment variables : `$GOPATH`, `$GOBIN`, `$PATH`.

Put the Pi-Bridge project in the specific path，switch to `master` branch，download related dependencies, then make install:

Check if the installation is successful:

```
build/pibridged --help
```

## Run & Build

### Single Node

Initialize the genesis.json file that will help you to bootstrap the network

```
 ./pibridge init my-node --chain-id my-chain
```

Create a key to hold your validator account

```
 ./pibridge keys add test
```

Add that key into the genesis.app_state.accounts array in the genesis file

NOTE: this command lets you set the number of coins. Make sure this account has some coins

with the genesis.app_state.staking.params.bond_denom denom, the default is staking

```
./pibridge add-genesis-account test 1000000000stake,1000000000validatortoken
```

Generate the transaction that creates your validator

```
./pibridge gentx test 1000000000stake --chain-id my-chain
```

Add the generated bonding transaction to the genesis file

```
 ./pibridge collect-gentxs
```

Now its safe to start `pibridge`

```
 ./pibridge start
```
