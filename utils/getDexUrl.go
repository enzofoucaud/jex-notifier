package utils

func GetDexUrl(source string) string {
	switch source {
	case "xexchange":
		return "https://xexchange.com/"
	case "jex_orderbook":
		return "https://app.jexchange.io/"
	case "onedex":
		return "https://swap.onedex.app/swap"
	case "ashswap":
		return "https://app.ashswap.io/swap/"
	case "vestadex":
		return "https://vestadex.com/"
	default:
		return ""
	}
}
