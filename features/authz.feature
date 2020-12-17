Feature: only admins
  In order to provide proper security
  As an operator
  I need to ensure a user is an admin

  Scenario: validate authorization
    Given There is a json file
    Given There is a rego file
    When I validate the files
    Then All users should be admins