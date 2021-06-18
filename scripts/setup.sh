set -ex
# Convert Solity smart contract to go library
docker run --rm -v $PWD:$PWD ethereum/solc:0.6.3 -o $PWD/build --abi --bin $PWD/solidity/*.sol # Get abi file
abigen --bin=$PWD/build/*.bin --abi=$PWD/build/*.abi --pkg=$PWD/crowdsourcing/lib -out=crowdsourcing.go
# Generate test Ethereum address
ganache-cli -m "enable laugh there club flower source fog attract giraffe trend light rival" -a 60 --account_keys_path=$PWD/pkg/crowdsourcing/platform/accounts.json
