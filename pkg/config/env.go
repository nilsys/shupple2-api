package config

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(dev, stg, prd)
*/
type Env int

// NOTE: prd2のような環境を作る可能性を考慮して、メソッドで判定するようにする
func (e Env) IsPrd() bool {
	return e == EnvPrd
}
