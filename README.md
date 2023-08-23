# KVDB
Key-Value Database implemented in Golang

Design and develop a simple persistent key-value database in Golang. Your project should
produce a `kvdb` binary that at least implements the following commands.
- kvdb set {key} {value} : Associates `key` with `value`.
- kvdb get {key} : Fetches the value associated with `key`.
- kvdb del {key} : Removes `key` from the database.
- kvdb ts {key} : Returns the timestamp when `key` was first set and when it was last set. Expected timestamp format is
"YYYY-MM-DD HH:MM:SS.SSS".

Your implementation should store data on disk, so that kvdb get can retrieve a value that
was previously set with kvdb set.

### Rules:
- Your program should be written in Golang.
- Assume that your database can be accessed by multiple processes concurrently.
- Keep performance and storage efficiency in mind.
- You are expected to not spend much more than 4 hours on this project.
- You are allowed to use libraries, but don't just write a wrapper around an existing
database.
- Please provide a short write-up of the limitations of your program. (20 lines max).
- Hand in a tarball containing your project folder and the write-up.

### Write-up:

# Datastructure
The datastructure used is a map of maps. The outer map is a map of keys to a map of timestamps to values. The inner map is a map of timestamps to values. The timestamp is a string in the format "YYYY-MM-DD HH:MM:SS.SSS". The value is a string.

# Limitations   
- The database is not thread safe. If multiple processes try to access the database at the same time, the database may become corrupted.

# Github CI
The project is built and tested on Github CI. 

# How to run
- Clone the repository
- Run `go build` to build the binary
- Run `./kvdb` to see the usage instructions
- Run `./kvdb set key value` to set a key value pair
- Run `./kvdb get key` to get the value for a key
- Run `./kvdb del key` to delete a key
- Run `./kvdb ts key` to get the timestamp for a key