// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Provides information about a bot alias.
type BotAliasMetadata struct {

	// The name of the bot to which the alias points.
	BotName *string

	// The version of the Amazon Lex bot to which the alias points.
	BotVersion *string

	// Checksum of the bot alias.
	Checksum *string

	// Settings that determine how Amazon Lex uses conversation logs for the alias.
	ConversationLogs *ConversationLogsResponse

	// The date that the bot alias was created.
	CreatedDate *time.Time

	// A description of the bot alias.
	Description *string

	// The date that the bot alias was updated. When you create a resource, the
	// creation date and last updated date are the same.
	LastUpdatedDate *time.Time

	// The name of the bot alias.
	Name *string
}

// Represents an association between an Amazon Lex bot and an external messaging
// platform.
type BotChannelAssociation struct {

	// An alias pointing to the specific version of the Amazon Lex bot to which this
	// association is being made.
	BotAlias *string

	// Provides information necessary to communicate with the messaging platform.
	BotConfiguration map[string]string

	// The name of the Amazon Lex bot to which this association is being made.
	// Currently, Amazon Lex supports associations with Facebook and Slack, and Twilio.
	BotName *string

	// The date that the association between the Amazon Lex bot and the channel was
	// created.
	CreatedDate *time.Time

	// A text description of the association you are creating.
	Description *string

	// If status is FAILED, Amazon Lex provides the reason that it failed to create the
	// association.
	FailureReason *string

	// The name of the association between the bot and the channel.
	Name *string

	// The status of the bot channel.
	//
	// * CREATED - The channel has been created and is
	// ready for use.
	//
	// * IN_PROGRESS - Channel creation is in progress.
	//
	// * FAILED -
	// There was an error creating the channel. For information about the reason for
	// the failure, see the failureReason field.
	Status ChannelStatus

	// Specifies the type of association by indicating the type of channel being
	// established between the Amazon Lex bot and the external messaging platform.
	Type ChannelType
}

// Provides information about a bot. .
type BotMetadata struct {

	// The date that the bot was created.
	CreatedDate *time.Time

	// A description of the bot.
	Description *string

	// The date that the bot was updated. When you create a bot, the creation date and
	// last updated date are the same.
	LastUpdatedDate *time.Time

	// The name of the bot.
	Name *string

	// The status of the bot.
	Status Status

	// The version of the bot. For a new bot, the version is always $LATEST.
	Version *string
}

// Provides metadata for a built-in intent.
type BuiltinIntentMetadata struct {

	// A unique identifier for the built-in intent. To find the signature for an
	// intent, see Standard Built-in Intents
	// (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	Signature *string

	// A list of identifiers for the locales that the intent supports.
	SupportedLocales []Locale
}

// Provides information about a slot used in a built-in intent.
type BuiltinIntentSlot struct {

	// A list of the slots defined for the intent.
	Name *string
}

// Provides information about a built in slot type.
type BuiltinSlotTypeMetadata struct {

	// A unique identifier for the built-in slot type. To find the signature for a slot
	// type, see Slot Type Reference
	// (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/slot-type-reference)
	// in the Alexa Skills Kit.
	Signature *string

	// A list of target locales for the slot.
	SupportedLocales []Locale
}

// Specifies a Lambda function that verifies requests to a bot or fulfills the
// user's request to a bot..
type CodeHook struct {

	// The version of the request-response that you want Amazon Lex to use to invoke
	// your Lambda function. For more information, see using-lambda.
	//
	// This member is required.
	MessageVersion *string

	// The Amazon Resource Name (ARN) of the Lambda function.
	//
	// This member is required.
	Uri *string
}

