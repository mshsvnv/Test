# language: en

Feature: User Login with OTP

  Scenario: Successful login and sent OTP code for login
    When User send "POST" request to "/login"
    Then the response code on /login should be 200
    And the response on /login should match json:
        """
        {
          "msg": "OTP code to \"Login\" was sent to your email"
        }
        """
    And User send "POST" request to "/login/verify"
    Then the response code on /login/verify should be 200
    And the response on /login/verify should match json:
        """
        {
          "access_token": "your_access_token"
        }
        """