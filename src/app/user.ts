import { Service } from "./service";

export class User {
    constructor(
        public ID:string,
        public firstName : string,
        public lastName : string,
        public email : string,
        public password : string,
        public accountType : string,

        public classes : Service[]

    ) {}
}
