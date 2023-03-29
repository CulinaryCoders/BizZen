import { ServiceOffering } from "./service-offering";

export class Service {
    constructor(
        public serviceId : string,
        public name : string,
        public description : string,

        public offering : ServiceOffering

    ) {}
}