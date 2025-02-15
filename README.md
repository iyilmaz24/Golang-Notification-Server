# Golang Notification Server
This project implements a notification server using Go, designed to send email and SMS notifications based on various triggers. It integrates with AWS Systems Manager Parameter Store for configuration and incorporates robust error handling and logging.

## Features
*   **Multiple Notification Types:** Supports daily analytics reports, internal alerts, cron job notifications, and administrative notifications.
*   **Email and SMS Delivery:** Sends notifications via email and SMS, with fallback mechanisms for email delivery failures.
*   **AWS Systems Manager Integration:** Configuration is managed via AWS SSM Parameter Store for secure storage of sensitive information.
*   **Comprehensive Logging:** Detailed logging of successful and failed notifications for debugging and monitoring.
*   **CORS Middleware:** Includes CORS middleware to handle requests from different origins securely.
*   **Error Handling:** Robust error handling to gracefully manage issues during notification delivery.
*   **Customizable Email Templates:** Uses customizable HTML templates for email notifications.

## Usage
The server exposes several endpoints for triggering notifications:

*   `/dailyAnalyticsReport` (POST): Sends a daily analytics report.  Requires a JSON payload containing report details.
*   `/internalNotification` (POST): Sends an internal alert notification. Requires a JSON payload containing notification details.
*   `/cronNotification` (POST): Logs a cron job execution. Requires a JSON payload containing notification details.
*   `/adminNotification` (POST): Allows sending test notifications from outside the VPC.  Requires a JSON payload containing notification details.

## Technologies Used
*   **Go:** The programming language used for the server.
*   **AWS SDK for Go v2:** Used for interacting with AWS Systems Manager Parameter Store.
*   **net/smtp:** Go's built-in package for sending emails via SMTP.
*   **log:** Go's built-in package for logging.
*   **encoding/json:** Go's built-in package for JSON encoding and decoding.
*   **net/http:** Go's built-in package for HTTP handling.
*   **github.com/pkg/errors:**  For error wrapping and handling.

## Configuration
Configuration is managed through AWS Systems Manager Parameter Store.  The following parameters are used:

*   `/backend/internal/admin-cors-origin`: Comma-separated list of allowed CORS origins (StringList).
*   `/backend/internal/db_dsn`: Database connection string (SecureString).
*   `/backend/ports/admin`: Server port (String, defaults to ":3100").
*   `/backend/internal/admin-password`: Admin password (SecureString).
*   `/backend/internal/gmail-app-password`: Gmail App Password (SecureString).
*   `/backend/internal/gmail-address`: Gmail address (String).
*   `/backend/internal/alert-phone-numbers`: Comma-separated list of alert phone numbers (StringList).

## Dependencies
The project dependencies are listed in `go.mod`.

*README.md was made with [Etchr](https://etchr.dev)*