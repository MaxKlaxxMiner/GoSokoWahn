module sokoServer

go 1.21

require (
	github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib v0.0.0-20240904174251-4f8a4394a8e3
	github.com/rs/cors v1.11.1
)

require golang.org/x/exp v0.0.0-20240904232852-e7e105dedf7e // indirect

replace github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib => ../sokoLib
