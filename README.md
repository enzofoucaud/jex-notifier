## Configurations

```sh
cp example.config.json config.json
```

Edit `config.json` with your own values :

- `log_level` can be one of `debug` and `info`.
- `discord_id` and `discord_token` are extracted from the url. [Discord webhook url is in the format of: https://discord.com/api/webhooks/{DISCORD_ID}/{DISCORD_TOKEN}]
- `sleep_time` is the time in seconds to sleep between each check.
- `tokens` is a list of tokens to monitor.
  - `token` is the token name
  - `price` is the price to monitor
  - `is_below` is a boolean value to indicate whether to monitor if the price is below the given price
  - `is_above` is a boolean value to indicate whether to monitor if the price is above the given price.

### Tokens name list

- Elrond eGold
- WrappedEGLD
- Exrond
- LAUNCH
- Water
- DEAD
- ThankYou
- JOY
- Munchkin
- HODL
- VITAL
- Mermaid
- EleRon
- ZorgCoin
- LOOT
- ElrondBuds
- ELRONDCULTCOIN
- versX
- Caviar [SCAM alert](https://twitter.com/PepeVersX/status/1654854429592506369?s=20)
- Nexus
- MYSTERYMINT
- WAGMI
- ECITY
- INFRA
