package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type AocSlackbotStackProps struct {
	awscdk.StackProps
}

func NewAocSlackbotStack(scope constructs.Construct, id string, props *AocSlackbotStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	leaderboardEnvVars := map[string]*string{
		"SESSION_COOKIE":     aws.String(os.Getenv("SESSION_COOKIE")),
		"AOC_LEADERBOARD_ID": aws.String(os.Getenv("AOC_LEADERBOARD_ID")),
		"SLACK_WEBHOOK":      aws.String(os.Getenv("SLACK_WEBHOOK")),
		"TIMEZONE":           aws.String(os.Getenv("TIMEZONE")),
	}

	leaderboardLambda := awslambda.NewFunction(stack, jsii.String("aocLeaderboardFunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("aocLeaderboardFunction"),
		Code:         awslambda.Code_FromAsset(jsii.String("../src/aocLeaderboardFunction"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("index.lambdaHandler"),
		Runtime:      awslambda.Runtime_NODEJS_14_X(),
		Environment:  &leaderboardEnvVars,
	})

	leaderboardCronRule := awsevents.NewRule(stack, jsii.String("aocLeaderboardCron"), &awsevents.RuleProps{
		RuleName: jsii.String("aocLeaderboardCron"),
		Schedule: awsevents.Schedule_Expression(jsii.String("cron(0 17 * * ? *)")),
		Enabled: jsii.Bool(true),
	})

	leaderboardCronRule.AddTarget(awseventstargets.NewLambdaFunction(leaderboardLambda, &awseventstargets.LambdaFunctionProps{}))

	solutionEnvVars := map[string]*string{
		"SLACK_WEBHOOK": aws.String(os.Getenv("SLACK_WEBHOOK")),
		"TIMEZONE":      aws.String(os.Getenv("TIMEZONE")),
	}

	solutionLambda := awslambda.NewFunction(stack, jsii.String("dailySolutionThreadFunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("dailySolutionThreadFunction"),
		Code:         awslambda.Code_FromAsset(jsii.String("../src/dailySolutionThreadFunction"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("index.lambdaHandler"),
		Runtime:      awslambda.Runtime_NODEJS_14_X(),
		Environment:  &solutionEnvVars,
	})

	solutionCronRule := awsevents.NewRule(stack, jsii.String("aocSolutionCron"), &awsevents.RuleProps{
		RuleName: jsii.String("aocSolutionCron"),
		Schedule: awsevents.Schedule_Expression(jsii.String("cron(0 5 * * ? *)")),
		Enabled: jsii.Bool(true),
	})

	solutionCronRule.AddTarget(awseventstargets.NewLambdaFunction(solutionLambda, &awseventstargets.LambdaFunctionProps{}))

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewAocSlackbotStack(app, "AocSlackbotStack", &AocSlackbotStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
