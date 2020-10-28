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

	subject := "Test Email"
	//text := "This is the message content"
	//html := "<p>This is the html content 3</p>"

	from := mailersend.From{
		Name:  "MailerSend",
		Email: "test@ycode.com",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Bob Gordon",
			Email: "robert@mailerlite.com",
		},
	}

	variables := []mailersend.Variables{
		{
			Email: "robert@mailerlite.com",
			Substitutions: []mailersend.Substitution{
				{
					Var:   "test",
					Value: "Test",
				},
			},
		},
	}

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	//message.SetHTML(html)
	//message.SetText(text)
	message.SetTemplateID("7z3m5jgrvd4dpyo6")
	message.SetSubstitutions(variables)

	ms.Send(ctx, message)

}

```

<a name="testing"></a>

# Testing

[pkg/testing](https://golang.org/pkg/testing/)

```
$ go test
```
