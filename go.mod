module gogocache

go 1.15

require (
	gocache v0.0.0
	lru v0.0.0
)

replace (
	gocache => ./gocache
	lru => ./gocache/lru
)
