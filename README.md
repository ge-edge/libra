[![Build Status](https://travis-ci.com/ge-edge/libra.svg?branch=v0.1.0-reiwa)](https://travis-ci.com/ge-edge/libra)

# libra
The project libra (Linear-Inversion Blockchain Research Assembly) is a subscription library implemented with De-centralised-logic approach.

## Table
* [Requirement](#requirement) 
* [How to compile](#how-to-compile) 
* [Resources](#resources)
* [Licensing](#licensing)
* [Libra Nodes Interface For Popular Brands](#libra-nodes-interface-for-popular-brands)
* [Reference](#reference)

## Requirement
- CMake
- GNU g++
- CLang
- Go 1.12+
- [wzlib crypto library](https://github.com/EC-Release/sdk/tree/v1.1beta/lib/go/pkg)

## How to compile
```
# clone the repo
git clone https://github.com/ge-edge/libra.git
cd libra
mkdir build
cd build
cmake ./..
make
```
upon the compiling, cmake will generate a c-shared library, a header file, the libra library, and a test artifact. You may verify the compiing detail in [libra ci](https://travis-ci.com/ge-edge/libra)

## Resources
- [Lucidcharts (Design, Diagram)](https://app.lucidchart.com/invitations/accept/ea4f9e3e-349a-4044-af09-264aa0cc5ec7)
- [Internal Resource (Slides, Licenses, Legal, etc.)](https://ge.box.com/s/z7czuln44wvk898t513jwpl6t3q0wguz)

## Licensing
Developer Licenses are free to acquire for beta usage and development purpose [in certifactory](https://github.com/EC-Release/certifactory/tree/beta).

## Libra Nodes Interface For Popular Brands
- Bitcoin Core[1]
- Hyperledger[2]
- AWS BC Service[3]
- Azure BC Service[4]

## Marketplace
Libra Transactional Node owners may share their services at global scale, it allows the owners to produtise or monetise their services over the node. 

## Collaboration
Libra utilised [the slack channel (look for #libra and request join)](enterprisecon-j2w6229.slack.com) for better communication across the board.

Libra is also avaialble [on twitter](https://twitter.com/LibraRelease)

## Project Management
Libra invites contributors whoever interested in managing the project scope, and/or customer-relation to [the current Salesforce CRM development](https://na139.salesforce.com/home/home.jsp?tsid=02u4W000001HD2F).

## Reference
<sup>[1]Bitcoin is a consensus network that enables a new payment system and a completely digital money. It is the first decentralized peer-to-peer payment network that is powered by its users with no central authority or middlemen. From a user perspective, Bitcoin is pretty much like cash for the Internet. Bitcoin can also be seen as the most prominent triple entry bookkeeping system in existence.</sup><br />
<sup>[2]Hyperledger is an open source community focused on developing a suite of stable frameworks, tools and libraries for enterprise-grade blockchain deployments. Hyperledger topology is currently adopted and monetised by companies such as Oracle, IBM, etc.</sup><br />
<sup>[3]Amazon trademark Blockchain service provision both centalised and de-centralised governmentship. The service partner can also promote their service amongst the subscribers in the AWS marketplace.</sup><br />
<sup>[4]Azure Blockchain Service is a fully managed ledger service that enables users the ability to grow and operate blockchain networks at scale in Azure</sup>

<sup>Author: A. Yasuda. Rev: v0.1.0</sup>
