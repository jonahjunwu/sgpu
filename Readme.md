# spd gpu
gpuspd.go (main code)
rjson.go (read jsonfile)
input.json (contains whisper files as a list )

It contains concurrency and pipeline.
pipeline use p1python.py p2python.py p3python.py

## Execute command
go run gpuspd.go rjson.go
