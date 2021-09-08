go: Platform accounts.json
	abigen --bin=$(shell pwd)/build/Platform.bin --abi=$(shell pwd)/build/Platform.abi --pkg=cplatform -out=$(shell pwd)/pkg/crowdsourcing/smartcontract/cplatform/platform.go
	abigen --bin=$(shell pwd)/build/Task.bin --abi=$(shell pwd)/build/Task.abi --pkg=ctask -out=$(shell pwd)/pkg/crowdsourcing/smartcontract/ctask/task.go
	(ganache-cli -m "enable laugh there club flower source fog attract giraffe trend light rival" -a 60 --account_keys_path=$(shell pwd)/accounts.json --verbose &)

# Task.abi Task.bin:
#	docker run --rm -v $(shell pwd):$(shell pwd) ethereum/solc:0.6.3 -o $(shell pwd)/build --abi --bin $(shell pwd)/solidity/framework/contracts/Task.sol # Get abi file

Platform:
	docker run --rm -v $(shell pwd):$(shell pwd) ethereum/solc:0.6.3 -o $(shell pwd)/build --abi --bin $(shell pwd)/solidity/framework/contracts/Platform.sol

accounts.json:
	touch $(shell pwd)/accounts.json

clean:
	rm -f $(shell pwd)/build/*
	rm $(shell pwd)/accounts.json
	rm -rf $(shell pwd)/pkg/crowdsourcing/smartcontract/task
	pkill node
