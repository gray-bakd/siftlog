# siftlog

A fast CLI tool for filtering and colorizing structured JSON logs from stdin with field-based queries.

---

## Installation

```bash
go install github.com/yourusername/siftlog@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/siftlog.git && cd siftlog && go build -o siftlog .
```

---

## Usage

Pipe JSON logs into `siftlog` and filter by field values:

```bash
# Show only error-level logs
cat app.log | siftlog --filter 'level=error'

# Filter by multiple fields
kubectl logs my-pod | siftlog --filter 'level=error' --filter 'service=auth'

# Colorized output with pretty printing
tail -f app.log | siftlog --pretty --filter 'status=500'
```

### Flags

| Flag | Description |
|------|-------------|
| `--filter`, `-f` | Field query in `key=value` format (repeatable) |
| `--pretty`, `-p` | Pretty-print and colorize output |
| `--fields` | Comma-separated list of fields to display |
| `--no-color` | Disable colorized output |

### Example Output

```
[ERROR] 2024-01-15T10:32:01Z  service=auth  msg="invalid token"  user_id=42
[ERROR] 2024-01-15T10:33:18Z  service=auth  msg="session expired"  user_id=99
```

---

## Requirements

- Go 1.21+

---

## License

MIT © 2024 yourusername