package discord

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/automuteus/utils/pkg/game"

	"github.com/bwmarrin/discordgo"
)

const (
	UnlinkEmojiName = "auunlink"
	X               = "❌"
	ThumbsUp        = "👍"
	Hourglass       = "⌛"
)

// Emoji struct for discord
type Emoji struct {
	Name string
	ID   string
}

// FormatForInline does what it sounds like
func (e *Emoji) FormatForInline() string {
	return "<:" + e.Name + ":" + e.ID + ">"
}

// GetDiscordCDNUrl does what it sounds like
func (e *Emoji) GetDiscordCDNUrl() string {
	return "https://cdn.discordapp.com/emojis/" + e.ID + ".png"
}

// DownloadAndBase64Encode does what it sounds like
func (e *Emoji) DownloadAndBase64Encode() string {
	url := e.GetDiscordCDNUrl()
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	encodedStr := base64.StdEncoding.EncodeToString(bytes)
	return "data:image/png;base64," + encodedStr
}

func emptyStatusEmojis() AlivenessEmojis {
	topMap := make(AlivenessEmojis)
	topMap[true] = make([]Emoji, 34) // 34 colors for alive/dead
	topMap[false] = make([]Emoji, 34)
	return topMap
}

func (bot *Bot) addAllMissingEmojis(s *discordgo.Session, guildID string, alive bool, serverEmojis []*discordgo.Emoji) {
	for i, emoji := range GlobalAlivenessEmojis[alive] {
		alreadyExists := false
		for _, v := range serverEmojis {
			if v.Name == emoji.Name {
				emoji.ID = v.ID
				log.Println("Found emoji in guild. i=" + strconv.Itoa(i) + " alive=" + strconv.FormatBool(alive) + " length=" + strconv.Itoa(len(bot.StatusEmojis[alive])))
				bot.StatusEmojis[alive][i] = emoji
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			log.Println("Creating new emoji: " + emoji.Name)
			b64 := emoji.DownloadAndBase64Encode()
			em, err := s.GuildEmojiCreate(guildID, emoji.Name, b64, nil)
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Added emoji %s successfully!\n", emoji.Name)
				emoji.ID = em.ID
				bot.StatusEmojis[alive][i] = emoji
			}
		}
	}
}

func EmojisToSelectMenuOptions(emojis []Emoji, unlinkEmoji string) (arr []discordgo.SelectMenuOption) {
	for i, v := range emojis {
		// can only add 25 emojis (24 + unlink) on the menu, others can be added via command
		if i >= 24 {
			continue
		}
		arr = append(arr, v.toSelectMenuOption(game.GetColorStringForInt(i)))
	}
	arr = append(arr, discordgo.SelectMenuOption{
		Label:   "unlink",
		Value:   UnlinkEmojiName,
		Emoji:   discordgo.ComponentEmoji{Name: unlinkEmoji},
		Default: false,
	})
	return arr
}

func (e Emoji) toSelectMenuOption(displayName string) discordgo.SelectMenuOption {
	return discordgo.SelectMenuOption{
		Label:   displayName,
		Value:   displayName, // use the Name for listen events later
		Emoji:   discordgo.ComponentEmoji{ID: e.ID},
		Default: false,
	}
}

// AlivenessEmojis map
type AlivenessEmojis map[bool][]Emoji

