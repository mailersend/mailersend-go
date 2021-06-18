<a href="https://www.mailersend.com"><img src="https://www.mailersend.com/images/logo.svg" width="200px"/></a>

MailerSend Golang SDK

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE.md)
[![Test](https://github.com/mailersend/mailersend-go/actions/workflows/test.yaml/badge.svg)](https://github.com/mailersend/mailersend-go/actions/workflows/test.yaml)
# Table of Contents
- [Installation](#installation)
- [Usage](#usage)
    - [Email](#email)
        - [Sending a basic email](#send-a-basic-email)
        - [Sending an email with CC and BCC](#sending-an-email-with-cc-and-bcc)
        - [Sending an email with variables (simple personalisation)](#sending-an-email-with-variables-simple-personalization)
        - [Sending an email with personalization (advanced personalisation)](#sending-an-email-with-personalization-advanced-personalization)
        - [Sending a template-based email](#send-a-template-based-email)
        - [Sending an email with attachment](#sending-an-email-with-attachment)
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

var APIKey string = "Api Key Here"

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

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)

	res, _ := ms.Send(ctx, message)

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

var APIKey string = "Api Key Here"

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

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)
	message.SetCc(cc)
	message.SetBcc(bcc)

	res, _ := ms.Send(ctx, message)

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

var APIKey string = "Api Key Here"

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
	
	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetText(text)
	message.SetHTML(html)
	
	message.SetSubstitutions(variables)
	
	res, _ := ms.Send(ctx, message)

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

var APIKey string = "Api Key Here"

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

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetText(text)
	message.SetHTML(html)
	
	message.SetPersonalization(personalization)

	res, _ := ms.Send(ctx, message)

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

var APIKey string = "Api Key Here"

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
	
	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID("testtemplateid")
	message.SetSubstitutions(variables)
	
	res, _ := ms.Send(ctx, message)

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

var APIKey string = "Api Key Here"

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

	message := ms.NewMessage()

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

	res, _ := ms.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
```

<a name="testing"></a>

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
