#!/bin/bash -eu

RUN="go run"

while [ $# -gt 0 ]; do
	case "$1" in
		-v) RUN="go run -v" ;;
		*)  echo "$0: unknown option $1" >&2; exit 1;;
	esac
	shift
done

# Test examples
${RUN} ./test/artist/main.go artist.example.yml
${RUN} ./test/crew/main.go crew.example.yml
${RUN} ./test/pack/main.go pack.example.yml

# Test contributed metadata
${RUN} ./test/artist/main.go artist/*.yml
${RUN} ./test/crew/main.go crew/*.yml
${RUN} ./test/pack/main.go pack/*/*.yml
