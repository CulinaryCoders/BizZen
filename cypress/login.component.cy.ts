import { TestBed } from '@angular/core/testing';
import { LoginComponent } from "src/app/login/login.component";
import { RouterTestingModule } from "@angular/router/testing";
import { HttpClientTestingModule } from "@angular/common/http/testing";
import { FormsModule } from '@angular/forms';


describe('LoginComponent', () => {
  
  beforeEach(async () => {
    await TestBed.configureTestingModule({

      //importing stuff that login uses
      imports: [
        RouterTestingModule,
        HttpClientTestingModule,
        FormsModule
      ],

    }).compileComponents();
  });

  it('button should be disabled to start', () => {
    cy.mount(LoginComponent)   //mount (add) the component
    cy.get('#submitButton').should('be.disabled')
  })

  it('button should be enabled when forms are filled', () => {
    cy.mount(LoginComponent)   //mount (add) the component
    cy.get('#username').type('test username')
    cy.get('#password').type('test password')
    cy.get('#submitButton').should('be.enabled')
  })

})