// Provides the settings needed for conversation logs.
type ConversationLogsRequest struct {

	// The Amazon Resource Name (ARN) of an IAM role with permission to write to your
	// CloudWatch Logs for text logs and your S3 bucket for audio logs. If audio
	// encryption is enabled, this role also provides access permission for the AWS KMS
	// key used for encrypting audio logs. For more information, see Creating an IAM
	// Role and Policy for Conversation Logs
	// (https://docs.aws.amazon.com/lex/latest/dg/conversation-logs-role-and-policy.html).
	//
	// This member is required.
	IamRoleArn *string

	// The settings for your conversation logs. You can log the conversation text,
	// conversation audio, or both.
	//
	// This member is required.
	LogSettings []LogSettingsRequest
}

// Contains information about conversation log settings.
type ConversationLogsResponse struct {

	// The Amazon Resource Name (ARN) of the IAM role used to write your logs to
	// CloudWatch Logs or an S3 bucket.
	IamRoleArn *string

	// The settings for your conversation logs. You can log text, audio, or both.
	LogSettings []LogSettingsResponse
}

// Each slot type can have a set of values. Each enumeration value represents a
// value the slot type can take. For example, a pizza ordering bot could have a
// slot type that specifies the type of crust that the pizza should have. The slot
// type could include the values
//
// * thick
//
// * thin
//
// * stuffed
type EnumerationValue struct {

	// The value of the slot type.
	//
	// This member is required.
	Value *string

	// Additional values related to the slot type value.
	Synonyms []string
}

// A prompt for additional activity after an intent is fulfilled. For example,
// after the OrderPizza intent is fulfilled, you might prompt the user to find out
// whether the user wants to order drinks.
type FollowUpPrompt struct {

	// Prompts for information from the user.
	//
	// This member is required.
	Prompt *Prompt

	// If the user answers "no" to the question defined in the prompt field, Amazon Lex
	// responds with this statement to acknowledge that the intent was canceled.
	//
	// This member is required.
	RejectionStatement *Statement
}

// Describes how the intent is fulfilled after the user provides all of the
// information required for the intent. You can provide a Lambda function to
// process the intent, or you can return the intent information to the client
// application. We recommend that you use a Lambda function so that the relevant
// logic lives in the Cloud and limit the client-side code primarily to
// presentation. If you need to update the logic, you only update the Lambda
// function; you don't need to upgrade your client application. Consider the
// following examples:
//
// * In a pizza ordering application, after the user provides
// all of the information for placing an order, you use a Lambda function to place
// an order with a pizzeria.
//
// * In a gaming application, when a user says "pick up
// a rock," this information must go back to the client application so that it can
// perform the operation and update the graphics. In this case, you want Amazon Lex
// to return the intent data to the client.
type FulfillmentActivity struct {

	// How the intent should be fulfilled, either by running a Lambda function or by
	// returning the slot data to the client application.
	//
	// This member is required.
	Type FulfillmentActivityType

	// A description of the Lambda function that is run to fulfill the intent.
	CodeHook *CodeHook
}

// The name of a context that must be active for an intent to be selected by Amazon
// Lex.
type InputContext struct {

	// The name of the context.
	//
	// This member is required.
	Name *string
}

// Identifies the specific version of an intent.
type Intent struct {

	// The name of the intent.
	//
	// This member is required.
	IntentName *string

	// The version of the intent.
	//
	// This member is required.
	IntentVersion *string
}

// Provides information about an intent.
type IntentMetadata struct {

	// The date that the intent was created.
	CreatedDate *time.Time

	// A description of the intent.
	Description *string

	// The date that the intent was updated. When you create an intent, the creation
	// date and last updated date are the same.
	LastUpdatedDate *time.Time

	// The name of the intent.
	Name *string

	// The version of the intent.
	Version *string
}

