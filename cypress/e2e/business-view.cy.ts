import { Appointment } from "src/app/appointment"
import { Service } from "src/app/service"
import { ServiceAppointment } from "src/app/service-appointment"
import { User } from "src/app/user"

describe('template spec', () => {

    it('should edit classes', () => {

        cy.visit('http://localhost:4200/')

        let testUser:User = new User("1", "Business", "Last Name", "test@email.com", "pass", "Business", [])
        let testService:Service = new Service("1", "Test Service", "A service for testing", 
            new Date("4/19/2023 12:00:00"), 120, 10, 200)
        
        let testServices:Service[] = [testService]
        let testServiceUsers:User[] = [testUser]
        let testUserServices:ServiceAppointment[] = []

        let appointment:Appointment = new Appointment("1", testService, testUser);

        //http interceptions
        cy.intercept('POST', 'login', {
            body:testUser
        }).as('loginUser')

        cy.intercept('GET', 'services', {
            body: testServices
        }).as('getServices')

        cy.intercept('GET', 'service/1/users', {
            body:testServiceUsers
        }).as('getServiceUsers')

        cy.intercept('PUT', 'service/1', {
            body:testServiceUsers
        }).as('updateService')

        //logging in
        cy.contains('Log In').click()
        
        cy.get('#username').type(testUser.email)
        cy.get('#password').type(testUser.password)
        cy.contains('Submit').click()

        cy.wait('@loginUser')

        //dashboard
        cy.get('#classes').click()

        cy.wait('@getServices')

        //service page
        cy.contains('More Info').click()
        cy.wait('@getServiceUsers')
        cy.contains('Edit').click()
        cy.get('#serviceDescription').type(' new stuff')
        cy.contains('Save Edit').click()
        cy.wait('@updateService')
        
    })

    
    it('should add a new class', () => {

        cy.visit('http://localhost:4200/')

        let testUser:User = new User("1", "Business", "Last Name", "test@email.com", "pass", "Business", [])
        let testService:Service = new Service("1", "Test Service", "A service for testing", 
            new Date("4/19/2023 12:00:00"), 120, 10, 200)

        let newTestService:Service = new Service("2", "Test Service Two", "A new test service to create", 
            new Date("4/20/2023 12:00:00"), 120, 10, 100)
        
        let testServices:Service[] = [testService]

        //http interceptions
        cy.intercept('POST', 'login', {
            body:testUser
        }).as('loginUser')

        cy.intercept('GET', 'services', {
            body: testServices
        }).as('getServices')


        //logging in
        cy.contains('Log In').click()
        
        cy.get('#username').type(testUser.email)
        cy.get('#password').type(testUser.password)
        cy.contains('Submit').click()

        cy.wait('@loginUser')

        //dashboard
        cy.get('#classes').click()
        cy.wait('@getServices')
        cy.get('#addService').click()

        //add service
        cy.get('#business-name').type(newTestService.name)
        cy.get('#business-description').type(newTestService.desc)

        cy.get('#opening-time').invoke('removeAttr', 'type').type('2022-12-01 11:00:AM').trigger('change');

        cy.get('#closing-time').type(newTestService.length.toString())
        cy.get('#num-participants').type(newTestService.capacity.toString())
        cy.get('#price').type(newTestService.price.toString())
        cy.get('#cancellation-fee').type('10')
        testServices.push(newTestService)

        cy.get('#addClass').click()

        //dashboard
        cy.wait('@getServices')
    })

  })