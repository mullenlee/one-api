package perplexity

import (
	"fmt"
	"github.com/songquanpeng/one-api/relay/meta"
	"github.com/songquanpeng/one-api/relay/relaymode"
)

func GetRequestURL(meta *meta.Meta) (string, error) {
	if meta.Mode == relaymode.ChatCompletions {
		return "https://api.perplexity.ai/chat/completions", nil
	}
	return "", fmt.Errorf("unsupported relay mode %d for perplexity", meta.Mode)
}
