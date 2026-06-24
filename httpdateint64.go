package httpdateint64

import "net/http"

var (
	days   = [...]string{"Thu", "Fri", "Sat", "Sun", "Mon", "Tue", "Wed"}
	months = [...]string{"Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec", "Jan", "Feb"}
)

// Conv mengubah UNIX timestamp (wajib >= 0) menjadi string HTTP Date Header (RFC 7231)
func Conv(unixTime int64) (buf [len(http.TimeFormat)]byte) {

	// 1. Hitung Hari dalam Seminggu (Epoch 1970-01-01 adalah Kamis(Thu))
	wday := (unixTime / 86400) % 7

	dayStr := days[wday]

	// 2. Hitung Waktu (Jam, Menit, Detik) dalam 1 hari (86400 detik)
	tod := unixTime % 86400
	h := tod / 3600
	m := (tod % 3600) / 60
	s := tod % 60

	// 3. Algoritma Kalender Howard Hinnant (Hanya untuk z >= 0)
	z := unixTime/86400 + 719468
	era := z / 146097
	doe := uint32(z - era*146097)
	yoe := (doe - doe/1460 + doe/36524 - doe/146096) / 365
	y := int64(yoe) + era*400
	doy := doe - (365*yoe + yoe/4 - yoe/100)
	mp := (5*doy + 2) / 153
	d := doy - (153*mp+2)/5 + 1

	if mp > 9 {
		y++
	}

	monthStr := months[mp]

	// 4. Konstruksi String langsung ke Buffer Array 29 Byte (Zero-Alloc)
	buf[0] = dayStr[0]
	buf[1] = dayStr[1]
	buf[2] = dayStr[2]
	buf[3] = ','
	buf[4] = ' '
	buf[5] = byte('0' + d/10)
	buf[6] = byte('0' + d%10)
	buf[7] = ' '
	buf[8] = monthStr[0]
	buf[9] = monthStr[1]
	buf[10] = monthStr[2]
	buf[11] = ' '
	buf[12] = byte('0' + (y/1000)%10)
	buf[13] = byte('0' + (y/100)%10)
	buf[14] = byte('0' + (y/10)%10)
	buf[15] = byte('0' + y%10)
	buf[16] = ' '
	buf[17] = byte('0' + h/10)
	buf[18] = byte('0' + h%10)
	buf[19] = ':'
	buf[20] = byte('0' + m/10)
	buf[21] = byte('0' + m%10)
	buf[22] = ':'
	buf[23] = byte('0' + s/10)
	buf[24] = byte('0' + s%10)
	buf[25] = ' '
	buf[26] = 'G'
	buf[27] = 'M'
	buf[28] = 'T'
	return
}
