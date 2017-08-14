package replaying

import (
	"github.com/v2pro/koala/recording"
)

type ReplayedTalk struct {
	MatchedTalk   *recording.Talk
	MatchedTalkIndex int
	ReplayedRequestTime int64
	ReplayedRequest []byte
	ReplayedResponseTime int64
}