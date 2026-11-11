# freqr

**freqr** (frequrency ranger) is a command-line tool for programming Uniden radio scanners. It allows you to export, import, and list programmed frequencies.

## Requirements
- Go 1.25 or later

## Usage

Run `freqr -h` to display help and a list of all available commands.

## Most useful commands
1. Export all programmed channels to a CSV file: `freqr export-channels -p filename.csv`
2. Import channels from a CSV file (overwrites existing radio channels): `freqr import-channels -p filename.csv`
3. List all programmed channels in a table: `freqr list-channels`

## Configuration

The application configuration is stored in `config/config.yaml`.

Available configuration parameters:

| Parameter | Description |
|---|---|
| `port_name` | Name of the USB serial port connected to the radio. |
| `baud_rate` | Serial communication speed, expressed in bits per second (baud). |
| `max_retries` | Maximum number of connection attempts to the radio. |
| `read_sleep_time` | Delay between consecutive channel read operations. |
| `write_sleep_time` | Delay between consecutive channel write operations. |
| `total_channels` | Maximum number of channels supported by the radio. |
| `break_when_err` | Whether to abort the entire operation if an error occurs while reading frequencies from the radio. |

## CSV Format
CSV files and terminal tables use the same column format:

`Index,Name,Frequency,Mode,Delay,Lockout,Priority`

- Index, Frequency: Integer values
- Mode: AM or FM
- Delay, Lockout, Priority: Boolean values (true / false)

An example CSV file is available at examples/frequencies.csv.