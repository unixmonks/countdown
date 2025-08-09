# countdown

A simple CLI countdown timer written in Go.

## Usage

```bash
countdown duration
```

## Examples

```bash
countdown 30s        # 30 seconds
countdown 5m         # 5 minutes  
countdown 1h30m      # 1 hour 30 minutes
countdown 1m30s      # 1 minute 30 seconds
```

## Installation

### From source

```bash
git clone https://github.com/username/countdown.git
cd countdown
make install
```

### User installation (no sudo required)

```bash
make install-user
```

This installs to `~/.local/bin` - make sure this directory is in your `$PATH`.

## Building

```bash
make build
```

For an optimized release build:

```bash
make release
```

## License

MIT License - see [LICENSE](LICENSE) file for details.