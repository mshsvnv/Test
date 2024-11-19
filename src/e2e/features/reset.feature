# language: en

Feature: User reset password with OTP

  Scenario: Successful login and sent OTP code for reset
    When User send "POST" request to "/reset_password"
    Then the response code on /reset_password should be 200
    And the response on /reset_password should match json:
        """
        {
          "msg": "OTP code to \"Reset Password\" was sent to your email"
        }
        """
    And User send "POST" request to "/reset_password/verify"
    Then the response code on /reset_password/verify should be 200
    And the response on /reset_password/verify should match json:
        """
        {
          "msg": "Your password has been already updated!"
        }
        """