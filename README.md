# fuel-calc :fuelpump:

Calculate the fuel price in kilometers.

# Versions

- `v1` is a basic terminal application.
- `v2` is a CLI tool, that can be used in your terminal.

# Installation

Make sure you have `Go` installed on your machine.
You can download it by visiting the official [downloads page](https://golang.org/dl/).

## CLI (v2)

> :information_source: Make sure you are in the root directory of the project.

If you want to be able to use **fuel-calc** from any terminal, you can install it by executing

```sh
go install
```

Now, you can use the command `fuel-calc` from any terminal window :rocket:

## Terminal (v1)

To use the `v1` version of **fuel-calc**, you can run the following command in your terminal

```sh
go run main.go
```

You'll be prompted to enter the fuel consumption and the fuel price.

# Usage

## CLI (v2)

To use the CLI tool, you can run the following command in your terminal

```sh
fuel-calc -h # Show's the help message
```

Example usage:

```sh
fuel-calc 17 1.99
```

Output: `The price per kilometer is: 0.34 €`