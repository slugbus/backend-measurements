// Package measurements measures distances
package measurements

import (
	"math"
)

const CenterLat = 36.99250
const CenterLong = -112.060569

// Speed returns the speed of a bus in mph
// where d is in miles and t is in milli
func Speed(distance float64, time float64) float64 {
	seconds := time / 1000
	hours := seconds / 60 / 60
	return distance / hours
}

// Dir gives direction relative
func Angle(lat1, long1, lat2, long2 float64) float64 {
	top := math.Log(math.Tan((lat2 / 2) + (math.Pi / 4)))
	bottom := math.Log(math.Tan((lat1 / 2) + (math.Pi / 4)))
	dTeta := top - bottom

	dLong := math.Abs(long1 - long2)
	teta := math.Atan2(dLong, dTeta)

	return math.Round((180 / math.Pi) * teta)
}

// Dist returns distance traveled in miles
func GetDistance(lat1, long1, lat2, long2 float64) float64 {
	lat1 = lat1 * math.Pi / 180
	long1 = long1 * math.Pi / 180

	lat2 = lat2 * math.Pi / 180
	long2 = long2 * math.Pi / 180

	dlon := long2 - long1
	dlat := lat2 - lat1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))

	r := float64(6371)

	kilo := c * r
	miles := kilo * 0.621371

	return miles
}


func getDistanceFromStopToStop(firstStopID int, secondStopID int) float64 {
	outDistances := [17]float64{0.3, 0.4, 0.4, 0.2, 0.2, 0.1, 0.2, 0.2, 0.1, 0.5, 0.1, 0.2, 0.5, 0.4, 0.1}
	innerDistances := [17]float64{0.2, 0.4, 0.7, 0.2, 0.3, 0.2, 0.1, 0.2, 0.4, 0.4, 0.4, 0.4, 0.4}
	distance := 0
	if firstBusStop.IsInner == true {
		for firstStopID < secondStopId {
			distance += innerDistances[firstStopID + 1]
			firstStopID++
		}
	}
	else if firstBusStop.isOutter == true {
		for firstStopID < secondStopID{
			distance += outDistances[firstStopID + 1]
			firstStopID++
		}
	}
}

// GetETA func
func GetETA(firstStopID int, secondStopID int, speed float64) float64 {
	// get amount of time in hours
	distance := getDistanceFromStopToStop(firstStopID, secondStopID)
	t := distance / speed
	// convert to seconds from hours
	t = t * 60 * 60
	return t
}

// returns the quadrant of a given lat and long coordinate
func getCurrentQuad(Lat, Long float64) string {
	// if in the upper left of campus return Quad 1
	if Lat < CenterLat && Long < CenterLong {
		return "Q1"
	}
	// if in the upper right of campus return Quad 2
	if Lat > CenterLat && Long < CenterLong {
		return "Q2"
	}
	// if in the bottom right of campus return Qaud 3
	if Lat > CenterLat && Long > CenterLong {
		return "Q3"
	}
	// if in the bottom left return Quad4
	return "Q4"
}
