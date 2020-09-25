package work_request

import (
	"testing"
)

func TestBuilderImpl_FilterUrl(t *testing.T) {
	b := &builderImpl{}
	urls := map[string]string{
		"https://slack.com/api/conversations.info?token=xxxx-xxxxxxxxx-xxxx&channel=C1234567890": "https://slack.com/api/conversations.info?" + slackTokenReplace + "&channel=C1234567890",
		"https://slack.com/api/conversations.info?channel=C1234567890":                           "https://slack.com/api/conversations.info?channel=C1234567890",
	}

	for target, expected := range urls {
		actual := b.FilterUrl(target)
		if actual != expected {
			t.Error(actual, expected)
		}
	}
}
