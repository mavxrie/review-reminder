# review-reminder

A small tool to post PR waiting to get reviewed into slack.

## Sample use case

```sh
$ export GITHUB_TOKEN=your-github-token
$ export SLACK_TOKEN=app-slack-token
$ ./review-reminder -how-many 3
Hello team. There are currently 5 PR currently awaiting reviews:
AvxSre/blabla1#192 by chocapic: AVXSRE-292: Add useless stuff updated 7 days ago - https://github.com/AvxSre/blabla1/pull/192
AvxSre/coolstuff#205 by mielpops: AVXSRE-333: Removing milk from bowl updated 4 days ago - https://github.com/AvxSre/coolstuff/pull/205
AvxSre/blabla1#209 by rrey-aviatrix: AVXSRE-374: Muting the bot for saying things updated 3 days ago - https://github.com/AvxSre/blabla1/pull/209
... and other can be viewed https://github.com/pulls/review-requested
```

Using `slack-target` to send this to slack will create links.

## Usage

```sh
$ go get -v
$ go build
$ ./review-reminder -help
Usage of ./review-reminder:
  -github-org string
        The github organization to fetch PR from (default "avxSre")
  -github-token string
        The github token to use (or use GITHUB_TOKEN)
  -greeting string
        A custom greeting, only for jokes.
  -how-many int
        Number of PR to list. (default 3)
  -jira-browse-url string
        Your Jira browse url (default "https://jira.com/browse/")
  -slack-target string
        The slack target (channel, user) to send notification to.
  -slack-token string
        The slack token to use (or use SLACK_TOKEN)
```
