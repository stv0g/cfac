package lanuv

// https://www.lanuv.nrw.de/fileadmin/lanuv/luft/temes/heut/VACW.htm

// https://github.com/mpfeil/qualitySCHU
// https://github.com/mpfeil/lanuv-now

// https://luadb.lds.nrw.de/LUA/hygon/pegel.php?rohdaten=ja
// https://github.com/Piccar/NRW-Wetter/blob/master/new1.py
// https://github.com/smatt/LanuvParser/blob/master/LanuvParser.java

const (
	UrlBase      = "https://www.lanuv.nrw.de/fileadmin/lanuv/luft/immissionen/aktluftqual"
	UrlCsv       = UrlBase + "/eu_luftqualitaet.csv"
	UrlCsvHeader = UrlBase + "/header_eu_luftqualitaet.csv"
)
