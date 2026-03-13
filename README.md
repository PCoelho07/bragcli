# brag

A command-line tool to keep your brag doc up to date — without breaking your flow.

Track your wins, shipped features, and impact as they happen. When review season comes, you'll have everything ready.

---

## Installation

```sh
make install
```

Installs the `brag` binary to `/usr/local/bin`.

---

## Usage

### Initialize a bragdoc

```sh
brag init -n work
```

Creates a new bragdoc named `work` at `~/.brag/work/`.

### Add an entry

```sh
brag create -n work -t "shipped payment redesign" -d "reduced checkout drop-off by 18%"
```

Each entry is saved as a timestamped JSON file inside `~/.brag/<name>/`.

---

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `init`  | `i`   | Initialize a new bragdoc |
| `create` | `c`  | Add a new entry to a bragdoc |

### Flags

**`init`**

| Flag | Short | Required | Description |
|------|-------|----------|-------------|
| `--name` | `-n` | yes | Name of the bragdoc |

**`create`**

| Flag | Short | Required | Description |
|------|-------|----------|-------------|
| `--name` | `-n` | yes | Name of the bragdoc |
| `--title` | `-t` | yes | Title of the entry |
| `--description` | `-d` | no | Details about the entry |

---

## Development

```sh
make test   # run tests
make fmt    # format code
make build  # build binary to dist/
```

---

## Storage

Entries are stored locally as JSON files:

```
~/.brag/
└── work/
    ├── 20260313T142500.json
    └── 20260314T091200.json
```

Each file contains the title, description, and timestamp of the entry.
