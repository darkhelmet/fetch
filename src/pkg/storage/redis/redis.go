package redis

import (
    "strings"
    "container/vector"
    "tideland-rdc.googlecode.com/hg"
    "fetch/tokenizer"
)

type Redis struct {
    redis *rdc.RedisDatabase
}

func buildKey(index, scope, field string, token *tokenizer.Token) string {
    return strings.Join([]string{index, scope, field, token.Backing()}, ":")
}

func (r *Redis) getSetKeys(index, scope, field string, tc tokenizer.TokenChan) (keys vector.Vector) {
    var ts vector.Vector
    for token := range tc {
        ts.Push(buildKey(index, scope, field, token))
    }
    rs := r.redis.Command("keys", ts...)
    rs.ValuesDo(func(rv rdc.ResultValue) {
        keys.Push(rv.String())
    })
    return keys
}

func pumpValues(rs *rdc.ResultSet) chan string {
    output := make(chan string, 10)
    go func() {
        rs.ValuesDo(func(rv rdc.ResultValue) {
            output <- rv.String()
        })
        close(output)
    }()
    return output
}

func (r *Redis) Store(index, scope, id, field string, tc tokenizer.TokenChan) bool {
    return r.redis.MultiCommand(func(mc *rdc.MultiCommand) {
        for token := range tc {
            mc.Command("sadd", buildKey(index, scope, field, token), id)
        }
    }).IsOK()
}

func (r *Redis) SearchField(index, scope, field string, tc tokenizer.TokenChan) chan string {
    keys := r.getSetKeys(index, scope, field, tc)
    return pumpValues(r.redis.Command("sinter", keys...))
}

func (r *Redis) SearchScope(index, scope string, tc tokenizer.TokenChan) chan string {
    return r.SearchField(index, scope, "*", tc)
}

func Build() (r *Redis) {
    return &Redis{redis: rdc.NewRedisDatabase(rdc.Configuration{})}
}
