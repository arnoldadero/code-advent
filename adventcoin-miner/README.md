# AdventCoin Miner

This Go program helps Santa mine AdventCoins by finding the smallest positive number that, when combined with a given secret key, results in an MD5 hash starting with a specified number of leading zeroes.

## Problem

The task is to find the lowest number `n` such that the MD5 hash of the concatenation of a secret key and `n` produces a hash with at least five leading zeroes in its hexadecimal form.

### Example:

- Secret key: `abcdef`
- The lowest number `n` that produces a valid hash is `609043`, because the MD5 hash of `abcdef609043` starts with `000001dbbfa...`.

## How to Run

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/adventcoin-miner.git
    cd adventcoin-miner
    ```

2. Install Go modules:
    ```bash
    go mod tidy
    ```

3. Run the program:
    ```bash
    go run main.go
    ```

## Tests

Run the tests using:

```bash
go test ./...

'''

## Part 2: Mining with Six Zeroes

In this part, you need to find the lowest number such that the MD5 hash of the secret key combined with the number starts with **six leading zeroes**.

### Example:

- Secret key: `iwrupvqb`
- The lowest number that produces a valid hash is `9958218`, because the MD5 hash of `iwrupvqb9958218` starts with `000000ab871...`.

### How to Run

Simply run the program again after updating it to search for six leading zeroes:

```bash
go run main.go
