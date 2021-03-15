<p align="center">
  <a href="https://github.com/yogamuris/balbalancli">
    <img src="assets/img.png" alt="Logo" width="500" height="250">
  </a>

  <h3 align="center">BalbalanCLI</h3>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="##setup-the-api-key">Setup the API Key</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
    <a href="#usage">Usage</a>
        <ul>
            <li><a href="#available-league">Available League</a></li>
            <li><a href="#available-command">Available Command</a></li>
        </ul>
    </li>
    <li><a href="#license">License</a></li>
    <li><a href="#go-package">Go Package</a></li>
  </ol>
</details>


## About The Project

BalbalanCLI is a simple CLI application that can show the latest football score and league standing.


## Getting Started

Ensure that Go is installed in your computer.

### Setup the API Key

To get the API Key, please register to [football-data.org](https://www.football-data.org/). After getting the API Key or Token, then add your API Key through ```balbalancli config --token YOUR_API_KEY``` command.


### Installation

Run
   ```shell
   go get -u github.com/yogamuris/balbalancli
   ```

## Usage
### Available league
| Code| League|
| :---: | :---: |
| PL | English Premier League |
| PD | Primera Division La Liga |
| SA | Serie A |
| BL1 | Bundesliga |
| FL1 | Ligue One |
| DED| Eredivisie |
| PPL | Primeira Liga |
| CL | UEFA Champions League |

### Available Command
 ```shell
balbalancli [command]

Available Commands:
  config      Setting the BalbalanCLI Configuration.
  help        Help about any command
  list        Get list of available league.
  score       Show the latest score.
  standing    Show standing of the League.

Flags:
  -h, --help      help for balbalancli
  -v, --version   version for balbalancli
 ```

#### `config` command
Setup the API Key / Token
```shell
balbalancli config [flags]

Examples:

        balbalan config -t YOUR_TOKEN_VALUE
        balbalan config --token YOUR_TOKEN_VALUE

Flags:
  -h, --help           help for config
  -t, --token string   set API Token
```
#### `help` command
Get help for specific command
```shell
   balbalancli help <command>
```
#### `list` command
Get list of available league
```shell
balbalancli list [flags]

Flags:
  -h, --help   help for list
```
#### `score` command
Get the latest football score
```shell
balbalancli score [flags]

Examples:

        balbalan score -l PL
        balbalan score --league PL

Flags:
  -h, --help            help for score
  -l, --league string   get the league latest score.
```
#### `standing` command
Get the latest league standing
```shell
balbalancli standing [flags]

Examples:

        balbalan standing -l PL
        balbalan standing --league PL

Flags:
  -h, --help            help for standing
  -l, --league string   get the league standing.
```


## License

Distributed under the Apache. See `LICENSE` for more information.


## Go Package

* [cobra](https://github.com/spf13/cobra) 
* [tablewriter](https://github.com/olekukonko/tablewriter)
* [go-dotenv](https://github.com/profclems/go-dotenv)
* [fatih/color](https://github.com/fatih/color/)