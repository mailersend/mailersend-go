<a href="https://www.mailersend.com"><img src="https://www.mailersend.com/images/logo.svg" width="200px"/></a>

MailerSend Golang SDK

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE.md)

# Table of Contents
* [Installation](#installation)
* [Usage](#usage)
* [Testing](#testing)
* [Support and Feedback](#support-and-feedback)
* [License](#license)

<a name="installation"></a>
# Installation
We recomend using this package with golang [modules](https://github.com/golang/go/wiki/Modules)

```
$ go get github.com/mailersend/mailersend-go/v1
```

<a name="usage"></a>
# Usage

Sending a basic email.

``` go
package main

import (
    "github.com/mailersend/mailersend-go/v1"
)

var APIKey string = "Api Key Here"

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	//text := "This is the text content"
	//html := "<p>This is the HTML content</p>"

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

	tags := []string{"foo", "bar"}

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	//message.SetHTML(html)
	//message.SetText(text)
	message.SetTemplateID("testtemplateid")
	message.SetSubstitutions(variables)

	message.SetTags(tags)

	ms.Send(ctx, message)

}

```

<a name="testing"></a>

# Testing

[pkg/testing](https://golang.org/pkg/testing/)

```
$ go test
```