// GlobalAlivenessEmojis keys are IsAlive, Color
var GlobalAlivenessEmojis = AlivenessEmojis{
	true: []Emoji{
		game.Red: {
			Name: "aured",
			ID:   "866558066921177108",
		},
		game.Blue: {
			Name: "aublue",
			ID:   "866558066484183060",
		},
		game.Green: {
			Name: "augreen",
			ID:   "866558066568986664",
		},
		game.Pink: {
			Name: "aupink",
			ID:   "866558067004538891",
		},
		game.Orange: {
			Name: "auorange",
			ID:   "866558066902958090",
		},
		game.Yellow: {
			Name: "auyellow",
			ID:   "866558067243221002",
		},
		game.Black: {
			Name: "aublack",
			ID:   "866558066442895370",
		},
		game.White: {
			Name: "auwhite",
			ID:   "866558067026165770",
		},
		game.Purple: {
			Name: "aupurple",
			ID:   "866558066966396928",
		},
		game.Brown: {
			Name: "aubrown",
			ID:   "866558066564136970",
		},
		game.Cyan: {
			Name: "aucyan",
			ID:   "866558066525601853",
		},
		game.Lime: {
			Name: "aulime",
			ID:   "866558066963382282",
		},
		game.Maroon: {
			Name: "aumaroon",
			ID:   "866558066917113886",
		},
		game.Rose: {
			Name: "aurose",
			ID:   "866558066921439242",
		},
		game.Banana: {
			Name: "aubanana",
			ID:   "866558065917558797",
		},
		game.Gray: {
			Name: "augray",
			ID:   "866558066174459905",
		},
		game.Tan: {
			Name: "autan",
			ID:   "866558066820382721",
		},
		game.Coral: {
			Name: "aucoral",
			ID:   "866558066552209448",
		},
		game.Watermelon: {
			Name: "auwatermelon",
			ID:   "1066789583804702953",
		},
		game.Chocolate: {
			Name: "auchocolate",
			ID:   "1066789585608249415",
		},
		game.SkyBlue: {
			Name: "auskyblue",
			ID:   "1066789587701202944",
		},
		game.Beige: {
			Name: "aubeige",
			ID:   "1066789589433466960",
		},
		game.Magenta: {
			Name: "aumagenta",
			ID:   "1066749135157469234",
		},
		game.Turquoise: {
			Name: "auturquoise",
			ID:   "1066789592864395344",
		},
		game.Lilac: {
			Name: "aulilac",
			ID:   "1066789594948980746",
		},
		game.Olive: {
			Name: "auolive",
			ID:   "1066789597478125608",
		},
		game.Azure: {
			Name: "auazure",
			ID:   "1066789599172644915",
		},
		game.Plum: {
			Name: "auplum",
			ID:   "1066749303646863391",
		},
		game.Jungle: {
			Name: "aujungle",
			ID:   "1067882143646236773",
		},
		game.Mint: {
			Name: "aumint",
			ID:   "1066749219844673566",
		},
		game.Chartreuse: {
			Name: "auchartreuse",
			ID:   "1066748547648725122",
		},
		game.Macau: {
			Name: "aumacau",
			ID:   "1066748685750386698",
		},
		game.Tawny: {
			Name: "autawny",
			ID:   "1066749360018296962",
		},
		game.Gold: {
			Name: "augold",
			ID:   "1066748597309292545",
		},
		// dont even think of adding rainbow
	},
	false: []Emoji{
		game.Red: {
			Name: "aureddead",
			ID:   "866558067255279636",
		},
		game.Blue: {
			Name: "aubluedead",
			ID:   "866558066660999218",
		},
		game.Green: {
			Name: "augreendead",
			ID:   "866558067088949258",
		},
		game.Pink: {
			Name: "aupinkdead",
			ID:   "866558066945556512",
		},
		game.Orange: {
			Name: "auorangedead",
			ID:   "866558067508510730",
		},
		game.Yellow: {
			Name: "auyellowdead",
			ID:   "866558067206520862",
		},
		game.Black: {
			Name: "aublackdead",
			ID:   "866558066668339250",
		},
		game.White: {
			Name: "auwhitedead",
			ID:   "866558067231293450",
		},
		game.Purple: {
			Name: "aupurpledead",
			ID:   "866558067223298048",
		},
		game.Brown: {
			Name: "aubrowndead",
			ID:   "866558066945163304",
		},
		game.Cyan: {
			Name: "aucyandead",
			ID:   "866558067051200512",
		},
		game.Lime: {
			Name: "aulimedead",
			ID:   "866558067344408596",
		},
		game.Maroon: {
			Name: "aumaroondead",
			ID:   "866558067238895626",
		},
		game.Rose: {
			Name: "aurosedead",
			ID:   "866558067083444225",
		},
		game.Banana: {
			Name: "aubananadead",
			ID:   "866558066342625350",
		},
		game.Gray: {
			Name: "augraydead",
			ID:   "866558067049758740",
		},
		game.Tan: {
			Name: "autandead",
			ID:   "866558067230638120",
		},
		game.Coral: {
			Name: "aucoraldead",
			ID:   "866558067024723978",
		},
		game.Watermelon: {
			Name: "auwatermelondead",
			ID:   "1066789611466137651",
		},
		game.Chocolate: {
			Name: "auchocolatedead",
			ID:   "1066789613689114634",
		},
		game.SkyBlue: {
			Name: "auskybluedead",
			ID:   "1066789615652057259",
		},
		game.Beige: {
			Name: "aubeigedead",
			ID:   "1066789617480769658",
		},
		game.Magenta: {
			Name: "aumagentadead",
			ID:   "1066749137837637682",
		},
		game.Turquoise: {
			Name: "auturquoisedead",
			ID:   "1066789623034040320",
		},
		game.Lilac: {
			Name: "aulilacdead",
			ID:   "1066789625726763058",
		},
		game.Olive: {
			Name: "auolivedead",
			ID:   "1066789627832307813",
		},
		game.Azure: {
			Name: "auazuredead",
			ID:   "1066789629623275570",
		},
		game.Plum: {
			Name: "auplumdead",
			ID:   "1066749306658361374",
		},
		game.Jungle: {
			Name: "aujungledead",
			ID:   "1067882146934575164",
		},
		game.Mint: {
			Name: "aumintdead",
			ID:   "1066749222046683137",
		},
		game.Chartreuse: {
			Name: "auchartreusedead",
			ID:   "1066748549372596315",
		},
		game.Macau: {
			Name: "aumacaudead",
			ID:   "1066748688501846138",
		},
		game.Tawny: {
			Name: "autawnydead",
			ID:   "1066749363252101201",
		},
		game.Gold: {
			Name: "augolddead",
			ID:   "1066748600383717386",
		},
		// dont even think of adding rainbow
	},
}

/*
Helpful for copy/paste into Discord to get new emoji IDs when they are re-uploaded...
\:aured:
\:aublue:
\:augreen:
\:aupink:
\:auorange:
\:auyellow:
\:aublack:
\:auwhite:
\:aupurple:
\:aubrown:
\:aucyan:
\:aulime:
\:aumaroon:
\:aurose:
\:aubanana:
\:augray:
\:autan:
\:aucoral:

\:aureddead:
\:aubluedead:
\:augreendead:
\:aupinkdead:
\:auorangedead:
\:auyellowdead:
\:aublackdead:
\:auwhitedead:
\:aupurpledead:
\:aubrowndead:
\:aucyandead:
\:aulimedead:
\:aumaroondead:
\:aurosedead:
\:aubananadead:
\:augraydead:
\:autandead:
\:aucoraldead:
*/
