export const registerUser = (email:string, password:string) =>
{
    let endpoint = '/login';
    cy.intercept({
        method: 'POST',
        path: 'register'
    }).as('registerUser');
};