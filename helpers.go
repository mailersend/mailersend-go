package mailersend

// Activity-related event constants
const (
	EventActivitySent            = "activity.sent"             // Fired when your email is sent from our sending servers.
	EventActivityDelivered       = "activity.delivered"        // Fired when your email is successfully delivered with no errors.
	EventActivitySoftBounced     = "activity.soft_bounced"     // Fired when your email is not delivered because it soft bounced.
	EventActivityHardBounced     = "activity.hard_bounced"     // Fired when your email is not delivered.
	EventActivityDeferred        = "activity.deferred"         // Fired when your email is temporarily deferred by the receiving server.
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

// Email verification-related event constants
const (
	EventEmailSingleVerified = "email_single.verified" // Fired when a single email address has been successfully verified.
	EventEmailListVerified   = "email_list.verified"   // Fired when an email list has been successfully verified.
)

// Bulk email-related event constants
const (
	EventBulkEmailCompleted = "bulk_email.completed" // Fired when a bulk email sending has been completed.
)

// Recipient-related event constants
const (
	EventRecipientOnHoldAdded   = "recipient.on_hold_added"   // Fired when a recipient is added to the on-hold list.
	EventRecipientOnHoldRemoved = "recipient.on_hold_removed" // Fired when a recipient is removed from the on-hold list.
)
