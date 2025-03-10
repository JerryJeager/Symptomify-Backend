package emails

import (
	"fmt"

	"github.com/JerryJeager/Symptomify-Backend/utils"
)

func CreateUserMail(name, email, otp string) string {
	resetLink := fmt.Sprintf(`/verify-email/%s?otp=%s`, utils.GetClientBaseUrl(), email, otp)
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Reset Your Password - symptomify</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
        .container {
            width: 100%%;
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
            text-align: center;
        }
        .header {
            font-size: 24px;
            font-weight: bold;
            color: #2C3650;
        }
        .otp {
            font-size: 28px;
            font-weight: bold;
            color: #2C3650;
            background-color: #f0f0f0;
            padding: 10px;
            border-radius: 5px;
            display: inline-block;
            margin: 15px 0;
        }
        .message {
            font-size: 16px;
            color: #555;
            margin: 20px 0;
        }
        .footer {
            font-size: 14px;
            color: #888;
            margin-top: 20px;
        }
        .btn {
            display: inline-block;
            background-color: #2C3650;
            color: #ffffff;
            text-decoration: none;
            padding: 10px 20px;
            border-radius: 5px;
            font-size: 16px;
            margin-top: 20px;
        }
        .btn:hover {
            background-color: #1d263b;
        }
    </style>
</head>
<body>
    <div class="container">
        <p class="header">Verify your Email</p>
        <p class="message">Hey <strong>%s</strong>,</p>
        <p class="message">We received a request to reset your password. Use the OTP below to verify your account:</p>
        <p class="otp">%s</p>
        <p class="message">If you did not request this, you can safely ignore this email.</p>
        <a href="%s" class="btn">Verify Email</a>
        <p class="footer">If you need help, contact our support team at <a href="mailto:support@symptomify.com">support@symptomify.com</a></p>
    </div>
</body>
</html>
`, name, otp, resetLink)
}