// Provides configuration information for the AMAZON.KendraSearchIntent intent.
// When you use this intent, Amazon Lex searches the specified Amazon Kendra index
// and returns documents from the index that match the user's utterance. For more
// information, see  AMAZON.KendraSearchIntent
// (http://docs.aws.amazon.com/lex/latest/dg/built-in-intent-kendra-search.html).
type KendraConfiguration struct {

	// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the
	// AMAZON.KendraSearchIntent intent to search. The index must be in the same
	// account and Region as the Amazon Lex bot. If the Amazon Kendra index does not
	// exist, you get an exception when you call the PutIntent operation.
	//
	// This member is required.
	KendraIndex *string

	// The Amazon Resource Name (ARN) of an IAM role that has permission to search the
	// Amazon Kendra index. The role must be in the same account and Region as the
	// Amazon Lex bot. If the role does not exist, you get an exception when you call
	// the PutIntent operation.
	//
	// This member is required.
	Role *string

	// A query filter that Amazon Lex sends to Amazon Kendra to filter the response
	// from the query. The filter is in the format defined by Amazon Kendra. For more
	// information, see Filtering queries
	// (http://docs.aws.amazon.com/kendra/latest/dg/filtering.html). You can override
	// this filter string with a new filter string at runtime.
	QueryFilterString *string
}

// Settings used to configure delivery mode and destination for conversation logs.
type LogSettingsRequest struct {

	// Where the logs will be delivered. Text logs are delivered to a CloudWatch Logs
	// log group. Audio logs are delivered to an S3 bucket.
	//
	// This member is required.
	Destination Destination

	// The type of logging to enable. Text logs are delivered to a CloudWatch Logs log
	// group. Audio logs are delivered to an S3 bucket.
	//
	// This member is required.
	LogType LogType

	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or S3 bucket
	// where the logs should be delivered.
	//
	// This member is required.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the AWS KMS customer managed key for
	// encrypting audio logs delivered to an S3 bucket. The key does not apply to
	// CloudWatch Logs and is optional for S3 buckets.
	KmsKeyArn *string
}

// The settings for conversation logs.
type LogSettingsResponse struct {

	// The destination where logs are delivered.
	Destination Destination

	// The Amazon Resource Name (ARN) of the key used to encrypt audio logs in an S3
	// bucket.
	KmsKeyArn *string

	// The type of logging that is enabled.
	LogType LogType

	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or S3 bucket
	// where the logs are delivered.
	ResourceArn *string

	// The resource prefix is the first part of the S3 object key within the S3 bucket
	// that you specified to contain audio logs. For CloudWatch Logs it is the prefix
	// of the log stream name within the log group that you specified.
	ResourcePrefix *string
}

// The message object that provides the message text and its type.
type Message struct {

	// The text of the message.
	//
	// This member is required.
	Content *string

	// The content type of the message string.
	//
	// This member is required.
	ContentType ContentType

	// Identifies the message group that the message belongs to. When a group is
	// assigned to a message, Amazon Lex returns one message from each group in the
	// response.
	GroupNumber *int32
}

// Provides information about alerts and warnings that Amazon Lex sends during a
// migration. The alerts include information about how to resolve the issue.
type MigrationAlert struct {

	// Additional details about the alert.
	Details []string

	// A message that describes why the alert was issued.
	Message *string

	// A link to the Amazon Lex documentation that describes how to resolve the alert.
	ReferenceURLs []string

	// The type of alert. There are two kinds of alerts:
	//
	// * ERROR - There was an issue
	// with the migration that can't be resolved. The migration stops.
	//
	// * WARN - There
	// was an issue with the migration that requires manual changes to the new Amazon
	// Lex V2 bot. The migration continues.
	Type MigrationAlertType
}

// Provides information about migrating a bot from Amazon Lex V1 to Amazon Lex V2.
type MigrationSummary struct {

	// The unique identifier that Amazon Lex assigned to the migration.
	MigrationId *string

	// The status of the operation. When the status is COMPLETE the bot is available in
	// Amazon Lex V2. There may be alerts and warnings that need to be resolved to
	// complete the migration.
	MigrationStatus MigrationStatus

	// The strategy used to conduct the migration.
	MigrationStrategy MigrationStrategy

	// The date and time that the migration started.
	MigrationTimestamp *time.Time

	// The locale of the Amazon Lex V1 bot that is the source of the migration.
	V1BotLocale Locale

	// The name of the Amazon Lex V1 bot that is the source of the migration.
	V1BotName *string

	// The version of the Amazon Lex V1 bot that is the source of the migration.
	V1BotVersion *string

	// The unique identifier of the Amazon Lex V2 that is the destination of the
	// migration.
	V2BotId *string

	// The IAM role that Amazon Lex uses to run the Amazon Lex V2 bot.
	V2BotRole *string
}

