<a href="https://www.mailersend.com"><img src="https://www.mailersend.com/images/logo.svg" width="200px"/></a>

MailerSend Golang SDK

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Test](https://github.com/mailersend/mailersend-go/actions/workflows/test.yaml/badge.svg)](https://github.com/mailersend/mailersend-go/actions/workflows/test.yaml)

# Table of Contents
- [Installation](#installation)
- [Usage](#usage)
    - [Email](#email)
        - [Sending a basic email](#sending-a-basic-email)
        - [Sending an email with CC and BCC](#sending-an-email-with-cc-and-bcc)
        - [Sending an email with variables (simple personalisation)](#sending-an-email-with-variables-simple-personalization)
        - [Sending an email with personalization (advanced personalisation)](#sending-an-email-with-personalization-advanced-personalization)
        - [Sending a template-based email](#send-a-template-based-email)
        - [Sending an email with attachment](#sending-an-email-with-attachment)
    - [Activity](#activity)
        - [List activities](#list-activities)
    - [Analytics](#analytics)
        - [Activity data by date](#activity-data-by-date)
        - [Opens by country](#opens-by-country)
        - [Opens by user-agent name](#opens-by-user-agent-name)
        - [Opens by reading environment](#opens-by-reading-environment)
    - [Domains](#domains)
        - [Get all domains](#get-all-domains)
        - [Get a single domain](#get-a-single-domain)
        - [Delete a domain](#delete-a-domain)
        - [Get recipients for a domain](#get-recipients-for-a-domain)
        - [Update domain settings](#update-domain-settings)
    - [Messages](#messages)
        - [Get all messages](#get-all-messages)
        - [Get a single message](#get-a-single-message)
    - [Recipients](#recipients)
      - [Get all recipients](#get-all-recipients)
      - [Get all recipients for a specific domain](#get-all-recipients-for-a-specific-domain)
      - [Get a single recipients](#get-a-single-recipient)
      - [Delete a recipients](#delete-a-recipient)
    - [Tokens](#tokens)
      - [Create a token](#create-a-token)
      - [Pause / Unpause Token](#pause--unpause-token)
      - [Delete a token](#delete-a-token)
    - [Webhooks](#webhooks)
      - [Get all webhooks](#get-all-webhooks)
      - [Get a single webhook](#get-a-single-webhook)
      - [Create a webhook](#create-a-webhook)
      - [Update a Webhook](#update-a-webhook)
      - [Delete a Webhook](#delete-a-webhook)
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

### Sending a basic email

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}

```

### Sending an email with CC and BCC

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

### Sending an email with variables (simple personalization)

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	subject := "Subject {$var}"
	text := "This is the text version with a {$var}."
	html := "<p>This is the HTML version with a {$var}.</p>"

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
					Var: "value",
				},
			},
		},
	}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetText(text)
	message.SetHTML(html)
	
	message.SetSubstitutions(variables)
	
	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

### Sending an email with personalization (advanced personalization)

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

### Send a template-based email

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

### Sending an email with attachment

```go
package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"

    "github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	attachment := mailersend.Attachment{Filename: "file.jpg", Content: encoded}

	message.AddAttachment(attachment)

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

<a name="activity"></a>

## Activity

### List activities

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := time.Now().Add(-24 * time.Hour).Unix()
	to := time.Now().Unix()
	domainID := "domain-id"
	events := []string{"processed", "sent", "queued"}

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
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

## Domains

### Get all domains

```go
package main

import (
	"context"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

### Get recipients for a domain

```go
package main

import (
	"context"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

### Get all messages

```go
package main

import (
	"context"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

## Recipients

### Get all recipients

```go
package main

import (
	"context"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := &mailersend.ListRecipientOptions{
		Page:  1,
		Limit: 25,
	}
	
	_, _, err := ms.Recipient.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get all recipients for a specific domain

```go
package main

import (
	"context"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	domainID := "domain-id"

	options := &mailersend.ListRecipientOptions{
		DomainID: domainID,
		Page:     1,
		Limit:    25,
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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

## Tokens

### Create a token

```go
package main

import (
	"context"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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

### Get all webhooks

```go
package main

import (
	"context"
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	domainID := "domain-id"
	events := []string{"activity.sent", "activity.opened"}

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
	"log"
	"time"
	
	"github.com/mailersend/mailersend-go"
)

var APIKey = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

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
