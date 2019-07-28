module my-local-test

go 1.12

require (
	buy v0.0.0-00010101000000-000000000000
	github.com/line/line-bot-sdk-go v6.3.0+incompatible
)

replace buy => ./buy
