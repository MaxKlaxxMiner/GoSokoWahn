module sokoServer

go 1.22

require (
	github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.11.1
)

require golang.org/x/exp v0.0.0-20240823005443-9b4947da3948 // indirect

replace github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib => ../sokoLib
