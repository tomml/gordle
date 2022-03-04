package words

import (
	"math/rand"
	"sort"
	"time"
)

const WordMaxLength = 5

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomWordFromList() string {
	return wordsList[rand.Intn(len(wordsList))]
}

func CheckGuessIsInList(guess string) bool {
	sort.Strings(wordsList)
	i := sort.SearchStrings(wordsList, guess)
	return (i < len(wordsList) && wordsList[i] == guess)
}

var (
	wordsList = []string{
		"CIGAR", "REBUT", "SISSY", "DWARF", "QUIET", "STAFF", "FINAL", "SHEEN",
		"NAVAL", "SERVE", "HEATH", "DWARF", "MODEL", "KARMA", "STINK", "GRADE",
		"QUIET", "BENCH", "ABATE", "FEIGN", "MAJOR", "DEATH", "FRESH", "CRUST",
		"STOOL", "COLON", "ABASE", "MARRY", "REACT", "BATTY", "PRIDE", "FLOSS",
		"HELIX", "CROAK", "STAFF", "PAPER", "UNFED", "WHELP", "TRAWL", "OUTDO",
		"ADOBE", "CRAZY", "SOWER", "REPAY", "DIGIT", "CRATE", "CLUCK", "SPIKE",
		"MIMIC", "POUND", "MAXIM", "LINEN", "UNMET", "FLESH", "BOOBY", "FORTH",
		"FIRST", "STAND", "BELLY", "IVORY", "SEEDY", "PRINT", "YEARN", "DRAIN",
		"BRIBE", "STOUT", "PANEL", "CRASS", "FLUME", "OFFAL", "AGREE", "ERROR",
		"SWIRL", "ARGUE", "BLEED", "DELTA", "FLICK", "TOTEM", "WOOER", "FRONT",
		"SHRUB", "PARRY", "BIOME", "LAPEL", "START", "GREET", "GONER", "GOLEM",
		"LUSTY", "LOOPY", "ROUND", "AUDIT", "LYING", "GAMMA", "LABOR", "ISLET",
		"CIVIC", "FORGE", "CORNY", "MOULT", "BASIC", "SALAD", "AGATE", "SPICY",
		"SPRAY", "ESSAY", "FJORD", "SPEND", "KEBAB", "GUILD", "ABACK", "MOTOR",
		"ALONE", "HATCH", "HYPER", "THUMB", "DOWRY", "OUGHT", "BELCH", "DUTCH",
		"PILOT", "TWEED", "COMET", "JAUNT", "ENEMA", "STEED", "ABYSS", "GROWL",
		"FLING", "DOZEN", "BOOZY", "ERODE", "WORLD", "GOUGE", "CLICK", "BRIAR",
		"GREAT", "ALTAR", "PULPY", "BLURT", "COAST", "DUCHY", "GROIN", "FIXER",
		"GROUP", "ROGUE", "BADLY", "SMART", "PITHY", "GAUDY", "CHILL", "HERON",
		"VODKA", "FINER", "SURER", "RADIO", "ROUGE", "PERCH", "RETCH", "WROTE",
		"CLOCK", "TILDE", "STORE", "PROVE", "BRING", "SOLVE", "CHEAT", "GRIME",
		"EXULT", "USHER", "EPOCH", "TRIAD", "BREAK", "RHINO", "VIRAL", "CONIC",
		"MASSE", "SONIC", "VITAL", "TRACE", "USING", "PEACH", "CHAMP", "BATON",
		"BRAKE", "PLUCK", "CRAZE", "GRIPE", "WEARY", "PICKY", "ACUTE", "FERRY",
		"ASIDE", "TAPIR", "TROLL", "UNIFY", "REBUS", "BOOST", "TRUSS", "SIEGE",
		"TIGER", "BANAL", "SLUMP", "CRANK", "GORGE", "QUERY", "DRINK", "FAVOR",
		"ABBEY", "TANGY", "PANIC", "SOLAR", "SHIRE", "PROXY", "POINT", "ROBOT",
		"PRICK", "WINCE", "CRIMP", "KNOLL", "SUGAR", "WHACK", "MOUNT", "PERKY",
		"COULD", "WRUNG", "LIGHT", "THOSE", "MOIST", "SHARD", "PLEAT", "ALOFT",
		"SKILL", "ELDER", "FRAME", "HUMOR", "PAUSE", "ULCER", "ULTRA", "ROBIN",
		"CYNIC", "AGORA", "AROMA", "CAULK", "SHAKE", "PUPAL", "DODGE", "SWILL",
		"TACIT", "OTHER", "THORN", "TROVE", "BLOKE", "VIVID", "SPILL", "CHANT",
		"CHOKE", "RUPEE", "NASTY", "MOURN", "AHEAD", "BRINE", "CLOTH", "HOARD",
		"SWEET", "MONTH", "LAPSE", "WATCH", "TODAY", "FOCUS", "SMELT", "TEASE",
		"CATER", "MOVIE", "LYNCH", "SAUTE", "ALLOW", "RENEW", "THEIR", "SLOSH",
		"PURGE", "CHEST", "DEPOT", "EPOXY", "NYMPH", "FOUND", "SHALL", "HARRY",
		"STOVE", "LOWLY", "SNOUT", "TROPE", "FEWER", "SHAWL", "NATAL", "FIBRE",
		"COMMA", "FORAY", "SCARE", "STAIR", "BLACK", "SQUAD", "ROYAL", "CHUNK",
		"MINCE", "SLAVE", "SHAME", "CHEEK", "AMPLE", "FLAIR", "FOYER", "CARGO",
		"OXIDE", "PLANT", "OLIVE", "INERT", "ASKEW", "HEIST", "SHOWN", "ZESTY",
		"HASTY", "TRASH", "FELLA", "LARVA", "FORGO", "STORY", "HAIRY", "TRAIN",
		"HOMER", "BADGE", "MIDST", "CANNY", "FETUS", "BUTCH", "FARCE", "SLUNG",
		"TIPSY", "METAL", "YIELD", "DELVE", "BEING", "SCOUR", "GLASS", "GAMER",
		"SCRAP", "MONEY", "HINGE", "ALBUM", "VOUCH", "ASSET", "TIARA", "CREPT",
		"BAYOU", "ATOLL", "MANOR", "CREAK", "SHOWY", "PHASE", "FROTH", "DEPTH",
		"GLOOM", "FLOOD", "TRAIT", "GIRTH", "PIETY", "PAYER", "GOOSE", "FLOAT",
		"DONOR", "ATONE", "PRIMO", "APRON", "BLOWN", "CACAO", "LOSER", "INPUT",
	}
)