// The specification of an output context that is set when an intent is fulfilled.
type OutputContext struct {

	// The name of the context.
	//
	// This member is required.
	Name *string

	// The number of seconds that the context should be active after it is first sent
	// in a PostContent or PostText response. You can set the value between 5 and
	// 86,400 seconds (24 hours).
	//
	// This member is required.
	TimeToLiveInSeconds *int32

	// The number of conversation turns that the context should be active. A
	// conversation turn is one PostContent or PostText request and the corresponding
	// response from Amazon Lex.
	//
	// This member is required.
	TurnsToLive *int32
}

// Obtains information from the user. To define a prompt, provide one or more
// messages and specify the number of attempts to get information from the user. If
// you provide more than one message, Amazon Lex chooses one of the messages to use
// to prompt the user. For more information, see how-it-works.
type Prompt struct {

	// The number of times to prompt the user for information.
	//
	// This member is required.
	MaxAttempts *int32

	// An array of objects, each of which provides a message string and its type. You
	// can specify the message string in plain text or in Speech Synthesis Markup
	// Language (SSML).
	//
	// This member is required.
	Messages []Message

	// A response card. Amazon Lex uses this prompt at runtime, in the PostText API
	// response. It substitutes session attributes and slot values for placeholders in
	// the response card. For more information, see ex-resp-card.
	ResponseCard *string
}

// Describes the resource that refers to the resource that you are attempting to
// delete. This object is returned as part of the ResourceInUseException exception.
type ResourceReference struct {

	// The name of the resource that is using the resource that you are trying to
	// delete.
	Name *string

	// The version of the resource that is using the resource that you are trying to
	// delete.
	Version *string
}

// Identifies the version of a specific slot.
type Slot struct {

	// The name of the slot.
	//
	// This member is required.
	Name *string

	// Specifies whether the slot is required or optional.
	//
	// This member is required.
	SlotConstraint SlotConstraint

	// A list of default values for the slot. Default values are used when Amazon Lex
	// hasn't determined a value for a slot. You can specify default values from
	// context variables, session attributes, and defined values.
	DefaultValueSpec *SlotDefaultValueSpec

	// A description of the slot.
	Description *string

	// Determines whether a slot is obfuscated in conversation logs and stored
	// utterances. When you obfuscate a slot, the value is replaced by the slot name in
	// curly braces ({}). For example, if the slot name is "full_name", obfuscated
	// values are replaced with "{full_name}". For more information, see  Slot
	// Obfuscation  (https://docs.aws.amazon.com/lex/latest/dg/how-obfuscate.html).
	ObfuscationSetting ObfuscationSetting

	// Directs Amazon Lex the order in which to elicit this slot value from the user.
	// For example, if the intent has two slots with priorities 1 and 2, AWS Amazon Lex
	// first elicits a value for the slot with priority 1. If multiple slots share the
	// same priority, the order in which Amazon Lex elicits values is arbitrary.
	Priority *int32

	// A set of possible responses for the slot type used by text-based clients. A user
	// chooses an option from the response card, instead of using text to reply.
	ResponseCard *string

	// If you know a specific pattern with which users might respond to an Amazon Lex
	// request for a slot value, you can provide those utterances to improve accuracy.
	// This is optional. In most cases, Amazon Lex is capable of understanding user
	// utterances.
	SampleUtterances []string

	// The type of the slot, either a custom slot type that you defined or one of the
	// built-in slot types.
	SlotType *string

	// The version of the slot type.
	SlotTypeVersion *string

	// The prompt that Amazon Lex uses to elicit the slot value from the user.
	ValueElicitationPrompt *Prompt
}

