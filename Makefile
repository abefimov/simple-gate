
GO=go
TRUFFLE=./node_modules/truffle/build/cli.bundled.js
TESTRPC=./node_modules/ethereumjs-testrpc/build/cli.node.js
SOLIUM=./node_modules/solium/bin/solium.js
GAS_LIMIT=15000000

.PHONY: all test clean node_modules coverage

all: default lint compile migrate test

default: compile

compile:
	@echo "+ $@"
	@${TRUFFLE} compile

node_modules:
	@echo "+ $@"
	@npm install

testrpc:
	@echo "+ $@"
	@nohup ${TESTRPC} -l ${GAS_LIMIT} &

test:
	@echo "+ $@"
	@scripts/test.sh

ci: node_modules lint fmt compile test

lint:
	@echo "+ $@"
	@${SOLIUM} --dir ./contracts

migrate:
	${TRUFFLE} migrate

coverage:
	scripts/coverage.sh

deploy:
	${TRUFFLE} migrate --network private

fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v ^vendor/ | tee /dev/stderr)" || \
		(echo >&2 "+ please format Go code with 'gofmt -s'" && false)
