// Code generated by go-enum
// DO NOT EDIT!

package model

import (
	"fmt"
)

const (
	// ContextKeyTransaction is a ContextKey of type Transaction
	ContextKeyTransaction ContextKey = iota + 1
)

const _ContextKeyName = "Transaction"

var _ContextKeyMap = map[ContextKey]string{
	1: _ContextKeyName[0:11],
}

// String implements the Stringer interface.
func (x ContextKey) String() string {
	if str, ok := _ContextKeyMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ContextKey(%d)", x)
}

var _ContextKeyValue = map[string]ContextKey{
	_ContextKeyName[0:11]: 1,
}

// ParseContextKey attempts to convert a string to a ContextKey
func ParseContextKey(name string) (ContextKey, error) {
	if x, ok := _ContextKeyValue[name]; ok {
		return x, nil
	}
	return ContextKey(0), fmt.Errorf("%s is not a valid ContextKey", name)
}

// MarshalText implements the text marshaller method
func (x ContextKey) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *ContextKey) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseContextKey(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// GenderMale is a Gender of type Male
	GenderMale Gender = iota + 1
	// GenderFemale is a Gender of type Female
	GenderFemale
)

const _GenderName = "MaleFemale"

var _GenderMap = map[Gender]string{
	1: _GenderName[0:4],
	2: _GenderName[4:10],
}

// String implements the Stringer interface.
func (x Gender) String() string {
	if str, ok := _GenderMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Gender(%d)", x)
}

var _GenderValue = map[string]Gender{
	_GenderName[0:4]:  1,
	_GenderName[4:10]: 2,
}

// ParseGender attempts to convert a string to a Gender
func ParseGender(name string) (Gender, error) {
	if x, ok := _GenderValue[name]; ok {
		return x, nil
	}
	return Gender(0), fmt.Errorf("%s is not a valid Gender", name)
}

// MarshalText implements the text marshaller method
func (x Gender) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Gender) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseGender(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// MainMatchingStatusUndefined is a MainMatchingStatus of type Undefined
	MainMatchingStatusUndefined MainMatchingStatus = iota
	// MainMatchingStatusArrangeSchedule is a MainMatchingStatus of type ArrangeSchedule
	MainMatchingStatusArrangeSchedule
)

const _MainMatchingStatusName = "UndefinedArrangeSchedule"

var _MainMatchingStatusMap = map[MainMatchingStatus]string{
	0: _MainMatchingStatusName[0:9],
	1: _MainMatchingStatusName[9:24],
}

// String implements the Stringer interface.
func (x MainMatchingStatus) String() string {
	if str, ok := _MainMatchingStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("MainMatchingStatus(%d)", x)
}

var _MainMatchingStatusValue = map[string]MainMatchingStatus{
	_MainMatchingStatusName[0:9]:  0,
	_MainMatchingStatusName[9:24]: 1,
}

// ParseMainMatchingStatus attempts to convert a string to a MainMatchingStatus
func ParseMainMatchingStatus(name string) (MainMatchingStatus, error) {
	if x, ok := _MainMatchingStatusValue[name]; ok {
		return x, nil
	}
	return MainMatchingStatus(0), fmt.Errorf("%s is not a valid MainMatchingStatus", name)
}

// MarshalText implements the text marshaller method
func (x MainMatchingStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *MainMatchingStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseMainMatchingStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// MatchingReasonUndefined is a MatchingReason of type Undefined
	MatchingReasonUndefined MatchingReason = iota
	// MatchingReasonRenai is a MatchingReason of type Renai
	MatchingReasonRenai
	// MatchingReasonAsobi is a MatchingReason of type Asobi
	MatchingReasonAsobi
	// MatchingReasonImakaraNomitai is a MatchingReason of type ImakaraNomitai
	MatchingReasonImakaraNomitai
	// MatchingReasonSyumatsuDate is a MatchingReason of type SyumatsuDate
	MatchingReasonSyumatsuDate
)

const _MatchingReasonName = "UndefinedRenaiAsobiImakaraNomitaiSyumatsuDate"

var _MatchingReasonMap = map[MatchingReason]string{
	0: _MatchingReasonName[0:9],
	1: _MatchingReasonName[9:14],
	2: _MatchingReasonName[14:19],
	3: _MatchingReasonName[19:33],
	4: _MatchingReasonName[33:45],
}

// String implements the Stringer interface.
func (x MatchingReason) String() string {
	if str, ok := _MatchingReasonMap[x]; ok {
		return str
	}
	return fmt.Sprintf("MatchingReason(%d)", x)
}

var _MatchingReasonValue = map[string]MatchingReason{
	_MatchingReasonName[0:9]:   0,
	_MatchingReasonName[9:14]:  1,
	_MatchingReasonName[14:19]: 2,
	_MatchingReasonName[19:33]: 3,
	_MatchingReasonName[33:45]: 4,
}

