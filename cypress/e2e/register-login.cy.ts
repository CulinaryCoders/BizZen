import { User } from "src/app/user"

describe('template spec', () => {

  it('should register and log in as a user', () => {

      cy.visit('http://localhost:4200/')

      let testUser:User = new User("1", "First Name", "Last Name", "test@email.com", "pass", "User", [])

      cy.intercept('POST', 'register', {
        user:testUser
      }).as('registerUser')

      cy.intercept('POST', 'login', {
        fixture:'mock-services.json'
      }).as('loginUser')

      cy.contains('Register').click()

      cy.get('#firstName').type(testUser.first_name)
      cy.get('#lastName').type(testUser.last_name)
      cy.get('#email').type(testUser.email)
      cy.get('#pass').type(testUser.password)
      cy.get('#confPass').type(testUser.password)

      cy.contains('Submit').click()
      cy.wait('@registerUser')

      cy.contains('[bizZen]').click()

      //logging in
      cy.contains('Log In').click()
      
      cy.get('#username').type(testUser.email)
      cy.get('#password').type(testUser.password)
      cy.contains('Submit').click()

      cy.wait('@loginUser')


  })

  it('should register & login as a business', () => {
    cy.visit('http://localhost:4200/')

    let testUser:User = new User("2", "Business", "Last Name", "business@email.com", "pass", "Business", [])

    cy.intercept('POST', 'register', {
      user:testUser
    }).as('registerUser')

    cy.intercept('POST', 'login', {
      user:testUser,
      first_name: testUser.first_name
    }).as('loginUser')

    cy.contains('Register').click()

    cy.get('#firstName').type(testUser.first_name)
    cy.get('#lastName').type(testUser.last_name)
    cy.get('#email').type(testUser.email)
    cy.get('#pass').type(testUser.password)
    cy.get('#confPass').type(testUser.password)
    cy.get('[type="checkbox"]').check({force:true})

    cy.contains('Submit').click()
    cy.wait('@registerUser')

    cy.contains('[bizZen]').click()

    //logging in
    cy.contains('Log In').click()
    
    cy.get('#username').type(testUser.email)
    cy.get('#password').type(testUser.password)
    cy.contains('Submit').click()

    cy.wait('@loginUser')

  })
})