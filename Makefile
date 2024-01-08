

smartcontracts:
	cd web3/ && rm -rf ../backend/build
	cd web3/ && solc --optimize --abi ./contracts/ICredits.sol -o ../backend/build
	cd web3/ && solc --optimize --allow-paths ./web3 --bin ./contracts/ICredits.sol -o ../backend/build
	cd web3/ && abigen --abi=../backend/build/ICredits.abi --bin=../backend/build/ICredits.bin --pkg=credits --out=../backend/smartcontractinterfaces/ICredits.go

start backend:
	cd backend && go run ./cmd/api/

