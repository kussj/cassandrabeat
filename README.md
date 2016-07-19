# Cassandrabeat

Cassandrabeat is a [Beat](https://www.elastic.co/products/beats) 
for monitoring Cassandra database nodes and lag. This beat uses 
Cassandra's `nodetool cfstats` utility to retrieve table read and 
write latencies.

## Installation

Requires go version go1.6.2+

```bash
cd $GOPATH/src
go get github.com/goomzee/cassandrabeat
cd github.com/goomzee/cassandrabeat
cp beater/awkscript.sh $GOPATH/bin
go install
```

## Run

Note: you must be in the same directory as `awkscript.sh`

```bash
cd $GOPATH/bin
./cassandrabeat -c ../src/github.com/goomzee/cassandrabeat.yml
```

## Exported fields

There is only one type of document exported:
- `type: stats` for table read and write latencies

Table statistics:

<pre>
{
    "beat": {
        "name": "cassandrabeat",
        "host": "localhost"
    },
    "@timestamp": "2016-07-12T16:28:19.670Z",
    "@version": "1",
    "type": "stats",
    "count": 1,
    "table_name": "system.local", 
    "write_latency": 1.074, 
    "read_latency": 3.343
}
</pre>

## Testing

Note: currently, testing is being done from the elastic/beats
directory.

Copy cassandrabeat to elastic/beats directory:
```bash
cp -r $GOPATH/src/github.com/goomzee/cassandrabeat $GOPATH/src/github.com/elastic/beats
```

From the elastic/beats/cassandrabeat directory:

1. Prepare and build python environment
```bash
cd /path/to/elastic/beats/cassandrabeat
make python-env
```

2. Activate python test environment
```bash
source build/python-env/bin/activate
```

3. Build test-beat. Creates a `cassandrabeat.test` binary.
```bash
make buildbeat.test
```

4. Go to tests/system
```bash
cd tests/system
```

5. Run nosetests (`-x` = stop on first failure, `-v` = verbose)
```bash
nosetests --with-timer -v -x test_stats.py
```

### Known issues
The integration test is not functional at the moment. There are no
errors in the log file `cassandrabeat/build/system-tests/last_run/cassandrabeat.log`
and the .yml file `cassandrabeat/build/system-tests/last_run/cassandrabeat.yml`
is being generated correctly, but there is no output file being
created and written to, which is causing the test to fail.

The reasoning behind running the tests from the elastic/beats
directory is that there are errors when attempting to build the
python environment from the goomzee directory, and this is a way
to step around the issue.
