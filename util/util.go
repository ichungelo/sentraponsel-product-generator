package util

import (
	"encoding/json"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/geo/s2"
	"github.com/oklog/ulid"
)

const (
	earthRadiusM = 6371000 // per https://nssdc.gsfc.nasa.gov/planetary/factsheet/earthfact.html
	hashLength = 8
)

func GetUlid() (id string) {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	_id := ulid.MustNew(ulid.Timestamp(t), entropy)

	id = _id.String()

	return
}

func GenerateHash(lat float64, long float64) s2.CellID {
	latLng := s2.LatLngFromDegrees(lat, long)
	cell := s2.CellFromLatLng(latLng)
	return cell.ID()
}


func GenCellIntPrefix(cell s2.CellID) (hash uint64, prefix uint64) {
	hash = uint64(cell)
	geohashString := strconv.FormatUint(hash, 10)
	denominator := math.Pow10(len(geohashString) - hashLength)
	prefix = hash / uint64(denominator)
	return
}

func ProsesNameToStandard(name string) (n string) {
	re, _ := regexp.Compile(`[^\w]`)

	n = re.ReplaceAllString(name, "")

	n = strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(n, " ", ""), "_", ""))

	return
}

func MapToStruct(m interface{}, s interface{}) (err error) {
	bytes, err := json.Marshal(m)
	if nil != err {
		return err
	}

	err = json.Unmarshal(bytes, s)

	return
}