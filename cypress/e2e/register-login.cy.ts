import { User } from "src/app/user"
import 'cypress-plugin-api'
describe('template spec', () => {

  it('should register and log in as a user', () => {

    cy.visit('http://localhost:4200/')

    cy.contains('Register').click()

    cy.get('#firstName').type("First Name")
    cy.get('#lastName').type("Last Name")
    cy.get('#email').type("test@email.com")
    cy.get('#pass').type("pass")
    cy.get('#confPass').type("pass")

    cy.contains('Submit').click()
  
    
    cy.request('POST', 'register',
      {first_name: "First Name", last_name: "Last Name", email: "test@email.com", password: "pass", account_type: "User"})
  
    })
})