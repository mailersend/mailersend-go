<a href="https://www.mailersend.com"><img src="https://www.mailersend.com/images/logo.svg" width="200px"/></a>

MailerSend Golang SDK

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Test](https://github.com/mailersend/mailersend-go/actions/workflows/test.yaml/badge.svg)](https://github.com/mailersend/mailersend-go/actions/workflows/test.yaml)

# Table of Contents
- [Installation](#installation)
- [Usage](#usage)
    - [Email](#email)
       - [Send an email](#send-an-email)
       - [Add CC, BCC recipients](#add-cc-bcc-recipients)
       - [Send a template-based email](#send-a-template-based-email)
       - [Personalization](#personalization)
       - [Send email with attachment](#send-email-with-attachment)
       - [Send email with inline attachment](#send-email-with-inline-attachment)
    - [Bulk Email](#bulk-email)
       - [Send bulk email](#send-bulk-email)
       - [Get bulk email status](#get-bulk-email-status)
    - [Activity](#activity)
       - [Get a list of activities](#get-a-list-of-activities)
    - [Analytics](#analytics)
       - [Activity data by date](#activity-data-by-date)
       - [Opens by country](#opens-by-country)
       - [Opens by user-agent name](#opens-by-user-agent-name)
       - [Opens by reading environment](#opens-by-reading-environment)
    - [Inbound Routes](#inbound-routes)
       - [Get a list of inbound routes](#get-a-list-of-inbound-routes)
       - [Get a single inbound route](#get-a-single-inbound-route)
       - [Add an inbound route](#add-an-inbound-route)
       - [Update an inbound route](#update-an-inbound-route)
       - [Delete an inbound route](#delete-an-inbound-route)
    - [Domains](#domains)
       - [Get a list of domains](#get-a-list-of-domains)
       - [Get a single domain](#get-a-single-domain)
       - [Delete a domain](#delete-a-domain)
       - [Add a Domain](#add-a-domain)
       - [Get DNS Records](#get-dns-records)
       - [Get verification status](#get-verification-status)
       - [Get a list of recipients per domain](#get-a-list-of-recipients-per-domain)
       - [Update domain settings](#update-domain-settings)
    - [Messages](#messages)
       - [Get a list of messages](#get-a-list-of-messages)
       - [Get a single message](#get-a-single-message)
    - [Scheduled Messages](#scheduled-messages)
       - [Get a list of scheduled messages](#get-a-list-of-scheduled-messages)
       - [Get a single scheduled message](#get-a-single-scheduled-message)
       - [Delete a scheduled message](#delete-a-scheduled-message)
    - [Recipients](#recipients)
       - [Get a list of recipients](#get-a-list-of-recipients)
       - [Get a single recipients](#get-a-single-recipient)
       - [Delete a recipients](#delete-a-recipient)
       - [Get recipients from a suppression list](#get-recipients-from-a-suppression-list)
       - [Add recipients to a suppression list](#add-recipients-to-a-suppression-list)
       - [Delete recipients from a suppression list](#delete-recipients-from-a-suppression-list)
    - [Tokens](#tokens)
       - [Create a token](#create-a-token)
       - [Pause / Unpause Token](#pause--unpause-token)
       - [Delete a token](#delete-a-token)
    - [Webhooks](#webhooks)
       - [Get a list of webhooks](#get-a-list-of-webhooks)
       - [Get a single webhook](#get-a-single-webhook)
       - [Create a webhook](#create-a-webhook)
       - [Update a Webhook](#update-a-webhook)
       - [Delete a Webhook](#delete-a-webhook)
    - [Templates](#templates)
       - [Get a list of templates](#get-a-list-of-templates)
       - [Get a single template](#get-a-single-template)
       - [Delete a template](#delete-a-template)
    - [Email Verification](#email-verification)
       - [Verify a single email](#verify-single-email)
       - [Get all email verification lists](#get-all-email-verification-lists)
       - [Get an email verification list](#get-an-email-verification-list)
       - [Create an email verification list](#create-an-email-verification-list)
       - [Verify an email list](#verify-an-email-list)
       - [Get email verification list results](#get-email-verification-list-results)
    - [SMS](#sms)
       - [Send an SMS](#send-an-sms)
    - [SMS Messages](#sms-messages)
       - [Get a list of SMS messages](#get-a-list-of-sms-messages)
       - [Get info on an SMS message](#get-info-on-an-sms-message)
    - [SMS Activity](#sms-activity)
       - [Get a list of activities](#get-a-list-of-sms-activities)
       - [Get activity of a single SMS message](#get-activity-of-a-single-sms-message)
    - [SMS Phone Numbers](#sms-phone-numbers)
       - [Get a list of SMS phone numbers](#get-a-list-of-sms-phone-numbers)
       - [Get an SMS phone number](#get-an-sms-phone-number)
       - [Update a single SMS phone number](#update-a-single-sms-phone-number)
       - [Delete an SMS phone number](#delete-an-sms-phone-number)
    - [SMS Recipients](#sms-recipients)
       - [Get a list of SMS recipients](#get-a-list-of-sms-recipients)
       - [Get an SMS recipient](#get-an-sms-recipient)
       - [Update a single SMS recipient](#update-a-single-sms-recipient)
    - [SMS Inbounds](#sms-inbounds)
       - [Get a list of SMS inbound routes](#get-a-list-of-sms-inbound-routes)
       - [Get a single SMS inbound route](#get-a-single-inbound-route)
       - [Create an SMS inbound route](#create-an-sms-inbound-route)
       - [Update an SMS inbound route](#update-an-sms-inbound-route)
       - [Delete an SMS inbound route](#delete-an-sms-inbound-route)
    - [SMS Webhooks](#sms-webhook)
       - [Get a list of SMS webhooks](#get-a-list-of-sms-webhooks)
       - [Get an SMS webhook](#get-an-sms-webhook)
       - [Create an SMS webhook](#create-an-sms-webhook)
       - [Update an SMS webhook](#update-an-sms-webhook)
       - [Delete an SMS webhook](#delete-an-sms-webhook)
    - [Sender Identities](#sender-identities)
      - [Get a list of Sender Identities](#get-a-list-of-sender-identities)
      - [Get a single Sender Identity](#get-a-single-sender-identity)
      - [Get a single Sender Identity By Email](#get-a-single-sender-identity-by-email)
      - [Add a Sender_Identity](#create-a-sender-identity)
      - [Update a Sender Identity](#update-a-sender-identity)
      - [Update a Sender Identity By Email](#update-a-sender-identity-by-email)
      - [Delete a Sender Identity](#delete-a-sender-identity)
      - [Delete a Sender Identity By Email](#delete-a-sender-identity-by-email)
	- [Other Endpoints](#other-endpoints)
	  - [Get an API Quota](#get-an-api-quota)
- [Types](#types)
- [Helpers](#helpers)   
- [Testing](#testing)
- [Support and Feedback](#support-and-feedback)
- [License](#license)

<a name="installation"></a>

# Installation
We recommend using this package with golang [modules](https://github.com/golang/go/wiki/Modules)

```
$ go get github.com/mailersend/mailersend-go
```

# Usage

## Email 

### Send an email

```go
package main

import (
    "context"
    "os"
    "fmt"
    "time"

    "github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	text := "This is the text content"
	html := "<p>This is the HTML content</p>"

	from := mailersend.From{
		Name:  "Your Name",
		Email: "your@domain.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: "your@client.com",
		},
	}
	
	// Send in 5 minute
	sendAt := time.Now().Add(time.Minute * 5).Unix()

	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)
	message.SetSendAt(sendAt)
	message.SetInReplyTo("client-id")

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}

```

### Add CC, BCC recipients

```go
package main

import (
	"context"
	"os"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	subject := "Subject"
	text := "This is the text content"
	html := "<p>This is the HTML content</p>"

	from := mailersend.From{
		Name:  "Your Name",
		Email: "your@domain.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: "your@client.com",
		},
	}

	cc := []mailersend.Recipient{
		{
			Name:  "CC",
			Email: "cc@client.com",
		},
	}

	bcc := []mailersend.Recipient{
		{
			Name:  "BCC",
			Email: "bcc@client.com",
		},
	}

	replyTo := mailersend.ReplyTo{
		Name:  "Reply To",
		Email: "reply_to@client.com",
	}
	
	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)
	message.SetCc(cc)
	message.SetBcc(bcc)
	message.SetReplyTo(replyTo)

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

### Send a template-based email

```go
package main

import (
    "context"
    "os"
    "fmt"
    "time"

    "github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	subject := "Subject"

	from := mailersend.From{
		Name:  "Your Name",
		Email: "your@domain.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: "your@client.com",
		},
	}

	variables := []mailersend.Variables{
		{
			Email: "your@client.com",
			Substitutions: []mailersend.Substitution{
				{
					Var:   "foo",
					Value: "bar",
				},
			},
		},
	}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID("template-id")
	message.SetSubstitutions(variables)
	
	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

### Personalization

```go
package main

import (
	"context"
	"os"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	subject := "Subject {{ var }}"
	text := "This is the text version with a {{ var }}."
	html := "<p>This is the HTML version with a {{ var }}.</p>"

	from := mailersend.From{
		Name:  "Your Name",
		Email: "your@domain.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: "your@client.com",
		},
	}

	personalization := []mailersend.Personalization{
		{
			Email: "your@client.com",
			Data: map[string]interface{}{
				"Var":   "value",
			},
		},
	}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetText(text)
	message.SetHTML(html)
	
	message.SetPersonalization(personalization)

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

### Send email with attachment

```go
package main

import (
	"bufio"
	"context"
	"os"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

    "github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	text := "This is the text content"
	html := "<p>This is the HTML content</p>"

	from := mailersend.From{
		Name:  "Your Name",
		Email: "your@domain.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: "your@client.com",
		},
	}

	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)

	// Open file on disk.
	f, _ := os.Open("./file.jpg")

	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	attachment := mailersend.Attachment{Filename: "file.jpg", Content: encoded}

	message.AddAttachment(attachment)

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```


### Send email with inline attachment

```go
package main

import (
	"bufio"
	"context"
	"os"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

    "github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	text := "This is the text content"
	html := "<p>This is the HTML content</p> <p><img src=\"cid:image.jpeg\"/></p>"

	from := mailersend.From{
		Name:  "Your Name",
		Email: "your@domain.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: "your@client.com",
		},
	}

	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)

	// Open file on disk.
	f, _ := os.Open("./image.jpeg")

	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Inside template add <img src="cid:image.jpg"/> should match ID 
	attachment := mailersend.Attachment{Filename: "image.jpeg", ID: "image.jpeg", Content: encoded, Disposition: mailersend.DispositionInline}

	message.AddAttachment(attachment)

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

<a name="activity"></a>

## Bulk Email

### Send bulk email

```go
package main

import (
    "context"
    "os"
	"time"
	"log"
	"fmt"
	
    "github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))
	
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	text := "This is the text content"
	html := "<p>This is the HTML content</p>"

	from := mailersend.From{
		Name:  "Your Name",
		Email: "your@domain.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: "your@client.com",
		},
	}

	var messages []*mailersend.Message
	
	for i := range [2]int{} {
		msg := &mailersend.Message{
			From:       from,
			Recipients: recipients,
			Subject:    fmt.Sprintf("%s %v", subject, i),
			Text:       text,
			HTML:       html,
		}
		messages = append(messages, msg)
	}
	
	_, _, err := ms.BulkEmail.Send(ctx, messages)
	if err != nil {
		log.Fatal(err)
	}

}

```

### Get bulk email status

```go
package main

import (
	"context"
	"os"
	"time"
	"log"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	_, _, err := ms.BulkEmail.Status(ctx, "bulk-email-id")
	if err != nil {
		log.Fatal(err)
	}
	
}
```

## Activity

### Get a list of activities

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := time.Now().Add(-24 * time.Hour).Unix()
	to := time.Now().Unix()
	domainID := "domain-id"

	options := &mailersend.ActivityOptions{
		DomainID: domainID,
		DateFrom: from, 
		DateTo: to,
	}
	
	_, _, err := ms.Activity.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Analytics

### Activity data by date

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := time.Now().Add(-24 * time.Hour).Unix()
	to := time.Now().Unix()
	domainID := "domain-id"
	events := []string{"sent", "queued"}

	options := &mailersend.AnalyticsOptions{
		DomainID:    domainID,
		DateFrom:    from,
		DateTo:      to,
		Event:       events,
	}

	_, _, err := ms.Analytics.GetActivityByDate(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Opens by country

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := time.Now().Add(-24 * time.Hour).Unix()
	to := time.Now().Unix()
	domainID := "domain-id"

	options := &mailersend.AnalyticsOptions{
		DomainID: domainID,
		DateFrom: from,
		DateTo:   to,
	}

	_, _, err := ms.Analytics.GetOpensByCountry(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Opens by user-agent name

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := time.Now().Add(-24 * time.Hour).Unix()
	to := time.Now().Unix()
	domainID := "domain-id"

	options := &mailersend.AnalyticsOptions{
		DomainID: domainID,
		DateFrom: from,
		DateTo:   to,
	}

	_, _, err := ms.Analytics.GetOpensByUserAgent(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Opens by reading environment

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := time.Now().Add(-24 * time.Hour).Unix()
	to := time.Now().Unix()
	domainID := "domain-id"

	options := &mailersend.AnalyticsOptions{
		DomainID: domainID,
		DateFrom: from,
		DateTo:   to,
	}

	_, _, err := ms.Analytics.GetOpensByReadingEnvironment(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Inbound Routes

### Get a list of inbound routes

```go
package main

import (
	"context"
	"os"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()

	domainID := "domain-id"

	listOptions := &mailersend.ListInboundOptions{
		DomainID: domainID,
	}
	
	_, _, _ = ms.Inbound.List(ctx, listOptions)
}
```

### Get a single inbound route

```go
package main

import (
	"context"
	"os"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()

	inboundID := "inbound-id"

	_, _, _ = ms.Inbound.Get(ctx, inboundID)
}
```


### Add an inbound route

```go
package main

import (
	"context"
	"os"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()

	domainID := "domain-id"

	createOptions := &mailersend.CreateInboundOptions{
		DomainID:      domainID,
		Name:          "Example Route",
		DomainEnabled: *mailersend.Bool(false),
		MatchFilter: &mailersend.MatchFilter{
			Type: "match_all",
		},
		InboundPriority: 1,
		CatchFilter: &mailersend.CatchFilter{},
		Forwards: []mailersend.Forwards{
			{
				Type:  "webhook",
				Value: "https://example.com",
			},
		},
	}

	_, _, _ = ms.Inbound.Create(ctx, createOptions)
}
```

### Update an inbound route

```go
package main

import (
	"context"
	"os"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()

	inboundID := "inbound-id"

	updateOptions := &mailersend.UpdateInboundOptions{
		Name:          "Example Route",
		DomainEnabled: *mailersend.Bool(true),
		InboundDomain: "inbound.example.com",
		InboundPriority: 1,
		MatchFilter: &mailersend.MatchFilter{
			Type: "match_all",
		},
		CatchFilter: &mailersend.CatchFilter{
			Type: "catch_recipient",
			Filters: []mailersend.Filter{
				{
					Comparer: "equal",
					Value:    "email",
				},
				{
					Comparer: "equal",
					Value:    "emails",
				},
			},
		},
		Forwards: []mailersend.Forwards{
			{
				Type:  "webhook",
				Value: "https://example.com",
			},
		},
	}

	_, _, _ = ms.Inbound.Update(ctx, inboundID, updateOptions)
}
```

### Delete an inbound route

```go
package main

import (
	"context"
	"os"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()

	inboundID := "inbound-id"

	_, _ = ms.Inbound.Delete(ctx, inboundID)
}
```

## Domains

### Get a list of domains

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListDomainOptions{
		Page:  1,
		Limit: 25,
	}

	_, _, err := ms.Domain.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single domain

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	domainID := "domain-id"

	_, _, err := ms.Domain.Get(ctx, domainID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a domain

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	domainID := "domain-id"

	_, err := ms.Domain.Delete(ctx, domainID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Add a domain

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.CreateDomainOptions{
		Name: "domain.test",
	}

	_, _, err := ms.Domain.Create(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get DNS Records

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"

	_, _, err := ms.Domain.GetDNS(ctx, domainID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get verification status

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"

	_, _, err := ms.Domain.Verify(ctx, domainID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a list of recipients per domain

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"
	
	options := &mailersend.GetRecipientsOptions{
	 	DomainID: domainID,
	 	Page:     1,
	 	Limit:    25,
	}
	
	_, _, err := ms.Domain.GetRecipients(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update domain settings

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"
	
	options := &mailersend.DomainSettingOptions{
		DomainID:    domainID,
		SendPaused:  mailersend.Bool(false),
		TrackClicks: mailersend.Bool(true),
	}
	
	_, _, err := ms.Domain.Update(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Messages

### Get a list of messages

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListMessageOptions{
		Page:  1,
		Limit: 25,
	}

	_, _, err := ms.Message.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single message

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	messageID := "message-id"

	_, _, err := ms.Message.Get(ctx, messageID)
	if err != nil {
		log.Fatal(err)
	}
}
```


## Scheduled messages

### Get a list of scheduled messages

```go
package main

import (
	"context"
	"os"
	"log"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()

	domainID := "domain-id"

	_, _, err := ms.ScheduleMessage.List(ctx, domainID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single scheduled message

```go
package main

import (
	"context"
	"os"
	"log"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()
	
	messageID := "message-id"

	_, _, err := ms.ScheduleMessage.Get(ctx, messageID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a scheduled message

```go
package main

import (
	"context"
	"os"
	"log"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.TODO()

	messageID := "message-id"
	
	_, err := ms.ScheduleMessage.Delete(ctx, messageID)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Recipients

### Get a list of recipients

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListRecipientOptions{
		//DomainID: domainID,
		Page:  1,
		Limit: 25,
	}
	
	_, _, err := ms.Recipient.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single recipient

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	recipientID := "recipient-id"

	_, _, err := ms.Recipient.Get(ctx, recipientID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a recipient

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	recipientID := "recipient-id"
	
	_, err := ms.Recipient.Delete(ctx, recipientID)
	if err != nil {
		log.Fatal(err)
	}
}
```


### Get recipients from a suppression list

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	listOptions := &mailersend.SuppressionOptions{
		DomainID: "domain-id",
		Page:     1,
		Limit:    25,
	}

	// List Block List Recipients 
	_, _, err := ms.Suppression.ListBlockList(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}

	// List Hard Bounces 
	_, _, _ = ms.Suppression.ListHardBounces(ctx, listOptions)

	// List Spam Complaints 
	_, _, _ = ms.Suppression.ListSpamComplaints(ctx, listOptions)
	
	// List Unsubscribes
	_, _, _ = ms.Suppression.ListUnsubscribes(ctx, listOptions)


}
```

### Add recipients to a suppression list

```go
package main

import (
	"context"
	"os"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Add Recipient to Block List
	createSuppressionBlockOptions := &mailersend.CreateSuppressionBlockOptions{
		DomainID:   "domain-id",
		Recipients: []string{"test@example.com"},
		Patterns: []string{".*@example.com"},
	}
	
	_, _, _ = ms.Suppression.CreateBlock(ctx, createSuppressionBlockOptions)

	// Add Recipient to Hard Bounces
	createSuppressionHardBounceOptions := &mailersend.CreateSuppressionOptions{
		DomainID:   "domain-id",
		Recipients: []string{"test@example.com"},
	}

	_, _, _ = ms.Suppression.CreateHardBounce(ctx, createSuppressionHardBounceOptions)

	// Add Recipient to Spam Complaints
	createSuppressionSpamComplaintsOptions := &mailersend.CreateSuppressionOptions{
		DomainID:   "domain-id",
		Recipients: []string{"test@example.com"},
	}

	_, _, _ = ms.Suppression.CreateHardBounce(ctx, createSuppressionSpamComplaintsOptions)

	// Add Recipient to Unsubscribes
	createSuppressionUnsubscribesOptions := &mailersend.CreateSuppressionOptions{
		DomainID:   "domain-id",
		Recipients: []string{"test@example.com"},
	}
	
	_, _, _ = ms.Suppression.CreateHardBounce(ctx, createSuppressionUnsubscribesOptions)

}
```

### Delete recipients from a suppression list

```go
package main

import (
	"context"
	"os"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"

	// Delete All {type}
	
	// mailersend.BlockList 
	// mailersend.HardBounces 
	// mailersend.SpamComplaints 
	// mailersend.Unsubscribes
	
	_, _ = ms.Suppression.DeleteAll(ctx, domainID, mailersend.Unsubscribes)
	
	// Delete 

	deleteSuppressionOption := &mailersend.DeleteSuppressionOptions{
		DomainID: domainID,
		Ids:      []string{"suppression-id"},
	}

	_, _ = ms.Suppression.Delete(ctx, deleteSuppressionOption, mailersend.Unsubscribes)


}
```

## Tokens

### Create a token

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"
	
	scopes := []string{
		"tokens_full", 
		"email_full",
		"domains_full",
		"activity_full",
		"analytics_full",
		"webhooks_full",
		"templates_full",
	}

	options := &mailersend.CreateTokenOptions{
		Name:     "token name",
		DomainID: domainID,
		Scopes:   scopes,
	}

	newToken, _, err := ms.Token.Create(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
	
	// Make sure you keep your access token secret
	log.Print(newToken.Data.AccessToken)
}
```

### Pause / Unpause Token

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tokenID := "token-id"
	
	updateOptions := &mailersend.UpdateTokenOptions{
		TokenID: tokenID,
		Status:  "pause/unpause",
	}

	_, _, err := ms.Token.Update(ctx, updateOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a Token

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	tokenID := "token-id"
	
	_, err := ms.Token.Delete(ctx, tokenID)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Webhooks

### Get a list of webhooks

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"
	
	options := &mailersend.ListWebhookOptions{
		DomainID: domainID,
		Limit:    25,
	}

	_, _, err := ms.Webhook.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	webhookID := "webhook-id"
	
	_, _, err := ms.Webhook.Get(ctx, webhookID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create a Webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"
	events := []string{"activity.opened", "activity.clicked"}

	createOptions := &mailersend.CreateWebhookOptions{
		Name:     "Webhook",
		DomainID: domainID,
		URL:      "https://test.com",
		Enabled:  mailersend.Bool(false),
		Events:   events,
	}
	
	_, _, err := ms.Webhook.Create(ctx, createOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a Webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	webhookID := "webhook-id"
	events := []string{"activity.clicked"}

	updateOptions := &mailersend.UpdateWebhookOptions{
		WebhookID: webhookID,
		Enabled:   mailersend.Bool(true),
		Events:    events,
	}
	
	_, _, err := ms.Webhook.Update(ctx, updateOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a Webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	webhookID := "webhook-id"
	
	_, err := ms.Webhook.Delete(ctx, webhookID)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Templates

### Get a list of templates

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	options := &mailersend.ListTemplateOptions{
		Page:  1,
		Limit: 25,
	}
	
	_, _, err := ms.Template.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single template

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	templateID := "template-id"
	
	_, _, err := ms.Template.Get(ctx, templateID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a template

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	templateID := "template-id"
	
	_, err := ms.Template.Delete(ctx, templateID)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Email Verification

### Verify a single email

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.SingleEmailVerificationOptions{
		Email: "john@doe.com"
	}

	_, _, err := ms.EmailVerification.VerifySingle(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get all email verification lists

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListEmailVerificationOptions{
		Page:  1,
		Limit: 25,
	}

	_, _, err := ms.EmailVerification.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an email verification list

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.EmailVerification.Get(ctx, "email-verification-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create an email verification list

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.CreateEmailVerificationOptions{
		Name:   "Email Verification List ",
		Emails: []string{"your@client.com", "your@client.eu"},
	}
	
	_, _, err := ms.EmailVerification.Create(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Verify an email list

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	_, _, err := ms.EmailVerification.Verify(ctx, "email-verification-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get email verification list results

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.GetEmailVerificationOptions{
		EmailVerificationId: "email-verification-id",
		Page:                1,
		Limit:               25,
	}

	_, _, err := ms.EmailVerification.GetResults(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

## SMS 

### Send an SMS

```go
package main

import (
	"context"
	"os"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	message := ms.Sms.NewMessage()
	message.SetFrom("your-number")
	message.SetTo([]string{"client-number"})
	message.SetText("This is the message content {{ var }}")

	personalization := []mailersend.SmsPersonalization{
		{
			PhoneNumber: "client-number",
			Data: map[string]interface{}{
				"var": "foo",
			},
		},
	}

	message.SetPersonalization(personalization)

	res, _ := ms.Sms.Send(context.TODO(), message)
	fmt.Printf(res.Header.Get("X-SMS-Message-Id"))
}
```

## SMS Messages

### Get a list of SMS messages

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListSmsMessageOptions{
		Limit: 10,
	}

	_, _, err := ms.SmsMessage.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get info on an SMS message

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.SmsMessage.Get(ctx, "sms-message-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

## SMS Activity

### Get a list of SMS activities

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.SmsActivityOptions{}
	
	_, _, err := ms.SmsActivityService.List(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get activity of a single SMS message

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.SmsActivityService.Get(context.TODO(), "message-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

## SMS phone numbers

### Get a list of SMS phone numbers

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.SmsNumberOptions{}

	_, _, err := ms.SmsNumber.List(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an SMS phone number

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.SmsNumber.Get(context.TODO(), "number-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a single SMS phone number

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.SmsNumberSettingOptions{
		Id:     "number-id",
		Paused: mailersend.Bool(false),
	}

	_, _, err := ms.SmsNumber.Update(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
}
```


### Delete an SMS phone number

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	numberID := "number-id"

	_, err := ms.SmsNumber.Delete(ctx, numberID)
	if err != nil {
		log.Fatal(err)
	}
}
```

## SMS recipients

### Get a list of SMS recipients

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.SmsRecipientOptions{SmsNumberId: "sms-number-id"}
	
	_, _, err := ms.SmsRecipient.List(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an SMS recipient

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.SmsRecipient.Get(context.TODO(), "sms-recipient-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a single SMS recipient

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.SmsRecipientSettingOptions{
		Id:     "sms-recipient-id",
		Status: "opt_out",
	}

	_, _, err := ms.SmsRecipient.Update(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
}
```


## SMS inbounds

### Get a list of SMS inbound routes

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	listOptions := &mailersend.ListSmsInboundOptions{
		SmsNumberId: "sms-number-id",
	}
	
	_, _, err := ms.SmsInbound.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single SMS inbound route

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.SmsInbound.Get(ctx, "sms-inbound-id")
	if err != nil {
		log.Fatal(err)
	}
	
}
```


### Create an SMS inbound route

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.CreateSmsInboundOptions{
		SmsNumberId: "sms-number-id",
		Name:        "Example Route",
		ForwardUrl:  "https://example.com",
		Filter: mailersend.Filter{
			Comparer: "equal",
			Value:    "START",
		},
		Enabled: mailersend.Bool(true),
	}
	
	_, _, err := ms.SmsInbound.Create(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update an SMS inbound route

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	options := &mailersend.UpdateSmsInboundOptions{
		Id:          "sms-inbound-id",
		SmsNumberId: "sms-number-id",
		Name:        "Example Route",
		ForwardUrl:  "https://example.com",
		Filter: mailersend.Filter{
			Comparer: "equal",
			Value:    "START",
		},
		Enabled: mailersend.Bool(false),
	}

	_, _, err := ms.SmsInbound.Update(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete an SMS inbound route

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := ms.SmsInbound.Delete(ctx, "sms-inbound-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

## SMS webhooks

### Get a list of SMS webhooks

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListSmsWebhookOptions{
		SmsNumberId: "sms-number-id",
	}

	_, _, err := ms.SmsWebhook.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an SMS webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.SmsWebhook.Get(ctx, "sms-webhook-id")
	if err != nil {
		log.Fatal(err)
	}
	
}
```


### Create an SMS webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	events := []string{"sms.sent"}
	
	options := &mailersend.CreateSmsWebhookOptions{
		SmsNumberId: "sms-number-id",
		Name:        "Webhook",
		Events:      events,
		URL:         "https://test.com",
		Enabled:  mailersend.Bool(false),
	}
	
	_, _, err := ms.SmsWebhook.Create(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update an SMS webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	events := []string{"sms.sent"}

	options := &mailersend.UpdateSmsWebhookOptions{
		Id:   "sms-webhook-id",
		Name: "Webhook",
		Events: events,
		Enabled: mailersend.Bool(true),
		URL:    "https://test.com",
	}
	
	_, _, err := ms.SmsWebhook.Update(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete an SMS webhook

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	_, err := ms.SmsWebhook.Delete(ctx, "sms-webhook-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Sender identities

### Get a list of Sender Identities

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListIdentityOptions{
		DomainID: "domain-id",
	}

	_, _, err := ms.Identity.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single Sender Identity

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.Identity.Get(ctx, "identity-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a single Sender Identity By Email

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.Identity.GetByEmail(ctx, "identity-email")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create a Sender Identity

```go
package main

import (
	"context"
	"os"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	options := &mailersend.CreateIdentityOptions{
		DomainID: "domain-id",
		Name:     "Sender Name",
		Email:    "Sender Email",
	}
	
	_, _, err := ms.Identity.Create(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a Sender Identity

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.UpdateIdentityOptions{
		Name:            "Sender Name",
		ReplyToEmail:    "Reply To Email",
	}

	_, _, err := ms.Identity.Update(ctx, "identity-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a Sender Identity By Email

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.UpdateIdentityOptions{
		Name:            "Sender Name",
		ReplyToEmail:    "Reply To Email",
	}

	_, _, err := ms.Identity.UpdateByEmail(ctx, "identity-email", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a Sender Identity

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := ms.Identity.Delete(ctx, "identity-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a Sender Identity By Email

```go
package main

import (
	"context"
	"os"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := ms.Identity.DeleteByEmail(ctx, "identity-email")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Other Endpoints

### Get an API Quota

```go
package main

import (
	"context"
	"log"
	"time"
	"fmt"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, _, err := ms.ApiQuota.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
```

# Types

Most API responses are Unmarshalled into their corresponding types.

You can see all available types on pkg.go.dev

https://pkg.go.dev/github.com/mailersend/mailersend-go#pkg-types

# Helpers

We provide a few helpers to help with development.

```go

// Create a pointer to a true boolean 
mailersend.Bool(true)

// Create a pointer to a false boolean
mailersend.Bool(false)

// Create a pointer to a Int
mailersend.Int(2)

// Create a pointer to a Int64
mailersend.Int64(2)

// Create a pointer to a String
mailersend.String("string")

```

# Testing

We provide interfaces for all services to help with testing

```go
type mockDomainService struct {
	mailersend.DomainService
}

func (m *mockDomainService) List(ctx context.Context, options *ListDomainOptions) (*DomainRoot, *Response, error) {
	return &mailersend.DomainRoot{Data: []mailersend.Domain{{Name: "example.com"}}}, nil, nil
}

func TestListDomains(t *testing.T) {
	client := &mailersend.Client{}
	client.Domain = &mockDomainService{}

	ctx := context.Background()
	result, _, err := client.Domain.List(ctx, nil)
	if err != nil || len(result.Data) == 0 || result.Data[0].Name != "example.com" {
		t.Fatalf("mock failed")
	}
}
```

[pkg/testing](https://golang.org/pkg/testing/)

```
$ go test
```

<a name="endpoints"></a>
# Available endpoints

| Feature group | Endpoint    | Available |
| ------------- | ----------- | --------- |
| Email         | `POST send` | ✅         |

*If, at the moment, some endpoint is not available, please use other available tools to access it. [Refer to official API docs for more info](https://developers.mailersend.com/).*


<a name="support-and-feedback"></a>
# Support and Feedback

In case you find any bugs, submit an issue directly here in GitHub.

You are welcome to create SDK for any other programming language.

If you have any troubles using our API or SDK free to contact our support by email [info@mailersend.com](mailto:info@mailersend.com)

The official documentation is at [https://developers.mailersend.com](https://developers.mailersend.com)

<a name="license"></a>
# License

[The MIT License (MIT)](LICENSE)
