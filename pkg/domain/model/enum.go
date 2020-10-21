package model

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Transaction = 1)
*/
type ContextKey int

/*
ENUM(Male, Female)
*/
type Gender int

/*
ENUM(Hokkaido = 1, Aomori, Iwate, Miyagi, Akita, Yamagata, Hukushima, Ibaragi, Totigi, Gunma, Saitama, Tiba, Tokyo, Kanagawa, Nigata, Toyama, Ishikawa, Hukui, Yamanashi, Nagano, Gihu, Sizuoka, Aiti, Mie, Shiga, Kyoto, Osaka, Hyogo, Nara, Wakayama, Tottori, Shimane, Okayama, Hiroshima, Yamaguchi, Tokushima, Kagawa, Ehime, Koti, Hukuoka, Saga, Nagasaki, Kumamoto, Oita, Miyazaki, Kagoshima, Okinawa)
*/
type Prefecture int
