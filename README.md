# Golang Notification Server
This project is a notification server built with Go. It sends email and SMS notifications triggered by various events.  It uses AWS Systems Manager Parameter Store for secure configuration and includes robust error handling and logging.

## Features
* Multiple Notification Types: Sends daily analytics reports, internal alerts, cron job notifications, and administrative notifications.
* Email and SMS Delivery: Delivers notifications via email and SMS, with fallback mechanisms.
* Secure Configuration: Uses AWS SSM Parameter Store for secure configuration management.
* CORS Enabled: Includes CORS middleware for secure cross-origin requests.
* Implements rate limiting to prevent abuse.  Uses a token bucket algorithm to limit requests per IP address to 15 per minute.
* Includes API key authentication middleware to protect endpoints.  Requires the `X-API-Key` header to match the admin password configured in AWS SSM Parameter Store.  Uses constant-time comparison to mitigate timing attacks.
* Customizable Email Templates: Supports customizable HTML email templates.
* Robust Logging: Provides detailed logs for monitoring and debugging.
* Comprehensive Error Handling: Gracefully manages errors during notification delivery.



## Usage
The server provides several endpoints for triggering notifications:

* `/dailyAnalyticsReport` (POST): Sends a daily analytics report. Requires a JSON payload with report details.
* `/internalNotification` (POST): Sends an internal alert notification. Requires a JSON payload with notification details.
* `/cronNotification` (POST): Logs a cron job execution. Requires a JSON payload with notification details.
* `/adminNotification` (POST): Sends test notifications from outside the VPC. Requires a JSON payload with notification details.

## Technologies Used
* **Go:** The primary programming language.
* **AWS SDK for Go v2:** Interacts with AWS Systems Manager Parameter Store.
* **net/smtp:** Go's built-in package for sending emails via SMTP.
* **log:** Go's built-in package for logging.
* **encoding/json:** Go's built-in package for JSON encoding and decoding.
* **net/http:** Go's built-in package for HTTP handling.
* **github.com/pkg/errors:** For enhanced error handling.
* **AWS Systems Manager Parameter Store:** Securely stores configuration parameters.

## Configuration
Configuration is managed through AWS Systems Manager Parameter Store.  The following parameters are required:

* `/backend/internal/admin-cors-origin`: Comma-separated list of allowed CORS origins (StringList).
* `/backend/internal/db_dsn`: Database connection string (SecureString).
* `/backend/ports/admin`: Server port (String, defaults to ":3100").
* `/backend/internal/admin-password`: Admin password (SecureString).
* `/backend/internal/gmail-app-password`: Gmail App Password (SecureString).
* `/backend/internal/gmail-address`: Gmail address (String).
* `/backend/internal/alert-phone-numbers`: Comma-separated list of alert phone numbers (StringList).

*README.md was made with [Etchr](https://etchr.dev)*