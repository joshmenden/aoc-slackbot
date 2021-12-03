const axios = require('axios')
const moment = require('moment-timezone')

exports.lambdaHandler = async () => {
    let now = moment().tz(process.env.TIMEZONE)
    if (now.date() > 24 || (now.month() + 1) !== 12) return {}

    let payload = {
        "blocks": [
            {
                "type": "section",
                "text": {
                    "type": "mrkdwn",
                    "text": `*Day ${now.date() + 1}* solutions thread :point_down:`
                }
            }
        ]
    }

    await axios.post(process.env.SLACK_WEBHOOK, payload, {
        headers: {
            "Content-type": "application/json"
        }
    })

    return { statusCode: 200, body: JSON.stringify({}) }
}
