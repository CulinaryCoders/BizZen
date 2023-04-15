import { Service } from "./service";

export class User {
    constructor(
        public ID:string,
        public first_name : string,
        public last_name : string,
        public email : string,
        public password : string,
        public accountType : string,

        public classes : Service[]

    ) {}
}