// ParseMatchingReason attempts to convert a string to a MatchingReason
func ParseMatchingReason(name string) (MatchingReason, error) {
	if x, ok := _MatchingReasonValue[name]; ok {
		return x, nil
	}
	return MatchingReason(0), fmt.Errorf("%s is not a valid MatchingReason", name)
}

// MarshalText implements the text marshaller method
func (x MatchingReason) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *MatchingReason) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseMatchingReason(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// PrefectureHokkaido is a Prefecture of type Hokkaido
	PrefectureHokkaido Prefecture = iota + 1
	// PrefectureAomori is a Prefecture of type Aomori
	PrefectureAomori
	// PrefectureIwate is a Prefecture of type Iwate
	PrefectureIwate
	// PrefectureMiyagi is a Prefecture of type Miyagi
	PrefectureMiyagi
	// PrefectureAkita is a Prefecture of type Akita
	PrefectureAkita
	// PrefectureYamagata is a Prefecture of type Yamagata
	PrefectureYamagata
	// PrefectureHukushima is a Prefecture of type Hukushima
	PrefectureHukushima
	// PrefectureIbaragi is a Prefecture of type Ibaragi
	PrefectureIbaragi
	// PrefectureTotigi is a Prefecture of type Totigi
	PrefectureTotigi
	// PrefectureGunma is a Prefecture of type Gunma
	PrefectureGunma
	// PrefectureSaitama is a Prefecture of type Saitama
	PrefectureSaitama
	// PrefectureTiba is a Prefecture of type Tiba
	PrefectureTiba
	// PrefectureTokyo is a Prefecture of type Tokyo
	PrefectureTokyo
	// PrefectureKanagawa is a Prefecture of type Kanagawa
	PrefectureKanagawa
	// PrefectureNigata is a Prefecture of type Nigata
	PrefectureNigata
	// PrefectureToyama is a Prefecture of type Toyama
	PrefectureToyama
	// PrefectureIshikawa is a Prefecture of type Ishikawa
	PrefectureIshikawa
	// PrefectureHukui is a Prefecture of type Hukui
	PrefectureHukui
	// PrefectureYamanashi is a Prefecture of type Yamanashi
	PrefectureYamanashi
	// PrefectureNagano is a Prefecture of type Nagano
	PrefectureNagano
	// PrefectureGihu is a Prefecture of type Gihu
	PrefectureGihu
	// PrefectureSizuoka is a Prefecture of type Sizuoka
	PrefectureSizuoka
	// PrefectureAiti is a Prefecture of type Aiti
	PrefectureAiti
	// PrefectureMie is a Prefecture of type Mie
	PrefectureMie
	// PrefectureShiga is a Prefecture of type Shiga
	PrefectureShiga
	// PrefectureKyoto is a Prefecture of type Kyoto
	PrefectureKyoto
	// PrefectureOsaka is a Prefecture of type Osaka
	PrefectureOsaka
	// PrefectureHyogo is a Prefecture of type Hyogo
	PrefectureHyogo
	// PrefectureNara is a Prefecture of type Nara
	PrefectureNara
	// PrefectureWakayama is a Prefecture of type Wakayama
	PrefectureWakayama
	// PrefectureTottori is a Prefecture of type Tottori
	PrefectureTottori
	// PrefectureShimane is a Prefecture of type Shimane
	PrefectureShimane
	// PrefectureOkayama is a Prefecture of type Okayama
	PrefectureOkayama
	// PrefectureHiroshima is a Prefecture of type Hiroshima
	PrefectureHiroshima
	// PrefectureYamaguchi is a Prefecture of type Yamaguchi
	PrefectureYamaguchi
	// PrefectureTokushima is a Prefecture of type Tokushima
	PrefectureTokushima
	// PrefectureKagawa is a Prefecture of type Kagawa
	PrefectureKagawa
	// PrefectureEhime is a Prefecture of type Ehime
	PrefectureEhime
	// PrefectureKoti is a Prefecture of type Koti
	PrefectureKoti
	// PrefectureHukuoka is a Prefecture of type Hukuoka
	PrefectureHukuoka
	// PrefectureSaga is a Prefecture of type Saga
	PrefectureSaga
	// PrefectureNagasaki is a Prefecture of type Nagasaki
	PrefectureNagasaki
	// PrefectureKumamoto is a Prefecture of type Kumamoto
	PrefectureKumamoto
	// PrefectureOita is a Prefecture of type Oita
	PrefectureOita
	// PrefectureMiyazaki is a Prefecture of type Miyazaki
	PrefectureMiyazaki
	// PrefectureKagoshima is a Prefecture of type Kagoshima
	PrefectureKagoshima
	// PrefectureOkinawa is a Prefecture of type Okinawa
	PrefectureOkinawa
)

const _PrefectureName = "HokkaidoAomoriIwateMiyagiAkitaYamagataHukushimaIbaragiTotigiGunmaSaitamaTibaTokyoKanagawaNigataToyamaIshikawaHukuiYamanashiNaganoGihuSizuokaAitiMieShigaKyotoOsakaHyogoNaraWakayamaTottoriShimaneOkayamaHiroshimaYamaguchiTokushimaKagawaEhimeKotiHukuokaSagaNagasakiKumamotoOitaMiyazakiKagoshimaOkinawa"

