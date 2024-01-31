# testbot

#### Run
```shell
git submodule init
git submodule update --recursive --checkout
make build
./testbot
```
#### Genesis Contracts
| Contract Address | Contract Name |  Path | classHash |
| ----------- | ----------- | --------- | ---------- |
|     0       | NoValidateAccount | data/genesis/NoValidateAccount.sierra.json | 0x24a22c925997bba54caa6ddf38b4d1616c9968e757d6846890b46caa88907e3  |
|     1       | UniversalDeployer | data/genesis/UniversalDeployer.json |  0x4569ffd48c2a3d455437c16dc843801fb896b1af845bc8bc7ba83ebc4358b7f   |
|     2       | cool_sierra_contract_class| data/genesis/cool_sierra_contract_class.json  | 0x35eb1d3593b1fe9a8369a023ffa5d07d3b2050841cb75ad6ef00698d9307d10  |