// A default value for a slot.
type SlotDefaultValue struct {

	// The default value for the slot. You can specify one of the following:
	//
	// *
	// #context-name.slot-name - The slot value "slot-name" in the context
	// "context-name."
	//
	// * {attribute} - The slot value of the session attribute
	// "attribute."
	//
	// * 'value' - The discrete value "value."
	//
	// This member is required.
	DefaultValue *string
}

// Contains the default values for a slot. Default values are used when Amazon Lex
// hasn't determined a value for a slot.
type SlotDefaultValueSpec struct {

	// The default values for a slot. You can specify more than one default. For
	// example, you can specify a default value to use from a matching context
	// variable, a session attribute, or a fixed value. The default value chosen is
	// selected based on the order that you specify them in the list. For example, if
	// you specify a context variable and a fixed value in that order, Amazon Lex uses
	// the context variable if it is available, else it uses the fixed value.
	//
	// This member is required.
	DefaultValueList []SlotDefaultValue
}

// Provides configuration information for a slot type.
type SlotTypeConfiguration struct {

	// A regular expression used to validate the value of a slot.
	RegexConfiguration *SlotTypeRegexConfiguration
}

// Provides information about a slot type..
type SlotTypeMetadata struct {

	// The date that the slot type was created.
	CreatedDate *time.Time

	// A description of the slot type.
	Description *string

	// The date that the slot type was updated. When you create a resource, the
	// creation date and last updated date are the same.
	LastUpdatedDate *time.Time

	// The name of the slot type.
	Name *string

	// The version of the slot type.
	Version *string
}

// Provides a regular expression used to validate the value of a slot.
type SlotTypeRegexConfiguration struct {

	// A regular expression used to validate the value of a slot. Use a standard
	// regular expression. Amazon Lex supports the following characters in the regular
	// expression:
	//
	// * A-Z, a-z
	//
	// * 0-9
	//
	// * Unicode characters ("\ u")
	//
	// Represent Unicode
	// characters with four digits, for example "\u0041" or "\u005A". The following
	// regular expression operators are not supported:
	//
	// * Infinite repeaters: *, +, or
	// {x,} with no upper bound.
	//
	// * Wild card (.)
	//
	// This member is required.
	Pattern *string
}

// A collection of messages that convey information to the user. At runtime, Amazon
// Lex selects the message to convey.
type Statement struct {

	// A collection of message objects.
	//
	// This member is required.
	Messages []Message

	// At runtime, if the client is using the PostText
	// (http://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostText.html) API, Amazon
	// Lex includes the response card in the response. It substitutes all of the
	// session attributes and slot values for placeholders in the response card.
	ResponseCard *string
}

// A list of key/value pairs that identify a bot, bot alias, or bot channel. Tag
// keys and values can consist of Unicode letters, digits, white space, and any of
// the following symbols: _ . : / = + - @.
type Tag struct {

	// The key for the tag. Keys are not case-sensitive and must be unique.
	//
	// This member is required.
	Key *string

	// The value associated with a key. The value may be an empty string but it can't
	// be null.
	//
	// This member is required.
	Value *string
}

// Provides information about a single utterance that was made to your bot.
type UtteranceData struct {

	// The number of times that the utterance was processed.
	Count *int32

	// The total number of individuals that used the utterance.
	DistinctUsers *int32

	// The date that the utterance was first recorded.
	FirstUtteredDate *time.Time

	// The date that the utterance was last recorded.
	LastUtteredDate *time.Time

	// The text that was entered by the user or the text representation of an audio
	// clip.
	UtteranceString *string
}

// Provides a list of utterances that have been made to a specific version of your
// bot. The list contains a maximum of 100 utterances.
type UtteranceList struct {

	// The version of the bot that processed the list.
	BotVersion *string

	// One or more UtteranceData objects that contain information about the utterances
	// that have been made to a bot. The maximum number of object is 100.
	Utterances []UtteranceData
}
