import { Service } from "./service";

export class User {
    constructor(
        public userId : string,
        public username : string,
        public password : string,
        public accountType : string,

        public classes : Service[]

    ) {}
}