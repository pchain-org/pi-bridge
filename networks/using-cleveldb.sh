#!/bin/bash

make install PBRIDGE_BUILD_OPTIONS="cleveldb"

pibridged init "t6" --home ./t6 --chain-id t6

pibridged unsafe-reset-all --home ./t6

mkdir -p ./t6/data/snapshots/metadata.db

pibridged keys add validator --keyring-backend test --home ./t6

pibridged add-genesis-account $(pibridged keys show validator -a --keyring-backend test --home ./t6) 100000000stake --keyring-backend test --home ./t6

pibridged gentx validator 100000000stake --keyring-backend test --home ./t6 --chain-id t6

pibridged collect-gentxs --home ./t6

pibridged start --db_backend cleveldb --home ./t6
