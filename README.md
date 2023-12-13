# jex-notifier

## Description

This tool is used to monitor the price of a token on the [JEXCHANGE](https://jexchange.io/) exchange and send a notification to a discord channel when the price is above or below a given price.

## Installation

Install [Golang](https://go.dev/doc/install). Then run:

```bash
go get github.com/enzofoucaud/node_taiko-jolnir-tx
```

## Configurations

```sh
cp example.config.json config.json
```

Edit `config.json` with your own values :

- `log_level` can be one of `debug` and `info`.
- `discord_id` and `discord_token` are extracted from the url. [Discord webhook url is in the format of: https://discord.com/api/webhooks/{DISCORD_ID}/{DISCORD_TOKEN}]
- `sleep_time` is the time in seconds to sleep between each check.
- `tokens` is a list of tokens to monitor.
  - `token` is the token identifier
  - `price` is the price to monitor
  - `is_below` is a boolean value to indicate whether to monitor if the price is below the given price
  - `is_above` is a boolean value to indicate whether to monitor if the price is above the given price.

### Tokens name list

- AERO-458bbf
- AIR-317920
- ALP-afc922
- ASH-a642d1
- BEAR-f9c271
- BEE-cb37b6
- BHAT-c1fde3
- BNI-bb0a30
- BONEZ-ff9a73
- BTX-0f676d
- BSK-baa025
- BUSD-40b57e
- C23-1325d3
- CGO-5e9528
- CHARGED-703583
- CHECKR-60108b
- CPA-97530a
- CRT-52decf
- CRU-a5f4aa
- CTP-298075
- CYBER-489c1c
- DICE-9749cc
- DKM-592719
- DNA-b144d1
- EBUD-d29cce
- ECITY-2a383a
- EARTH-cac0a1
- EFFORT-a13513
- EFOO-8fe2d4
- EFTR-315177
- ENERGY-36fb1d
- EPUNKS-dc0f59
- ESSENCE-67531c
- ESTAR-461bab
- EVLD-43f56f
- FIRE-a9a32a
- FIXX-074526
- GEMS-5fbf7d
- GIANT-761d27
- GNG-8735ae
- GSC-08cbdb
- HBUSD-ac1fca
- HERBS-1c45fb
- HEGLD-d61095
- HRD-71df2d
- HSEGLD-c13a4e
- HTM-f51d55
- HUSDC-d80042
- HUSDT-6f0914
- HUTK-4fa4b2
- HWBTC-49ca31
- HWETH-b3d17e
- HYPE-619661
- INFRA-43985c
- INSTANT-086961
- ISET-84e55e
- ITHEUM-df6f26
- JWLASH-f362b9
- JEX-9040ca
- JIASH-a6ae7a
- JIBUSD-9d7082
- JIEGLD-f6f180
- JIUSDC-6e2818
- JIUSDT-bcef75
- KOSON-5dd4fa
- KVRI-743439
- KRO-df97ec
- LEGLD-d74da9
- LKASH-10bd00
- LKMEX-aab910
- MAGIC-4798ea
- MEX-455c57
- MID-ecb7bf
- MOOVE-875539
- MPH-f8ea2b
- MUNCHKIN-3865e6
- NUTS-8ad81a
- ODIN-4d429b
- OFE-29eb54
- ONE-f9954f
- PLATA-9ba6c3
- PRICK-744592
- PROTEO-0c7311
- QWT-46ac01
- RARE-99e8b0
- REALM-8ead17
- RIDE-7d18e9
- RISA-c115c7
- SEGLD-3ad2d0
- SFIT-aebc90
- SING-35e8da
- SPORES-2f0887
- SPROTEO-c2dffe
- SUPER-507aa6
- TOLKEN-a9eb7f
- TTG-1abcec
- UPARK-982dd6
- USDC-c76f1f
- USDT-f8c08c
- UTK-2f80e9
- VEGLD-2b9319
- VITAL-ab7917
- VITAL-bc0917
- WAM-510e42
- WATER-9ed400
- WBTC-5349b3
- WEGLD-bd4d79
- WETH-b4ca29
- WFLS-5385b4
- WHALE-b018f0
- XAPES-1d15a5
- XBID-c7e360
- XBIT-630969
- XBONK-7cde03
- XBURN-33f7f6
- XLH-8daa50
- ZOG-c66239
- ZPAY-247875
- GCC-3194ab
- WAGMI-3f803d
- HODL-b8bd81
- MERMAID-9c388a
- MEME-2101aa
- ECPX-5cbfeb
- TCX-8d448d
- NEAT-811b8d
- PEPE-8ef042
- ORDER-f58c5b
- BEER-77650d
- LBT-1b7754
- TRO-94c925
- INFRA-758365
- ROF-c85ab7
- GROGU-b40e75
- REWARD-cf6eac
- EDIA-60b86f
- AAR-80c00b
- BFY-8344ff
- ANT-dada1a
- MINNIE-b45403
- HONK-bd4e39
- QAKAI-2d4840
- TIME-84518f
- TGR-68da1e
- XMPH-3af949
- LEGACY-64a1dc
- XAI-f10327
- PADAWAN-a17f58
- YFI-29c2dd
- XTW-78700a
- FOG-12657e
- FGG-9326d0
- MMLG-9a907b
- MGK-853e2f
- TCHAIN-414e2d
- TCH-cb0dfc
- RONE-bb2e69
- OURO-9ecd6a
- VST-c40502
- SSC-dcc54f
- HHTM-e03ba5
- WEB-b61d65
- SUPER-764d8d
- LAND-40f26f
- LAUNCH-3e2258
- HODL-d7f4b5
- DEAD-4c133a
- MARE-63e515
- LPAD-84628f
- ESTAR-afaaf0
- JOY-43bad3
- CUMB-8b7006
- EXR-401c82
- FMX-5dc275
- THANKS-27b918
- ELEROMWATT-121373
- LOOT-af0e6f
- ECULT-d81333
- VERSX-4e6329
- CAVIAR-2106e0
- NEXUS-71d1d6
- WAGMI-76c711
- MYMINT-e0b3a9
- NDO-691004
- JWLEGLD-023462
- JWLHTM-8e3cd5
- JWLUSD-62939e
- WSDAI-277fee
- ZAYA-024622

## Run

```sh
go run .
```
