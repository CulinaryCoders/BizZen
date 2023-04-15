import { Appointment } from "./appointment";
import { Service } from "./service";

export class ServiceAppointment {
    constructor(
        public appointment:Appointment,
        public service:Service

    ) {}
}
