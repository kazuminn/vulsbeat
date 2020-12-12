# vulsbeat

Welcome to vulsbeat.

This software allows you Vulnerability scan results of [vuls](https://github.com/future-architect/vuls) can be imported to Elastic Stack.
You can do various things with elasticsearch. For example, analyze or detect complex threats with SIEM.


Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/kazuminn/vulsbeat`

## Getting Started with {Beat}

### Requirements

* [Golang](https://golang.org/dl/) 1.7 later
* [vuls](https://github.com/future-architect/vuls) v0.13.9 later
* [mage](https://github.com/magefile/mage)

### Config

change path: in vulsbeat.yml.

```
path: "/path/to/results/"
```

### Init Project
To get running with vulsbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push {Beat} in the git repository, run the following commands:

```
git remote set-url origin https://github.com/kazuminn/vulsbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for vulsbeat run the command below. This will generate a binary
in the same directory with the name vulsbeat.

```
make
```


### Run

To run vulsbeat with debugging output enabled, run:

```
./vulsbeat -c vulsbeat.yml -e -d "*"
```


### Test

To test vulsbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  vulsbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone vulsbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/kazuminn/vulsbeat
git clone https://github.com/kazuminn/vulsbeat ${GOPATH}/src/github.com/kazuminn/vulsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
