Feature: get ping
  In order to know the server health
  As an API user
  I need to be able to ping the server

  Scenario: does not allow POST method
    When I send "POST" request to "/ping"
    Then the response code should be 405
    And the response should match text:
      """
      Method Not Allowed
      """

  Scenario: should get pong response
    When I send "GET" request to "/ping"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "message": "pong"
      }
      """
