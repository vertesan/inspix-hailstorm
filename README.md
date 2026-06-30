# inspix-hailstorm 

An all-in-one tool suite aims to analyze/decrypt resources and the master database of Link Like Love Live.


## LLLL EoS

With the officially shutdown of Link Like Love Live on June 30, 2026, this project no longer works.


## Building

Simply run the following command in your terminal to build an executable:

```bash
go build .
```

Once build completed, run `./hailstorm -h` to get instructions. Here are some hints:

- (without options): Download and decrypt all new assets and database since the last run. Note this operation can take from minutes to hours.
- `--analyze`: For developers to analyze the structure of the database.
- `--dbonly`: Ignore asset files, only download and decrypt master database. This operation can take several seconds.


## License

AGPL-3.0 license
