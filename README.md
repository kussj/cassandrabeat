# Cassandrabeat

Welcome to Cassandrabeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/goomzee`

## Getting Started with Cassandrabeat

### Requirements

* [Golang](https://golang.org/dl/) 1.6.2


### Init Project

To get running with Cassandrabeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Cassandrabeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/goomzee/cassandrabeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


### Build

To build the binary for Cassandrabeat run the command below. This will generate a binary
in the same directory with the name cassandrabeat.

```
make
```


### Run

Before running Cassandrabeat, execute the following command:

```
cp $GOPATH/src/github.com/goomzee/cassandrabeat/beater/awkscript.sh $GOPATH/bin
```

To run Cassandrabeat with debugging output enabled, run:

```
./cassandrabeat -c cassandrabeat.yml -e -d "*"
```


### Exported fields

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


### Test

From $GOPATH/src/github.com/goomzee/cassandrabeat:

1. Prepare and build python environment
   ```
   make python-env
   ```

2. Activate python test environment
   ```
   source build/python-env/bin/activate
   ```

3. Build test-beat. Creates a `cassandrabeat.test` binary.
   ```
   make buildbeat.test
   ```

4. Go to tests/system
   ```
   cd tests/system
   ```

5. Run nosetests (`-x` = stop on first failure, `-v` = verbose)
   ```
   nosetests --with-timer -v -x test_stats.py
   ```

6. Deactivate python environment
   ```
   deactivate
   ```


### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/cassandrabeat.template.json and etc/cassandrabeat.asciidoc

```
make update
```


### Cleanup

To clean  Cassandrabeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Cassandrabeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/goomzee
cd ${GOPATH}/github.com/goomzee
git clone https://github.com/goomzee/cassandrabeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
