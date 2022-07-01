# What's HackerGo?

### HackerGo is a [Golang](https://go.dev/) CLI tool built with [Cobra (& cobra-cli)](https://github.com/spf13/cobra) and [Goquery](https://github.com/PuerkitoBio/goquery).
HackerGo is a CLI tool that aids developers stay productive and up-to-date with the technology world.
You'll find HackerGo useful if you're struggling with productivity;
You can use HackerGo to read tech news seamlessly from your CLI, and seeking for Job opportunities.

You don't need to leave your IDE/Editor; The world is waiting; Keep hacking hacker 🦹🏾‍ 🦸🏼

## Pre-requisites
- [The Go SDK](https://go.dev/dl/)
- You; Yes, you!


## Getting started

There are two options available if you want to get started using HackerGo.

- Clone the repository
```
$ git clone https://github.com/Goodnessuc/HackerGo.git
```
After you've cloned the repository, you can build the tool by running:

 ```
 $ go build
 $ go install
 ```
- Or just install tool directly in your project.
```
$ go install github.com/Goodnessuc/HackerGo.git
```



## HackerGo Usage
### There are three commands, `jobs`, `page`, and  `random` (so far).
<br>

#### The `jobs` command scrapes recent jobs from the [Hacker-news website](https://news.ycombinator.com/jobs) based on the additional arguments you provide

```
$ HackerGo jobs  <all | integer>
```
<br>

The  `all` argument scrapes all jobs on the first page of the website.
If you specify an integer less than 30 after specifying the page, you'll get the first "integer" jobs from the website.

<br>

Here's an example of the `jobs` command:

```
<br>
$ HackerGo jobs 5

Fetching you those gigs; This might take a while...
Smarking (YC W15) Is Hiring Lead ML Eng to scale tech for $655B parking industry (lever.co) 
 https://jobs.lever.co/smarking/01b7a4c5-28ce-4a4c-9c88-d4cad6c01c76

UPchieve (Edtech Nonprofit, YC W21) is hiring our first product lead (welcomekit.co) 
 https://upchieve.welcomekit.co/

SimpleHash (YC W22) is hiring back end engineers to build NFT APIs (ycombinator.com) 
 https://www.ycombinator.com/companies/simplehash/jobs/ID7qnlS-senior-backend-api-engineer

Padlet (YC W13) Is Hiring in Singapore to Build Software for a Good Education (padlet.jobs) 
 https://padlet.jobs

Reverie Labs (YC W18) is hiring software engineers to develop cancer therapies (reverielabs.com) 
 https://www.reverielabs.com/careers

GiveCampus (YC S15) is hiring a senior site reliability engineer (lever.co) 
 https://jobs.lever.co/givecampus/e7ae4695-e4d6-4ed0-b16f-f489a8f2382d

```
<br>


### The `page` scrapes the page you specify as an argument, from the hacker-news website.
<br>

 ```
 $ HackerGo page <integer> <all | integer>
 ```
<br>

The  `all` argument scrapes all news on the page of the website.
If you specify an integer less than 30 after the `page` command, you'll get the first "integer" news from the website.



Here's an example of the `page` command:

<br>

```

$ HackerGo page all 4

Fetching you those buzzing tech gists; This might take a while...
91.RISC-V is getting MSIs (stephenmarz.com) 
 https://blog.stephenmarz.com/2022/06/30/msi/

92.Some Macs are getting fewer updates than they used to. Here’s why it’s a problem (arstechnica.com) 
 https://arstechnica.com/gadgets/2022/07/some-macs-are-getting-fewer-updates-than-they-used-to-heres-why-its-a-problem/

93.Jenkins discloses dozens of zero-day bugs in multiple plugins (bleepingcomputer.com) 
 https://www.bleepingcomputer.com/news/security/jenkins-discloses-dozens-of-zero-day-bugs-in-multiple-plugins/

94.Long live software Easter eggs (acm.org) 
 https://queue.acm.org/detail.cfm?id=3534857

95.Vim 9.0 (vim.org) 
 https://www.vim.org/vim90.php


----------------------------------------------------------------more----------------------------------------------------
 ```
<br><br>

### The `random` command scrapes a random page from the hacker-news website and returns news based on the additional argument.
<br>

 ```
 $ HackerGo random <all | integer>
 ```
<br>
The  `all` argument scrapes all the news on the  page of the website.
If you specify an integer less than 30 after the `random` command, you'll get the first "integer" news from the website.



Here's an example of the `random` command:

```
HackerGo random 6

Fetching you those buzzing tech gists; This might take a while...

451.Ask HN: As a data scientist/engineer/analyst, what pisses you off in your job? 
 item?id=31928176

452.Economic Inequality Inflicts Real Biological Harm (2018) (scientificamerican.com) 
 https://www.scientificamerican.com/article/how-economic-inequality-inflicts-real-biological-harm/

453.Why This Crypto Crash Is Different (coindesk.com) 
 https://www.coindesk.com/layer2/futureofmoney/2022/06/29/why-this-crypto-crash-is-different/

454.NPM Registry is currently experiencing an outage (npmjs.org) 
 https://status.npmjs.org/incidents/6wr25yb0b2dd

455.I think Bitcoin has room for one more bubble (noahpinion.substack.com) 
 https://noahpinion.substack.com/p/i-think-bitcoin-has-room-for-one

456.Energy Dome’s approach to long-duration energy storage (canarymedia.com) 
 https://www.canarymedia.com/articles/long-duration-energy-storage/energy-dome-is-on-the-brink-of-a-long-duration-storage-breakthrough

```



## Tutorials

Here's [a tutorial I wrote](https://www.makeuseof.com/go-goquery-scrape-website/) on how to scrape websites using Goquery.
<br>
I'm working on a tutorial on how to build CLI apps in Go, you'll find it here and on my socials soon.

## Contributions
Pull requests and contributions are welcome on this repository and any other repositories on my page 💙
 
 