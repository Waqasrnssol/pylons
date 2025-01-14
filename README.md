
# Pylons

Pylons is a blockchain network that allows artists, game studios, and other creatives to sell unique items using Stripe, Apple Pay, and Google Pay.  These items can then move freely through the IBC network.

This repository contains:

* The Pylons blockchain software, written in Go using the Cosmos-SDK.
* Big Dipper, a block explorer from Forbole that we have customized for use with Pylons.  Thank you Forbole!
* Our genesis state

## Documentation

### Docker build
Add your favorite peers to scripts/peers.sh and then you can run
```
docker-compose up -- build
```
You can use -d if you want it in the background
Should be fun!

### Outdated stuff
To learn how to use run a Pylons node, check out the [documentation pages](./docs/README.md).

### SDK
To learn more about developing NFTs on Pylons, see the [Pylons SDK](https://github.com/Pylons-tech/pylons_dart_sdk) project.

## Our Research

* We are actively researching new ways to allow for NFTs to be listed for sale on different IBC enabled blockchain networks after purchase.
* We are the first permissionless layer one chain to integrate with traditional credit card payment systems.
* We are working to achieve 500ms block times, to allow for seamless integration with mobile games.

## Talk to us!

Pylons is maintained by Pylons, LLC in collaboration with RNS and Notional Labs.

* [Discord](https://discord.gg/dZgUGy32j7)
* [Twitter](https://twitter.com/pylonstech)
* [Jobs](https://www.linkedin.com/company/pylons/jobs/)

## Contributing

We welcome contributions from everyone.  For more information about contributing, please review our [guidelines](CONTRIBUTING.md). Thank you to our Pylons contributors!
