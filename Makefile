go: Task.abi Task.bin accounts.json
	abigen --bin=$(shell pwd)/build/Task.bin --abi=$(shell pwd)/build/Task.abi --pkg=crowdsourcing -out=$(shell pwd)/pkg/crowdsourcing/taskContract.go
	(ganache-cli -m "enable laugh there club flower source fog attract giraffe trend light rival" -a 60 --account_keys_path=$(shell pwd)/pkg/crowdsourcing/platform/accounts.json &)

Task.abi Task.bin:
	docker run --rm -v $(shell pwd):$(shell pwd) ethereum/solc:0.6.3 --overwrite @openzeppelin=$(shell pwd)/solidity/@openzeppelin -o $(shell pwd)/build --abi --bin $(shell pwd)/solidity/Task.sol # Get abi file

accounts.json:
	touch $(shell pwd)/pkg/crowdsourcing/platform/accounts.json

clean:
	pkill node
	rm -f $(shell pwd)/build/*
	rm $(shell pwd)/pkg/crowdsourcing/platform/accounts.json
