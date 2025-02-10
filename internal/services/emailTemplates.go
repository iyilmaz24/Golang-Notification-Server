package services

import (
	"fmt"
	"strings"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type failureInfo struct {
	FailedAttempts int
	ErrorTime      string
	ErrorCode      string
	ErrorMessage   string
}

func getEmailContent(notificationObj models.Notification) []byte {
	urgencyColor := map[string]string{
		"high":   "#DC2626",
		"medium": "#F59E0B",
		"low":    "#10B981",
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
    <div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <!-- Header -->
        <div style="background-color: #1E40AF; padding: 20px; border-radius: 8px 8px 0 0; margin: -20px -20px 20px -20px;">
            <h1 style="color: #ffffff; margin: 0; font-size: 24px;">%s</h1>
            <p style="color: #E5E7EB; margin: 10px 0 0 0;">%s</p>
        </div>

        <!-- Notification Details -->
        <div style="margin-bottom: 20px;">
            <div style="background-color: %s; color: white; display: inline-block; padding: 5px 10px; border-radius: 4px; font-size: 14px; margin-bottom: 15px;">
                %s Priority
            </div>
            
            <table style="width: 100%%; border-collapse: collapse; margin-bottom: 20px;">
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280; width: 140px;">Source</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Type</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Time</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s %s (%s)</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Status</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">ID</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
            </table>
        </div>

        <!-- Message Content -->
        <div style="background-color: #F3F4F6; padding: 20px; border-radius: 4px; margin-bottom: 20px;">
            <p style="margin: 0; line-height: 1.6;">%s</p>
        </div>

        <!-- Footer -->
        <div style="text-align: center; color: #6B7280; font-size: 12px; margin-top: 20px; padding-top: 20px; border-top: 1px solid #e5e7eb;">
            <p>This is an automated notification from your monitoring system.</p>
        </div>
    </div>
</body>
</html>`,
		notificationObj.NotificationSubject,
		notificationObj.NotificationType,
		urgencyColor[strings.ToLower(notificationObj.NotificationUrgency)],
		strings.Title(strings.ToLower(notificationObj.NotificationUrgency)),
		notificationObj.NotificationSource,
		notificationObj.NotificationType,
		notificationObj.NotificationTime,
		notificationObj.NotificationDate,
		notificationObj.NotificationTimezone,
		notificationObj.NotificationStatus,
		notificationObj.NotificationID,
		notificationObj.NotificationMessage,
	)

	return []byte(fmt.Sprintf(
		"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n%s",
		notificationObj.NotificationSubject,
		html,
	))
}

func getDailyAnalyticsEmailContent(analytics models.DailyAnalytics) []byte {
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
    <div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <!-- Header -->
        <div style="background-color: #1E40AF; padding: 20px; border-radius: 8px 8px 0 0; margin: -20px -20px 20px -20px;">
            <h1 style="color: #ffffff; margin: 0; font-size: 24px;">%s</h1>
            <p style="color: #E5E7EB; margin: 10px 0 0 0;">Daily Analytics Report</p>
        </div>

        <!-- Analytics Details -->
        <div style="margin-bottom: 20px;">
            <table style="width: 100%%; border-collapse: collapse; margin-bottom: 20px;">
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280; width: 140px;">Source</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Type</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Time</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s %s (%s)</td>
                </tr>
            </table>
        </div>

        <!-- Message Content -->
        <div style="background-color: #F3F4F6; padding: 20px; border-radius: 4px; margin-bottom: 20px;">
            <p style="margin: 0; line-height: 1.6;">%s</p>
        </div>

        <!-- Footer -->
        <div style="text-align: center; color: #6B7280; font-size: 12px; margin-top: 20px; padding-top: 20px; border-top: 1px solid #e5e7eb;">
            <p>This is an automated analytics report.</p>
        </div>
    </div>
</body>
</html>`,
		analytics.NotificationSubject,
		analytics.NotificationSource,
		analytics.NotificationType,
		analytics.NotificationTime,
		analytics.NotificationDate,
		analytics.NotificationTimezone,
		analytics.NotificationMessage,
	)

	return []byte(fmt.Sprintf(
		"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n%s",
		analytics.NotificationSubject,
		html,
	))
}

func getEmailServiceFailureContent(failureInfo failureInfo) []byte {
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
    <div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <div style="background-color: #991B1B; padding: 20px; border-radius: 8px 8px 0 0; margin: -20px -20px 20px -20px;">
            <h1 style="color: #ffffff; margin: 0; font-size: 24px;">⚠️ Email Service Failure</h1>
            <p style="color: #E5E7EB; margin: 10px 0 0 0;">System Alert</p>
        </div>

        <div style="margin-bottom: 20px;">
            <div style="background-color: #DC2626; color: white; display: inline-block; padding: 5px 10px; border-radius: 4px; font-size: 14px; margin-bottom: 15px;">
                High Priority
            </div>
            
            <table style="width: 100%%; border-collapse: collapse; margin-bottom: 20px;">
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280; width: 140px;">Service</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">Email Notification System</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Status</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #DC2626; font-weight: bold;">FAILED</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Error Time</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Failed Attempts</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%d</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Error Code</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
            </table>
        </div>

        <div style="background-color: #FEE2E2; padding: 20px; border-radius: 4px; margin-bottom: 20px; border: 1px solid #DC2626;">
            <p style="margin: 0; line-height: 1.6; color: #991B1B;">
                %s
                <br><br>
                Possible causes include:
                <br><br>
                • Invalid SMTP credentials
                <br>
                • Gmail security settings changes
                <br>
                • Network connectivity issues
                <br><br>
                Please check SMTP configuration and Gmail App Password settings.
            </p>
        </div>

        <div style="text-align: center; color: #6B7280; font-size: 12px; margin-top: 20px; padding-top: 20px; border-top: 1px solid #e5e7eb;">
            <p>This is a critical system alert. Immediate attention required.</p>
        </div>
    </div>
</body>
</html>`,
		failureInfo.ErrorTime,
		failureInfo.FailedAttempts,
		failureInfo.ErrorCode,
		failureInfo.ErrorMessage,
	)

	return []byte(fmt.Sprintf(
		"Subject: ⚠️ ALERT: Email Service Failure\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n%s",
		html,
	))
}

func getSMSServiceFailureContent(failureInfo failureInfo) []byte {
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
    <div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <div style="background-color: #991B1B; padding: 20px; border-radius: 8px 8px 0 0; margin: -20px -20px 20px -20px;">
            <h1 style="color: #ffffff; margin: 0; font-size: 24px;">⚠️ SMS Service Failure</h1>
            <p style="color: #E5E7EB; margin: 10px 0 0 0;">System Alert</p>
        </div>

        <div style="margin-bottom: 20px;">
            <div style="background-color: #DC2626; color: white; display: inline-block; padding: 5px 10px; border-radius: 4px; font-size: 14px; margin-bottom: 15px;">
                High Priority
            </div>
            
            <table style="width: 100%%; border-collapse: collapse; margin-bottom: 20px;">
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280; width: 140px;">Service</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">SMS Notification System</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Status</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #DC2626; font-weight: bold;">FAILED</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Error Time</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Failed Attempts</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%d</td>
                </tr>
                <tr>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb; color: #6B7280;">Error Code</td>
                    <td style="padding: 8px; border-bottom: 1px solid #e5e7eb;">%s</td>
                </tr>
            </table>
        </div>

        <div style="background-color: #FEE2E2; padding: 20px; border-radius: 4px; margin-bottom: 20px; border: 1px solid #DC2626;">
            <p style="margin: 0; line-height: 1.6; color: #991B1B;">
                %s
                <br><br>
                Possible causes include:
                <br><br>
                • Invalid API credentials
                <br>
                • API rate limit exceeded
                <br>
                • Service provider outage
                <br><br>
                Please check SMS service configuration and API status.
            </p>
        </div>

        <div style="text-align: center; color: #6B7280; font-size: 12px; margin-top: 20px; padding-top: 20px; border-top: 1px solid #e5e7eb;">
            <p>This is a critical system alert. Immediate attention required.</p>
        </div>
    </div>
</body>
</html>`,
		failureInfo.ErrorTime,
		failureInfo.FailedAttempts,
		failureInfo.ErrorCode,
		failureInfo.ErrorMessage,
	)

	return []byte(fmt.Sprintf(
		"Subject: ⚠️ ALERT: SMS Service Failure\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n%s",
		html,
	))
}
