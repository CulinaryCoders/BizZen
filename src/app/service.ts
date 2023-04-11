
export class Service {

    constructor(
        public serviceId : string,
        public name : string,
        public description : string,
        public start_date_time: Date,
        public length: number,
        public capacity: number,
        public price: number

    ) {}

}