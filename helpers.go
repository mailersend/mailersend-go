package mailersend

// Activity-related event constants
const (
	EventActivitySent            = "activity.sent"             // Fired when your email is sent from our sending servers.
	EventActivityDelivered       = "activity.delivered"        // Fired when your email is successfully delivered with no errors.
	EventActivitySoftBounced     = "activity.soft_bounced"     // Fired when your email is not delivered because it soft bounced.
	EventActivityHardBounced     = "activity.hard_bounced"     // Fired when your email is not delivered.
	EventActivityOpened          = "activity.opened"           // Fired when the recipient receives your email and opens it.
	EventActivityOpenedUnique    = "activity.opened_unique"    // Fired when the recipient receives your email and opens it only for the first time.
	EventActivityClicked         = "activity.clicked"          // Fired when the recipient clicks a link in your email.
	EventActivityClickedUnique   = "activity.clicked_unique"   // Fired when the recipient clicks a link in your email only for the first time.
	EventActivityUnsubscribed    = "activity.unsubscribed"     // Fired when the recipient unsubscribes from your emails.
	EventActivitySpamComplaint   = "activity.spam_complaint"   // Fired when the recipient marks your emails as spam or junk.
	EventActivitySurveyOpened    = "activity.survey_opened"    // Fired when the recipient opens an email containing a survey for the first time.
	EventActivitySurveySubmitted = "activity.survey_submitted" // Fired when the recipient answers all available questions in a survey-based email or after an idle time of 30 minutes.
)

// Sender identity-related event constants
const (
	EventSenderIdentityVerified = "sender_identity.verified" // Fired when the sender identity has been successfully verified.
)

// Maintenance-related event constants
const (
	EventMaintenanceStart = "maintenance.start" // Fired when the maintenance period begins.
	EventMaintenanceEnd   = "maintenance.end"   // Fired when the maintenance period ends.
)

// Inbound forward-related event constants
const (
	EventInboundForwardFailed = "inbound_forward.failed" // Fired when an inbound message fails to forward.
)