var _PrefectureMap = map[Prefecture]string{
	1:  _PrefectureName[0:8],
	2:  _PrefectureName[8:14],
	3:  _PrefectureName[14:19],
	4:  _PrefectureName[19:25],
	5:  _PrefectureName[25:30],
	6:  _PrefectureName[30:38],
	7:  _PrefectureName[38:47],
	8:  _PrefectureName[47:54],
	9:  _PrefectureName[54:60],
	10: _PrefectureName[60:65],
	11: _PrefectureName[65:72],
	12: _PrefectureName[72:76],
	13: _PrefectureName[76:81],
	14: _PrefectureName[81:89],
	15: _PrefectureName[89:95],
	16: _PrefectureName[95:101],
	17: _PrefectureName[101:109],
	18: _PrefectureName[109:114],
	19: _PrefectureName[114:123],
	20: _PrefectureName[123:129],
	21: _PrefectureName[129:133],
	22: _PrefectureName[133:140],
	23: _PrefectureName[140:144],
	24: _PrefectureName[144:147],
	25: _PrefectureName[147:152],
	26: _PrefectureName[152:157],
	27: _PrefectureName[157:162],
	28: _PrefectureName[162:167],
	29: _PrefectureName[167:171],
	30: _PrefectureName[171:179],
	31: _PrefectureName[179:186],
	32: _PrefectureName[186:193],
	33: _PrefectureName[193:200],
	34: _PrefectureName[200:209],
	35: _PrefectureName[209:218],
	36: _PrefectureName[218:227],
	37: _PrefectureName[227:233],
	38: _PrefectureName[233:238],
	39: _PrefectureName[238:242],
	40: _PrefectureName[242:249],
	41: _PrefectureName[249:253],
	42: _PrefectureName[253:261],
	43: _PrefectureName[261:269],
	44: _PrefectureName[269:273],
	45: _PrefectureName[273:281],
	46: _PrefectureName[281:290],
	47: _PrefectureName[290:297],
}

// String implements the Stringer interface.
func (x Prefecture) String() string {
	if str, ok := _PrefectureMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Prefecture(%d)", x)
}

var _PrefectureValue = map[string]Prefecture{
	_PrefectureName[0:8]:     1,
	_PrefectureName[8:14]:    2,
	_PrefectureName[14:19]:   3,
	_PrefectureName[19:25]:   4,
	_PrefectureName[25:30]:   5,
	_PrefectureName[30:38]:   6,
	_PrefectureName[38:47]:   7,
	_PrefectureName[47:54]:   8,
	_PrefectureName[54:60]:   9,
	_PrefectureName[60:65]:   10,
	_PrefectureName[65:72]:   11,
	_PrefectureName[72:76]:   12,
	_PrefectureName[76:81]:   13,
	_PrefectureName[81:89]:   14,
	_PrefectureName[89:95]:   15,
	_PrefectureName[95:101]:  16,
	_PrefectureName[101:109]: 17,
	_PrefectureName[109:114]: 18,
	_PrefectureName[114:123]: 19,
	_PrefectureName[123:129]: 20,
	_PrefectureName[129:133]: 21,
	_PrefectureName[133:140]: 22,
	_PrefectureName[140:144]: 23,
	_PrefectureName[144:147]: 24,
	_PrefectureName[147:152]: 25,
	_PrefectureName[152:157]: 26,
	_PrefectureName[157:162]: 27,
	_PrefectureName[162:167]: 28,
	_PrefectureName[167:171]: 29,
	_PrefectureName[171:179]: 30,
	_PrefectureName[179:186]: 31,
	_PrefectureName[186:193]: 32,
	_PrefectureName[193:200]: 33,
	_PrefectureName[200:209]: 34,
	_PrefectureName[209:218]: 35,
	_PrefectureName[218:227]: 36,
	_PrefectureName[227:233]: 37,
	_PrefectureName[233:238]: 38,
	_PrefectureName[238:242]: 39,
	_PrefectureName[242:249]: 40,
	_PrefectureName[249:253]: 41,
	_PrefectureName[253:261]: 42,
	_PrefectureName[261:269]: 43,
	_PrefectureName[269:273]: 44,
	_PrefectureName[273:281]: 45,
	_PrefectureName[281:290]: 46,
	_PrefectureName[290:297]: 47,
}

// ParsePrefecture attempts to convert a string to a Prefecture
func ParsePrefecture(name string) (Prefecture, error) {
	if x, ok := _PrefectureValue[name]; ok {
		return x, nil
	}
	return Prefecture(0), fmt.Errorf("%s is not a valid Prefecture", name)
}

// MarshalText implements the text marshaller method
func (x Prefecture) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Prefecture) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePrefecture(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
