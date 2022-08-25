package stratempo

var magnitudeMap = map[string]int{
	"cent":        100,
	"cento":       100,
	"mille":       powInt(10, 3),
	"mila":        powInt(10, 3),
	"milione":     powInt(10, 6),
	"milioni":     powInt(10, 6),
	"miliardo":    powInt(10, 9),
	"migliardo":   powInt(10, 9),
	"miliardi":    powInt(10, 9),
	"migliardi":   powInt(10, 9),
	"trilione":    powInt(10, 12),
	"trilioni":    powInt(10, 12),
	"quadrilione": powInt(10, 15),
	"quadrilioni": powInt(10, 15),

	"dozzina": 12,
	"dozzine": 12,
	"decina":  10,
	"decine":  10,
}
