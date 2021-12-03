const axios = require('axios')
const moment = require('moment-timezone')

const aocLeaderboardURL = `https://adventofcode.com/2021/leaderboard/private/view/${process.env.AOC_LEADERBOARD_ID}.json`

exports.lambdaHandler = async () => {
    let now = moment().tz(process.env.TIMEZONE)
    if (now.date() > 24 || (now.month() + 1) !== 12) return {}

    let { data } = await axios.get(aocLeaderboardURL, {
        headers: {
            Cookie: `session=${process.env.SESSION_COOKIE};`
        }
    })

    let participants = []

    Object.keys(data.members).forEach(memberKey => {
        let member = data.members[memberKey]

        participants.push({
            name: member.name,
            localScore: member.local_score,
            totalStars: member.stars,
            daysCompleted: Object.keys(member.completion_day_level).length
        })
    })

    participants = participants.sort((x, y) => {
        if (x.localScore > y.localScore) {
            return -1;
        }
        if (x.localScore < y.localScore) {
            return 1;
        }
        return 0;
    }).filter(participant => participant.daysCompleted > 0)

    let payload = formatSlackPayload(participants)

    await axios.post(process.env.SLACK_WEBHOOK, payload, {
        headers: {
            "Content-type": "application/json"
        }
    })

    return { statusCode: 200, body: JSON.stringify(participants) }
}

function formatSlackPayload(participants) {
   let payload = {
       blocks: [
           {
               "type": "section",
               "text": {
                   "type": "mrkdwn",
                   "text": `*Advent of Code Leaderboard, ${moment().format('dddd MMM Do YYYY')} :trophy:*`
               }
           },
           {
               "type": "divider"
           },
       ]
   }

   participants.forEach((part, index) => {
       payload.blocks.push({
           "type": "section",
           "text": {
               "type": "mrkdwn",
               "text": `*${index + 1}. ${part.name} (${part.localScore})* — ${part.totalStars} x :star:, ${part.daysCompleted}/25 days complete`
           }
       })
   })

    return payload
}