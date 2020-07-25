package mo_message

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"testing"
)

const (
	sampleJson = `{
  "id": "xxxxxxxxxxxxxxxx",
  "threadId": "xxxxxxxxxxxxxxxx",
  "labelIds": [
    "UNREAD",
    "CATEGORY_PERSONAL",
    "INBOX"
  ],
  "snippet": "お使いの xxxxxx アカウントへのアクセスが xxxxxxxxx xxxxxxx に許可されました xxxxxxxxxxxxx@xxxxx.xxx アクセスを許可した覚えがない場合は、このアクティビティをご確認のうえ、アカウントを保護してください。 アクティビティを確認 このメールは xxxxxx のアカウントやサービスの重要な変更についてお知らせするためにお送りしています。 © xxxx",
  "payload": {
    "mimeType": "multipart/alternative",
    "headers": [
      {
        "name": "Delivered-To",
        "value": "xxxxxxxxxxxxx@xxxxx.xxx"
      },
      {
        "name": "Received",
        "value": "by xxxx:xxx:xxxx:x:x:x:x:x with SMTP id xxxxxxxxxxxxxxxx;        Fri, 24 Jul 2020 17:19:18 -0700 (PDT)"
      },
      {
        "name": "X-Received",
        "value": "by xxxx:xxx:xxxx:: with SMTP id xxxxxxxxxxxxxxxx.xxx.xxxxxxxxxxxxx;        Fri, 24 Jul 2020 17:19:18 -0700 (PDT)"
      },
      {
        "name": "ARC-Seal",
        "value": "i=1; a=rsa-sha256; t=1595636358; cv=none;        d=google.com; s=arc-20160816;        b=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/x         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         H1A+xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/xxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/e9         FWUg=="
      },
      {
        "name": "ARC-Message-Signature",
        "value": "i=1; a=rsa-sha256; c=relaxed/relaxed; d=google.com; s=arc-20160816;        h=to:from:subject:message-id:feedback-id:date:mime-version         :dkim-signature;        bh=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=;        b=xxxxxxxxxxxxxxxxxxxxxxxxx+xxx+xxxxxxxxxxxxx/xx+xxxxxxx/xxxxx+xxxxx         xxxxxxxxxxxxxxxxxxxxxxxxx+xxxxxxxxxxxxxxxxxxxxxxxxxx+xxxxxxxxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxx/xxxxx+xxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxx         /xxxxxxxxxxxxxxxx+xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx++j         3i1g=="
      },
      {
        "name": "ARC-Authentication-Results",
        "value": "i=1; mx.google.com;       dkim=pass header.i=@accounts.google.com header.s=20161025 header.b=xxxxxxxx;       spf=pass (google.com: domain of xxxxxxxxxxxxxx-xxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxx.xxx@gaia.bounces.google.com designates xxx.xx.xxx.xx as permitted sender) smtp.mailfrom=xxxxxxxxxxxxxx-xxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxx.xxx@gaia.bounces.google.com;       dmarc=pass (p=REJECT sp=REJECT dis=NONE) header.from=accounts.google.com"
      },
      {
        "name": "Return-Path",
        "value": "<xxxxxxxxxxxxxx-xxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxx.xxx@gaia.bounces.google.com>"
      },
      {
        "name": "Received",
        "value": "from mail-sor-f73.google.com (xxxx-xxx-xxx.google.com. [xxx.xx.xxx.xx])        by mx.google.com with SMTPS id xxxxxxxxxxxxxxxx.xx.xxxx.xx.xx.xx.xx.xx        for <xxxxxxxxxxxxx@xxxxx.xxx>        (Google Transport Security);        Fri, 24 Jul 2020 17:19:18 -0700 (PDT)"
      },
      {
        "name": "Received-SPF",
        "value": "pass (google.com: domain of xxxxxxxxxxxxxx-xxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxx.xxx@gaia.bounces.google.com designates xxx.xx.xxx.xx as permitted sender) client-ip=xxx.xx.xxx.xx;"
      },
      {
        "name": "Authentication-Results",
        "value": "mx.google.com;       dkim=pass header.i=@accounts.google.com header.s=20161025 header.b=MoNDJuNP;       spf=pass (google.com: domain of xxxxxxxxxxxxxx-xxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxx.xxx@gaia.bounces.google.com designates xxx.xx.xxx.xx as permitted sender) smtp.mailfrom=xxxxxxxxxxxxxx-xxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxx.xxx@gaia.bounces.google.com;       dmarc=pass (p=REJECT sp=REJECT dis=NONE) header.from=accounts.google.com"
      },
      {
        "name": "DKIM-Signature",
        "value": "v=1; a=rsa-sha256; c=relaxed/relaxed;        d=accounts.google.com; s=20161025;        h=mime-version:date:feedback-id:message-id:subject:from:to;        bh=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=;        b=xxxxxxxx+xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx+xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx+xxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxx=="
      },
      {
        "name": "X-Google-DKIM-Signature",
        "value": "v=1; a=rsa-sha256; c=relaxed/relaxed;        d=1e100.net; s=20161025;        h=x-gm-message-state:mime-version:date:feedback-id:message-id:subject         :from:to;        bh=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=;        b=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx+xxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxxxxxxxxxxxxxxxxxxxx++xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxxxxxxx+xxxxxxxxxxxxx+xxxxxxxxxxxxxxxxxxxxxxxxxxxxx/xxxxxxxxxxxxxx         xxxxxxxx+xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx         xxxx=="
      },
      {
        "name": "X-Gm-Message-State",
        "value": "xxxxxxxx+xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxx"
      },
      {
        "name": "X-Google-Smtp-Source",
        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxx/x/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/xxxxxxxxxxxxx=="
      },
      {
        "name": "MIME-Version",
        "value": "1.0"
      },
      {
        "name": "X-Received",
        "value": "by xxxx:xxx:xxxx:: with SMTP id xxxxxxxxxxxxxxxx.xxx.xxxxxxxxxxxxx; Fri, 24 Jul 2020 17:19:18 -0700 (PDT)"
      },
      {
        "name": "Date",
        "value": "Sat, 25 Jul 2020 00:19:17 GMT"
      },
      {
        "name": "X-Account-Notification-Type",
        "value": "127"
      },
      {
        "name": "Feedback-ID",
        "value": "127:account-notifier"
      },
      {
        "name": "X-Notifications",
        "value": "xxxxxxxxxxxxxxxx"
      },
      {
        "name": "Message-ID",
        "value": "<xxxxxxxxxxxxxxxxxxxxxx.x@notifications.google.com>"
      },
      {
        "name": "Subject",
        "value": "セキュリティ通知"
      },
      {
        "name": "From",
        "value": "Google <no-reply@accounts.google.com>"
      },
      {
        "name": "To",
        "value": "xxxxxxxxxxxxx@xxxxx.xxx"
      },
      {
        "name": "Content-Type",
        "value": "multipart/alternative; boundary=\"xxxxxxxxxxxxxxxxxxxxxxxxxxxx\""
      }
    ]
  },
  "sizeEstimate": 12186,
  "historyId": "2073",
  "internalDate": "1595636357000"
}
`
)

func TestMessage_Processed(t *testing.T) {
	j, err := es_json.ParseString(sampleJson)
	if err != nil {
		t.Error(err)
		return
	}
	msg := &Message{}
	if err := j.Model(msg); err != nil {
		t.Error(err)
		return
	}

	p, err := msg.Processed()
	if err != nil {
		t.Error(err)
	}

	if p.From.Address != "no-reply@accounts.google.com" {
		t.Error(p.From)
	}
}
