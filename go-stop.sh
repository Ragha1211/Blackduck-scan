kill -9 $(pgrep -f "make run-go")
kill -9 $(pgrep -f "bash ./go-run.sh")
kill -9 $(pgrep -f "CompileDaemon --command=./order")
kill -9 $(pgrep -f "./order")