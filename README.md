# Advent of Code Slackbot

This repo has everything you need to spin up a serverless slackbot to play along with the [Advent of Code](https://adventofcode.com/) challenge on Slack.

When deployed to AWS with the proper Slack information, this bot will do 2 things 
1. Post a "Solutions Thread" message each day of the challenge immediately after the challenge becomes available
2. Post a replica of a Private Leaderboard to Slack 12 hours after the previous challenge is made available

## Setup

1. Clone repo
2. Create a file located at `./cdk/.env` that looks likes this (more info on the variables below)

```
SLACK_WEBHOOK=
AOC_LEADERBOARD_ID=
SESSION_COOKIE=
TIMEZONE=
```
3. Make sure you've installed & bootstrapped [aws cdk](https://docs.aws.amazon.com/cdk/latest/guide/getting_started.html#getting_started_install)
4. Run `cd cdk && cdk deploy`

That's it! Happy coding!

### Env Variables

| Variable Name | Info |
| --- | --- |
| `SLACK_WEBHOOK` | Can be found after creating a new Slack App and enabling webhooks |
| `AOC_LEADERBOARD_ID` | This is **not** the ID you use to join the private leaderboard, rather, it's the ID in the URL when you are looking at the private leaderboard. Something like `https://adventofcode.com/2021/leaderboard/private/view/{{ ID LOCATED HERE }}` |
| `SESSION_COOKIE` | This is how the creator of Advent of Code instructs one to programmatically hit his website's API. Inspect your browser when logged in to find it. They are valid for about a month, so you should only need to do this once.
| `TIMEZONE` | Pick a [`momentjs` timezone](https://momentjs.com/timezone/)
