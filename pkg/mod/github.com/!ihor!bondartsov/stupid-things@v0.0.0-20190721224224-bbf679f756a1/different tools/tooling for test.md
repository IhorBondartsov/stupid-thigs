**This tutorial I create using Alex Edwards article: 'An Overview of Go's Tooling'.**


    $ go test .          # Run all tests in the current directory
    $ go test ./...      # Run all tests in the current directory and sub-directories
    $ go test ./foo/bar  # Run all tests in the ./foo/bar directory 
    
    $ go test -race ./...
    
    $ go test -count=1 ./...    # Bypass the test cache when running tests
    $ go clean -testcache       # Delete all cached test results
    
    $ go test -v -run=^TestFooBar$ .          # Run the test with the exact name TestFooBar
    $ go test -v -run=^TestFoo .              # Run tests whose names start with TestFoo
    $ go test -v -run=^TestFooBar$/^Baz$ .    # Run the Baz subtest of the TestFooBar test only
    
    $ go test -short ./...      # Skip long running tests
    $ go test -failfast ./...   # Don't run further tests after a failure.

    $ go test -cover ./...
    
    $ go test -coverprofile=/tmp/profile.out ./...
    $ go tool cover -html=/tmp/profile.out
    
`-covermode=count` flag to make the coverage profile record the exact number of times that each statement is executed during the tests.

    $ go test -covermode=count -coverprofile=/tmp/profile.out ./...
    $ go tool cover -html=/tmp/profile.out
    
  **Stress Testing**
  
    $ go test -run=^TestFooBar$ -count=500 .
    
 **Diagnosing Problems and Making Optimizations**
 
    $ go test -bench=. ./...                        # Run all benchmarks and tests
    $ go test -run=^$ -bench=. ./...                # Run all benchmarks (and no tests)
    $ go test -run=^$ -bench=^BenchmarkFoo$ ./...   # Run only the BenchmarkFoo benchmark (and no tests)
    
   `-benchmem `flag, which forces memory allocation statistics to be included in the output.
   
    $ go test -bench=. -benchmem ./...
    $ go test -bench=. -benchtime=5s ./...       # Run each benchmark test for at least 5 seconds
    $ go test -bench=. -benchtime=500x /....     # Run each benchmark test for exactly 500 iterations
    $ go test -bench=. -count=3 ./...            # Repeat each benchmark test 3 times over
    $ go test -bench=. -cpu=1,4,8 ./...
    
   **Profiling and Tracing**

If you have a web application you can import the net/http/pprof package. This will register some handlers with the http.DefaultServeMux which you can then use to generate and download profiles for your running application. This post provides a good explanation and some sample code.

For other types of applications, you can profile your running application using the pprof.StartCPUProfile() and pprof.WriteHeapProfile() functions. See the runtime/pprof documentation for sample code.

Or you can generate profiles while running benchmarks or tests by using the various -***profile flags like so:

    $ go test -run=^$ -bench=^BenchmarkFoo$ -cpuprofile=/tmp/cpuprofile.out .
    $ go test -run=^$ -bench=^BenchmarkFoo$ -memprofile=/tmp/memprofile.out .
    $ go test -run=^$ -bench=^BenchmarkFoo$ -blockprofile=/tmp/blockprofile.out .
    $ go test -run=^$ -bench=^BenchmarkFoo$ -mutexprofile=/tmp/mutexprofile.out .
    
    go test -run=^$ -bench=^BenchmarkFoo$ -o=/tmp/foo.test -cpuprofile=/tmp/cpuprofile.out .
    
    $ go tool pprof -http=:5000 /tmp/cpuprofile.out  # Command to open it in a web browser
    $ go tool pprof --nodefraction=0.1 -http=:5000 /tmp/cpuprofile.out
    
    $ go test -run=^$ -bench=^BenchmarkFoo$ -trace=/tmp/trace.out .
    $ go tool trace /tmp/trace.out
    

    
   