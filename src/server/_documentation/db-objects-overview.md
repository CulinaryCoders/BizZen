| **MVP** | **Object Name**      | **Description**                                                                                                                                   |
|----------|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| x        | Address              | Address info for a particular record (User or Business)                                                                                           |
| x        | Appointment          | Service appointments that can be scheduled between users and businesses                                                                           |
| x        | Business             | Businesses / Organizations                                                                                                                        |
| x        | ContactInfo          | Contact info for a particular record (User or Business)                                                                                           |
| x        | Office               | Each business has at least one office, but larger ones can have multiple                                                                          |
| x        | OperatingHours       | Operating hours for a particular Office                                                                                                           |
| x        | Payment              | Payment records associated with each transaction                                                                                                  |
| x        | ProfilePic           | Used to reference profile pic path data                                                                                                           |
| x        | Resource             | Equipment, rooms, etc. that are limited and need to be taken into consideration for appointment availability and booking                          |
| x        | ResourceAvailability | Schedule of availability blocks for each resource (can have multiple blocks of availability per day)                                              |
| x        | Service              | Services offered by each business                                                                                                                 |
| x        | ServiceOffering      | Used for appointment booking and tied to specific staff member and/or resource                                                                    |
| x        | StaffShifts          | Schedule of availability blocks for each staff member (can have multiple blocks of availability per day)                                          |
| x        | TimeSlots            | Generic table of time slots used for scheduling, appointments, and availability logic (10 minute increments across a full 24 hour day - 144 rows) |
| x        | Transaction          | Payment-for-service transactions between customers and businesses                                                                                 |
| x        | User                 | User accounts                                                                                                                                     |
| x        | UserPermissions      | Generic permission levels that can be assigned to user accounts                                                                                   |
|          | Interaction          | Email, SMS, chat, etc. interactions between businesses and customers (for logging/reporting purposes)                                             |
|          | Invoice              | Invoice made up of all the transactions being billed                                                                                              |
|          | Review               | Customer review of business, service, staff, and/or transaction                                                                                   |
|          | UserPreferences      | Preference settings set/managed by each user that are specific to their user account                                                              |
