package utils

var ConnId uint32

func GenerateConnID() uint32 {
	ConnId++
	return ConnId
}
