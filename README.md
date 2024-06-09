# inspix-hailstorm 

All-in-one tool suite aims to analyze/decrypt resources and the master database of Link Like Love Live.


## Usage

You can choose to build from source on your own or utilize our pre-built docker image to use.

### Build

Simply run the following code in your terminal to build an executable:

```bash
go build .
```

Once build is completed, run `./hailstorm -h` to get instructions. Here are some hints:

- (without options): Download and decrypt all new assets and database since the last run. Note this operation can take from minutes to hours.
- `--analyze`: For developers to analyze the structure of the database.
- `--dbonly`: Ignore asset files, only download and decrypt master database. This operation can take several seconds.


### Docker

Get the docker image from https://github.com/vertesan/inspix-hailstorm/pkgs/container/inspix-hailstorm and run `run_docker.sh`.  
By default the container will be executed in `dbonly` mode, you can overwrite this action by handing over `--entrypoint` to docker CLI if you need to change the default mode.


## License

AGPL-3.0 license
