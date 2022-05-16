package rule

const (
	// CmstTime обстоятельство времени
	CmstTime uint32 = 0x01
	// CmstPlaceAbs обстоятельство места абстрактное.
	// Например: в лесу, в школе. Для уточнения требуют координат.
	CmstPlaceAbs uint32 = 0x02
	// CmstPlaceX обстоятельство места долгота
	CmstPlaceX uint32 = 0x04
	// CmstPlaceZ обстоятельство места широта
	CmstPlaceZ uint32 = 0x08
	// CmstPlaceY обстоятельство места по вертикали
	CmstPlaceY uint32 = 0x10
	// CmstAction обстоятельства образа действия
	CmstAction uint32 = 0x20
)

func IsValidCmstElement(value uint32) bool {
	allCmst := CmstTime | CmstPlaceAbs | CmstPlaceX |
		CmstPlaceZ | CmstPlaceY | CmstAction

	return (allCmst & value) > 0
}
