go: Crowdsourcing.abi Crowdsourcing.bin accounts.json
	abigen --bin=$(shell pwd)/build/Crowdsourcing.bin --abi=$(shell pwd)/build/Crowdsourcing.abi --pkg=crowdsourcing -out=$(shell pwd)/pkg/crowdsourcing/crowdsourcing.go
	(ganache-cli -m "enable laugh there club flower source fog attract giraffe trend light rival" -a 60 --account_keys_path=$(shell pwd)/pkg/crowdsourcing/plantform/accounts.json &)

Crowdsourcing.abi Crowdsourcing.bin:
	docker run --rm -v $(shell pwd):$(shell pwd) ethereum/solc:0.6.3 --overwrite @openzeppelin=$(shell pwd)/solidity/@openzeppelin -o $(shell pwd)/build --abi --bin $(shell pwd)/solidity/crowdsourcing.sol # Get abi file

accounts.json:
	touch $(shell pwd)/pkg/crowdsourcing/plantform/accounts.json	

clean:
	rm -f $(shell pwd)/build/*
	rm $(shell pwd)/pkg/crowdsourcing/plantform/accounts.json
