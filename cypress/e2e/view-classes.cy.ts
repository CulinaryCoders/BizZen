import { Appointment } from "src/app/appointment"
import { Service } from "src/app/service"
import { ServiceAppointment } from "src/app/service-appointment"
import { User } from "src/app/user"

describe('template spec', () => {

    it('should log in as a user and join a class', () => {

        cy.visit('http://localhost:4200/')

        let testUser:User = new User("1", "First Name", "Last Name", "test@email.com", "pass", "User", [])
        let testService:Service = new Service("1", "Test Service", "A service for testing", 
            new Date("4/19/2023 12:00:00"), 120, 10, 200)
        
        let testServices:Service[] = [testService]
        let testServiceUsers:User[] = []
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

        cy.intercept('POST', 'appointment', {
            appointment: appointment,
            service: testService
        }).as('joinService')

        cy.intercept('GET', 'user/1/service-appointments', {
            body:testUserServices
        }).as('getUserServices')


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
        cy.contains('Join').click()
        testUserServices.push(new ServiceAppointment(appointment, testService));
        cy.wait('@joinService')

        //profile page
        cy.get('#profile').click({force:true})
        cy.get('#appointments').click()
        cy.wait('@getUserServices')

    
    })

  })