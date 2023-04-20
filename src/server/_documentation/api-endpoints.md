| **URI**                                 | **DB Object**          | **Function Called**            | **Request Type** | **Description **                                 |
|-----------------------------------------|------------------------|--------------------------------|------------------|--------------------------------------------------|
| **/**                                   | N/A (Static HTTP Page) | serveTableOfContents           | GET              | Backend / API reference links and documentation  |
| **/home**                               | N/A (Static HTTP Page) | serveTableOfContents           | GET              | Backend / API reference links and documentation  |
| **/index**                              | N/A (Static HTTP Page) | serveTableOfContents           | GET              | Backend / API reference links and documentation  |
| **/register**                           | User                   | CreateUser                     | POST             |                                                  |
| **/login**                              | User                   | Authenticate                   | POST             |                                                  |
| **/user/{id}**                          | User                   | GetUser                        | GET              |                                                  |
| **/user/{id}**                          | User                   | UpdateUser                     | PUT              |                                                  |
| **/user/{id}**                          | User                   | DeleteUser                     | DELETE           |                                                  |
| **/user/{id}/service-appointments**     | User                   | GetUserServiceAppointments     | GET              |                                                  |
| **/business**                           | Business               | CreateBusiness                 | POST             |                                                  |
| **/business/{id}**                      | Business               | GetBusiness                    | GET              |                                                  |
| **/business/{id}**                      | Business               | UpdateBusiness                 | PUT              |                                                  |
| **/business/{id}**                      | Business               | DeleteBusiness                 | DELETE           |                                                  |
| **/business/{id}/services**             | Business               | GetBusinessServices            | GET              |                                                  |
| **/business/{id}/service-appointments** | Business               | GetBusinessServiceAppointments | GET              |                                                  |
| **/service**                            | Service                | CreateService                  | POST             |                                                  |
| **/service/{id}**                       | Service                | GetService                     | GET              |                                                  |
| **/service/{id}**                       | Service                | UpdateService                  | PUT              |                                                  |
| **/service/{id}**                       | Service                | DeleteService                  | DELETE           |                                                  |
| **/services**                            | Service     | GetServices                  | GET    |                                                                                 |
| **/service/{id}/users**                  | Service     | GetListOfEnrolledUsers       | GET    |                                                                                 |
| **/service/{id}/user-count**             | Service     | GetEnrolledUsersCount        | GET    |                                                                                 |
| **/service/{service-id}/user/{user-id}** | Service     | GetUserEnrolledStatus        | GET    |                                                                                 |
| **/service/{id}/appointments**           | Service     | GetActiveServiceAppointments | GET    |                                                                                 |
| **/service/{id}/appointments/active**    | Service     | GetActiveServiceAppointments | GET    |                                                                                 |
| **/service/{id}/appointments/all**       | Service     | GetServiceAppointments       | GET    |                                                                                 |
| **/appointment**                         | Appointment | CreateAppointment            | POST   |                                                                                 |
| **/appointment/{id}**                    | Appointment | GetAppointment               | GET    |                                                                                 |
| **/appointment/{id}**                    | Appointment | UpdateAppointment            | UPDATE |                                                                                 |
| **/appointment/{id}**                    | Appointment | DeleteAppointment            | DELETE |                                                                                 |
| **/appointment/{id}/cancel**             | Appointment | CancelAppointment            | POST   |                                                                                 |
| **/appointments**                        | Appointment | GetActiveAppointments        | GET    |                                                                                 |
| **/appointments/active**                 | Appointment | GetActiveAppointments        | GET    | Same as /appointments, just added for consistent naming convention alternative  |
| **/appointments/all**                    | Appointment | GetAppointments              | GET    |                                                                                 |
| **/invoice**                             | Invoice     | CreateInvoice                | POST   |                                                                                 |
| **/invoice/{id}**                        | Invoice     | GetInvoice                   | GET    |                                                                                 |
| **/invoice/{id}**                        | Invoice     | UpdateInvoice                | UPDATE |                                                                                 |
| **/invoice/{id}**                        | Invoice     | DeleteInvoice                | DELETE |                                                                                 |
| **/invoices**                            | Invoice     | GetInvoices                  | GET    |                                                                                 |
