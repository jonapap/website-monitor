mkdir -p bin
for CMD in cmd/*
do
	echo "Building $CMD"
	go build -o bin/ ./$CMD 
done
