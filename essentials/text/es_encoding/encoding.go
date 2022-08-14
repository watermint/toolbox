package es_encoding

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"strings"
)

var (
	EncodingNames = []string{
		"utf8",
		"utf8bom",
		"utf16",
		"utf16bom",
		"utf16le",
		"utf16be",
		"sjis", "shift-jis",
		"iso-2022-jp",
		"euc-jp",
		"euc-kr",
		"gb18030",
		"gbk", "cp936",
		"hzgb2312",
		"big5",
		"cp037", "cp37", "codepage037",
		"cp1047", "codepage1047",
		"cp1140", "codepage1140",
		"cp437", "codepage437",
		"cp850", "codepage850",
		"cp852", "codepage852",
		"cp855", "codepage855",
		"cp858", "codepage858",
		"cp860", "codepage860",
		"cp862", "codepage862",
		"cp863", "codepage863",
		"cp865", "codepage865",
		"cp866", "codepage866",
		"iso8859-1",
		"iso8859-10",
		"iso8859-13",
		"iso8859-14",
		"iso8859-15",
		"iso8859-16",
		"iso8859-2",
		"iso8859-3",
		"iso8859-4",
		"iso8859-5",
		"iso8859-6",
		"iso8859-7",
		"iso8859-8",
		"iso8859-9",
		"koi8r",
		"koi8u",
		"macintosh",
		"macintoshcyrillic",
		"windows1250",
		"windows1251",
		"windows1252",
		"windows1253",
		"windows1254",
		"windows1255",
		"windows1256",
		"windows1257",
		"windows1258",
		"windows874",
	}
)

// SelectEncoding selects encoding instance from encoding name.
// Returns nil if no encoding instance for the name
func SelectEncoding(encName string) encoding.Encoding {
	switch strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(encName, "-", ""), "_", "")) {
	// # Unicode encodings

	case "utf8":
		return unicode.UTF8
	case "utf8bom":
		return unicode.UTF8BOM
	case "utf16":
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	case "utf16bom":
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
	case "utf16le":
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	case "utf16be":
		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)

		// # Japanese encodings

	case "sjis", "shiftjis":
		return japanese.ShiftJIS
	case "iso2022jp":
		return japanese.ISO2022JP
	case "eucjp":
		return japanese.EUCJP

		// # Korean encodings

	case "euckr":
		return korean.EUCKR

		// # Simplified Chinese encodings

	case "gb18030":
		return simplifiedchinese.GB18030
	case "gbk", "cp936":
		return simplifiedchinese.GBK
	case "hzgb2312":
		return simplifiedchinese.HZGB2312

		// # Traditional Chinese encodings

	case "big5":
		return traditionalchinese.Big5

		// # 8 bit character maps

	case "cp037", "cp37", "codepage037": // CodePage037 is the IBM Code Page 037 encoding.
		return charmap.CodePage037

	case "cp1047", "codepage1047": // CodePage1047 is the IBM Code Page 1047 encoding.
		return charmap.CodePage1047

	case "cp1140", "codepage1140": // CodePage1140 is the IBM Code Page 1140 encoding.
		return charmap.CodePage1140

	case "cp437", "codepage437": // CodePage437 is the IBM Code Page 437 encoding.
		return charmap.CodePage437

	case "cp850", "codepage850": // CodePage850 is the IBM Code Page 850 encoding.
		return charmap.CodePage850

	case "cp852", "codepage852": // CodePage852 is the IBM Code Page 852 encoding.
		return charmap.CodePage852

	case "cp855", "codepage855": // CodePage855 is the IBM Code Page 855 encoding.
		return charmap.CodePage855

	case "cp858", "codepage858": // CodePage858 is the Windows Code Page 858 encoding.
		return charmap.CodePage858

	case "cp860", "codepage860": // CodePage860 is the IBM Code Page 860 encoding.
		return charmap.CodePage860

	case "cp862", "codepage862": // CodePage862 is the IBM Code Page 862 encoding.
		return charmap.CodePage862

	case "cp863", "codepage863": // CodePage863 is the IBM Code Page 863 encoding.
		return charmap.CodePage863

	case "cp865", "codepage865": // CodePage865 is the IBM Code Page 865 encoding.
		return charmap.CodePage865

	case "cp866", "codepage866": // CodePage866 is the IBM Code Page 866 encoding.
		return charmap.CodePage866

	case "iso88591": // ISO8859_1 is the ISO 8859-1 encoding.
		return charmap.ISO8859_1

	case "iso885910": // ISO8859_10 is the ISO 8859-10 encoding.
		return charmap.ISO8859_10

	case "iso885913": // ISO8859_13 is the ISO 8859-13 encoding.
		return charmap.ISO8859_13

	case "iso885914": // ISO8859_14 is the ISO 8859-14 encoding.
		return charmap.ISO8859_14

	case "iso885915": // ISO8859_15 is the ISO 8859-15 encoding.
		return charmap.ISO8859_15

	case "iso885916": // ISO8859_16 is the ISO 8859-16 encoding.
		return charmap.ISO8859_16

	case "iso88592": // ISO8859_2 is the ISO 8859-2 encoding.
		return charmap.ISO8859_2

	case "iso88593": // ISO8859_3 is the ISO 8859-3 encoding.
		return charmap.ISO8859_3

	case "iso88594": // ISO8859_4 is the ISO 8859-4 encoding.
		return charmap.ISO8859_4

	case "iso88595": // ISO8859_5 is the ISO 8859-5 encoding.
		return charmap.ISO8859_5

	case "iso88596": // ISO8859_6 is the ISO 8859-6 encoding.
		return charmap.ISO8859_6

	case "iso88597": // ISO8859_7 is the ISO 8859-7 encoding.
		return charmap.ISO8859_7

	case "iso88598": // ISO8859_8 is the ISO 8859-8 encoding.
		return charmap.ISO8859_8

	case "iso88599": // ISO8859_9 is the ISO 8859-9 encoding.
		return charmap.ISO8859_9

	case "koi8r": // KOI8R is the KOI8-R encoding.
		return charmap.KOI8R

	case "koi8u": // KOI8U is the KOI8-U encoding.
		return charmap.KOI8U

	case "macintosh": // Macintosh is the Macintosh encoding.
		return charmap.Macintosh

	case "macintoshcyrillic": // MacintoshCyrillic is the Macintosh Cyrillic encoding.
		return charmap.MacintoshCyrillic

	case "windows1250": // Windows1250 is the Windows 1250 encoding.
		return charmap.Windows1250

	case "windows1251": // Windows1251 is the Windows 1251 encoding.
		return charmap.Windows1251

	case "windows1252": // Windows1252 is the Windows 1252 encoding.
		return charmap.Windows1252

	case "windows1253": // Windows1253 is the Windows 1253 encoding.
		return charmap.Windows1253

	case "windows1254": // Windows1254 is the Windows 1254 encoding.
		return charmap.Windows1254

	case "windows1255": // Windows1255 is the Windows 1255 encoding.
		return charmap.Windows1255

	case "windows1256": // Windows1256 is the Windows 1256 encoding.
		return charmap.Windows1256

	case "windows1257": // Windows1257 is the Windows 1257 encoding.
		return charmap.Windows1257

	case "windows1258": // Windows1258 is the Windows 1258 encoding.
		return charmap.Windows1258

	case "windows874": // Windows874 is the Windows 874 encoding
		return charmap.Windows874

	default:
		return nil
	}
}
