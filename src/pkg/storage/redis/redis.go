package redis

import (
    "strings"
    "container/vector"
    "tideland-cgl.googlecode.com/hg"
    "tideland-rdc.googlecode.com/hg"
    "fetch/tokenizer"
)

type Redis struct {
    redis *rdc.RedisDatabase
}

func buildKey(index, scope, field, token string) string {
    return strings.Join([]string{index, scope, field, token}, ":")
}

func (r *Redis) getSetKeys(index, scope, field string, tc tokenizer.TokenChan) (keys vector.Vector) {
    var ts vector.Vector
    for token := range tc {
        ts.Push(buildKey(index, scope, field, token.Backing()))
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

func emptyResult() chan string {
    output := make(chan string)
    close(output)
    return output
}

func (r *Redis) Store(index, scope, id, field string, tc tokenizer.TokenChan) bool {
    scores := make(map[string]int)
    for token := range tc {
        scores[token.Backing()] += 1
    }
    return r.redis.MultiCommand(func(mc *rdc.MultiCommand) {
        for token, score := range scores {
            mc.Command("ZADD", buildKey(index, scope, field, token), score, id)
        }
    }).IsOK()
}

func (r *Redis) Delete(index, scope, id string) bool {
    panic("TODO: Implement")
    return false
}

func unionArgs(keys *vector.Vector) (destination string, args vector.Vector) {
    destination = cgl.NewUUID().String()
    args.Push(destination)
    args.Push(len(*keys))
    args.AppendVector(keys)
    return
}

func revRangeArgs(destination string) (args vector.Vector) {
    args.Push(destination)
    args.Push(0)
    args.Push(-1)
    return
}

func (r *Redis) SearchField(index, scope, field string, tc tokenizer.TokenChan) chan string {
    // var intersectArgs, rangeArgs vector.Vector
    keys := r.getSetKeys(index, scope, field, tc)
    if 0 == len(keys) {
        return emptyResult()
    }
    destination, uArgs := unionArgs(&keys)
    r.redis.Command("ZUNIONSTORE", uArgs...)
    return pumpValues(r.redis.Command("ZREVRANGE", revRangeArgs(destination)...))
}

func (r *Redis) SearchScope(index, scope string, tc tokenizer.TokenChan) chan string {
    return r.SearchField(index, scope, "*", tc)
}

func Build() (r *Redis) {
    return &Redis{redis: rdc.NewRedisDatabase(rdc.Configuration{})}
}
