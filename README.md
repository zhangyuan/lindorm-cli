# lindorm-cli

Unoffical Lindorm CLI, written in Go and with better developer experience and less bugs.

## Usage

### Download the binary

Go to the [Releases](https://github.com/zhangyuan/lindorm-cli/releases) page to download the binary for specific platform. Optionally, rename the binary name to `lindorm-cli`.

### Setup the credentials

Set the database credentials via environment variables or `.env` file. e.g.

```.env
ENDPOINT=http://...
USERNAME=username
PASSWORD=password
DATABASE=database
```

### Run interactive command line interface

```bash
lindorm-cli
```

Note that current the SQL statement must be provided as a single line without line break. Type `exit` to exit the interface.

### Run single statement with subcommand `query`

#### Provide the SQL statement via argument

Pass the SQL statement with `-s` option:

```bash
lindorm-cli query -s "select 1"
```

Note that when the statement contains the backtick character(`` ` ``), it's tricky to provide the correct argument in Linux termimal. So use `-f` option below is a better option.

#### Provide the SQL statement via file

Pass the file path containing a SQL statement via `-f` option:

```bash
lindorm-cli query -f sql.sql
```


