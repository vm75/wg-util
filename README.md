# WireGuard util

A command-line util for [WireGuard](https://www.wireguard.com/), presently supporting generating keys to be used with wireguard-go.

## Features

- Generates compliant [curve25519](https://cr.yp.to/ecdh.html) private and public keys


## Usage options

```
Usage: wg-util <genkey|genpsk|pubkey>

```


## Example

```
$ wg-util genkey | tee abc.key | wg-util pubkey > abc.pub
```

<!-- ## Installing

Download the [latest binary release](https://github.com/vm75/wg-util/releases/latest) for your system,
or build from source `go install github.com/vm75/wg-util@latest`. -->
