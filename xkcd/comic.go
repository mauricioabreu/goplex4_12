package xkcd

// Comic comic data
// {
// 		"month": "9",
// 		"num": 2354,
// 		"link": "",
// 		"year": "2020",
// 		"news": "",
// 		"safe_title": "Stellar Evolution",
// 		"transcript": "",
// 		"alt": "It may remain in equilibrium for some time, slowly growing, and then suddenly become significantly redder.",
// 		"img": "https://imgs.xkcd.com/comics/stellar_evolution.png",
// 		"title": "Stellar Evolution",
// 		"day": "2"
//  }
type Comic struct {
	Day        string
	Month      string
	Year       string
	Num        int
	Link       string
	News       string
	SafeTitle  string
	Transcript string
	Alt        string
	Img        string
	Title      string
}
