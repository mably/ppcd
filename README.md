ppcd
====

[![Build Status](https://travis-ci.org/mably/ppcd.png?branch=master)]
(https://travis-ci.org/mably/ppcd)

ppcd is an alternative full node peercoin implementation written in Go (golang).

This project is currently under active development and is in a Beta state.  It
is extremely stable and has been in production use for over 6 months as of May
2014, however there are still a couple of major features we want to add before
we come out of beta.

It properly downloads, validates, and serves the block chain using the exact
rules (including bugs) for block acceptance as Peercoin Core.  We have taken
great care to avoid ppcd causing a fork to the block chain.

It also relays newly mined blocks, maintains a transaction pool, and relays
individual transactions that have not yet made it into a block.  It ensures all
transactions admitted to the pool follow the rules required by the block chain
and also includes the same checks which filter transactions based on
miner requirements ("standard" transactions) as Bitcoin Core.

One key difference between ppcd and Peercoin Core is that ppcd does *NOT* include
wallet functionality and this was a very intentional design decision.  See the
blog entry [here](https://blog.conformal.com/btcd-not-your-moms-bitcoin-daemon)
for more details.  This means you can't actually make or receive payments
directly with btcd.  That functionality is provided by the
[btcwallet](https://github.com/conformal/btcwallet) and
[btcgui](https://github.com/conformal/btcgui) projects which are both under
active development.

## Requirements

[Go](http://golang.org) 1.2 or newer.

## Installation

#### Windows - MSI Available

https://github.com/mably/ppcd/releases

#### Linux/BSD/POSIX - Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Run the following command to obtain btcd, all dependencies, and install it:
  ```$ go get github.com/mably/ppcd/...```

- btcd (and utilities) will now be installed in either ```$GOROOT/bin``` or
  ```$GOPATH/bin``` depending on your configuration.  If you did not already
  add the bin directory to your system path during Go installation, we
  recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/POSIX - Build from Source

- Run the following command to update btcd, all dependencies, and install it:
  ```$ go get -u -v github.com/mably/ppcd/...```

## Getting Started

ppcd has several configuration options avilable to tweak how it runs, but all
of the basic operations described in the intro section work with zero
configuration.

#### Windows (Installed from MSI)

Launch ppcd from your Start menu.

#### Linux/BSD/POSIX/Source

```bash
$ ./ppcd
````

## IRC server

- chat.freenode.net:6697
- channel #peercoin

## Forum

- http://www.peercointalk.org

## Issue Tracker

The [integrated github issue tracker](https://github.com/mably/ppcd/issues)
is used for this project.

## Documentation

The documentation is a work-in-progress.  It uses the [github wiki](https://github.com/mably/ppcd/wiki) facility.

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from Conformal.  To verify the
signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

btcd is licensed under the [copyfree](http://copyfree.org) ISC License.
