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
